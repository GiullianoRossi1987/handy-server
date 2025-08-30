package routes

import (
	"net/http"
	services "services/users"
	"strconv"
	requests "types/requests/users"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jackc/pgx/v5/pgxpool"
)

func checkWorkerExistsUUID(pool *pgxpool.Pool, uuid string) bool {
	data, err := services.GetWorkerByUUID(pool, uuid)
	return data != nil && err == nil
}

// put
func AddWorkerHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid ID received")
			return
		}
		content := requests.UpdateUserRequest{}
		if err := c.ShouldBindBodyWith(&content, binding.JSON); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		WorkerId, err := services.AddWorker(pool, content, int32(id))
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, struct{ Id int32 }{
			Id: *WorkerId,
		})
	}
	return gin.HandlerFunc(fn)
}

func UpdateWorkerHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		uuid := c.Param("uuid")
		if !checkWorkerExistsUUID(pool, uuid) {
			c.String(http.StatusNotFound, "UUID not found")
			return
		}
		content := requests.UpdateUserRequest{}
		if err := c.ShouldBindBodyWith(&content, binding.JSON); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		err := services.UpdateWorker(pool, content, uuid)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.Status(http.StatusOK)
	}
	return gin.HandlerFunc(fn)
}

func GetWorkerHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		uuid := c.Param("uuid")
		if !checkWorkerExistsUUID(pool, uuid) {
			c.String(http.StatusNotFound, "UUID not found")
			return
		}
		Worker, err := services.GetWorkerByUUID(pool, uuid)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, Worker)
	}
	return gin.HandlerFunc(fn)
}

func DeleteWorkerHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		uuid := c.Param("uuid")
		if !checkWorkerExistsUUID(pool, uuid) {
			c.String(http.StatusNotFound, "UUID not found")
			return
		}
		if err := services.DeleteWorker(pool, uuid); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.Status(http.StatusOK)
	}
	return gin.HandlerFunc(fn)
}

func RouteWorkers(router gin.IRouter, pool *pgxpool.Pool) {
	router.PUT("/worker/add/:id", AddWorkerHandler(pool))
	router.GET("/worker/:uuid", GetWorkerHandler(pool))
	router.DELETE("/worker/delete/:uuid", DeleteWorkerHandler(pool))
	router.PUT("/worker/update/:uuid", UpdateWorkerHandler(pool))
}

// TODO: DEFINE A PATTERN FOR THE URLs
