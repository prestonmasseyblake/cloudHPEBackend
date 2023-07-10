package main

import (
	"fmt"
	"net/http"
	"io"
	"os"
	"log"
	"reflect"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// owm "github.com/briandowns/openweathermap"
)

func main() {
	r := gin.New()
	fmt.Println("Hello World!")


	err := godotenv.Load(".env")
	if err != nil {
    	log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("API_KEY")

	// var weatherData map[string]float64
	// weatherData = map[string]float64{
	// 	"temperature": 72,
		
	// }
	// fmt.Println(weatherData)
	url := "https://weatherapi-com.p.rapidapi.com/current.json?q=53.1%2C-0.13"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", apiKey)
	req.Header.Add("X-RapidAPI-Host", "weatherapi-com.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)
	// fmt.Println("this is res ",res)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	bod := string(body)
	fmt.Println("this is bod",bod)
	// valueType := reflect.TypeOf(body)

	// Print the type
	// fmt.Println("Type:", valueType)

	r.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": bod,
		})
	})
	// bod := string(body)
	// temp_f := bod["temp_f"]
	// fmt.Println(res)
	// fmt.Println(bod)
	// url := "https://weatherapi-com.p.rapidapi.com/current.json?q=53.1%2C-0.13"
	// req, _ := http.NewRequest("GET", url, nil)
	// req.Header.Add("X-RapidAPI-Key", "d4e0ad73a6a74e1c929174306231007")
	// req.Header.Add("X-RapidAPI-Host", "weatherapi-com.p.rapidapi.com")

	// res, _ := http.DefaultClient.Do(req)
	// defer res.Body.Close()
	// body, _ := io.ReadAll(res.Body)
	// fmt.Println(res)
	// fmt.Println(string(body))
	
	// r.GET("/weather", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		weatherData
	// 	})
	// })

	r.Run()
}
