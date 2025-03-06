package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	files "github.com/agustinfreda/FiacaMercato/cmd/manipulateFiles"
	"github.com/gorilla/mux"
)

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API")
}

func getAllProducts(w http.ResponseWriter, r *http.Request) {
	records := files.ReadCsvFile("/home/agustin/Documentos/workspace/FiacaMercato/data/23-02-2025.csv")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(records)
}

func Rutas() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/productos", getAllProducts).Methods("GET")

	// Servidor HTTP
	log.Fatal(http.ListenAndServe(":4567", router))
}
