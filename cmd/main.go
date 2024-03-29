package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"service-profile/internal/handlers"
	"service-profile/pkg/database"
	"syscall"
)

func main() {
	// Inisialisasi koneksi database
	db := database.NewDatabase()

	defer db.Close()

	// Menginisialisasi handler dengan koneksi database
	handler := handlers.NewHandler(db)

	// Routing endpoints
	http.HandleFunc("/users", handler.GetUsers)
	http.HandleFunc("/users/create", handler.CreateUser)
	http.HandleFunc("/users/update", handler.UpdateUser)
	http.HandleFunc("/users/delete", handler.DeleteUser)

	// Menjalankan server dalam goroutine
	go func() {
		log.Println("Server started on :8080")
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Menangkap sinyal SIGINT (Ctrl+C) dan SIGTERM
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	// Menutup koneksi database sebelum keluar
	log.Println("Shutting down server...")
}
