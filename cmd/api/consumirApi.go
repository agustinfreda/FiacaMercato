package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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
