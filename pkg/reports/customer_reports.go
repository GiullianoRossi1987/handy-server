package pkg

import (
	"context"
	"fmt"
	types "types/database/reports"
	errors "types/errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetCustomerReports(customerId int32, conn *pgxpool.Conn) ([]types.CustomerReport, error) {
	rows, err := conn.Query(context.Background(), "SELECT * FROM reports_customer WHERE id_reported_customer = $1", customerId)
	if err != nil {
		return nil, err
	}
	reports, err := pgx.CollectRows(rows, pgx.RowToStructByPos[types.CustomerReport])
	if err != nil {
		return nil, err
	}
	return reports, nil
}

func AddCustomerReport(report types.CustomerReport, conn *pgxpool.Conn) (*int32, error) {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	var id int32
	if err := conn.QueryRow(
		context.Background(),
		`INSERT INTO reports_customer (id_reported, tags, rating, description) VALUES ($1, $2, $3, $4) RETURNING id;`,
		report.Id_Customer, report.Tags, report.Rating, report.Description,
	).Scan(&id); err != nil {
		tx.Rollback(context.Background())
		return nil, err
	}
	if err := tx.Commit(context.Background()); err != nil {
		tx.Rollback(context.Background())
		return nil, err
	}
	return &id, nil
}

func DeleteCustomerReportById(reportId int32, conn *pgxpool.Conn) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(context.Background(),
		"DELETE FROM reports_customer WHERE id = $1;", reportId,
	)
	if commandTag.RowsAffected() != 1 {
		tx.Rollback(context.Background())
		return &errors.UnexpectedDBChangeBehaviourError{
			Operation:            "delete",
			Table:                "reports_customer",
			ExpectedChangedLines: 1,
			ChangedLines:         int(commandTag.RowsAffected()),
			Identifier:           fmt.Sprintf("%d", reportId),
		}
	}
	if err != nil {
		return err
	}
	if err := tx.Commit(context.Background()); err != nil {
		tx.Rollback(context.Background())
		return err
	}
	return nil
}

func GetCustomerReportById(reportId int32, conn *pgxpool.Conn) (*types.CustomerReport, error) {
	row, err := conn.Query(
		context.Background(),
		"SELECT * FROM reports_customer WHERE id = $1;",
		reportId,
	)
	if err != nil {
		return nil, err
	}
	report, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[types.CustomerReport])
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func UpdateCustomerReport(report types.CustomerReport, conn *pgxpool.Conn) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(
		context.Background(),
		`UPDATE reports_workers SET tags = $1, rating = $2, revoked = $3, updated_at = CURRENT_TIMESTAMP WHERE id = $4;`,
		report.Tags,
		report.Rating,
		report.Revoked,
		report.Id,
	)
	if commandTag.RowsAffected() != 1 {
		tx.Rollback(context.Background())
		return &errors.UnexpectedDBChangeBehaviourError{
			Operation:            "update/revoke",
			Table:                "reports_workers",
			ExpectedChangedLines: 1,
			ChangedLines:         int(commandTag.RowsAffected()),
			Identifier:           fmt.Sprintf("%d", report.Id),
		}
	}
	if err := tx.Commit(context.Background()); err != nil {
		tx.Rollback(context.Background())
		return err
	}
	return nil
}
