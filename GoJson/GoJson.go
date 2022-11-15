package main //penggunaan package main

//import library
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// making struct of mahasiswa
type Mahasiswa struct {
	ID   int    `json:"id"`
	NIM  int    `json:"nim"`
	Name string `json:"name"`
}

// make a function of struct
func NewMahasiswa() []Mahasiswa {
	mhs := []Mahasiswa{ 
		Mahasiswa{
			ID:   1,
			NIM:  123454,
			Name: "Didik Prabowo",
		},
		Mahasiswa{
			ID:   2,
			NIM:  923454,
			Name: "Joni Gunawan",
		},
		Mahasiswa{
			ID:   3,
			NIM:  923454,
			Name: "Muhammad Irwan",
		},
	}
	return mhs
}


// make get function for mahasiswa
func GetMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		mhs := NewMahasiswa()
		datamahasiswa, err := json.Marshal(mhs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(datamahasiswa)
		return
	}

	http.Error(w, "hayo mau ngapain", http.StatusNotFound)
}
//main function for handlefunc
func main() {
	http.HandleFunc("/mahasiswa", GetMahasiswa)
	fmt.Println("Server running...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
