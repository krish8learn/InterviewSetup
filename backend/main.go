package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/krish8learn/InterviewSetup/backend/db"
	"github.com/krish8learn/InterviewSetup/backend/handlers"
	"github.com/krish8learn/InterviewSetup/backend/model"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	env     *model.EnvKeys
	client  *mongo.Client
	mysqlDB *sql.DB
)

func init() {
	// Load environment configuration
	env = model.LoadEnv()

	var err error
	client, err = db.Connect(env.MongoURI)
	if err != nil {
		log.Fatalf("db.Connect failed: %v", err)
	}
	log.Println("MongoDB connected")

	mysqlDB, err = db.ConnectMySQL(env.MySQLDSN)
	if err != nil {
		log.Fatalf("db.ConnectMySQL failed: %v", err)
	}
	log.Println("MySQL connected")
}

func main() {
	defer mysqlDB.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.Root())
	mux.HandleFunc("GET /mongo-health", handlers.MongoHealth(client))
	mux.HandleFunc("GET /mysql-health", handlers.MySQLHealth(mysqlDB))

	// ---------------------------------------------------------------
	// TODO (interview): register your API routes here, e.g.
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
