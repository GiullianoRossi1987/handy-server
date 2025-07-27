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

func GetWPhonesHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		uuid := c.Param("uuid")
		addrs, err := services.GetWorkerPhones(pool, uuid)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, addrs)
	}
	return gin.HandlerFunc(fn)
}

func GetCPhonesHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		uuid := c.Param("uuid")
		addrs, err := services.GetWorkerPhones(pool, uuid)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, addrs)
	}
	return gin.HandlerFunc(fn)
}

func AddPhoneHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		content := requests.PhoneBody{}
		if err := c.ShouldBindBodyWithJSON(content); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		id, err := services.AddPhone(pool, content)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, serial.IdReturing{Id: *id})
	}
	return gin.HandlerFunc(fn)
}

func DeletePhoneHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		addr, err := (strconv.Atoi(c.Param("id")))
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		if exists, err := services.GetPhoneById(pool, int32(addr)); err != nil || exists == nil {
			c.String(http.StatusNotFound, "Phone not found")
			return
		}
		if err := services.DeletePhone(pool, int32(addr)); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, "Phone deleted")
	}
	return gin.HandlerFunc(fn)
}

func UpdatePhoneHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		addr, err := (strconv.Atoi(c.Param("id")))

		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		if exists, err := services.GetPhoneById(pool, int32(addr)); err != nil || exists == nil {
			c.String(http.StatusNotFound, "Phone not found")
			return
		}
		content := requests.PhoneBody{}
		if err := c.ShouldBindBodyWithJSON(content); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		if err := services.UpdatePhone(pool, content, int32(addr)); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, "Phone deleted")
	}
	return gin.HandlerFunc(fn)
}

func RoutePhones(router gin.IRouter, pool *pgxpool.Pool) {
	router.GET("/worker/phones/:id/", GetWPhonesHandler(pool))
	router.GET("/customer/phones/:id/", GetCPhonesHandler(pool))
	router.POST("/phone/", AddPhoneHandler(pool))
	router.PUT("/phone/:id", UpdatePhoneHandler(pool))
	router.DELETE("/phone/:id", DeletePhoneHandler(pool))
}
