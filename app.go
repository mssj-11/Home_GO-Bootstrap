package main // Paquete principal

// Importando Librerias
import (
	"fmt"           // Paquete para formatear la entrada y la salida.
	"net/http"      // Paquete para realizar la creacion y conexion del servidor
	"text/template" // Devolviendo archivos HTML
)

// Manejando Datos
type Users struct {
	Name   string
	Skills string
	Age    int
}

// Funcion Index
func Index(rw http.ResponseWriter, r *http.Request) {
	// Agregando el template
	template, error := template.ParseFiles("templates/index.html")

	// Creando una estructura
	user := Users{"Michael S.S.", "desarrollador backend", 24}

	// Si error es diferente de nil
	if error != nil {
		panic(error)
	} else { // Si no lo hay
		// Renderizando un Template
		template.Execute(rw, user)
	}
}

// Funcion Principal
func main() {
	// Agregando la ruta, retornara la funcion Index
	http.HandleFunc("/", Index)
	// CREACION DEL SERVIDOR
	// Creando el mensaje de la conexion del servidor
	fmt.Println("El servidor esta corriendo en el puerto: 3000")
	fmt.Println("Corriendo en el Servidor: http://localhost:3000")
	// Indicando en que puerto correra el servidor(direccion:puerto), valor
	http.ListenAndServe("localhost:3000", nil)
}
