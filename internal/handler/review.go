package handler

import (
	"net/http"
	"pizzaria/internal/data"
	"pizzaria/internal/models"
	"pizzaria/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PostReview(c *gin.Context) {
	pizzaId, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var newReview models.Review
	if err := c.ShouldBind(&newReview); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.ValidateReviewRating(newReview); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	for i, p := range data.Pizzas {
		if p.ID == pizzaId {
			p.Review = append(p.Review, newReview)
			data.Pizzas[i] = p
			data.SavePizzas()
			c.JSON(http.StatusCreated, p)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"Message": "Pizza not found"})
}
