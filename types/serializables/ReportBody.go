package types

import (
	"encoding/json"
	"time"

	db "types/database/reports"
)

type ReportBody struct {
	Id          int32     `json:"id,omitempty"`
	IdReported  int32     `json:"reported" binding:"required"`
	Tags        []string  `json:"tags,omitempty" binding:"required"`
	Rating      float32   `json:"rating" binding:"required"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at" binding:"required"`
	UpdatedAt   time.Time `json:"updated_at" binding:"required"`
	Revoked     bool      `json:"revoked" binding:"required"`
}

func (rb *ReportBody) CustomerRecord() *db.CustomerReport {
	if rb == nil {
		return nil
	}
	return &db.CustomerReport{
		Id:          rb.Id,
		Id_Customer: rb.IdReported,
		Tags:        rb.Tags,
		Rating:      rb.Rating,
		Description: rb.Description,
		CreatedAt:   rb.CreatedAt,
		UpdatedAt:   rb.UpdatedAt,
		Revoked:     rb.Revoked,
	}
}

func (rb *ReportBody) WorkerRecord() *db.WorkerReport {
	if rb == nil {
		return nil
	}
	return &db.WorkerReport{
		Id:          rb.Id,
		Id_Worker:   rb.IdReported,
		Tags:        rb.Tags,
		Rating:      rb.Rating,
		Description: rb.Description,
		CreatedAt:   rb.CreatedAt,
		UpdatedAt:   rb.UpdatedAt,
		Revoked:     rb.Revoked,
	}
}

func (rb *ReportBody) ToJSON() (string, error) {
	val, err := json.Marshal(rb)
	if err != nil {
		return "nil", nil
	}
	return string(val), nil
}

func SerializeCustomerReport(record *db.CustomerReport) *ReportBody {
	if record == nil {
		return nil
	}
	return &ReportBody{
		Id:          record.Id,
		IdReported:  record.Id_Customer,
		Tags:        record.Tags,
		Rating:      record.Rating,
		Description: record.Description,
		CreatedAt:   record.CreatedAt,
		UpdatedAt:   record.UpdatedAt,
		Revoked:     record.Revoked,
	}
}

func SerializeWorkerReport(record *db.WorkerReport) *ReportBody {
	if record == nil {
		return nil
	}
	return &ReportBody{
		Id:          record.Id,
		IdReported:  record.Id_Worker,
		Tags:        record.Tags,
		Rating:      record.Rating,
		Description: record.Description,
		CreatedAt:   record.CreatedAt,
		UpdatedAt:   record.UpdatedAt,
		Revoked:     record.Revoked,
	}
}
