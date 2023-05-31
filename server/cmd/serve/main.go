package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/cors"

	"github.com/bcinnovationlabs/Apps/Chatbot-AI/server/internal/controllers"
	"github.com/bcinnovationlabs/Apps/Chatbot-AI/server/pkg/utils"
)

func main() {
	// Load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the database connection
	db := initializeDB()
	defer db.Close()

	// Initialize the controllers with the database
	c := controllers.New(db)

	// Create a new HTTP router
	mux := http.NewServeMux()

	// Register middleware and handlers
	mux.HandleFunc("/", controllers.ChainMiddleware(c.HandleRequests))

	// Create a CORS-enabled handler
	handler := cors.AllowAll().Handler(mux)

	// Create the HTTP server
	server := &http.Server{
		Addr:    getServerAddress(),
		Handler: handler,
	}

	// Start the server in a separate goroutine
	go runServer(server)

	// Wait for termination signal
	waitForTerminationSignal(server)
}

// initializeDB initializes the database connection and returns the DB instance
func initializeDB() *sql.DB {
	databaseHost := os.Getenv("DATABASE_HOST")
	databasePort := os.Getenv("DATABASE_PORT")
	databaseUser := os.Getenv("DATABASE_USER")
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	databaseName := os.Getenv("DATABASE_NAME")

	db, err := utils.ConnectDB(databaseHost, databasePort, databaseUser, databasePassword, databaseName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db)

	return db
}

// getServerAddress returns the address for the HTTP server
func getServerAddress() string {
	serverDomain := os.Getenv("SERVER_API_DOMAIN")
	serverPort := os.Getenv("SERVER_PORT")

	return fmt.Sprintf("%s:%s", serverDomain, serverPort)
}

// runServer starts the HTTP server
func runServer(server *http.Server) {
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server error: %s\n", err)
	}
}

// waitForTerminationSignal waits for the termination signal and gracefully shuts down the server
func waitForTerminationSignal(server *http.Server) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %s", err)
	}

	log.Print("Server gracefully shutdown")
	log.Print("Server exited")
}
