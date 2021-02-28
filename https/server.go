package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var (
	cache = make([]string, 0)
	mux   sync.Mutex
)

type Resp struct {
	Response []bool `json: "response"`
}

func StringSliceHandler(w http.ResponseWriter, r *http.Request) {
	var resp Resp
	resp.Response = make([]bool, 0)
	var Req struct {
		Request []string `json:"request"`
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "%d", 500)
	}

	err = json.Unmarshal(body, &Req)
	if err != nil {
		fmt.Fprintf(w, "%d", 500)
	}

	var flag = true
	mux.Lock()
	for _, v := range Req.Request {
		for _, c := range cache {
			if c == v {
				resp.Response = append(resp.Response, true)
				flag = false
				break
			}
		}
		if flag {
			resp.Response = append(resp.Response, false)
			cache = append(cache, v)
		} else {
			flag = true
		}
	}
	mux.Unlock()

	response, err := json.Marshal(resp)
	if err != nil {
		fmt.Fprintf(w, "%d", 500)
	}
	w.Write(response)
}
