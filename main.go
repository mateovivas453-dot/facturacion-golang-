package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Estructura de la Matrícula/Factura
type FacturaMatricula struct {
	ID           int       `json:"id"`
	Estudiante   string    `json:"estudiante"`
	Materia      string    `json:"materia"`
	Monto        float64   `json:"monto"`
	FechaEmision time.Time `json:"fecha_emision"`
}

// "Base de datos" en memoria (Slice)
var historialFacturas []FacturaMatricula
var contadorID = 1

func main() {
	// Rutas de la API
	http.HandleFunc("/listar", listarFacturas)   // GET: Ver todas las matrículas
	http.HandleFunc("/matricular", crearFactura) // POST: Registrar nueva y facturar

	port := ":8080"
	fmt.Printf("Servidor local iniciado en http://localhost%s\n", port)
	fmt.Println("Usa Ctrl+C en PowerShell para detenerlo.")
	
	log.Fatal(http.ListenAndServe(port, nil))
}

// Función para listar (GET)
func listarFacturas(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(historialFacturas)
}

// Función para crear (POST)
func crearFactura(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Usa POST para enviar datos", http.StatusMethodNotAllowed)
		return
	}

	var nueva FacturaMatricula
	if err := json.NewDecoder(r.Body).Decode(&nueva); err != nil {
		http.Error(w, "Error en formato JSON", http.StatusBadRequest)
		return
	}

	// Lógica de negocio automática
	nueva.ID = contadorID
	nueva.FechaEmision = time.Now()
	contadorID++

	historialFacturas = append(historialFacturas, nueva)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"mensaje": "Matrícula y Factura generada con éxito",
		"estado":  "Pagado",
	})
}
