package response

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, response HTTPResponse) error {
	w.WriteHeader(response.statusCode)
	if err := json.NewEncoder(w).Encode(response.payload); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return json.NewEncoder(w).Encode(map[string]string{"message": "internal server error"})
	}

	return nil
}
