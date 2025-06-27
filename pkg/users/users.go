package pkg

import (
	"context"
	types "types/database/users"
	errors "types/errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func AddUser(record types.UsersRecord, conn *pgxpool.Pool) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(
		context.Background(),
		"INSERT INTO users (login, password) VALUES ($1, $2)",
		record.Login,
		record.Password,
	)
	if commandTag.RowsAffected() != 1 {
		tx.Rollback(context.Background())
		return &errors.UnexpectedDBChangeBehaviourError{
			Operation:            "insert",
			Table:                "users",
			ExpectedChangedLines: 1,
			ChangedLines:         int(commandTag.RowsAffected()),
		}
	}
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	if err := tx.Commit(context.Background()); err != nil {
		tx.Rollback(context.Background())
		return err
	}
	return nil
}

func GetUserByLogin(login string, connection *pgxpool.Pool) (*types.UsersRecord, error) {
	var result *types.UsersRecord
	if err := connection.QueryRow(
		context.Background(),
		"SELECT * FROM users WHERE login = $1",
		login,
	).Scan(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func DeleteUserById(id int, conn *pgxpool.Pool) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(context.Background(), "DELETE FROM user WHERE id = $1", id)
	if commandTag.RowsAffected() != 1 {
		tx.Rollback(context.Background())
		return &errors.UnexpectedDBChangeBehaviourError{
			Operation:            "delete",
			Table:                "users",
			ExpectedChangedLines: 1,
			ChangedLines:         int(commandTag.RowsAffected()),
		}
	}
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	if err := tx.Commit(context.Background()); err != nil {
		tx.Rollback(context.Background())
		return err
	}
	return nil
}

func UpdateUserById(newDataRow types.UsersRecord, connection *pgxpool.Pool) error {
	tx, err := connection.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := connection.Exec(
		context.Background(),
		"UPDATE user SET login = $1, password = $2, updated_at = CURRENT_TIMESTAMP() WHERE id = $3",
		newDataRow.Login,
		newDataRow.Password,
		newDataRow.Id,
	)
	if commandTag.RowsAffected() != 1 {
		tx.Rollback(context.Background())
		return &errors.UnexpectedDBChangeBehaviourError{
			Operation:            "update",
			Table:                "users",
			ExpectedChangedLines: 1,
			ChangedLines:         int(commandTag.RowsAffected()),
		}
	}
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	if err := tx.Commit(context.Background()); err != nil {
		tx.Rollback(context.Background())
		return err
	}
	return nil
}
