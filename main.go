package main

import(
	"e/config"
	
	"e/database"
	"log"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"github.com/go-redis/redis/v8"
)

func main() {

	//connecting the database mysql and redis
	loadConfig, err := config.LoadConfig(".")
	if err!=nil{
		log.Fatal("cannot load environment variables", err)

	}
	//mysql
	db:= database.ConnectionMySQLDb(&loadConfig)
	rdb:= database.ConnectionRedisDb(&loadConfig)

	 startServer(db, rdb)
}

func startServer(db *gorm.DB, rdb *redis.Client){
	app := fiber.New()

	err :=app.Listen(":3400")
	if err !=nil{
		panic(err)
	}
}