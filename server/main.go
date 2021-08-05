package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"net/http"
	"os"
	"time"
	"way/pkg/db"
	"way/server/handler"
	"way/server/routes"
)

func main() {

	godotenv.Load("../.env")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// get port
	port, valid := os.LookupEnv("PORT")
	if !valid {
		log.Println("Invalid port")
		return
	}

	var databaseInfo db.DatabaseInfo

	// get database url
	databaseInfo.Host, valid = os.LookupEnv("DATABASE_URL")
	if !valid {
		log.Println("invalid database_url")
		return
	}
	// get database port
	databaseInfo.Port, valid = os.LookupEnv("DATABASE_PORT")
	if !valid {
		log.Println("invalid database_port")
		return
	}
	// get database user
	databaseInfo.User, valid = os.LookupEnv("DATABASE_USER")
	if !valid {
		log.Println("invalid database_user")
		return
	}
	// get database password
	databaseInfo.Password, valid = os.LookupEnv("DATABASE_PASSWORD")
	if !valid {
		log.Println("invalid database_password")
		return
	}
	// get database name
	databaseInfo.DatabaseName, valid = os.LookupEnv("DATABASE_NAME")
	if !valid {
		log.Println("invalid database_name")
		return
	}

	// connect to database
	err := databaseInfo.Connect()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Database connection established")

	// Close database connection when program exits
	defer func() {
		err = databaseInfo.Disconnect()
		if err != nil {
			log.Println(err)
			return
		}
	}()

	router := routes.Router()
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

	log.Println("starting http server on " + port)
	// on the testing, dev server
	server := &http.Server{
		Addr:         fmt.Sprintf(":%v", port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      handlers.CORS(origins, headers, methods)(router),
	}

	log.Fatal(server.ListenAndServe())
}
