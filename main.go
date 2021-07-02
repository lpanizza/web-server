package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	Id             int     `json:"id"`
	CodTransaccion string  `json:"codTransaccion"`
	Moneda         string  `json:"moneda"`
	Monto          float64 `json:"monto"`
	Emisor         int     `json:"emisor"`
	Receptor       int     `json:"receptor"`
	Fecha          string  `json:"fecha"`
}

type Handler struct{}

func main() {

	router := gin.Default()

	h := &Handler{}

	router.GET("/ej2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Luciano Panizza",
		})
	})

	router.GET("/transactions", h.GetAll)

	router.Run()

}

func (h *Handler) GetAll(c *gin.Context) {
	jsonData, err := ioutil.ReadFile("./transactions.json")
	ts := []Transaction{}

	if err != nil {
		fmt.Printf("Erro leyendo el json")
	}

	if err := json.Unmarshal(jsonData, &ts); err != nil {
		c.JSON(500, gin.H{
			"error": "No se pudo deserializar el archivo.",
		})
	}

	c.String(200, "%v", ts)

}
