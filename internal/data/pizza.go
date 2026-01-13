package data

import (
	"encoding/json"
	"fmt"
	"os"
	"pizzaria/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Pizzas []models.Pizza

func LoadPizzas() {
	file, err := os.Open("dados/pizza.json")
	if err != nil {
		fmt.Println("Error file", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&Pizzas); err != nil {
		fmt.Println("Error file", err)
		return
	}
}
func SavePizzas() {
	file, err := os.Create("dados/pizza.json")
	if err != nil {
		fmt.Println("Error file", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(Pizzas); err != nil {
		fmt.Println("Error file", err)
		return
	}

}

func DeletePizzaById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error()})
		return
	}

	for i, p := range Pizzas {
		if p.ID == id {
			Pizzas = append(Pizzas[:i], Pizzas[i+1:]...)
			SavePizzas()
			c.JSON(200, gin.H{
				"message": "Pizza deleted"})
			return
		}
	}
	c.JSON(404, gin.H{"message": `Pizza not found`})
}
