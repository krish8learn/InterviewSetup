package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/krish8learn/InterviewSetup/backend/db"
	"github.com/krish8learn/InterviewSetup/backend/handlers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, reading from environment")
	}

	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("MONGO_URI is not set")
	}

	client, err := db.Connect(uri)
	if err != nil {
		log.Fatalf("db.Connect failed: %v", err)
	}
	log.Println("MongoDB connected")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.Root())
	mux.HandleFunc("GET /db-health", handlers.DBHealth(client))

	// ---------------------------------------------------------------
	// TODO (interview): register your API routes here, e.g.
	//   mux.HandleFunc("POST /items", createItem)
	//   mux.HandleFunc("GET /items/{id}", getItem)
	// ---------------------------------------------------------------

	log.Println("server listening on :8080")
	if err := http.ListenAndServe(":8080", cors(mux)); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
