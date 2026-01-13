package main

import (
	"pizzaria/internal/data"
	"pizzaria/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	data.LoadPizzas()
	router := gin.Default()
	router.GET("/pizzas", handler.GetPizzas)
	router.POST("/pizzas", handler.PostPizzas)
	router.GET("/pizzas/:id", handler.GetPizzasbyId)
	router.DELETE("/pizzas/:id", data.DeletePizzaById)
	router.PUT("/pizzas/:id", handler.UpdatePizzaById)
	router.POST("/pizzas/:id/reviews", handler.PostReview)
	err := router.Run()
	if err != nil {
		return
	} // listens on 0.0.0.0:8080 by default
}
