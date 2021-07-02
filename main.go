package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	id             int     `json:"id"`
	codTransaccion string  `json:"codTransaccion"`
	moneda         string  `json:"moneda"`
	monto          float64 `json:"monto"`
	emisor         int     `json:"emisor"`
	receptor       int     `json:"receptor"`
	fecha          string  `json:"fecha"`
}

func main() {

	router := gin.Default()

	t := &Transaction{}

	router.GET("/ej2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Luciano Panizza",
		})
	})

	router.GET("/transactions", t.GetAll)

	router.Run()

}

func (t *Transaction) GetAll(c *gin.Context) {
	jsonData, err := ioutil.ReadFile("./transactions.json")
	var ts []Transaction

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
