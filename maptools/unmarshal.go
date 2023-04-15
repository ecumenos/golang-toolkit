package maptools

import (
	"encoding/json"

	"github.com/ecumenos/golang-toolkit/customerror"
)

// Unmarshal []byte into map[string]interface{}
func Unmarshal(in []byte) (map[string]interface{}, error) {
	out := make(map[string]interface{})
	if err := json.Unmarshal(in, &out); err != nil {
		return nil, customerror.NewToolkitFailure(err, "[Unmarshal] Can not unmarshal byte slice into map")
	}

	return out, nil
}
