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

func GetWEmailsHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		uuid := c.Param("uuid")
		addrs, err := services.GetWorkerEmails(pool, uuid)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, addrs)
	}
	return gin.HandlerFunc(fn)
}

func GetCEmailsHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		uuid := c.Param("uuid")
		addrs, err := services.GetWorkerEmails(pool, uuid)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, addrs)
	}
	return gin.HandlerFunc(fn)
}

func AddEmailHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		content := requests.EmailBody{}
		if err := c.ShouldBindBodyWithJSON(content); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		id, err := services.AddEmail(pool, content)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, serial.IdReturing{Id: *id})
	}
	return gin.HandlerFunc(fn)
}

func DeleteEmailHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		addr, err := (strconv.Atoi(c.Param("id")))
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		if exists, err := services.GetEmailById(pool, int32(addr)); err != nil || exists == nil {
			c.String(http.StatusNotFound, "Email not found")
			return
		}
		if err := services.DeleteEmail(pool, int32(addr)); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, "Email deleted")
	}
	return gin.HandlerFunc(fn)
}

func UpdateEmailHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		addr, err := (strconv.Atoi(c.Param("id")))

		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		if exists, err := services.GetEmailById(pool, int32(addr)); err != nil || exists == nil {
			c.String(http.StatusNotFound, "Email not found")
			return
		}
		content := requests.EmailBody{}
		if err := c.ShouldBindBodyWithJSON(content); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		if err := services.UpdateEmail(pool, content, int32(addr)); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, "Email deleted")
	}
	return gin.HandlerFunc(fn)
}

func RouteEmails(router gin.IRouter, pool *pgxpool.Pool) {
	router.GET("/worker/emails/:id", GetWEmailsHandler(pool))
	router.GET("/customer/emails/:id", GetCEmailsHandler(pool))
	router.POST("/emails/", AddEmailHandler(pool))
	router.PUT("/emails/:id", UpdateEmailHandler(pool))
	router.DELETE("/emails/:id", DeleteEmailHandler(pool))
}
