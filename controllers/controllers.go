package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Carlos-GitH/go-rest-api/database"
	"github.com/Carlos-GitH/go-rest-api/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	println("Home Page")
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-type", "aplication/json")
	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-type", "aplication/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personalidade

	database.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)
}

func CriaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-type", "aplication/json")
	var novaPersonalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&novaPersonalidade)
	database.DB.Create(&novaPersonalidade)
	json.NewEncoder(w).Encode(novaPersonalidade)
}

func DeletaPersonalidade(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-type", "aplication/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personalidade
	database.DB.Delete(&p, id)
	json.NewEncoder(w).Encode(p)
}

func EditaPersonalidade(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-type", "aplication/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personalidade
	database.DB.First(&p, id)
	json.NewDecoder(r.Body).Decode(&p)
	database.DB.Save(&p)
	json.NewEncoder(w).Encode(p)
}
