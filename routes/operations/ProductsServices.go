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

func GetWorkerPSHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		ps, err := services.GetWorkerCatalog(pool, int32(id))
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, ps)
	}
	return gin.HandlerFunc(fn)
}

func GetPSByIdHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		ps, err := services.GetProductServiceById(pool, int32(id))
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, ps)
	}
	return gin.HandlerFunc(fn)
}

func AddPSHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		content := requests.ProductServiceBody{}
		if err := c.ShouldBindBodyWithJSON(content); err != nil {
			c.String(http.StatusBadRequest, err.Error()) // TODO check the requests status codes in other routes
			return
		}
		id, err := services.AddProductService(pool, content)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, serial.IdReturing{Id: *id})
	}
	return gin.HandlerFunc(fn)
}

func UpdatePSHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		if ex, err := services.GetProductServiceById(pool, int32(id)); err != nil || ex == nil {
			strerr := string(err.Error())
			c.String(http.StatusNotFound, "Couldn't find Product/Service by Id. "+utils.Coalesce(&strerr, "No errors found"))
			return
		}
		content := requests.ProductServiceBody{}
		if err := c.ShouldBindBodyWithJSON(content); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		if err := services.UpdateProductService(pool, content, int32(id)); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.Status(http.StatusOK)
	}
	return gin.HandlerFunc(fn)
}

func DeletePSHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		if ex, err := services.GetProductServiceById(pool, int32(id)); err != nil || ex == nil {
			strerr := string(err.Error())
			c.String(http.StatusNotFound, "Couldn't find Product/Service by Id. "+utils.Coalesce(&strerr, "No errors found"))
			return
		}
		if err := services.DeleteProductService(pool, int32(id)); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.Status(http.StatusOK)
	}
	return gin.HandlerFunc(fn)
}

func RoutePS(router gin.IRouter, pool *pgxpool.Pool) {
	router.GET("/worker/catalog/:id", GetWorkerPSHandler(pool)) // [ ] Maybe change that to use the uuid in the future
	router.GET("/product-service/:id", GetPSByIdHandler(pool))
	router.POST("/product-service", AddPSHandler(pool))
	router.PUT("/product-service/:id", UpdatePSHandler(pool))
	router.DELETE("/product-service/:id", DeletePSHandler(pool))
}
