package routes

import (
	"DigiShop/server/handler"
	"DigiShop/tools"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

type Server struct {
	SC ServerConfig `json:"server"`
}

type ServerConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func App() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/", handler.MainHandler)

	config := loadConfig(tools.ConfigAddress)

	fmt.Println("App is running...")
	infoApp := fmt.Sprintf("Visit App: http://%s:%s\n", config.SC.Host, config.SC.Port)
	fmt.Println(infoApp)

	address := fmt.Sprintf("%s:%s", config.SC.Host, config.SC.Port)
	err := router.Run(address)
	tools.CheckError(err, "failed to start the server!")
}

func loadConfig(filename string) *Server {
	file, err := os.ReadFile(filename)
	tools.CheckError(err, "open database config!")

	var SC Server
	err = json.Unmarshal(file, &SC)
	tools.CheckError(err, "json unmarshal in load database config!")

	return &SC
}
