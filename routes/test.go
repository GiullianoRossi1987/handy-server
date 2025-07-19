package handlers

import (
	// "fmt"
	"net/http"

	// serializables "types/serializables"
	requests "types/requests/users"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func TestResponse(c *gin.Context) {
	content := requests.CreateUserRequest{}
	err := c.ShouldBindBodyWith(&content, binding.JSON)
	if err != nil {
		// response := fmt.Sprintf("Works %s", content.Login)
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, content)
}
