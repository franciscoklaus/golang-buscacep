package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ViaCep struct {
	Neighborhood string `json:"bairro"`
	ZipCode      string `json:"cep"`
	Complement   string `json:"complemento"`
	AreaCode     string `json:"ddd"`
	Gia          string `json:"gia"`
	Ibge         string `json:"ibge"`
	City         string `json:"localidade"`
	Street       string `json:"logradouro"`
	Siafi        string `json:"siafi"`
	State        string `json:"uf"`
}

func main() {
	// Using closures
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("Hello, World!\n"))
	//})

	http.HandleFunc("/", BuscaCepHandler)
	fmt.Println("Servidor rodando na porta: 8080...")

	http.ListenAndServe(":8080", nil)

}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cep, err := BuscaCep(cepParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	result, err := json.Marshal(cep)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(result)

	// Shortcut
	//json.NewEncoder(w).Encode(cep)

}

func BuscaCep(cep string) (*ViaCep, error) {
	req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err

	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var c ViaCep
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil

}
