package gopromax

import (
	"encoding/json"
)

// MarshalBeautyJSON return beauty json
func MarshalBeautyJSON(v any) ([]byte, error) {
	return json.MarshalIndent(v, "", "    ")
}
