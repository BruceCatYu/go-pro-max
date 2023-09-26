package gopromax

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Response struct {
	w *http.ResponseWriter
	r *http.Request
}

func NewResponse(w *http.ResponseWriter, r *http.Request) Response {
	return Response{w: w, r: r}
}

// func (res *Response) Upgrade(respHeader http.Header) (*websocket.Conn, error) {
// 	upgrader := websocket.Upgrader{}
// 	return upgrader.Upgrade(*res.w, res.r, respHeader)
// }

func (res *Response) PathVar(v string) []string {
	return res.r.URL.Query()[v]
}

func (res *Response) BodyBytes() ([]byte, error) {
	return io.ReadAll(res.r.Body)
}

func (res *Response) BodyString() (string, error) {
	if bodyBytes, err := res.BodyBytes(); err != nil {
		return "", err
	} else {
		return string(bodyBytes), nil
	}
}

func (res *Response) BodyJSON(v any) error {
	if bodyBytes, err := res.BodyBytes(); err != nil {
		return err
	} else {
		return json.Unmarshal(bodyBytes, v)
	}
}

func (res *Response) RespJSON(httpCode int, body any) {
	(*res.w).Header().Set("Content-Type", "application/json; charset=utf-8")
	bytes, err := json.Marshal(body)
	if err != nil {
		fmt.Fprintln(*res.w, err)
	}
	(*res.w).WriteHeader(httpCode)
	fmt.Fprintln(*res.w, string(bytes))
}

func (res *Response) RespOk(data any) {
	res.RespJSON(http.StatusOK, data)
}

func (res *Response) RespErr(data any) {
	res.RespJSON(http.StatusBadRequest, data)
}

func (res *Response) RespNoAuth(data any) {
	res.RespJSON(http.StatusUnauthorized, data)
}
