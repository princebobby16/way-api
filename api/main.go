package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"net/http"
	"os"
	"way/api/handler"
	"way/api/multiplexer"
	"way/pkg/db"
	"way/pkg/logger"
)

func main() {
	// get port
	port, valid := os.LookupEnv("PORT")
	if !valid {
		logger.Log("Invalid port")
		return
	}

	var databaseInfo db.DatabaseInfo

	// get database url
	databaseInfo.Host, valid = os.LookupEnv("DATABASE_URL")
	if !valid {
		logger.Log("invalid database_url")
		return
	}
	// get database port
	databaseInfo.Port, valid = os.LookupEnv("DATABASE_PORT")
	if !valid {
		logger.Log("invalid database_port")
		return
	}
	// get database user
	databaseInfo.User, valid = os.LookupEnv("DATABASE_USER")
	if !valid {
		logger.Log("invalid database_user")
		return
	}
	// get database password
	databaseInfo.Password, valid = os.LookupEnv("DATABASE_PASSWORD")
	if !valid {
		logger.Log("invalid database_password")
		return
	}
	// get database name
	databaseInfo.DatabaseName, valid = os.LookupEnv("DATABASE_NAME")
	if !valid {
		logger.Log("invalid database_name")
		return
	}

	// connect to database
	err := databaseInfo.Connect()
	if err != nil {
		logger.Log(err)
		return
	}
	logger.Log("Database connection established")

	// Close database connection when program exits
	defer func() {
		err = databaseInfo.Disconnect()
		if err != nil {
			logger.Log(err)
			return
		}
	}()

	router := multiplexer.Router()
	router.Use(handler.JSONMiddleware)

	// cors
	origins := handlers.AllowedOrigins([]string{"*"})
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-EventType"})
	methods := handlers.AllowedMethods([]string{
		http.MethodGet,
		http.MethodPut,
		http.MethodPost,
		http.MethodDelete,
		http.MethodOptions,
	})

	logger.Log("starting http server on " + port)
	// on the testing, dev server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: handlers.CORS(origins, headers, methods)(router),
	}


	log.Fatal(server.ListenAndServe())
}
