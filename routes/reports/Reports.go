package routes

import (
	"net/http"
	services "services/reports"
	"strconv"
	serial "types/serializables"
	"utils"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetWReports(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		w_id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.String(http.StatusInternalServerError, "Invalid ID received")
			return
		}
		reports, err := services.GetWorkerReports(pool, int32(w_id))
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, reports)
	}
	return gin.HandlerFunc(fn)
}

func GetCReports(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		w_id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.String(http.StatusInternalServerError, "Invalid ID received")
			return
		}
		reports, err := services.GetCustomerReports(pool, int32(w_id))
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, reports)
	}
	return gin.HandlerFunc(fn)
}

func AddWReportHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		content := serial.ReportBody{}
		if err := c.ShouldBindBodyWithJSON(content); err != nil {
			c.String(http.StatusInternalServerError, "Invalid ID received")
			return
		}
		id, err := services.AddWorkerReport(pool, content)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, serial.IdReturing{Id: *id})
	}
	return gin.HandlerFunc(fn)
}

func AddCReportHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		content := serial.ReportBody{}
		if err := c.ShouldBindBodyWithJSON(content); err != nil {
			c.String(http.StatusInternalServerError, "Invalid ID received")
			return
		}
		id, err := services.AddCustomerReport(pool, content)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, serial.IdReturing{Id: *id})
	}
	return gin.HandlerFunc(fn)
}

func DeleteWReportHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.String(http.StatusInternalServerError, "Invalid ID received")
			return
		}
		if ex, err := services.GetWorkerReportById(pool, int32(id)); err != nil || ex == nil {
			strerr := string(err.Error())
			c.String(http.StatusNotFound, "Worker Report Id not found. "+utils.Coalesce(&strerr, "No error returned"))
			return
		}
		if err := services.DeleteWorkerReport(pool, int32(id)); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.Status(http.StatusOK)
	}
	return gin.HandlerFunc(fn)
}

func DeleteCReportHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.String(http.StatusInternalServerError, "Invalid ID received")
			return
		}
		if ex, err := services.GetCustomerReportById(pool, int32(id)); err != nil || ex == nil {
			strerr := string(err.Error())
			c.String(http.StatusNotFound, "Customer Report Id not found. "+utils.Coalesce(&strerr, "No error returned"))
			return
		}
		if err := services.DeleteCustomerReport(pool, int32(id)); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.Status(http.StatusOK)
	}
	return gin.HandlerFunc(fn)
}

func UpdateCReport(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.String(http.StatusInternalServerError, "Invalid ID received")
			return
		}
		if ex, err := services.GetCustomerReportById(pool, int32(id)); err != nil || ex == nil {
			strerr := string(err.Error())
			c.String(http.StatusNotFound, "Customer Report Id not found. "+utils.Coalesce(&strerr, "No error returned"))
			return
		}
		content := serial.ReportBody{}
		if err := c.ShouldBindBodyWithJSON(content); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		if err := services.UpdateCustomerReport(pool, content, int32(id)); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.Status(http.StatusOK)
	}
	return gin.HandlerFunc(fn)
}

func UpdateWReport(pool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.String(http.StatusInternalServerError, "Invalid ID received")
			return
		}
		if ex, err := services.GetWorkerReportById(pool, int32(id)); err != nil || ex == nil {
			strerr := string(err.Error())
			c.String(http.StatusNotFound, "Customer Report Id not found. "+utils.Coalesce(&strerr, "No error returned"))
			return
		}
		content := serial.ReportBody{}
		if err := c.ShouldBindBodyWithJSON(content); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		if err := services.UpdateWorkerReport(pool, content, int32(id)); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.Status(http.StatusOK)
	}
	return gin.HandlerFunc(fn)
}

func RouteReports(router gin.IRouter, pool *pgxpool.Pool) {

}
