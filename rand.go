package gopromax

import (
	"math/rand"
	"reflect"
	"time"
)

func RandomNumeric[T NumericType](left, right T) (res T) {
	if right < left {
		return
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	switch reflect.ValueOf(left).Kind() {
	case reflect.Int:
		res = T(r.Intn(int(right-left)) + int(left))
	case reflect.Int32:
		res = T(r.Int31n(int32(right-left)) + int32(left))
	case reflect.Int64:
		res = T(r.Int63n(int64(right-left)) + int64(left))
	case reflect.Float32:
		res = T(r.Float32()*float32(right-left) + float32(left))
	case reflect.Float64:
		res = T(r.Float64()*float64(right-left) + float64(left))
	}

	return
}
