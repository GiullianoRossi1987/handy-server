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

func checkCustomerExistsUUID(pool *pgxpool.Pool, uuid string) bool {
	data, err := services.GetCustomerByUUID(pool, uuid)
	return data != nil && err == nil
}

// put
func AddCustomerHandler(pool *pgxpool.Pool) gin.HandlerFunc {
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
		customerId, err := services.AddCustomer(pool, content, int32(id))
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, struct{ Id int32 }{
			Id: *customerId,
		})
	}
	return gin.HandlerFunc(fn)
}

func UpdateCustomerHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		uuid := c.Param("uuid")
		if !checkCustomerExistsUUID(pool, uuid) {
			c.String(http.StatusNotFound, "UUID not found")
			return
		}
		content := requests.UpdateUserRequest{}
		if err := c.ShouldBindBodyWith(&content, binding.JSON); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		customer, err := services.UpdateCustomer(pool, content, uuid)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, customer)
	}
	return gin.HandlerFunc(fn)
}

func GetCustomerHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		uuid := c.Param("uuid")
		if !checkCustomerExistsUUID(pool, uuid) {
			c.String(http.StatusNotFound, "UUID not found")
			return
		}
		customer, err := services.GetCustomerByUUID(pool, uuid)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, customer)
	}
	return gin.HandlerFunc(fn)
}

func DeleteCustomerHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		uuid := c.Param("uuid")
		if !checkCustomerExistsUUID(pool, uuid) {
			c.String(http.StatusNotFound, "UUID not found")
			return
		}
		if err := services.DeleteCustomer(pool, uuid); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.Status(http.StatusOK)
	}
	return gin.HandlerFunc(fn)
}

func RouteCustomers(router gin.IRouter, pool *pgxpool.Pool) {
	router.PUT("/customer/add/:id", AddCustomerHandler(pool))
	router.GET("/customer/:uuid", GetCustomerHandler(pool))
	router.DELETE("/customer/delete/:uuid", DeleteCustomerHandler(pool))
	router.PUT("/customer/update/:uuid", UpdateCustomerHandler(pool))
}
