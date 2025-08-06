package services

import (
	op "pkg/operations"
	types "types/database/operations"
	requests "types/requests/operations"
	responses "types/responses/operations"
	"utils"

	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetProductServiceById(pool *pgxpool.Pool, id int32) (*responses.ProductServiceResponse, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	data, err := op.GetProductById(id, conn)
	conn.Release()
	if err != nil {
		return nil, err
	}
	return responses.SerializeProductService(data), nil
}

func GetWorkerCatalog(pool *pgxpool.Pool, workerId int32) ([]responses.ProductServiceResponse, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	records, err := op.GetWorkerProdSer(workerId, conn)
	conn.Release()
	if err != nil {
		return nil, err
	}
	data := utils.MapCar(
		records,
		func(record types.ProductService) responses.ProductServiceResponse {
			return *responses.SerializeProductService(&record)
		},
	)
	return data, nil
}

func SearchProdSer(pool *pgxpool.Pool, search string) ([]responses.ProductServiceResponse, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	records, err := op.SearchService(search, conn)
	conn.Release()
	if err != nil {
		return nil, err
	}
	data := utils.MapCar(
		records,
		func(record types.ProductService) responses.ProductServiceResponse {
			return *responses.SerializeProductService(&record)
		},
	)
	return data, nil
}

func AddProductService(pool *pgxpool.Pool, req requests.ProductServiceBody) (*int32, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	record := req.ToRecord()
	id, err := op.AddProdSer(*record, conn)
	conn.Release()
	if err != nil {
		return nil, err
	}
	return id, nil
}

func UpdateProductService(pool *pgxpool.Pool, req requests.ProductServiceBody, prodId int32) error {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	record := req.ToRecord()
	if record == nil {
		return nil
	}
	record.Id = prodId
	if err := op.UpdateProdSer(*record, conn); err != nil {
		return err
	}
	conn.Release()
	return nil
}

func DeleteProductService(pool *pgxpool.Pool, prodId int32) error {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	if err := op.DeleteProdSer(prodId, conn); err != nil {
		return err
	}
	conn.Release()
	return nil
}
