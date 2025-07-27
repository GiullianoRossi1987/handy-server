package services

import (
	reports "pkg/reports"
	types "types/database/reports"
	serial "types/serializables"
	"utils"

	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetCustomerReports(pool *pgxpool.Pool, customerId int32) ([]serial.ReportBody, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	records, err := reports.GetCustomerReports(customerId, conn)
	conn.Release()
	if err != nil {
		return nil, err
	}
	reports := utils.MapCar(
		records,
		func(record types.CustomerReport) serial.ReportBody {
			return *serial.SerializeCustomerReport(&record)
		},
	)
	return reports, nil
}

func GetCustomerReportById(pool *pgxpool.Pool, reportId int32) (*serial.ReportBody, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	report, err := reports.GetCustomerReportById(reportId, conn)
	conn.Release()
	if err != nil {
		return nil, err
	}
	return serial.SerializeCustomerReport(report), nil
}

func AddCustomerReport(pool *pgxpool.Pool, req serial.ReportBody) (*int32, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	record := req.CustomerRecord()
	record.Id = 0
	id, err := reports.AddCustomerReport(*record, conn)
	conn.Release()
	if err != nil {
		return nil, err
	}
	return id, err
}

func UpdateCustomerReport(pool *pgxpool.Pool, req serial.ReportBody, reportId int32) error {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	record := req.CustomerRecord()
	record.Id = reportId
	if err := reports.UpdateCustomerReport(*record, conn); err != nil {
		return err
	}
	conn.Release()
	return nil
}

func DeleteCustomerReport(pool *pgxpool.Pool, reportId int32) error {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	if err := reports.DeleteCustomerReportById(reportId, conn); err != nil {
		return err
	}
	conn.Release()
	return nil
}
