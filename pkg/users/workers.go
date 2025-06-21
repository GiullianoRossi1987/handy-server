package pkg

import (
	"context"
	types "types/database/users"
	"utils"

	"github.com/jackc/pgx/v5"
)

func GetWorkerById(id int32, conn *pgx.Conn) *types.WorkersRecord {
	var worker *types.WorkersRecord
	err := conn.QueryRow(context.Background(),
		"SELECT * FROM workers WHERE id = $1;", id).Scan(&worker)
	if err != nil {
		panic(err)
	}
	return worker
}

func GetWorkerByUserId(id int32, conn *pgx.Conn) *types.WorkersRecord {
	var worker *types.WorkersRecord
	err := conn.QueryRow(context.Background(),
		"SELECT * FROM workers WHERE id_user = $1;", id).Scan(&worker)
	if err != nil {
		panic(err)
	}
	return worker
}

func AddWorker(record types.WorkersRecord, conn *pgx.Conn) {
	commandTag, err := conn.Exec(context.Background(),
		"INSERT INTO workers (id_user, uuid, fullname) VALUES ($1, $2, $3);",
		record.UserId, record.UUID, record.Fullname)
	utils.CheckRowsAndError(commandTag, &err, 1)
}

func DeleteWorker(id int32, conn *pgx.Conn) {
	commandTag, err := conn.Exec(context.Background(), "DELETE FROM workers WHERE id = $1;", id)
	utils.CheckRowsAndError(commandTag, &err, 1)
}

func UpdateWorker(newDataRecord types.WorkersRecord, conn *pgx.Conn) types.WorkersRecord {
	commandTag, err := conn.Exec(context.Background(),
		"UPDATE FROM workers SET fullname = $1, active = $2, updated_at = CURRENT_TIMESTAMP() WHERE id = $3;",
		newDataRecord.Fullname, newDataRecord.Active, newDataRecord.Id)
	utils.CheckRowsAndError(commandTag, &err, 1)
	return newDataRecord
}

func UpdateWorkerRating(newDataRecord types.WorkersRecord, conn *pgx.Conn) types.WorkersRecord {
	commandTag, err := conn.Exec(context.Background(),
		"UPDATE FROM workers SET avg_rating = $1 WHERE id = $2;",
		newDataRecord.Avg_ratings, newDataRecord.Id)
	utils.CheckRowsAndError(commandTag, &err, 1)
	return newDataRecord
}

func DoesWorkerExists(workerId int32, conn *pgx.Conn) bool {
	return GetWorkerById(workerId, conn) != nil
}
