package http

import (
	"bytes"
	"encoding/json"
	ps "fase-4-hf-orch/external/strings"
	"fase-4-hf-orch/internal/core/domain/entity/dto"
	"fmt"
	"net/http"
	"strings"
)

func (h handler) getClientByID(rw http.ResponseWriter, req *http.Request) {
	id := getID("client", req.URL.Path)

	c, err := h.app.GetClientByID(id)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to get client by ID: %v"} `, err)
		return
	}

	if c == nil {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte(`{"error": "client not found"}`))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(ps.MarshalString(c)))
}

func (h handler) getClientByCPF(rw http.ResponseWriter, req *http.Request) {
	cpf := req.URL.Query().Get("cpf")

	if len(cpf) == 0 {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(rw, `{"error": "cpf is invalid or null"} `)
		return
	}

	c, err := h.app.GetClientByCPF(cpf)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to get client by ID: %v"} `, err)
		return
	}

	if c == nil {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte(`{"error": "client not found"}`))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(ps.MarshalString(c)))
}

func (h handler) saveClient(rw http.ResponseWriter, req *http.Request) {
	var (
		buff      bytes.Buffer
		reqClient dto.RequestClient
	)

	if _, err := buff.ReadFrom(req.Body); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to read data body: %v"} `, err)
		return
	}

	if err := json.Unmarshal(buff.Bytes(), &reqClient); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to Unmarshal: %v"} `, err)
		return
	}

	c, err := h.app.SaveClient(reqClient)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to save client: %v"} `, err)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(ps.MarshalString(c)))
}

func getCpf(url string) string {
	indexCpf := strings.Index(url, "client/")

	if indexCpf == -1 {
		return ""
	}

	return strings.ReplaceAll(url[indexCpf:], "client/", "")
}
