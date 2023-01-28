package webserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	cmd "github.com/projEduardoValencio/cli-golang-cobra/internal/operations"
)

// Struct para server
type Server struct {
	Port string
}

// Handler de soma
func (s Server) SumHandler(w http.ResponseWriter, r *http.Request) {
	// Tratando método de requisição
	if r.Method != "POST" {
		w.WriteHeader(400)
		w.Write([]byte("Método incorreto"))
		return
	}

	// Convertendo tipos
	a, err := strconv.ParseFloat(r.URL.Query().Get("a"), 64)
	if err != nil {
		a = 0
	}
	b, err := strconv.ParseFloat(r.URL.Query().Get("b"), 64)
	if err != nil {
		b = 0
	}

	// Operando com struct
	calc := cmd.NewCalc()
	calc.A = a
	calc.B = b

	// Criando map para montar response
	response := make(map[string]string)
	response["msg"] = fmt.Sprintf("Resultado: %f", calc.Sum())

	// Realizando Marshal para JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("error happened in JSON marshal: %s", err.Error())
	}

	// Definindo header para JSON
	w.Header().Set("Content-Type", "application/json")
	// Status de Sucesso
	w.WriteHeader(200)
	// Write response
	w.Write(jsonResponse)
}

// Iniciando Server
func (s Server) Serve() {
	// Definindo endpoint
	http.HandleFunc("/sum", s.SumHandler)
	fmt.Println("GET: /sum -> Query: a, b")

	// Start server na porta passada
	log.Fatal(http.ListenAndServe(":"+s.Port, nil))
}
