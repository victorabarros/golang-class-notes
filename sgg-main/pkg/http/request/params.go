package request

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

// Bind binds the request body of the http request to the entity
// that is used to receive the request content
func Bind(r *http.Request, requestEntity interface{}) error {
	return json.NewDecoder(r.Body).Decode(requestEntity)
}

// BindParams binds route params and query params to the entity that
// is used to receive the params of the request
func BindParams(r *http.Request, requestEntity interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	if err := schema.NewDecoder().Decode(requestEntity, r.Form); err != nil {
		if !strings.Contains(err.Error(), "schema: invalid path") {
			return err
		}
	}

	vars := mux.Vars(r)
	b, err := json.Marshal(vars)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, requestEntity)
}
