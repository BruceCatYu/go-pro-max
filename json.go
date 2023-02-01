package gopromax

import (
	"encoding/json"
)

// MarshalBeautyJSON return beauty json
func MarshalBeautyJSON(v any) ([]byte, error) {
	return json.MarshalIndent(v, "", "    ")
}

// MapToStruct marshal map and then unmarshal to struct
// s must be a pointer of variable
func MapToStruct(m, s any) error {
	content, err := json.Marshal(m)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(content, s); err != nil {
		return err
	}
	return nil
}
