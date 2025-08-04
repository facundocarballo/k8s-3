package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

// func main() {
// 	// URL a la que se le va a hacer la solicitud GET
// 	url := "http://k8s-2.microservices:8080/user"

// 	// Hacemos la solicitud
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		fmt.Println("Error al hacer la solicitud:", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	// Leemos el cuerpo de la respuesta
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Error al leer el cuerpo:", err)
// 		return
// 	}

// 	// Imprimimos la respuesta
// 	fmt.Println("Código de estado:", resp.StatusCode)
// 	fmt.Println("Respuesta:")
// 	fmt.Println(string(body))
// }

type User struct {
	ID    string `json:"id"`
	USERNAME string `json:"nombre"`
	PASSWORD string `json:"password"`
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	// Obtener el valor de la ruta: /user/{id}
	vars := mux.Vars(r)
	id := vars["id"]

	url := "http://k8s-2.microservices:8080/user/" + id

		// Hacemos la solicitud
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error al hacer la solicitud:", err)
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