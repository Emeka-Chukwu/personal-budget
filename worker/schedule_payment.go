package worker

import (
	"database/sql"
	schedule_payment_respositories "personal-budget/schedule/repositories"
	repositories_scheduled_transactions "personal-budget/schedule_transactions/repositories"
)

type WorkerScheduler struct {
	DB              *sql.DB
	scheduleTxnRepo repositories_scheduled_transactions.ScheduledTransactionRepo
	schedulerRepo   schedule_payment_respositories.SchedulePaymentRepositories
}

func NewWorkerScheduler(DB *sql.DB,
	scheduleTxnRepo repositories_scheduled_transactions.ScheduledTransactionRepo,
	schedulerRepo schedule_payment_respositories.SchedulePaymentRepositories) WorkerScheduler {
	return WorkerScheduler{DB: DB, scheduleTxnRepo: scheduleTxnRepo, schedulerRepo: schedulerRepo}
}

func (work *WorkerScheduler) execute() {
	// tx, err := work.DB.Begin()

}

// func VerifierScheduler() {

// 	batchSize := 100

// 	var totalRows int64
// 	database.Db.Model(&Employee{}).
// 		Where("NOW() - employees.updated_at > interval '6 hours' AND house_address_verification_status = ?", "processing").
// 		Joins("join survey_questions sq on sq.employee_id = employees.id").
// 		Count(&totalRows)
// 	var processedRows int64 = 0
// 	for processedRows < totalRows {
// 		tx := database.Db.Begin()

// 		var employees []Employee
// 		tx.Where("NOW() - employees.updated_at > interval '6 hours' AND house_address_verification_status = ?", "processing").
// 			Joins("join survey_questions sq on sq.employee_id = employees.id").
// 			Limit(batchSize).
// 			Offset(int(processedRows)).
// 			Find(&employees)

// 		for _, employee := range employees {
// 			employee.HouseAddressVerificationStatus = "pending"
// 			employee.VerifierID = 0
// 			tx.Save(&employee)

// 			tx.Exec("UPDATE survey_questions SET updated_at= NOW(), verifier_ward=?, verifier_ward_status=?, verifier_lga=?, verifier_lga_status=?, verifier_district=?, verifier_district_status=? where employee_id=?", 0, "", 0, "", 0, "", employee.ID)
// 		}
// 		if tx.Error != nil {
// 			tx.Rollback()
// 			log.Println(tx.Error)
// 		}

// 		tx.Commit()
// 		processedRows += int64(len(employees))
// 		fmt.Printf("Processed %d/%d rows\n", processedRows, totalRows)
// 	}
// }

// func ServeScheduler() {
// 	interval := 1 * time.Hour
// 	sigCh := make(chan os.Signal, 1)
// 	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
// 	stopCh := make(chan struct{})
// 	go func() {
// 		for {
// 			select {
// 			case <-time.After(interval):
// 				VerifierScheduler()
// 				fmt.Println("Task executed at", time.Now())
// 			case <-sigCh:
// 				fmt.Println("Received termination signal. Exiting...")
// 				close(stopCh)
// 				return
// 			}
// 		}
// 	}()

// 	select {
// 	case <-stopCh:
// 		fmt.Println("Task stopped. Exiting...")
// 		os.Exit(0)
// 	}
// }
