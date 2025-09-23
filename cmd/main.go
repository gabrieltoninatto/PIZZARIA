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
	router.GET("/pizzas/:id", handler.GetPizzaByID)
	//deletar uma pizza
	router.DELETE("/pizzas/:id", handler.DeletePizzaByID)
	//editar pizza ou atualizar uma pizza
	router.PUT("/pizzas/:id", handler.UpdatePizzaByID)
	router.POST("pizzas/:id/reviews", handler.PostReview)
	router.Run()
}
