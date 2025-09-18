package main

import (
	"encoding/json"
	"fmt"
	"os"
	"pizzaria/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var pizzas = []models.Pizza{
	{ID: 1, Nome: "Toscana", Preco: 49.5},
	{ID: 2, Nome: "Marguerita", Preco: 49.5},
	{ID: 3, Nome: "Atum com queijo", Preco: 49.5},
}

func main() {
	loadPizzas()
	router := gin.Default()
	router.GET("/pizzas", getPizzas)
	router.POST("/pizzas", postPizzas)
	router.GET("/pizzas/:id", getPizzaByID)
	//deletar uma pizza
	router.DELETE("/pizzas/:id", deletePizzaByID)
	//editar pizza ou atualizar uma pizza
	router.PUT("/pizzas/:id", updatePizzaByID)
	router.Run()
}

func getPizzas(c *gin.Context) {

	c.JSON(200, gin.H{
		"pizzas": pizzas,
	})
}

func postPizzas(c *gin.Context) {
	var newPizza models.Pizza
	if err := c.ShouldBindJSON(&newPizza); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error()})
		return
	}
	newPizza.ID = len(pizzas) + 1
	pizzas = append(pizzas, newPizza)
	savePizza()
	c.JSON(201, newPizza)
}

func getPizzaByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error()})
		return
	}
	for _, p := range pizzas {
		if p.ID == id {
			c.JSON(200, p)
			return
		}
	}
	c.JSON(404, gin.H{"message": "Pizza not found"})
}

func loadPizzas() {
	file, err := os.Open("dados/pizza.json")
	if err != nil {
		fmt.Println("Error file", err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&pizzas); err != nil {
		fmt.Println("Error decoding JSON", err)
	}
}

func savePizza() {
	file, err := os.Create("dados/pizza.json")
	if err != nil {
		fmt.Println("Error file", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(pizzas); err != nil {
		fmt.Println("Error encoding JSON", err)
	}
}

func deletePizzaByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error()})
		return
	}

	for i, p := range pizzas {
		if p.ID == id {
			pizzas = append(pizzas[:i], pizzas[i+1:]...)
			savePizza()
			c.JSON(200, gin.H{"message": "Pizza deleted"})
			return
		}
	}
	c.JSON(404, gin.H{"message": "Pizza not found"})
}

func updatePizzaByID(c *gin.Context) {

	//c.JSON(200, gin.H{"method": "put"})
}
