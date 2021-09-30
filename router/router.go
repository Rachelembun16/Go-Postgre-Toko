package router

import (
	"go-postgres-toko/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/barang", controller.AmbilSemuaBarang).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/barang/{id}", controller.AmbilBarang).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/barang", controller.TmbhBarang).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/barang/{id}", controller.UpdateBarang).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/barang/{id}", controller.HapusBarang).Methods("DELETE", "OPTIONS")
	return router
}
