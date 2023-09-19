package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var database *mongo.Database

// InitDB inisialisasi koneksi ke MongoDB
func InitDB() error {
	// Atur konfigurasi koneksi MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Atur timeout untuk koneksi ke MongoDB
	timeout := 10 * time.Second
	clientOptions.ConnectTimeout = &timeout

	// Buat koneksi ke MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Periksa apakah koneksi ke MongoDB berhasil
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Koneksi ke MongoDB berhasil")

	// Pilih database yang akan digunakan
	database = client.Database("employee_db")

	return nil
}

// GetDB mengembalikan instance database MongoDB
func GetDB() *mongo.Database {
	return database
}
