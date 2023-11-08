package coder

import (
	"encoding/json"
	"net/http"
)

// Раскодирование тела из запроса
func Decode[T any](r *http.Request) (T, error) {
	decoder := json.NewDecoder(r.Body)
	var decodeType T

	err := decoder.Decode(&decodeType)
	if err != nil {
		return decodeType, err
	}

	return decodeType, nil
}
