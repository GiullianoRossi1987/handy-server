package routes

import (
	"net/http"
	services "services/satellites"
	"strconv"
	requests "types/requests/satellites"
	serial "types/serializables"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetWAddressesHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		uuid := c.Param("uuid")
		addrs, err := services.GetWorkerAddresses(pool, uuid)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, addrs)
	}
	return gin.HandlerFunc(fn)
}

func GetCAddressesHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		uuid := c.Param("uuid")
		addrs, err := services.GetWorkerAddresses(pool, uuid)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, addrs)
	}
	return gin.HandlerFunc(fn)
}

func AddAddressHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		content := requests.AddressBody{}
		if err := c.ShouldBindBodyWithJSON(content); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		id, err := services.AddAddress(pool, content)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, serial.IdReturing{Id: *id})
	}
	return gin.HandlerFunc(fn)
}

func DeleteAddressHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		addr, err := (strconv.Atoi(c.Param("id")))
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		if exists, err := services.GetAddressById(pool, int32(addr)); err != nil || exists == nil {
			c.String(http.StatusNotFound, "Address not found")
			return
		}
		if err := services.DeleteAddress(pool, int32(addr)); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, "Address deleted")
	}
	return gin.HandlerFunc(fn)
}

func UpdateAddressHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		addr, err := (strconv.Atoi(c.Param("id")))

		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		if exists, err := services.GetAddressById(pool, int32(addr)); err != nil || exists == nil {
			c.String(http.StatusNotFound, "Address not found")
			return
		}
		content := requests.AddressBody{}
		if err := c.ShouldBindBodyWithJSON(content); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		if err := services.UpdateAddress(pool, content, int32(addr)); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, "Address deleted")
	}
	return gin.HandlerFunc(fn)
}

func RouteAddresses(router gin.IRouter, pool *pgxpool.Pool) {
	router.GET("/worker/addresses/:uuid", GetWAddressesHandler(pool))
	router.GET("/customer/addresses/:uuid", GetCAddressesHandler(pool))
	router.POST("/address/", AddAddressHandler(pool))
	router.PUT("/address/:id", UpdateAddressHandler(pool))
	router.DELETE("/address/:id", DeleteAddressHandler(pool))
}
