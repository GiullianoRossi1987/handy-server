package pkg

import (
	"context"
	"fmt"
	types "types/database/users"
	errors "types/errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetWorkerById(id int32, conn *pgxpool.Pool) (*types.WorkersRecord, error) {
	var worker *types.WorkersRecord
	err := conn.QueryRow(context.Background(),
		"SELECT * FROM workers WHERE id = $1;", id).Scan(&worker)
	if err != nil {
		return nil, err
	}
	return worker, nil
}

func GetWorkerByUserId(id int32, conn *pgxpool.Pool) (*types.WorkersRecord, error) {
	var worker *types.WorkersRecord
	err := conn.QueryRow(context.Background(),
		"SELECT * FROM workers WHERE id_user = $1;", id).Scan(&worker)
	if err != nil {
		return nil, err
	}
	return worker, nil
}

func AddWorker(record types.WorkersRecord, conn *pgxpool.Pool) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(context.Background(),
		"INSERT INTO workers (id_user, uuid, fullname) VALUES ($1, $2, $3);",
		record.UserId, record.UUID, record.Fullname)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	if commandTag.RowsAffected() != 1 {
		tx.Rollback(context.Background())
		return &errors.UnexpectedDBChangeBehaviourError{
			Operation:            "insert",
			Table:                "workers",
			ExpectedChangedLines: 1,
			ChangedLines:         int(commandTag.RowsAffected()),
			Identifier:           fmt.Sprintf("%d", record.Id),
		}
	}
	if err := tx.Commit(context.Background()); err != nil {
		tx.Rollback(context.Background())
		return err
	}
	return nil
}

func DeleteWorker(id int32, conn *pgxpool.Pool) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(context.Background(), "DELETE FROM workers WHERE id = $1;", id)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	if commandTag.RowsAffected() != 1 {
		tx.Rollback(context.Background())
		return &errors.UnexpectedDBChangeBehaviourError{
			Operation:            "delete",
			Table:                "workers",
			ExpectedChangedLines: 1,
			ChangedLines:         int(commandTag.RowsAffected()),
			Identifier:           fmt.Sprintf("%d", id),
		}
	}
	if err := tx.Commit(context.Background()); err != nil {
		tx.Rollback(context.Background())
		return err
	}
	return nil
}

func UpdateWorker(newDataRecord types.WorkersRecord, conn *pgxpool.Pool) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(context.Background(),
		"UPDATE FROM workers SET fullname = $1, active = $2, updated_at = CURRENT_TIMESTAMP() WHERE id = $3;",
		newDataRecord.Fullname, newDataRecord.Active, newDataRecord.Id)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	if commandTag.RowsAffected() != 1 {
		tx.Rollback(context.Background())
		return &errors.UnexpectedDBChangeBehaviourError{
			Operation:            "update",
			Table:                "workers",
			ExpectedChangedLines: 1,
			ChangedLines:         int(commandTag.RowsAffected()),
			Identifier:           fmt.Sprintf("%d", newDataRecord.Id),
		}
	}
	if err := tx.Commit(context.Background()); err != nil {
		tx.Rollback(context.Background())
		return err
	}
	return nil
}

func UpdateWorkerRating(newDataRecord types.WorkersRecord, conn *pgxpool.Pool) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(context.Background(),
		"UPDATE FROM workers SET avg_rating = $1 WHERE id = $2;",
		newDataRecord.Avg_ratings, newDataRecord.Id)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	if commandTag.RowsAffected() != 1 {
		tx.Rollback(context.Background())
		return &errors.UnexpectedDBChangeBehaviourError{
			Operation:            "update rating",
			Table:                "workers",
			ExpectedChangedLines: 1,
			ChangedLines:         int(commandTag.RowsAffected()),
			Identifier:           fmt.Sprintf("%d", newDataRecord.Id),
		}
	}
	if err := tx.Commit(context.Background()); err != nil {
		tx.Rollback(context.Background())
		return err
	}
	return nil
}

func DoesWorkerExists(workerId int32, conn *pgxpool.Pool) (bool, error) {
	ex, err := GetWorkerById(workerId, conn)
	if err != nil {
		return false, err
	}
	return ex != nil, nil
}
