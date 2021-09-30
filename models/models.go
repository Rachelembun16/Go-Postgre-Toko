package models

import (
	"database/sql"
	"fmt"
	"go-postgres-toko/config"
	"log"

	_ "github.com/lib/pq" // postgres golang driver
)

// Buku schema dari tabel Buku
// kita coba dengan jika datanya null
// jika return datanya ada yg null, silahkan pake NullString, contohnya dibawah
// Penulis       config.NullString `json:"penulis"`
type Barang struct {
	ID          int64  `json:"id_barang"`
	Nama        string `json:"nama"`
	Stok        string `json:"stok"`
	Harga       string `json:"harga"`
	Persen_laba string `json:"persen_laba"`
	Diskon      string `json:"diskon"`
}

func TambahBarang(barang Barang) int64 {

	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	// kita buat insert query
	// mengembalikan nilai id akan mengembalikan id dari buku yang dimasukkan ke db
	sqlStatement := `INSERT INTO barang (nama, stok, harga, persen_laba, diskon) VALUES ($1, $2, $3, $4, $5) RETURNING id_barang`

	// id yang dimasukkan akan disimpan di id ini
	var id int64

	// Scan function akan menyimpan insert id didalam id id
	err := db.QueryRow(sqlStatement, barang.Nama, barang.Stok, barang.Harga, barang.Persen_laba, barang.Diskon).Scan(&id)

	if err != nil {
		log.Fatalf("Tidak Bisa mengeksekusi query. %v", err)
	}

	fmt.Printf("Insert data single record %v ", id)

	// return insert id
	return id
}

// ambil satu buku
func AmbilSemuaBarang() ([]Barang, error) {
	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	var barangs []Barang

	// kita buat select query
	sqlStatement := `SELECT * FROM barang`

	// mengeksekusi sql query
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}

	// kita tutup eksekusi proses sql qeurynya
	defer rows.Close()

	// kita iterasi mengambil datanya
	for rows.Next() {
		var barang Barang

		// kita ambil datanya dan unmarshal ke structnya
		err = rows.Scan(&barang.ID, &barang.Nama, &barang.Stok, &barang.Harga, &barang.Persen_laba, &barang.Diskon)

		if err != nil {
			log.Fatalf("tidak bisa mengambil data. %v", err)
		}

		// masukkan kedalam slice bukus
		barangs = append(barangs, barang)

	}

	// return empty buku atau jika error
	return barangs, err
}

// mengambil satu buku
func AmbilSatuBarang(id int64) (Barang, error) {
	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	var barang Barang

	// buat sql query
	sqlStatement := `SELECT * FROM barang WHERE id_barang=$1`

	// eksekusi sql statement
	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&barang.ID, &barang.Nama, &barang.Stok, &barang.Harga, &barang.Persen_laba, &barang.Diskon)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Tidak ada data yang dicari!")
		return barang, nil
	case nil:
		return barang, nil
	default:
		log.Fatalf("tidak bisa mengambil data. %v", err)
	}

	return barang, err
}

// update user in the DB
func UpdateBarang(id int64, barang Barang) int64 {

	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	// kita buat sql query create
	sqlStatement := `UPDATE barnag SET nama=$2, stok=$3, harga=$4, persen_laba=$5, diskon=$6 WHERE id_barang=$1`

	// eksekusi sql statement
	res, err := db.Exec(sqlStatement, id, barang.Nama, barang.Stok, barang.Harga, barang.Persen_laba, barang.Diskon)

	if err != nil {
		log.Fatalf("Tidak bisa mengeksekusi query. %v", err)
	}

	// cek berapa banyak row/data yang diupdate
	rowsAffected, err := res.RowsAffected()

	//kita cek
	if err != nil {
		log.Fatalf("Error ketika mengecheck rows/data yang diupdate. %v", err)
	}

	fmt.Printf("Total rows/record yang diupdate %v\n", rowsAffected)

	return rowsAffected
}

func HapusBarang(id int64) int64 {

	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	// buat sql query
	sqlStatement := `DELETE FROM barang WHERE id_barang=$1`

	// eksekusi sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}

	// cek berapa jumlah data/row yang di hapus
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("tidak bisa mencari data. %v", err)
	}

	fmt.Printf("Total data yang terhapus %v", rowsAffected)

	return rowsAffected
}
