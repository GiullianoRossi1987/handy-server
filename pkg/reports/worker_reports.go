package pkg

import (
	"context"
	"fmt"
	types "types/database/reports"
	errors "types/errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetWorkerReportsById(workerId int32, conn *pgxpool.Pool) ([]types.WorkerReport, error) {
	rows, err := conn.Query(context.Background(), "SELECT * FROM reports_workers WHERE id_reported_worker = $1", workerId)
	if err != nil {
		return nil, err
	}
	reports, err := pgx.CollectRows(rows, pgx.RowToStructByPos[types.WorkerReport])
	if err != nil {
		return nil, err
	}
	return reports, nil
}

func AddWorkerReport(report types.WorkerReport, conn *pgxpool.Pool) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(
		context.Background(),
		"INSERT INTO reports_workers (id_reported, tags, description) VALUES ($1, $2, $3);",
		report.Id_Worker,
		report.Tags,
		report.Description,
	)
	if commandTag.RowsAffected() != 1 {
		tx.Rollback(context.Background())
		return &errors.UnexpectedDBChangeBehaviourError{
			Operation:            "insert",
			Table:                "reports_worker",
			ExpectedChangedLines: 1,
			ChangedLines:         int(commandTag.RowsAffected()),
			Identifier:           fmt.Sprintf("%d", report.Id),
		}
	}
	if err != nil {
		return err
	}
	if err := tx.Commit(context.Background()); err != nil {
		return err
	}
	return nil
}

func DeleteWorkerReportById(reportId int32, conn *pgxpool.Pool) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(context.Background(),
		"DELETE FROM reports_workers WHERE id = $1;", reportId,
	)
	if commandTag.RowsAffected() != 1 {
		tx.Rollback(context.Background())
		return &errors.UnexpectedDBChangeBehaviourError{
			Operation:            "delete",
			Table:                "reports_workers",
			ExpectedChangedLines: 1,
			ChangedLines:         int(commandTag.RowsAffected()),
			Identifier:           fmt.Sprintf("%d", reportId),
		}
	}
	if err != nil {
		return err
	}
	if err := tx.Commit(context.Background()); err != nil {
		return err
	}
	return nil
}

func GetWorkerReportById(reportId int32, conn *pgxpool.Pool) (*types.WorkerReport, error) {
	var row *types.WorkerReport
	if err := conn.QueryRow(
		context.Background(),
		"SELECT * FROM reports_workers WHERE id = $1;",
		reportId,
	).Scan(&row); err != nil {
		return nil, err
	}
	return row, nil
}

func RevokeWorkerReport(report types.WorkerReport, conn *pgxpool.Pool) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(
		context.Background(),
		"UPDATE reports_workers SET revoked = $1, updated_at = CURRENT_TIMESTAMP() WHERE id = $2;",
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
