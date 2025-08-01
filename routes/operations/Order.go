package routes

import (
	"net/http"
	services "services/operations"
	"strconv"
	requests "types/requests/operations"
	serial "types/serializables"
	"utils"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetOrderByIdHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.String(http.StatusBadRequest, "Received invalid id")
			return
		}
		order, err := services.GetOrderById(pool, int32(id))
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, order)
	}
	return gin.HandlerFunc(fn)
}

func GetCustomerOrdersHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.String(http.StatusBadRequest, "Received invalid id")
			return
		}
		order, err := services.GetCustomerOrders(pool, int32(id))
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, order)
	}
	return gin.HandlerFunc(fn)
}

func GetWorkerOrdersHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.String(http.StatusBadRequest, "Received invalid id")
			return
		}
		order, err := services.GetWorkerOrders(pool, int32(id))
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, order)
	}
	return gin.HandlerFunc(fn)
}

func AddOrderHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		content := requests.OrderBody{}
		if err := c.ShouldBindBodyWithJSON(content); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		id, err := services.PlaceOrder(pool, content)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, serial.IdReturing{Id: *id})
	}
	return gin.HandlerFunc(fn)
}

func UpdateOrderHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		if ex, err := services.GetOrderById(pool, int32(id)); err != nil || ex == nil {
			strerr := string(err.Error())
			c.String(http.StatusNotFound, utils.Coalesce(&strerr, "No errors found"))
			return
		}
		content := requests.OrderBody{}
		if err := c.ShouldBindBodyWithJSON(content); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		if err := services.UpdateOrder(pool, content, int32(id)); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.Status(http.StatusOK)
	}
	return gin.HandlerFunc(fn)
}

func DeleteOrderHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		if ex, err := services.GetOrderById(pool, int32(id)); err != nil || ex == nil {
			strerr := string(err.Error())
			c.String(http.StatusNotFound, utils.Coalesce(&strerr, "No errors found"))
			return
		}
		if err := services.DeleteOrder(pool, int32(id)); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.Status(http.StatusOK)
	}
	return gin.HandlerFunc(fn)
}

func RouteOrders(router gin.IRouter, pool *pgxpool.Pool) {
	router.GET("/order/:id", GetOrderByIdHandler(pool))
	router.GET("/customer/orders/:id/", GetCustomerOrdersHandler(pool))
	router.GET("/worker/orders/:id", GetWorkerOrdersHandler(pool))
	router.POST("/order/", AddOrderHandler(pool))
	router.PUT("/order/:id", UpdateOrderHandler(pool))
	router.DELETE("/order/:id", DeleteOrderHandler(pool))
}
