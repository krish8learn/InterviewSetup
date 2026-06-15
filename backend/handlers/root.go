package handlers

import (
	"net/http"

	"github.com/krish8learn/InterviewSetup/backend/utils"
)

// Root returns a handler that responds with a simple JSON greeting.
func Root() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Welcome to the Interview Setup API!"})
	}
}
