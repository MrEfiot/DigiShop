package config

import (
	"DigiShop/tools"
	"encoding/json"
	"os"
)

type Server struct {
	SC ServerConfig `json:"server"`
}

type ServerConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type Access struct {
	AL AccessLevel `json:"access_level"`
}

type AccessLevel struct {
	LevelOne   string `json:"level_one"`
	LevelTwo   string `json:"level_two"`
	LevelThree string `json:"level_three"`
	LevelFour  string `json:"level_four"`
}

func LoadServerConfig(filename string) *Server {
	file, err := os.ReadFile(filename)
	tools.CheckError(err, "failed to open database config")

	var SC Server
	err = json.Unmarshal(file, &SC)
	tools.CheckError(err, "failed to json unmarshal in load server config")

	return &SC
}

func LoadAccessLevelConfig(filename string) *Access {
	file, err := os.ReadFile(filename)
	tools.CheckError(err, "failed to open database config")

	var AL Access
	err = json.Unmarshal(file, &AL)
	tools.CheckError(err, "failed to json unmarshal in load access_level config")

	return &AL
}
