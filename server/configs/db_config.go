package configs

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type DBConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
}

var Configs DBConfig

func init() {
	conf_file, err := os.Open("server/db_properties.json")
	if err != nil {
		log.Fatalf("Error. Cant open config file: %v", err)
	}
	byte_value, _ := io.ReadAll(conf_file)
	json.Unmarshal(byte_value, &Configs)
	log.Println("Successfully parsed db config")
}
