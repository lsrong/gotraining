package web

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const jsonContentType = "application/json; charset=utf-8"

type D map[string]interface{}

func RespondJSON(ctx context.Context, w http.ResponseWriter, data interface{}, statusCode int) error {
	// Not contents
	if statusCode == http.StatusNoContent {
		w.WriteHeader(statusCode)
		return nil
	}

	streamJson, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal response data: %v", err)
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-type", jsonContentType)

	if _, err = w.Write(streamJson); err != nil {
		return fmt.Errorf("failed to write response data: %v", err)
	}

	return nil
}
