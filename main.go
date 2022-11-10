// main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type inl_s3_gdrive_mirror_log struct {
	id        int
	file_id   string
	file_name string
	owners    string
}

var db *gorm.DB
var err error

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func getData(w http.ResponseWriter, r *http.Request) {
	var inl_s3_gdrive_mirror_log inl_s3_gdrive_mirror_log
	rows := db.First(&inl_s3_gdrive_mirror_log)
	//fmt.Println("getData Hit:" + rows)
	json.NewEncoder(w).Encode(rows)
}

func deleteData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var deleteData inl_s3_gdrive_mirror_log
	result := db.Where("id = ?", id).Delete(&deleteData)

	return result.RowsAffected

}

func UpdatePayment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var updateData inl_s3_gdrive_mirror_log
	result := db.Model(&updateData).Where("id = ?", id).Updates(updateData)

	return result
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/getdata", getData)
	myRouter.HandleFunc("/updatedata/{id}", UpdatePayment).Methods("PATCH")
	myRouter.HandleFunc("/deletedata/{id}", deleteData).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {

	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable password=postgres")
	if err != nil {

		panic("failed to connect database")

	}

	defer db.Close()
	handleRequests()
}
