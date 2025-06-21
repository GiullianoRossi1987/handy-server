package pkg

import (
	"context"
	types "types/database/reports"
	"utils"

	"github.com/jackc/pgx/v5"
)

func GetWorkerReportsById(workerId int32, conn *pgx.Conn) []types.WorkerReport {
	rows, err := conn.Query(context.Background(), "SELECT * FROM reports_workers WHERE id_reported_worker = $1", workerId)
	if err != nil {
		panic(err)
	}
	reports, err := pgx.CollectRows(rows, pgx.RowTo[types.WorkerReport])
	if err != nil {
		panic(err)
	}
	return reports
}

func AddWorkerReport(report types.WorkerReport, conn *pgx.Conn) {
	commandTag, err := conn.Exec(context.Background(),
		"INSERT INTO reports_workers (id_reported, tags, description) VALUES ($1, $2, $3);",
		report.Id_Worker, report.Tags, report.Description)
	utils.CheckRowsAndError(commandTag, &err, 1)
}

func DeleteWorkerReportById(reportId int32, conn *pgx.Conn) {
	commandTag, err := conn.Exec(context.Background(),
		"DELETE FROM reports_workers WHERE id = $1;", reportId,
	)
	utils.CheckRowsAndError(commandTag, &err, 1)
}

func GetWorkerReportById(reportId int32, conn *pgx.Conn) *types.WorkerReport {
	var row *types.WorkerReport
	err := conn.QueryRow(context.Background(), "SELECT * FROM reports_workers WHERE id = $1;", reportId).Scan(&row)
	if err != nil {
		panic(err)
	}
	return row
}

func RevokeWorkerReport(report types.WorkerReport, conn *pgx.Conn) {
	commandTag, err := conn.Exec(context.Background(),
		"UPDATE reports_workers SET revoked = $1, updated_at = CURRENT_TIMESTAMP() WHERE id = $2;", report.Revoked, report.Id,
	)
	utils.CheckRowsAndError(commandTag, &err, 1)
}
