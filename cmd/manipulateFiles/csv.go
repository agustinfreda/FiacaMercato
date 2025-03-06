package manipulatefiles

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func ReadCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
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
