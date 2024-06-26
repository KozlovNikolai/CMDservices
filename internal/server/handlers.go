package server

import (
	"context"

	"net/http"
	"strconv"

	"github.com/KozlovNikolai/CMDservices/internal/model"
	"github.com/KozlovNikolai/CMDservices/internal/store"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	var service model.Service
	id, _ := strconv.Atoi(c.Param("id"))

	query := `
	SELECT services.id
	FROM services
	WHERE services.id=$1 LIMIT 100`
	row := store.DB.QueryRow(context.Background(), query, id)

	err := row.Scan(&service.ID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		return
	}

	query = `DELETE FROM services WHERE id=$1`
	_, err = store.DB.Exec(context.Background(), query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Deleted successfully"})
}

func Create(c *gin.Context) {
	var service model.Service
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	query := `
		INSERT INTO services (name,price)
		VALUES ($1,$2)
		RETURNING id`
	err := store.DB.QueryRow(context.Background(), query, service.Name, service.Price).Scan(&service.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, service)
}

func Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var service model.Service
	query := `
		SELECT id,name,price
		FROM services
		WHERE services.id=$1 LIMIT 100`
	row := store.DB.QueryRow(context.Background(), query, id)
	err := row.Scan(&service.ID, &service.Name, &service.Price)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		return
	}
	c.JSON(http.StatusOK, service)
}

func GetList(c *gin.Context) {
	var services []model.Service
	query := `
		SELECT id,name,price
		FROM services
		LIMIT 100`
	rows, err := store.DB.Query(context.Background(), query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()
	for rows.Next() {
		var service model.Service
		err := rows.Scan(&service.ID, &service.Name, &service.Price)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		services = append(services, service)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, services)
}
