package decode

import (
	"encoding/json"
	"io"
	"medods/pkg/validator"
	"reflect"
)

func Json(body io.ReadCloser, v any) error {
	err := json.NewDecoder(body).Decode(v)
	if err != nil {
		return err
	}

	if reflect.Slice != reflect.ValueOf(v).Kind() {
		return nil
	}

	return validator.Struct(v)
}
