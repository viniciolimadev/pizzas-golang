package handler

import (
	"net/http"
	"pizzaria/internal/data"
	"pizzaria/internal/models"
	"pizzaria/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPizzas(c *gin.Context) {
	c.JSON(http.StatusOK, data.Pizzas)
}

func PostPizzas(c *gin.Context) {
	var newPizza models.Pizza
	if err := c.ShouldBindJSON(&newPizza); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error()})
		return
	}

	if err := service.ValidatePizzaPrice(&newPizza); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	newPizza.ID = len(data.Pizzas) + 1
	data.Pizzas = append(data.Pizzas, newPizza)
	data.SavePizzas()
	c.JSON(http.StatusCreated, newPizza)
}

func GetPizzasbyId(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}
	for _, p := range data.Pizzas {
		if p.ID == id {
			c.JSON(http.StatusOK, p)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": `Pizza not found`})
}

func UpdatePizzaById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}

	var updatedPizza models.Pizza
	if err := c.ShouldBindJSON(&updatedPizza); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}

	if err := service.ValidatePizzaPrice(&updatedPizza); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	for i, p := range data.Pizzas {
		if p.ID == id {
			data.Pizzas[i] = updatedPizza
			data.Pizzas[i].ID = id
			data.SavePizzas()
			c.JSON(http.StatusOK, data.Pizzas[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": `Pizza not found`})
}
