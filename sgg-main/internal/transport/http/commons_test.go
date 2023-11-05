package http_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
)

const (
	dateLayout = "2006-01-02"

	smallImage   = "https://upload.wikimedia.org/wikipedia/commons/thumb/2/2d/Snake_River_%285mb%29.jpg/512px-Snake_River_%285mb%29.jpg"
	largeImage   = "https://upload.wikimedia.org/wikipedia/commons/2/2d/Snake_River_%285mb%29.jpg"
	invalidImage = "https://upload.wikimedia.org/wikipedia/commons/2/2d/Snake_River_%285mb%29.jp"
)

func readResponse(w *httptest.ResponseRecorder, target interface{}) error {
	b, err := ioutil.ReadAll(w.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, target)
}
