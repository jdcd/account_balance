package postgresSQL

import (
	"database/sql"
	"time"

	"github.com/jdcd/account_balance/internal/domain"
	"github.com/jdcd/account_balance/pkg"
	"github.com/lib/pq"
)

type Repository struct {
	Writer *sql.DB
}

const (
	rollBackError              = "error in rollback %s"
	insertIntoReportQueryError = `INSERT INTO report (file_name, date, isProcessed, error)
								VALUES ($1, $2, $3, $4)`
	insertIntoReportQueryOK = `INSERT INTO report (file_name, date, isProcessed, total, avrDebit, avrCredit, send_to, error)
								VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	insertIntoTransactionQuery = `INSERT INTO transactions (report_id, number, date, movement, value)
								VALUES ($1, $2, $3, $4, $5)`
	insertIntoIgTransactionQuery = `INSERT INTO ignored_transactions (report_id, ignored_id, date,transaction, reason)
								VALUES ($1, $2, $3, $4, $5)`
	insertIntoTransactionByMonthQuery = `INSERT INTO transactions_by_month (report_id, month, num_transactions)
								VALUES ($1, $2, $3)`
)

func (r *Repository) SaveSuccessReport(report domain.SuccessReport) error {
	tx, err := r.Writer.Begin()
	if err != nil {
		pkg.ErrorLogger().Printf("error starting save report transaction: %s", err)
		return err
	}

	id, err := r.saveReport(report, tx)
	if err != nil {
		rbErr := tx.Rollback()
		if rbErr != nil {
			pkg.WarningLogger().Printf(rollBackError, rbErr)
		}
		return err

	}

	err = r.saveTransaction(id, report.Transactions, tx)
	if err != nil {
		rbErr := tx.Rollback()
		if rbErr != nil {
			pkg.WarningLogger().Printf(rollBackError, rbErr)
		}
		return err
	}

	err = r.saveIgnoredTransaction(id, report.IgnoredTransaction, tx)
	if err != nil {
		rbErr := tx.Rollback()
		if rbErr != nil {
			pkg.WarningLogger().Printf(rollBackError, rbErr)
		}
		return err
	}

	err = r.saveTransactionByMonth(id, report.Summary.TransactionByMonth, tx)
	if err != nil {
		rbErr := tx.Rollback()
		if rbErr != nil {
			pkg.WarningLogger().Printf(rollBackError, rbErr)
		}
		return err
	}

	err = tx.Commit()
	if err != nil {
		pkg.ErrorLogger().Printf("error making report commit %s", err)
	}

	return nil
}

func (r *Repository) SaveErrorReport(fileName string, date time.Time, err error) {
	_, err = r.Writer.Exec(insertIntoReportQueryError, fileName, date, false, err.Error())
	if err != nil {
		pkg.ErrorLogger().Printf("error saving %s error-report : %s ", fileName, err)
	}

}

func (r *Repository) saveReport(rep domain.SuccessReport, tx *sql.Tx) (int, error) {
	stmt, err := tx.Prepare(insertIntoReportQueryOK)
	if err != nil {
		pkg.ErrorLogger().Printf("error in prepare report query to report table: %s", err)
		return 0, err
	}
	defer func(reportInsertStmt *sql.Stmt) {
		err := reportInsertStmt.Close()
		if err != nil {
			pkg.WarningLogger().Printf("error closing prepare report query to report table: %s", err)
		}
	}(stmt)

	var lastInsertId int
	err = stmt.QueryRow(rep.FileName, rep.Date, true, rep.Summary.Total, rep.Summary.AvrDebitAmount,
		rep.Summary.AvrCreditAmount, pq.Array(rep.SendTo), nil).Scan(&lastInsertId)
	if err != nil {
		pkg.ErrorLogger().Printf("error inserting report %v table: %s", rep.Summary, err)
		return 0, err
	}

	return lastInsertId, err
}

func (r *Repository) saveTransaction(id int, transaction []domain.Transaction, tx *sql.Tx) error {
	stmt, err := tx.Prepare(insertIntoTransactionQuery)
	if err != nil {
		pkg.ErrorLogger().Printf("error in prepare transaction query to report table: %s", err)
		return err
	}
	defer func(reportInsertStmt *sql.Stmt) {
		err := reportInsertStmt.Close()
		if err != nil {
			pkg.WarningLogger().Printf("error closing prepare transaction query to report table: %s", err)
		}
	}(stmt)

	for _, tr := range transaction {
		_, err = stmt.Exec(id, tr.Number, tr.Date, tr.Movement, tr.Value)
		if err != nil {
			pkg.ErrorLogger().Printf("error inserting transaction %v: %s", tr, err)
			return err
		}
	}

	return nil
}

func (r *Repository) saveIgnoredTransaction(id int, transaction []domain.IgnoredTransaction, tx *sql.Tx) error {
	stmt, err := tx.Prepare(insertIntoIgTransactionQuery)
	if err != nil {
		pkg.ErrorLogger().Printf("error in prepare ignored_transaction query to report table: %s", err)
		return err
	}
	defer func(reportInsertStmt *sql.Stmt) {
		err := reportInsertStmt.Close()
		if err != nil {
			pkg.WarningLogger().Printf("error closing prepare ignored transaction query to report table: %s", err)
		}
	}(stmt)

	for _, tr := range transaction {
		_, err = stmt.Exec(id, tr.ID, tr.Date, tr.Transaction, tr.Reason)
		if err != nil {
			pkg.ErrorLogger().Printf("error inserting ignored transaction %v: %s", tr, err)
			return err
		}
	}

	return nil
}

func (r *Repository) saveTransactionByMonth(id int, tbm map[time.Month]int, tx *sql.Tx) error {
	stmt, err := tx.Prepare(insertIntoTransactionByMonthQuery)
	if err != nil {
		pkg.ErrorLogger().Printf("error in prepare transaction_by_month query to report table: %s", err)
		return err
	}
	defer func(reportInsertStmt *sql.Stmt) {
		err := reportInsertStmt.Close()
		if err != nil {
			pkg.WarningLogger().Printf("error closing prepare transaction_by_month query to report table: %s", err)
		}
	}(stmt)

	for key, value := range tbm {
		_, err = stmt.Exec(id, key, value)
		if err != nil {
			pkg.ErrorLogger().Printf("error inserting transaction by month %s: %d", key, value)
			return err
		}
	}

	return nil
}
