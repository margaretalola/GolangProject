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
	mhs := []Mahasiswa{ //array; cause they're saving many data with the same type
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

/*
	func PostMahasiswa(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var mhs Mahasiswa
		if r.Method == "POST" {
			if r.Header.Get("Content-Type") == "application/json" {
				decodeJson := json.NewDecoder(r.Body)
				if err := decodeJson.Decode(&mhs); err != nil {
					log.Fatal(err)
				}
			} else {
				getID := r.PostFormValue("id")
				id, _ := strconv.Atoi(getID)
				getNim := r.PostFormValue("nim")
				nim, _ := strconv.Atoi(getNim)
				name := r.PostFormValue("name")
				mhs = Mahasiswa{
					ID:   id,
					NIM:  nim,
					Name: name,
				}
			}
			dataMahasiswa, _ := json.Marshal(mhs)
			w.Write(dataMahasiswa)
			return
		}
		http.Error(w, "hayo mau ngapain", http.StatusNotFound)
	}
*/

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

func main() {
	http.HandleFunc("/mahasiswa", GetMahasiswa)
	fmt.Println("Server running...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
