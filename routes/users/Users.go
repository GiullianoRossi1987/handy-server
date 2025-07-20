package routes

import (
	services "services/users"
	// responses "types/responses/users"
	"net/http"
	requests "types/requests/users"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IdReturing struct {
	Id int32
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
		c.JSON(http.StatusOK, IdReturing{
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
