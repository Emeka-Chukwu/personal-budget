package worker

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	repositories_account "personal-budget/accounts/repositories"
	"personal-budget/payment"
	schedule_payment_respositories "personal-budget/schedule/repositories"
	model_scheduled_transactions "personal-budget/schedule_transactions/model"
	repositories_scheduled_transactions "personal-budget/schedule_transactions/repositories"
	"personal-budget/shared"
	"syscall"
	"time"
)

type WorkerScheduler struct {
	DB              *sql.DB
	scheduleTxnRepo repositories_scheduled_transactions.ScheduledTransactionRepo
	schedulerRepo   schedule_payment_respositories.SchedulePaymentRepositories
	accountRepo     repositories_account.AccountRepository
	payService      payment.PaymentInterface
}

func NewWorkerScheduler(DB *sql.DB,
	scheduleTxnRepo repositories_scheduled_transactions.ScheduledTransactionRepo,
	schedulerRepo schedule_payment_respositories.SchedulePaymentRepositories, accountRepo repositories_account.AccountRepository,
	payService payment.PaymentInterface) WorkerScheduler {
	return WorkerScheduler{DB: DB, scheduleTxnRepo: scheduleTxnRepo, schedulerRepo: schedulerRepo, payService: payService, accountRepo: accountRepo}
}

func (work *WorkerScheduler) execute() {
	batchSize := 100
	total, _ := work.schedulerRepo.FetchPlansRecordsCounter()
	var processedRows int64 = 0
	for processedRows < int64(total) {
		tx, err := work.DB.Begin()
		if err != nil {
			tx.Rollback()
			log.Println(err)
		}
		resp, err := work.schedulerRepo.FetchPlansRecords(batchSize, int(processedRows), tx)
		if err != nil {
			tx.Rollback()
			log.Println(err)
		}

		for _, plan := range resp {
			var dateString string = time.Now().Local().UTC().GoString()
			reference := fmt.Sprintf("plan-payment-%s", dateString)
			if plan.RecipientCode != "" {
				amount := plan.Amount / plan.Periods
				paymentPayload := payment.InitiateTransfer{Source: "source",
					Reason: "Reason", Amount: 9999, Recipient: plan.RecipientCode}
				transModel := model_scheduled_transactions.ScheduledTransaction{
					Type:              "plan",
					Status:            "pending",
					UserID:            plan.UserId,
					SchedulePaymentID: plan.ID,
					PaidPeriod:        plan.PaidPeriods + 1,
					Reference:         reference,
					Amount:            amount,
				}
				_, err = work.payService.Create(paymentPayload)
				if err != nil {
					tx.Rollback()
					log.Println(err)
				}
				work.scheduleTxnRepo.CreateUserTransactionTx(transModel, tx)
				plan.PaidPeriods = plan.PaidPeriods + 1
				plan.Amount = plan.PaidPeriods + 1
				plan.PayDate = shared.GetNewDate(plan.Duration, plan.Periods)
				if plan.PaidPeriods == plan.Periods {
					plan.IsCompleted = true
				}
				work.schedulerRepo.UpdatePlanTx(plan.SchedulePayment, tx)
			}
		}
		if err = tx.Commit(); err != nil {
			tx.Rollback()
			log.Println(err)
		}
		processedRows += int64(len(resp))
		fmt.Printf("Processed %d/%d rows\n", processedRows, total)

	}

}

func (work *WorkerScheduler) ServeScheduler() {
	interval := 5 * time.Second
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	stopCh := make(chan struct{})
	go func() {
		for {
			select {
			case <-time.After(interval):
				work.execute()
				fmt.Println("Task executed at", time.Now())
			case <-sigCh:
				fmt.Println("Received termination signal. Exiting...")
				close(stopCh)
				return
			}
		}
	}()
	select {
	case <-stopCh:
		fmt.Println("Task stopped. Exiting...")
		os.Exit(0)
	}
}
