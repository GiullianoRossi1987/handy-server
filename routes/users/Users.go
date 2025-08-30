package routes

import (
	"net/http"
	services "services/users"
	"strconv"
	requests "types/requests/users"
	serial "types/serializables"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jackc/pgx/v5/pgxpool"
)

func checkUserExists(pool *pgxpool.Pool, user int) bool {
	data, err := services.GetUserById(pool, user)
	return data != nil && err == nil
}

func AddUserHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		content := requests.CreateUserRequest{}
		if err := c.ShouldBindBodyWith(&content, binding.JSON); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
		id, err := services.AddUser(pool, content)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
		c.JSON(http.StatusOK, serial.IdReturing{
			Id: *id,
		})
	}
	return gin.HandlerFunc(fn)
}

func GetUserByLogin(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		login := c.Param("login")
		data, err := services.GetUserByLogin(pool, login)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
		if data == nil {
			c.String(http.StatusNotFound, `user not found`)
			return
		}
		c.JSON(http.StatusOK, data)
	}
	return gin.HandlerFunc(fn)
}

func LoginWithUser(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		content := requests.LoginRequestBody{}
		if err := c.ShouldBindBodyWith(&content, binding.JSON); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
		response, err := services.Login(pool, content)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
		c.JSON(http.StatusOK, response)
	}
	return gin.HandlerFunc(fn)
}

func UpdateUser(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		content := requests.CreateUserRequest{}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.String(http.StatusInternalServerError, `Invalid user Id received`)
		}
		if !checkUserExists(pool, id) {
			c.String(http.StatusNotFound, "Couldn't find the user")
			return
		}
		if err := c.ShouldBindBodyWith(&content, binding.JSON); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		response, err := services.UpdateUser(pool, content, id)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, response)
	}
	return gin.HandlerFunc(fn)
}

func DeleteUser(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.String(http.StatusInternalServerError, `Invalid user Id received`)
		}
		if err := services.DeleteUser(pool, id); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
		c.Status(http.StatusOK)
	}
	return gin.HandlerFunc(fn)
}

func RouteUsers(router gin.IRouter, pool *pgxpool.Pool) {
	router.POST("/user/add", AddUserHandler(pool))
	router.GET("/user/get-login/:login", GetUserByLogin(pool))
	router.POST("/user/login", LoginWithUser(pool))
	router.PUT("/user/:id", UpdateUser(pool))
	router.DELETE(("/user/:id"), DeleteUser(pool))
}
