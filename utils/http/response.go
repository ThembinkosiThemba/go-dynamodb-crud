package http

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Status int         `json:"status"`
	Result interface{} `json:"result"`
}

func NewResponse(data interface{}, status int) *Response {
	return &Response{
		Status: status,
		Result: data,
	}
}

func (resp *Response) bytes() []byte {
	data, _ := json.Marshal(resp)
	return data

}

func (resp *Response) string() string {
	return string(resp.bytes())

}

func (resp *Response) SendResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(resp.Status)
	_, _ = w.Write(resp.bytes())
	log.Println(resp.string())

}

func StatusOk(w http.ResponseWriter, r *http.Request, data interface{}) {
	NewResponse(data, http.StatusOK).SendResponse(w, r)
}

func StatusNoContent(w http.ResponseWriter, r *http.Request) {
	NewResponse(nil, http.StatusNoContent).SendResponse(w, r)

}

func StatusBadRequest(w http.ResponseWriter, r *http.Request, err error) {
	data := map[string]interface{}{"error": err.Error()}
	NewResponse(data, http.StatusBadRequest).SendResponse(w, r)
}

func StatusNotFound(w http.ResponseWriter, r *http.Request, err error) {
	data := map[string]interface{}{"error": err.Error()}
	NewResponse(data, http.StatusNotFound).SendResponse(w, r)
}
func StatusMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	NewResponse(nil, http.StatusMethodNotAllowed).SendResponse(w, r)
}

func StatusConflict(w http.ResponseWriter, r *http.Request, err error) {
	data := map[string]interface{}{"error": err.Error()}
	NewResponse(data, http.StatusConflict).SendResponse(w, r)
}

func StatusInternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	data := map[string]interface{}{"error": err.Error()}
	NewResponse(data, http.StatusInternalServerError).SendResponse(w, r)
}
