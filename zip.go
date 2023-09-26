package gopromax

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
)

func Gzip(b []byte) []byte {
	var bf bytes.Buffer
	w := gzip.NewWriter(&bf)
	w.Write(b)
	w.Close()
	return bf.Bytes()
}

func UnGzip(b []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	return io.ReadAll(r)
}

func UnGzipToJson[T any](v *T, b []byte) (*T, error) {
	r, err := gzip.NewReader(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(r).Decode(v); err != nil {
		return nil, err
	}
	return v, nil
}
