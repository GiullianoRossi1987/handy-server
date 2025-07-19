package pkg

import (
	"context"
	"fmt"
	types "types/database/users"
	errors "types/errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetWorkerById(id int32, conn *pgxpool.Conn) (*types.WorkersRecord, error) {
	row, err := conn.Query(
		context.Background(),
		"SELECT * FROM workers WHERE id = $1;",
		id,
	)
	if err != nil {
		return nil, err
	}
	worker, err := pgx.CollectOneRow(row, pgx.RowToAddrOfStructByPos[types.WorkersRecord])
	if err != nil {
		return nil, err
	}
	return worker, nil
}

func GetWorkerByUUID(uuid string, conn *pgxpool.Conn) (*types.WorkersRecord, error) {
	row, err := conn.Query(
		context.Background(),
		`SELECT * FROM workers WHERE uuid = $1;`,
		uuid,
	)
	if err != nil {
		return nil, err
	}
	worker, err := pgx.CollectOneRow(row, pgx.RowToAddrOfStructByPos[types.WorkersRecord])
	if err != nil {
		return nil, err
	}
	return worker, nil
}

func GetWorkerByUserId(id int32, conn *pgxpool.Conn) (*types.WorkersRecord, error) {
	row, err := conn.Query(
		context.Background(),
		`SELECT * FROM workers WHERE id_user = $1;`,
		id,
	)
	if err != nil {
		return nil, err
	}
	worker, err := pgx.CollectOneRow(row, pgx.RowToAddrOfStructByPos[types.WorkersRecord])
	if err != nil {
		return nil, err
	}
	return worker, nil
}

func AddWorker(record types.WorkersRecord, conn *pgxpool.Conn) (*int32, error) {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	var id int32
	if err := conn.QueryRow(
		context.Background(),
		`INSERT INTO workers (id_user, uuid, fullname) VALUES ($1, $2, $3) RETURNING id;`,
		record.UserId, record.UUID, record.Fullname,
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

// TODO implement delete function using UUID instead of ID
// AND CHANGE THIS FUNCTION TO DEACTIVATE THE WORKER AND THE USER
func DeactivateWorker(uuid string, conn *pgxpool.Conn) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(
		context.Background(),
		`UPDATE workers SET active = FALSE, name = '', avg_ratings = 0, updated_at = CURRENT_TIMESTAMP
		WHERE uuid = $1::text;`,
		uuid,
	)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	if commandTag.RowsAffected() != 1 {
		tx.Rollback(context.Background())
		return &errors.UnexpectedDBChangeBehaviourError{
			Operation:            "deactivate",
			Table:                "workers",
			ExpectedChangedLines: 1,
			ChangedLines:         int(commandTag.RowsAffected()),
			Identifier:           uuid,
		}
	}
	if err := tx.Commit(context.Background()); err != nil {
		tx.Rollback(context.Background())
		return err
	}
	return nil
}

func DeleteWorker(uuid string, conn *pgxpool.Conn) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(
		context.Background(),
		`DELETE FROM workers WHERE uuid = $1::string;`,
		uuid,
	)
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
			Identifier:           uuid,
		}
	}
	if err := tx.Commit(context.Background()); err != nil {
		tx.Rollback(context.Background())
		return err
	}
	return nil
}

func UpdateWorker(newDataRecord types.WorkersRecord, conn *pgxpool.Conn) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(context.Background(),
		"UPDATE workers SET fullname = $1, active = $2, updated_at = CURRENT_TIMESTAMP WHERE id = $3;",
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

func UpdateWorkerRating(newDataRecord types.WorkersRecord, conn *pgxpool.Conn) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(context.Background(),
		"UPDATE workers SET avg_rating = $1 WHERE id = $2;",
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

func DoesWorkerExists(workerId int32, conn *pgxpool.Conn) (bool, error) {
	ex, err := GetWorkerById(workerId, conn)
	if err != nil {
		return false, err
	}
	return ex != nil && ex.Active, nil
}
