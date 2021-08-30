package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/katzien/go-structure-examples/domain-hex-actor/pkg/adding"
	"github.com/katzien/go-structure-examples/domain-hex-actor/pkg/backuping"
	"github.com/katzien/go-structure-examples/domain-hex-actor/pkg/listing"
)

func Handler(a adding.Service, l listing.Service, r backuping.Service) http.Handler {
	router := httprouter.New()

	router.GET("/users", getUsers(l))
	router.GET("/users/:id", getUser(l))
	router.GET("/users/:id/backups", getUserBackups(l))

	router.POST("/users", addUser(a))
	router.POST("/users/:id/backups", addUserBackup(r))

	return router
}

// addUser returns a handler for POST /users requests
func addUser(s adding.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		decoder := json.NewDecoder(r.Body)

		var newUser adding.User
		err := decoder.Decode(&newUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		s.AddUser(newUser)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("New user added.")
	}
}

// addUserBackup returns a handler for POST /users/:id/backups requests
func addUserBackup(s backuping.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ID, err := strconv.Atoi(p.ByName("id"))
		if err != nil {
			http.Error(w, fmt.Sprintf("%s is not a valid User ID, it must be a number.", p.ByName("id")), http.StatusBadRequest)
			return
		}

		var newBackup backuping.Backup
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&newBackup); err != nil {
			http.Error(w, "Failed to parse backup", http.StatusBadRequest)
		}

		newBackup.UserID = ID

		s.AddUserBackup(newBackup)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("New user backup added.")
	}
}

// getUsers returns a handler for GET /users requests
func getUsers(s listing.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		list := s.GetUsers()
		json.NewEncoder(w).Encode(list)
	}
}

// getUser returns a handler for GET /users/:id requests
func getUser(s listing.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ID, err := strconv.Atoi(p.ByName("id"))
		if err != nil {
			http.Error(w, fmt.Sprintf("%s is not a valid user ID, it must be a number.", p.ByName("id")), http.StatusBadRequest)
			return
		}

		user, err := s.GetUser(ID)
		if err == listing.ErrNotFound {
			http.Error(w, "The user you requested does not exist.", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

// getUserBackups returns a handler for GET /users/:id/backups requests
func getUserBackups(s listing.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ID, err := strconv.Atoi(p.ByName("id"))
		if err != nil {
			http.Error(w, fmt.Sprintf("%s is not a valid user ID, it must be a number.", p.ByName("id")), http.StatusBadRequest)
			return
		}

		backups := s.GetUserBackups(ID)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(backups)
	}
}
