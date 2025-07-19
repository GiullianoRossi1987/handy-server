package pkg

import (
	"context"
	"fmt"

	types "types/database/operations"
	errors "types/errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetProductById(prodserId int32, conn *pgxpool.Conn) (*types.ProductService, error) {
	var result *types.ProductService
	if err := conn.QueryRow(context.Background(), "SELECT * FROM products_services WHERE id = $1;", prodserId).Scan(result); err != nil {
		return nil, err
	}
	return result, nil
}

func GetWorkerProdSer(workerId int32, conn *pgxpool.Conn) ([]types.ProductService, error) {
	rows, err := conn.Query(context.Background(), "SELECT * FROM products_services WHERE id_worker = $1;", workerId)
	if err != nil {
		return nil, err
	}
	results, err := pgx.CollectRows(rows, pgx.RowToStructByPos[types.ProductService])
	if err != nil {
		return nil, err
	}
	return results, nil
}

func AddProdSer(prodser types.ProductService, conn *pgxpool.Conn) (*int32, error) {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	var id int32
	if err := conn.QueryRow(
		context.Background(),
		`INSERT INTO products_services (
			id_worker,
			name, 
			description, 
			available, 
			quantity_available, 
			service
		) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`,
		prodser.IdWorker,
		prodser.Name,
		prodser.Description,
		prodser.Available,
		prodser.QuantityAvailable,
		prodser.Service,
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

func DeleteProdSer(prodserId int32, conn *pgxpool.Conn) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(
		context.Background(),
		`DELETE FROM products_services WHERE id = $1;`,
		prodserId,
	)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	if commandTag.RowsAffected() != 1 {
		tx.Rollback(context.Background())
		return &errors.UnexpectedDBChangeBehaviourError{
			Operation:            "insert",
			Table:                "products_services",
			ExpectedChangedLines: 1,
			ChangedLines:         int(commandTag.RowsAffected()),
			Identifier:           fmt.Sprintf("%d", prodserId),
		}
	}
	if err := tx.Commit(context.Background()); err != nil {
		tx.Rollback(context.Background())
		return err
	}
	return nil
}

func UpdateProdSer(prodser types.ProductService, conn *pgxpool.Conn) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(
		context.Background(),
		`UPDATE products_services SET 
		 name = $2, 
		 description = $3, 
		 available = $4, 
		 quantity_available = $5,
		 service = $6,
		 updated_at = CURRENT_TIMESTAMP,
		 WHERE id = $1`, // [ ] Probably could change it to a database trigger
		prodser.Id,
		prodser.Name,
		prodser.Description,
		prodser.Available,
		prodser.QuantityAvailable,
		prodser.Service,
	)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	if commandTag.RowsAffected() != 1 {
		tx.Rollback(context.Background())
		return &errors.UnexpectedDBChangeBehaviourError{
			Operation:            "insert",
			Table:                "products_services",
			ExpectedChangedLines: 1,
			ChangedLines:         int(commandTag.RowsAffected()),
			Identifier:           fmt.Sprintf("%d", prodser.Id),
		}
	}
	if err := tx.Commit(context.Background()); err != nil {
		tx.Rollback(context.Background())
		return err
	}
	return nil
}
