package main

import (
	"api-go/internal/service"
	"api-go/internal/store"
	"api-go/internal/transport"
	"database/sql"

	"fmt"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
)

func main() {
	// Conectar a SQLite

	db, err := sql.Open("sqlite", "./books.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Crear el table si no existe

	q := `
	CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		author TEXT NOT NULL
	);
	`
	if _, err = db.Exec(q); err != nil {
		log.Fatal(err)
	}

	//Inyectar dependencias
	bookStore := store.New(db)
	bookService := service.New(bookStore)
	bookHandler := transport.New(bookService)

	//Configurar las rutas
	http.HandleFunc("/books", bookHandler.HandleBooks)
	http.HandleFunc("/books/", bookHandler.HandleBookByID)

	fmt.Println("Servicor ejecutandose en http://localhost:8080")
	fmt.Println("Enpoinds disponibles:")
	fmt.Println("Get /books - Obtener todos los libros")
	fmt.Println("Post /books - Crear un nuevo libro")
	fmt.Println("Get /books/{id} - Obtener un libro por ID")
	fmt.Println("Put /books/{id} - Actualizar un libro por ID")
	fmt.Println("Delete /books/{id} - Eliminar un libro por ID")

	//Iniciar y escuchar el Servidor
	log.Fatal(http.ListenAndServe(":8080", nil))

}
