package services

import (
	reports "pkg/reports"
	types "types/database/reports"
	serial "types/serializables"
	"utils"

	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetWorkerReports(pool *pgxpool.Pool, workerId int32) ([]serial.ReportBody, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	records, err := reports.GetWorkerReportsById(workerId, conn)
	conn.Release()
	if err != nil {
		return nil, err
	}
	reports := utils.MapCar(
		records,
		func(record types.WorkerReport) serial.ReportBody {
			return *serial.SerializeWorkerReport(&record)
		},
	)
	return reports, nil
}

func GetWorkerReportById(pool *pgxpool.Pool, reportId int32) (*serial.ReportBody, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	report, err := reports.GetWorkerReportById(reportId, conn)
	conn.Release()
	if err != nil {
		return nil, err
	}
	return serial.SerializeWorkerReport(report), nil
}

func AddWorkerReport(pool *pgxpool.Pool, req serial.ReportBody) (*int32, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	record := req.WorkerRecord()
	record.Id = 0
	id, err := reports.AddWorkerReport(*record, conn)
	conn.Release()
	if err != nil {
		return nil, err
	}
	return id, err
}

func UpdateWorkerReport(pool *pgxpool.Pool, req serial.ReportBody, reportId int32) error {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	record := req.WorkerRecord()
	record.Id = reportId
	if err := reports.UpdateWorkerReport(*record, conn); err != nil {
		return err
	}
	conn.Release()
	return nil
}

func DeleteWorkerReport(pool *pgxpool.Pool, reportId int32) error {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	if err := reports.DeleteWorkerReportById(reportId, conn); err != nil {
		return err
	}
	conn.Release()
	return nil
}
