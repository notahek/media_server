package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Data struct {
	Port         string `json:"port"`
	FileLocation string `json:"file"`
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()

	jsonFile, err := os.ReadFile("server.json")
	if err != nil {
		log.Panic("error in reading json file")
	}
	var data Data
	err = json.Unmarshal(jsonFile, &data)
	if err != nil {
		log.Panic("error while parsing json file")
	}
	fmt.Println("Server is started on port: ", data.Port)
	fmt.Println("server is hosting: ", data.FileLocation)
	app.Static("/static", data.FileLocation)

	app.Run(data.Port)
}
