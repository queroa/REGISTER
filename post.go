package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Coche struct {
	Id          int    `json:id`
	Brand       string `json:brand`
	Model       string `json:model`
	Horse_power int    `json:horse_power`
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/service/v1/cars", endpointFunc2JSONInput).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))

}

var config string

func initialize() {
	data, err := ioutil.ReadFile("/service/conf")
	if err != nil {
		fmt.Println("error reading configuration", err)
		return
	}
	config = string(data)
	fmt.Println("IP/NOMBRE DE LA DB:", config)
}
func endpointFunc2JSONInput(w http.ResponseWriter, r *http.Request) {
	var coche Coche

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 100))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &coche); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	///////////////// ACCESO SQL //////////////////////
	//	fmt.Println("POSTER MySQL Connection")

	db, err := sql.Open("mysql", "root:Miquelpiloto1@tcp("+config+":3306)/Concesionario")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	coche.Id = rand.Intn(2147483647)
	if coche.Brand != "" {
		sqlStatement := `INSERT INTO Coches (id,brand,model,horse_power) VALUES (?,?,?,?)`
		_, err = db.Exec(sqlStatement, coche.Id, coche.Brand, coche.Model, coche.Horse_power)

		if err != nil {
			panic(err.Error())
		}
	}

	b, err := json.Marshal(coche)

	//		fmt.Println("Conectado e insertado.")
	fmt.Fprintln(w, string(b))

}

func main() {
	initialize()
	handleRequests()

}
