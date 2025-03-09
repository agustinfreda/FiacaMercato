package manipulatefiles

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Quantity    string `json:"quantity"`
	Offer_Price string `json:"offer_price"`
	Price       string `json:"price"`
	PriceKG     string `json:"pricekg"`
	Category    string `json:"category"`
}

func ReadCsvFile(filePath string) ([]Product, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.TrimLeadingSpace = true

	_, e := csvReader.Read()
	if e != nil {
		return nil, err
	}
	var products []Product
	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		producto := Product{
			ID:          row[0],
			Name:        row[1],
			Quantity:    row[2],
			Offer_Price: row[3],
			Price:       row[4],
			PriceKG:     row[5],
			Category:    row[6],
		}

		products = append(products, producto)
	}

	return products, nil
}

func WriteCSV(filename string, products map[string]interface{}) error {
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
		if nombre == "" || nombre == "<nil>" { // Evita agregar productos sin nombre
			continue
		}

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
