package controllers

import (
	"net/http"
	"project/config"
	"project/models"
	"strconv"

	"github.com/labstack/echo"
)

// create admin
func CreateDendaController(c echo.Context) error {

	denda := models.Denda{}
	c.Bind(&denda)

	err := config.DB.Save(&denda).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Fail create JWT TOken",
			"status":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success menambahkan data denda",
		"denda":   denda,
	})
}

// get all denda

func GetDendaController(c echo.Context) error {
	var denda []models.Denda

	err := config.DB.Find(&denda).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	// get Id from JWT
	// myUserId := middleware.ExtractTokenUserId(c)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success menampilkan data semua denda",
		"denda":   denda,
	})
}

// get data buku by id
func GetDendaIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid denda id")
	}

	var denda models.Denda
	if err := config.DB.First(&denda, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success menampilkan data denda by id",
		"denda":   denda,
	})
}

// update data buku by id
func UpdateDendaIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid admin id")
	}

	var denda models.Denda

	if err := config.DB.First(&denda, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Bind(&denda)

	if err := config.DB.Save(&denda).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update denda by id",
		"denda":   denda,
	})
}

// delete denda by id
func DeleteDendaIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user id")
	}

	var denda models.Denda
	if err := config.DB.First(&denda, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Delete(&denda).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success menghapus data denda",
	})
}
