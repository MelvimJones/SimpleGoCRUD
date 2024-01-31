package main

import (
	"crudApp/controllers"
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
)

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bem-vindo à minha aplicação Go!")
}

func main() {
	http.HandleFunc("/", welcomeHandler)

	// Carregar variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
		fmt.Println("Erro ao carregar o arquivo .env")
	}

	// Definir rota para visualizar os produtos (GET)
	http.HandleFunc("/products", controllers.ListProductsFromMongoDB)

	// Definir rota para criar um novo produto (POST)
	http.HandleFunc("/products/create", controllers.CreateProductInMongoDB)

	// Definir rota para atualizar um produto (PUT)
	http.HandleFunc("/products/update", controllers.UpdateProductInMongoDB)

	// Definir rota para deletar um produto (DELETE)
	http.HandleFunc("/products/delete", controllers.DeleteProductFromMongoDB)

	fmt.Println("Servidor iniciado em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
