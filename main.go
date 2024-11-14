package main

import (
	"github.com/Ulpio/gin-api/database"
	"github.com/Ulpio/gin-api/routers"
)

func main() {
	database.ConnectDB()
	routers.DefineRotas()
}
