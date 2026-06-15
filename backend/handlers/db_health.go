package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/krish8learn/InterviewSetup/backend/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

const healthTimeout = 5 * time.Second

// DBHealth returns a handler that checks MongoDB connectivity.
func DBHealth(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), healthTimeout)
		defer cancel()

		if err := client.Ping(ctx, nil); err != nil {
			http.Error(w, `{"error":"MongoDB unreachable"}`, http.StatusServiceUnavailable)
			return
		}

		utils.WriteJSON(w, http.StatusOK, map[string]string{"status": "MongoDB connected"})
	}
}
