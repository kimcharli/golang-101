package main

import (
	"net/http"
	"encoding/json"
	"sync"
	"io/ioutil"
	"fmt"
	"time"
	"strings"
)

type OneList struct {
	Name    string `json:"name"`
	Manu    string `json:"manu"`
	ID      string `json:"id"`
	Num     int    `json:"num"`
}

type listHanders struct {
	sync.Mutex
	store map[string]OneList
}

func (h *listHanders) get(w http.ResponseWriter, r *http.Request) {
	oneList := make([]OneList, len(h.store))

	h.Lock()
	i := 0
	for _, aList := range h.store {
		oneList[i] = aList
		i++
	}
	h.Unlock()

	jsonBytes, err := json.Marshal(oneList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *listHanders) getList(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.String(), "/")
	if len(parts) != 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	h.Lock()
	lister, ok := h.store[parts[2]]
	h.Unlock()
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonBytes, err := json.Marshal(lister)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

	// oneList := make([]OneList, len(h.store))

	// h.Lock()
	// i := 0
	// for _, aList := range h.store {
	// 	oneList[i] = aList
	// 	i++
	// }
	// h.Unlock()

	// jsonBytes, err := json.Marshal(oneList)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte(err.Error()))
	// }

	// w.Header().Add("content-type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write(jsonBytes)
}


func (h *listHanders) post(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(fmt.Sprintf("need content-type 'application/json', but got '%s'", ct)))
		return
	}

	var lister OneList
	err = json.Unmarshal(bodyBytes, &lister)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	lister.ID = fmt.Sprintf("%d", time.Now().UnixNano())
	h.Lock()
	h.store[lister.ID] = lister
	defer h.Unlock()
}

func (h *listHanders) oneLists(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.get(w, r)
		return
	case "POST":
		h.post(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return

	}

}

func newListHanders() *listHanders {
	return &listHanders {
		store: map[string]OneList{
			// "id1": OneList {
			// 	Name:	"F 325",
			// 	Manu:	"Hyundai",
			// 	ID:		"id1",
			// 	Num:	33,
			// },
		},
	}
}

func main() {
	listHandlers := newListHanders()
	http.HandleFunc("/list", listHandlers.oneLists)
	http.HandleFunc("/list/", listHandlers.getList)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}