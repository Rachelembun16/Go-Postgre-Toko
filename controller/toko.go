package controller

import (
	"encoding/json"
	"fmt"
	"strconv"

	"log"
	"net/http"

	"go-postgres-toko/models"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type response struct {
	ID      int64  `json:"id,omitemty"`
	Message string `json:"message,omitempty"`
}

type Response struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Data    []models.Barang `json:"data"`
}

//tambah barang
func TmbhBarang(w http.ResponseWriter, r *http.Request) {
	var barang models.Barang

	err := json.NewDecoder(r.Body).Decode(&barang)

	if err != nil {
		log.Fatalf("Tidak bisa mengencoded dari recuest body. %v", err)
	}

	//panggil models nya lalu insert barang
	InsertID := models.TambahBarang(barang)

	//response objectnya
	res := response{
		ID:      InsertID,
		Message: "Data barang telah ditambahkan",
	}

	//kirim response
	json.NewEncoder(w).Encode(res)

}

func AmbilBarang(w http.ResponseWriter, r *http.Request) {
	//set header
	w.Header().Set("Context-Type", "application/x-www-form-urlecoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//mengambil idbarang dari parameter request, ketnya adalah "id"
	params := mux.Vars(r)

	//konversi id dari string ke int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Todak bisa mengubah dari String ke Int. %v", err)
	}

	//memanggil models ambilsatubarang dengan parameter id yang nantinya akan mengambil single data
	barang, err := models.AmbilSatuBarang(int64(id))

	if err != nil {
		log.Fatalf("Tidak dapat mengambil data barang. %v", err)
	}

	//kirim response
	json.NewEncoder(w).Encode(barang)
}

func AmbilSemuaBarang(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//memanggil models AmbilSemuaBuku()
	barangs, err := models.AmbilSemuaBarang()

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data. %v", err)
	}

	var response Response
	response.Status = 1
	response.Message = "Success"
	response.Data = barangs

	//kirim semua response
	json.NewEncoder(w).Encode(response)
}

func UpdateBarang(w http.ResponseWriter, r *http.Request) {
	//ambil request parameter idnya
	params := mux.Vars(r)

	//konversi string - int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int. %v", err)
	}

	//var barang dengan type models.barang
	var barang models.Barang

	//decode json request ke variable barang
	err = json.NewDecoder(r.Body).Decode(&barang)

	if err != nil {
		log.Fatalf("Tidak bisa decode request body. %v", err)
	}

	updateRows := models.UpdateBarang(int64(id), barang)

	msg := fmt.Sprintf("Buku telah berhasil diupdate. Jumlah yang diupdate %v rows/record", updateRows)

	//respond msg
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	//kirim berupa response
	json.NewEncoder(w).Encode(res)
}

func HapusBarang(w http.ResponseWriter, r *http.Request) {
	// kita ambil request parameter idnya
	params := mux.Vars(r)

	// konversikan ke int yang sebelumnya adalah string
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
	}

	// panggil fungsi hapusbarang , dan convert int ke int64
	deletedRows := models.HapusBarang(int64(id))

	// ini adalah format message berupa string
	msg := fmt.Sprintf("barang sukses di hapus. Total data yang dihapus %v", deletedRows)

	// ini adalah format reponse message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}
