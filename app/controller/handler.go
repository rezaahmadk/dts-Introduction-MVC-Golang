package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rezaahmadk/dts-Introduction-MVC-Golang/app/model"
)

// AddAntrianHandler untuk menambah data
func AddAntrianHandler(c *gin.Context) {
	flag, err := model.AddAntrian()
	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success",
		})
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error":  err,
		})
	}
}

// GetAntrianHandler untuk menampilkan data
func GetAntrianHandler(c *gin.Context) {
	flag, data, err := model.GetAntrian()
	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success",
			"data":   data,
		})
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error":  err,
		})
	}
}

// UpdateAntrianHandler untuk mengubah data
func UpdateAntrianHandler(c *gin.Context) {
	idAntrian := c.Param("idAntrian")
	flag, err := model.UpdateAntrian(idAntrian)
	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success",
		})
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error":  err,
		})
	}
}

// DeleteAntrianHandler untuk menghapus data
func DeleteAntrianHandler(c *gin.Context) {
	idAntrian := c.Param("idAntrian")
	flag, err := model.DeleteAntrian(idAntrian)
	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success",
		})
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error":  err,
		})
	}
}

// PageAntrianHandler untuk view
func PageAntrianHandler(c *gin.Context) {
	flag, result, err := model.GetAntrian()
	var currentAntrian map[string]interface{}

	for _, item := range result {
		if item != nil {
			currentAntrian = item
			break
		}
	}

	if flag && len(result) > 0 {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"antrian": currentAntrian["id"],
		})
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error":  err,
		})
	}
}
