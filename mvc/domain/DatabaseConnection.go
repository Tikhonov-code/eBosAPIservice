package domain

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)
type Config struct {
    DB struct {
        Host     string `json:"host"`
        Port     int    `json:"port"`
        Username string `json:"username"`
        Password string `json:"password"`
        Database string `json:"database"`
    } `json:"db"`
}
var DatabaseConnection string = "mongodb://localhost:27017"

func ReadConnectionString() string {
	file, err := os.Open("C:\\Users\\tix19\\Documents\\eBOStest1GoAPIservice\\config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}
	//"Server=62.109.17.140;Database=EbosDB;User Id=tix;Password=1325Nikita;"
	constring := fmt.Sprintf("Server=%s;Database=%s;User Id=%s;Password=%s;",
		config.DB.Host, config.DB.Database, config.DB.Username, config.DB.Password)
	return constring
}