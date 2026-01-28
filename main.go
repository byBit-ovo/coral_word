package main

import (
	_ "bytes"
	_ "context"
	_ "database/sql"
	_ "encoding/json"
	"log"
	_ "strconv"

	"github.com/byBit-ovo/coral_word/llm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/google/uuid"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Loading env file error")
	}
	if err = llm.InitModels(); err != nil {
		log.Fatal("InitModels error")
		return
	}
	if err = InitSQL(); err != nil {
		log.Fatal("Init SQL error")
	}
	
	if err = InitRedis(); err != nil {
		log.Fatal("Init Redis error")
	}
	if err = InitEs(); err != nil {
		log.Fatal("Init Es error")
	}

}


func main() {
	// scaleUpWords(100)
	// syncMissingFromLogs()
	// checkSyncLog()
}