package api

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func ConsumarApi(URL string) {
	data, err := fetchJSON(URL)
	if err != nil {
		fmt.Println(err)
		return
	}

	products, err := extractInteractivities(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := writeCSV("productos.csv", products); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Archivo productos.csv creado correctamente.")
}
func writeCSV(filename string, products map[string]interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error al crear el archivo CSV: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"id", "nombre", "cantidad", "precio_oferta", "precio", "precioPorKilo", "categoria"}
	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("error al escribir encabezado en CSV: %v", err)
	}

	for id, rawProduct := range products {
		product, ok := rawProduct.(map[string]interface{})
		if !ok {
			continue
		}

		nombre := fmt.Sprintf("%v", product["field1"])
		posible_cantidad := fmt.Sprintf("%v", product["field2"])
		precio_oferta := fmt.Sprintf("%v", product["offer_price"])
		precio := fmt.Sprintf("%v", product["price"])
		precioPorKilo := fmt.Sprintf("%v", product["price_for_kg"])
		categoria := fmt.Sprintf("%v", product["category"])

		if err := writer.Write([]string{id, nombre, posible_cantidad, precio_oferta, precio, precioPorKilo, categoria}); err != nil {
			return fmt.Errorf("error al escribir en CSV: %v", err)
		}
	}

	return nil
}

func extractInteractivities(data map[string]interface{}) (map[string]interface{}, error) {
	clientRaw, ok := data["client"]
	if !ok {
		return nil, fmt.Errorf("'client' no encontrado en el JSON")
	}

	client, ok := clientRaw.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("'client' no tiene el formato esperado")
	}

	interactivitiesRaw, ok := client["interactivities"]
	if !ok {
		return nil, fmt.Errorf("'interactivities' no encontrado en 'client'")
	}

	interactivities, ok := interactivitiesRaw.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("'interactivities' no tiene el formato esperado")
	}

	return interactivities, nil
}

func fetchJSON(url string) (map[string]interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error al obtener el JSON: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer el JSON: %v", err)
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("error al parsear el JSON: %v", err)
	}

	return data, nil
}
