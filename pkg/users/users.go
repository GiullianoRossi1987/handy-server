package pkg

import (
	"context"
	"fmt"
	types "types/database/users"
	errors "types/errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func AddUser(record types.UsersRecord, conn *pgxpool.Conn) (*int32, error) {
	var id int32
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	if err := conn.QueryRow(
		context.Background(),
		"INSERT INTO users (login, password) VALUES ($1, $2) RETURNING id;",
		record.Login,
		record.Password,
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

func GetUserByLogin(login string, connection *pgxpool.Conn) (*types.UsersRecord, error) {
	rows, err := connection.Query(
		context.Background(),
		`SELECT * FROM users WHERE login = $1`,
		login,
	)
	if err != nil {
		return nil, err
	}
	result, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByPos[types.UsersRecord])
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetUserById(id int, connection *pgxpool.Conn) (*types.UsersRecord, error) {
	rows, err := connection.Query(
		context.Background(),
		`SELECT * FROM users WHERE id = $1`,
		id,
	)
	if err != nil {
		return nil, err
	}
	result, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByPos[types.UsersRecord])
	if err != nil {
		return nil, err
	}
	return result, nil
}

func DeleteUserById(id int, conn *pgxpool.Conn) error {
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

func UpdateUserById(newDataRow types.UsersRecord, connection *pgxpool.Conn) error {
	tx, err := connection.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := connection.Exec(
		context.Background(),
		"UPDATE users SET login = $1, password = $2, updated_at = CURRENT_TIMESTAMP WHERE id = $3;",
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
			Identifier:           fmt.Sprintf("%d", newDataRow.Id),
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

func IsUserLoggable(userId int32, conn *pgxpool.Conn) (bool, error) {
	rows, err := conn.Query(
		context.Background(),
		`SELECT
			u.*,
			w.*
		FROM
			users as u,
			workers as w
		WHERE
			u.active == TRUE AND w.active AND $1 in (u.id_user, w.id_user)
		LIMIT 1;`,
	)
	if err != nil {
		return false, err
	}
	items, err := rows.Values()
	if err != nil {
		return false, err
	}
	return len(items) > 0, nil
}
