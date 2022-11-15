package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// struct matkul
type Matkul struct {
	Kode_MK     string `json:"kd_mk"`
	Nama_Matkul string `json:"nama_mk"`
	SKS         int    `json: "sks"`
}

type AllMatkul []Matkul

// isi data matkul
var Mk = AllMatkul{
	{
		Kode_MK:     "IT045203",
		Nama_Matkul: "Algoritma Pemrograman 3",
		SKS:         2,
	},
	{
		Kode_MK:     "IT045215",
		Nama_Matkul: "Matematika Informatika 3",
		SKS:         2,
	},
	{
		Kode_MK:     "IT045217",
		Nama_Matkul: "Matematika Lanjut 1",
		SKS:         2,
	},
}

// fungsi handle untuk Get dan Post
func matkulHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" { 	//get function
		datamatkul, err := json.Marshal(Mk)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(datamatkul)
		return

		//post function
	} else if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		var mkl Matkul

		if r.Header.Get("Content-Type") == "application/json" {
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&mkl); err != nil {
				log.Fatal(err)
			}
		} else {
			kd_mk := r.PostFormValue("kd_mk")
			nama_mk := r.PostFormValue("nama_mk")
			getSks := r.PostFormValue("sks")
			sks, _ := strconv.Atoi(getSks)
			mkl = Matkul{
				Kode_MK:     kd_mk,
				Nama_Matkul: nama_mk,
				SKS:         sks,
			}
		}
		//menambahkan data kedalam matkul
		Mk = append(Mk, mkl)
		dataMatkul, _ := json.Marshal(mkl)
		w.Write(dataMatkul)
		return
	}
	//pesan error yang akan tampil ketika terjadi kesalahan
	http.Error(w, "Hayo mau ngapain", http.StatusNotFound)
}

// main function untuk menjalankan matkulHandler
func main() {
	http.HandleFunc("/matkul", matkulHandler) //localhost:8000/matkul
	fmt.Println("Server running...")
	//using localhost:8000
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

//taruh sebelum func matkul handler
//func getMatkul(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)
// 		json.NewEncoder(w).Encode(Mk)
// 		return
// 	}
// 	http.Error(w, "Error!", http.StatusNotFound)
// }

// func postMatkul(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var mkl Matkul
// 	if r.Method == "POST" {
// 		reqbody, err := ioutil.ReadAll(r.Body)
// 		if err != nil {
// 			fmt.Fprintf(w, "Harap Masukkan Data")
// 		}
// 		json.Unmarshal(reqbody, &mkl)
// 		Mk = append(Mk, mkl)
// 		json.NewEncoder(w).Encode(mkl)
// 		return
// 	}
// 	http.Error(w, "Error!", http.StatusNotFound)
// }
