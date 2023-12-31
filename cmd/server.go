package cmd

import (
	"database/sql"
	"fmt"
	"net/http"
	repositories_account "personal-budget/accounts/repositories"
	usecase_account "personal-budget/accounts/usecase"
	account_v1 "personal-budget/accounts/v1"
	"personal-budget/middleware"
	"personal-budget/payment"
	schedule_payment_respositories "personal-budget/schedule/repositories"
	schedule_payment_usecase "personal-budget/schedule/usecase"
	schedule_payment_v1 "personal-budget/schedule/v1"
	repositories_scheduled_transactions "personal-budget/schedule_transactions/repositories"
	usecases_scheduled_transactions "personal-budget/schedule_transactions/usecases"
	transaction_scheduled_v1 "personal-budget/schedule_transactions/v1"
	"personal-budget/token"
	repositories_transaction "personal-budget/transactions/repositories"
	usecases_transaction "personal-budget/transactions/usecases"
	transaction_v1 "personal-budget/transactions/v1"
	repositories_users "personal-budget/users/repositories"
	usecase_user "personal-budget/users/usecase"
	users_v1 "personal-budget/users/v1"
	"personal-budget/util"
	repositories_wallet "personal-budget/wallet/repositories"
	usecase_wallet "personal-budget/wallet/usecase"
	wallet_v1 "personal-budget/wallet/v1"
	webhook_usecase "personal-budget/webhook/usecase"
	webhook_v1 "personal-budget/webhook/v1"
	"personal-budget/worker"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	config     util.Config
	conn       *sql.DB
	router     *gin.Engine
	tokenMaker token.Maker
}

//// server serves out http request for our backend service

func NewServer(config util.Config, conn *sql.DB) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{tokenMaker: tokenMaker, config: config, conn: conn}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}
	server.setupRouter()
	return server, nil
}

func errorrResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) setupRouter() {
	router := gin.Default()
	// Custom 404 handler

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Personal Budget app ruuning at %s", server.config.HTTPServerAddress),
		})
	})
	groupRouter := router.Group("/api/v1")

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error":       "Resource not found",
			"route":       c.Request.URL.Path,
			"status_code": 404,
		})
	})

	////payment
	payInterface := payment.NewPastackPayment(server.config)
	//////user
	userRepo := repositories_users.NewUserAuths(server.conn)
	userCase := usecase_user.NewUsecaseUser(server.config, server.tokenMaker, userRepo)
	users_v1.NewUserRoutes(groupRouter, userCase)

	//////////////
	transRepo := repositories_transaction.NewTransactionRepo(server.conn)
	walletRepo := repositories_wallet.NewWalletRepo(server.conn)
	// walletUse := usecase_wallet.NewAccountUsecase()

	//////// webhook
	webhookusecase := webhook_usecase.NewWebhookusecase(walletRepo, transRepo, server.config)
	webhook_v1.NewWebhooksRoutes(groupRouter, webhookusecase)

	//////middleware
	groupRouter.Use(middleware.AuthMiddleware(server.tokenMaker))

	///// accounts
	acctRepo := repositories_account.NewAccountRepository(server.conn)
	acctCase := usecase_account.NewAccountUsecase(server.tokenMaker, acctRepo, server.config, payInterface)
	account_v1.NewAccountsRoutes(groupRouter, acctCase)

	//////// transactions

	transUsecase := usecases_transaction.NewTransactionUsecase(transRepo)
	transaction_v1.NewTransactionRoutes(groupRouter, transUsecase)

	////// schedule transaction
	scheduledTransRepo := repositories_scheduled_transactions.NewTransactionRepo(server.conn)
	schedusecase := usecases_scheduled_transactions.NewTransactionUsecase(scheduledTransRepo)
	transaction_scheduled_v1.NewTransactionRoutes(groupRouter, schedusecase)

	//////// wallets
	walletUse := usecase_wallet.NewWalletUsecase(walletRepo, userRepo, server.config, payInterface)
	wallet_v1.NewWalletRoutes(groupRouter, transUsecase, walletUse, payInterface, acctCase, userCase)

	//////// Scheduled Payments
	plansRepo := schedule_payment_respositories.NewSchedulePaymentRepositories(server.conn)
	planUsecase := schedule_payment_usecase.NewScheduledPaymentsUsecase(plansRepo, walletRepo, acctRepo)
	schedule_payment_v1.NewScheduledPaymentRoutes(groupRouter, planUsecase)

	workerScheduler := worker.NewWorkerScheduler(server.conn, scheduledTransRepo, plansRepo, acctRepo, payInterface)
	go workerScheduler.ServeScheduler()

	server.router = router
}
