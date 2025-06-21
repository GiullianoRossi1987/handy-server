package pkg

import (
	"context"
	types "types/database/users"
	utils "utils"

	"github.com/jackc/pgx/v5"
)

func AddUser(record types.UsersRecord, connection *pgx.Conn) {
	commandTag, err := connection.Exec(context.Background(),
		"INSERT INTO", record.Login, record.Password)
	utils.CheckRowsAndError(commandTag, &err, 1)
}

func GetUserByLogin(login string, connection *pgx.Conn) *types.UsersRecord {
	var result *types.UsersRecord
	err := connection.QueryRow(context.Background(), "SELECT * FROM users WHERE login = $1", login).Scan(&result)
	if err != nil {
		panic(err)
	}
	return result
}

func DeleteUserById(id int, connection *pgx.Conn) {
	commandTag, err := connection.Exec(context.Background(), "DELETE FROM user WHERE id = $1", id)
	utils.CheckRowsAndError(commandTag, &err, 1)
}

func UpdateUserById(newDataRow types.UsersRecord, connection *pgx.Conn) types.UsersRecord {
	commandTag, err := connection.Exec(context.Background(),
		"UPDATE user SET login = $1, password = $2, updated_at = CURRENT_TIMESTAMP() WHERE id = $3",
		newDataRow.Login, newDataRow.Password, newDataRow.Id)
	utils.CheckRowsAndError(commandTag, &err, 1)
	return *GetUserByLogin(newDataRow.Login, connection)
}
