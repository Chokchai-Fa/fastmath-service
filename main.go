package main

import (
	"database/sql"
	"fastmath/routers"
	"fastmath/utils/dbpostgres"
	"fastmath/utils/logs"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load()
	if err != nil {
		logs.Error("Error loading .env file")
	}

	logs.Info("ENV: " + os.Getenv("ENV"))
	logs.Info("PORT: " + os.Getenv("PORT"))
}

func connectDBPostgres() (*dbpostgres.DBPG, *sql.DB) {

	port, err := strconv.Atoi(os.Getenv("PG_PORT"))
	if err != nil {
		logs.Error("get pg port error|:" + err.Error())
		os.Exit(0)
	}

	//CONNECT DB
	dbPG := dbpostgres.NewDBPostgres()
	dbPGCli, err := dbPG.Connect(os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_HOST"), os.Getenv("PG_DBNAME"), port)

	// set the maximum life time of a connection
	dbPGCli.SetConnMaxLifetime(5 * time.Minute)

	// set the maximum number of connections in the pool
	dbPGCli.SetMaxOpenConns(50)

	// set the maximum idle connections
	dbPGCli.SetMaxIdleConns(25)

	if err != nil {
		logs.Error("Postgres.Connect|:" + err.Error())
		os.Exit(0)
	}

	err = dbPGCli.Ping()
	if err != nil {
		logs.Error("Postgres.Ping|:" + err.Error())
		os.Exit(0)
	}

	return dbPG, dbPGCli

}

func main() {
	dbPG, dbPGCli := connectDBPostgres()
	router := gin.New()
	router.Use(gin.Recovery())
	routers.SetupRouter(router, dbPG, dbPGCli)
	port := os.Getenv("PORT")
	logs.Error(router.Run(":" + port))
}
