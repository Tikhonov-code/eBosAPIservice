package utils

//import (
//	"encoding/json"
//	"log"
//	"os"
//)

//func ReadConfig() (domain.Config, error) {
//	// read config file json
//	//config, err := os.ReadFile("C:/Users/tix19/go/src/github.com/tikhonovcode/mfxpings-service/mvc/config.json")
//	config, err := os.ReadFile("config.json")
//	var configObj domain.Config
//	if err != nil {
//		log.Println("Reading config file failed", err.Error())
//		return configObj, err	
//	}
//	
//	if err := json.Unmarshal(config, &configObj); err != nil {
//		return configObj, err
//	}
//	return configObj,nil
//}