package api

import (
	"fmt"
	"time"

	files "github.com/agustinfreda/FiacaMercato/cmd/manipulateFiles"
)

func obtenerFecha() string {
	fecha := time.Now().Format("02-01-2006")
	return fecha
}

func ConsumarApi(URL string) {
	data, err := files.FetchJSON(URL)
	if err != nil {
		fmt.Println(err)
		return
	}

	products, err := extractInteractivities(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	ruta := fmt.Sprintf("/home/agustin/Documentos/workspace/FiacaMercato/data/%v.csv", obtenerFecha())
	if err := files.WriteCSV(ruta, products); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Información almacenada correctamente.")
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
