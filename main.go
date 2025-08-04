package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	// Obtener el valor de la ruta: /user/{id}
	vars := mux.Vars(r)
	id := vars["id"]

	url := "http://k8s-2.microservices:8080/user/" + id

		// Hacemos la solicitud
		resp, err := http.Get(url)
		if err != nil {
			// Seteamos status y content-type
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadGateway)

			// Escribimos un body personalizado en JSON
			fmt.Fprintf(w, `{"error": "No se pudo contactar con el servicio externo"}`)
			return
		}
		defer resp.Body.Close()

    // Copiar el header Content-Type para que el cliente sepa que es JSON
    w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))

    // Copiar el status code
    w.WriteHeader(resp.StatusCode)

    // Escribir el body recibido directamente en la respuesta
    _, err = io.Copy(w, resp.Body)
    if err != nil {
        fmt.Println("Error al copiar el body:", err)
    }
}

func main() {
	// Crear el router
	r := mux.NewRouter()

	// Registrar la ruta con parámetro dinámico
	r.HandleFunc("/user/{id}", getUserHandler).Methods("GET")

	// Iniciar el servidor
	fmt.Println("Servidor corriendo en el puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}