package pkg

import (
	"context"
	types "types/database/reports"
	"utils"

	"github.com/jackc/pgx/v5"
)

func GetCustomerReportsById(workerId int32, conn *pgx.Conn) []types.CustomerReport {
	rows, err := conn.Query(context.Background(), "SELECT * FROM reports_customer WHERE id_reported_worker = $1", workerId)
	if err != nil {
		panic(err)
	}
	reports, err := pgx.CollectRows(rows, pgx.RowTo[types.CustomerReport])
	if err != nil {
		panic(err)
	}
	return reports
}

func AddCustomerReport(report types.CustomerReport, conn *pgx.Conn) {
	commandTag, err := conn.Exec(context.Background(),
		"INSERT INTO reports_customer (id_reported, tags, description) VALUES ($1, $2, $3);",
		report.Id_Customer, report.Tags, report.Description)
	utils.CheckRowsAndError(commandTag, &err, 1)
}

func DeleteCustomerReportById(reportId int32, conn *pgx.Conn) {
	commandTag, err := conn.Exec(context.Background(),
		"DELETE FROM reports_customer WHERE id = $1;", reportId,
	)
	utils.CheckRowsAndError(commandTag, &err, 1)
}

func GetCustomerReportById(reportId int32, conn *pgx.Conn) *types.CustomerReport {
	var row *types.CustomerReport
	err := conn.QueryRow(context.Background(), "SELECT * FROM reports_customer WHERE id = $1;", reportId).Scan(&row)
	if err != nil {
		panic(err)
	}
	return row
}

func RevokeCustomerReport(report types.CustomerReport, conn *pgx.Conn) {
	commandTag, err := conn.Exec(context.Background(),
		"UPDATE reports_customer SET revoked = $1, updated_at = CURRENT_TIMESTAMP() WHERE id = $2;", report.Revoked, report.Id,
	)
	utils.CheckRowsAndError(commandTag, &err, 1)
}
