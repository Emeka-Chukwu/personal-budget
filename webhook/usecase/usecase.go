package webhook_usecase

import (
	"crypto/hmac"
	"crypto/sha512"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"personal-budget/payment"
	model_transaction "personal-budget/transactions/model"
	repositories_transaction "personal-budget/transactions/repositories"
	"personal-budget/util"
	repositories_wallet "personal-budget/wallet/repositories"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ////// Amounts are all in kobos
type EventType string

type PaymentStatus string

type PaymentType string

var (
	ChargeSuccess   EventType = "charge.success"
	TranferSuccess  EventType = "transfer.success"
	TranferFailed   EventType = "transfer.failed"
	TranferReversed EventType = "transfer.reversed"
)
var (
	PaymentSuccess  PaymentStatus = "success"
	PaymentFailed   PaymentStatus = "failed"
	PaymentReversed PaymentStatus = "reversed"
)
var (
	PaymentCredit PaymentType = "credit"
	PaymentDebit  PaymentType = "debit"
)

type Webhookusecase interface {
	WithdrawalWebhook(payment.TransferWebHook) (any, error)
	FundWebhook(payment.ChargeWebhookResponse) (any, error)
	PayStackWebhook(ctx *gin.Context) (any, error)
	IsFromPaystackSource(ctx *gin.Context) (bool, []byte, error)
}
type webhookusecase struct {
	walletRepo      repositories_wallet.WalletRepo
	transactionRepo repositories_transaction.TransactionRepo
	config          util.Config
}

// FundWebhook implements Webhookusecase.
func (hook *webhookusecase) FundWebhook(req payment.ChargeWebhookResponse) (any, error) {
	var createTransaction bool
	transaction, err := hook.transactionRepo.GetUserTransactionByReference(req.Data.Reference)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
		createTransaction = true
	}
	userId, err := uuid.Parse(req.Data.Metadata.UserID)
	if err != nil {
		return nil, err
	}
	amount := int(req.Data.Amount) / 100
	if createTransaction {

		transactionParams := model_transaction.Transaction{Type: string(PaymentCredit), Status: req.Data.Status,
			Reference: req.Data.Reference, UserID: userId, Amount: amount,
		}
		_, err = hook.transactionRepo.CreateUserTransaction(transactionParams)
		if err != nil {
			return nil, err
		}
		if strings.EqualFold(req.Data.Status, string(PaymentSuccess)) {
			_, err = hook.walletRepo.FundAccount(amount, userId)
			if err != nil {
				return nil, err
			}
		}
		return "webhook processed", nil
	}
	if strings.EqualFold(transaction.Status, string(PaymentSuccess)) {
		return nil, errors.New("payment processed already")
	}
	if strings.EqualFold(transaction.Status, req.Data.Status) {
		return nil, errors.New("error ")
	}
	_, err = hook.transactionRepo.UpdateUserTransaction(req.Data.Reference, req.Data.Status)
	if err != nil {
		return nil, err
	}
	if strings.EqualFold(req.Data.Status, string(PaymentSuccess)) {
		_, err = hook.walletRepo.FundAccount(int(amount), userId)
		if err != nil {
			return nil, err
		}
	}
	return "webhook processed", nil
}

// IsFromPaystackSource implements Webhookusecase.
func (hook *webhookusecase) IsFromPaystackSource(c *gin.Context) (bool, []byte, error) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return false, body, err
	}
	hash := hmac.New(sha512.New, []byte(hook.config.PaystackKey))
	hash.Write(body)
	calculatedHash := hex.EncodeToString(hash.Sum(nil))
	return calculatedHash == c.GetHeader("x-paystack-signature"), body, nil
}

// PayStackWebhook implements Webhookusecase.
func (hook *webhookusecase) PayStackWebhook(ctx *gin.Context) (any, error) {
	isFromValidSourced, body, err := hook.IsFromPaystackSource(ctx)
	if err != nil || !isFromValidSourced {
		return nil, err
	}
	var req Event
	if err := json.Unmarshal(body, &req); err != nil {
		return false, err
	}
	if err != nil {
		return nil, err
	}
	var result any
	switch req.Event {
	case string(ChargeSuccess):
		var payload payment.ChargeWebhookResponse
		if err := json.Unmarshal(body, &payload); err != nil {
			return false, err
		}
		result, err = hook.FundWebhook(payload)
	case string(TranferSuccess), string(TranferReversed), string(TranferFailed):
		var payload payment.TransferWebHook
		if err := json.Unmarshal(body, &payload); err != nil {
			return false, err
		}
		result, err = hook.WithdrawalWebhook(payload)
	}
	return result, err
}

// WithdrawalWebhook implements Webhookusecase.
func (hook *webhookusecase) WithdrawalWebhook(req payment.TransferWebHook) (any, error) {
	/////// transaction record for this payout has already been created when the transfer api was called
	transaction, err := hook.transactionRepo.GetUserTransactionByReference(req.Data.Reference)
	if err != nil {
		return nil, err
	}
	if strings.EqualFold(transaction.Status, string(PaymentSuccess)) {
		return nil, errors.New("payment processed already")
	}
	if strings.EqualFold(transaction.Status, req.Data.Status) {
		return nil, errors.New("error ")
	}
	_, err = hook.transactionRepo.UpdateUserTransaction(req.Data.Reference, req.Data.Status)
	if err != nil {
		return nil, err
	}
	if strings.EqualFold(req.Data.Status, string(PaymentReversed)) {
		_, err = hook.transactionRepo.UpdateUserTransactionType(req.Data.Reference, string(PaymentCredit))
		if err != nil {
			return nil, err
		}
		_, err = hook.walletRepo.FundAccount(int(req.Data.Amount), transaction.UserID)
		if err != nil {
			return nil, err
		}
	}
	return "webhook processed", nil
}

func NewWebhookusecase(walletRepo repositories_wallet.WalletRepo,
	transactionRepo repositories_transaction.TransactionRepo,
	config util.Config) Webhookusecase {
	return &webhookusecase{walletRepo: walletRepo, transactionRepo: transactionRepo, config: config}
}

type Event struct {
	Event string `json:"event"`
}
