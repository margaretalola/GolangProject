// Buatlah API method GET pada Golang dengan mengambil NPM,
// NAMA, MATKUL dari
// kodingan berisikan data masing-masing dan cek lewat postman! (45 Point)
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// making struct of mahasiswa
type mataKuliah struct {
	NPM    int    `json:"npm"`
	NAMA   string `json:"nama"`
	MATKUL string `json:"matkul"`
}

// make a function of struct
func dataMatkul() []mataKuliah {
	mk := []mataKuliah{ //array; cause they're saving many data with the same type
		{
			NPM:    5042,
			NAMA:   "Mar",
			MATKUL: "Algoritma Pemrograman",
		},
	}
	return mk
}
func GetMatkul(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		mk := dataMatkul()
		datamatkul, err := json.Marshal(mk)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(datamatkul)
		return
	}

	http.Error(w, "ini error", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/matkul", GetMatkul)
	fmt.Println("Server running...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
