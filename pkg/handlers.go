package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"
)

func getRequestString(req generateRequest) string {
	return fmt.Sprintf("%s:%d", req.Type, req.Length)
}
func setJSONContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func Generate(w http.ResponseWriter, r *http.Request) {
	if http.MethodPost != r.Method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method not allowed"))
		return
	}
	req := generateRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	reqStr := getRequestString(req)
	log.Println(reqStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad Request"))
		return
	}
	if req.Length < 0 {
		w.WriteHeader(http.StatusLengthRequired)
		w.Write([]byte("411 - Length required"))
		return
	}
	if req.Length > math.MaxInt {
		w.WriteHeader(http.StatusRequestEntityTooLarge)
		w.Write([]byte("413 - Request entity too large"))
		return
	}
	// Идемпотентность на одну минуту
	if reqInfo, ok := requestStore[reqStr]; ok && time.Since(reqInfo.Time) < time.Minute {
		log.Println("Too many requests")
		reqInfo.Time = time.Now()
		respJSON, _ := json.Marshal(responseRequest{reqInfo.ID, store[reqInfo.ID]})
		setJSONContentType(w)
		w.Write(respJSON)
		return
	}
	mu.Lock()
	defer mu.Unlock()

	result, err := generateRandomValue(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	store[nextID] = result
	resp := responseRequest{nextID, result}
	requestStore[reqStr] = RequestInfo{nextID, time.Now()}
	defer func() { nextID += 1 }()

	respJSON, _ := json.Marshal(resp)
	setJSONContentType(w)
	w.Write(respJSON)
}

func Retrieve(w http.ResponseWriter, r *http.Request) {
	if http.MethodGet != r.Method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method not allowed"))
		return
	}
	vars := mux.Vars(r)["id"]
	id, err := strconv.Atoi(vars)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad request"))
		return
	}
	content, ok := store[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Id not found"))
		return
	}
	resp := responseRequest{id, content}
	respJSON, _ := json.Marshal(resp)
	setJSONContentType(w)
	w.Write(respJSON)
}
