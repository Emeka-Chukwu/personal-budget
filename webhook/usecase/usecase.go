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
	WithdrawalWebhook(uuid.UUID, payment.TransferWebHook) (any, error)
	FundWebhook(uuid.UUID, payment.ChargeWebhookResponse) (any, error)
	PayStackWebhook(userId uuid.UUID, ctx *gin.Context) (any, error)
	IsFromPaystackSource(ctx *gin.Context) (bool, error)
}
type webhookusecase struct {
	walletRepo      repositories_wallet.WalletRepo
	transactionRepo repositories_transaction.TransactionRepo
	config          util.Config
}

// FundWebhook implements Webhookusecase.
func (hook *webhookusecase) FundWebhook(userId uuid.UUID, req payment.ChargeWebhookResponse) (any, error) {
	var createTransaction bool
	transaction, err := hook.transactionRepo.GetUserTransactionByReference(req.Data.Reference)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
		createTransaction = true
	}
	if createTransaction {
		transactionParams := model_transaction.Transaction{Type: string(PaymentCredit), Status: req.Data.Status,
			Reference: req.Data.Reference, UserID: userId, Amount: int(req.Data.Amount),
		}
		_, err = hook.transactionRepo.CreateUserTransaction(transactionParams)
		if err != nil {
			return nil, err
		}
		if strings.EqualFold(req.Data.Status, string(PaymentSuccess)) {
			_, err = hook.walletRepo.FundAccount(int(req.Data.Amount), userId)
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
		_, err = hook.walletRepo.FundAccount(int(req.Data.Amount), userId)
		if err != nil {
			return nil, err
		}
	}
	return "webhook processed", nil
}

// IsFromPaystackSource implements Webhookusecase.
func (hook *webhookusecase) IsFromPaystackSource(c *gin.Context) (bool, error) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return false, err
	}
	hash := hmac.New(sha512.New, []byte(hook.config.PaystackKey))
	hash.Write(body)
	calculatedHash := hex.EncodeToString(hash.Sum(nil))
	return calculatedHash == c.GetHeader("x-paystack-signature"), nil
}

// PayStackWebhook implements Webhookusecase.
func (hook *webhookusecase) PayStackWebhook(userId uuid.UUID, ctx *gin.Context) (any, error) {
	isFromValidSourced, err := hook.IsFromPaystackSource(ctx)
	if err != nil || !isFromValidSourced {
		return nil, err
	}
	var req Event
	err = json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	var result any
	switch req.Event {
	case string(ChargeSuccess):
		payload := util.GetBody[payment.ChargeWebhookResponse](ctx)
		result, err = hook.FundWebhook(userId, payload)
	case string(TranferSuccess), string(TranferReversed), string(TranferFailed):
		payload := util.GetBody[payment.TransferWebHook](ctx)
		result, err = hook.WithdrawalWebhook(userId, payload)

	}
	return result, err
}

// WithdrawalWebhook implements Webhookusecase.
func (hook *webhookusecase) WithdrawalWebhook(userId uuid.UUID, req payment.TransferWebHook) (any, error) {
	var createTransaction bool
	typeWithdrawal := string(PaymentDebit)
	transaction, err := hook.transactionRepo.GetUserTransactionByReference(req.Data.Reference)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
		createTransaction = true
	}
	if createTransaction {
		if strings.EqualFold(req.Data.Status, string(PaymentReversed)) {
			typeWithdrawal = string(PaymentCredit)
		}
		transactionParams := model_transaction.Transaction{Type: typeWithdrawal, Status: req.Data.Status,
			Reference: req.Data.Reference, UserID: userId, Amount: int(req.Data.Amount),
		}
		_, err = hook.transactionRepo.CreateUserTransaction(transactionParams)
		if err != nil {
			return nil, err
		}
		if strings.EqualFold(req.Data.Status, string(PaymentSuccess)) {
			err = hook.walletRepo.Withdrawal(userId, int(req.Data.Amount), func() error {
				return nil
			})
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
		err = hook.walletRepo.Withdrawal(userId, int(req.Data.Amount), func() error {
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	if strings.EqualFold(req.Data.Status, string(PaymentReversed)) {
		_, err = hook.transactionRepo.UpdateUserTransactionType(req.Data.Reference, string(PaymentCredit))
		if err != nil {
			return nil, err
		}
		_, err = hook.walletRepo.FundAccount(int(req.Data.Amount), userId)
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
