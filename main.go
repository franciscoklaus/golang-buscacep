package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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
	for _, url := range os.Args[1:] {
		req, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error making request: %v\n", err)
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading response body: %v\n", err)
		}
		var data ViaCep
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing response: %v\n", err)
		}
		file, err := os.Create("zipCode.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating txt file: %v\n", err)
		}
		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("Neighborhood: %s,\nZipCode: %s,\nComplement: %s,"+
			"\nAreaCode: %s,\nGia: %s,\nIbge: %s,\nCity: %s,\nStreet:%s,\nSiafi: %s,\nState: %s",
			data.Neighborhood, data.ZipCode, data.Complement, data.AreaCode, data.Gia, data.Ibge, data.City,
			data.Street, data.Siafi, data.State))

		//fmt.Println(data)
	}
}
