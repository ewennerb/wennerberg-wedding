package routes

import (
	"github.com/gorilla/mux"
	"wennerbeg-wedding/controllers"
)

// Routes -> define endpoints
func Routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/household/{guest_name}", controllers.GetHouseholdByGuest).Methods("GET")
	router.HandleFunc("/household/{id}", controllers.UpdateHousehold).Methods("POST")
	router.HandleFunc("/info", controllers.GetGuestInfo).Methods("GET")
	return router
}
