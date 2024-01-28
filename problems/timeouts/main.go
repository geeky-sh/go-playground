package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

/*
What is done here?
- Used a API using beeceptor to simulate timeouts
- Create Timeout Handlers which make the request to the above API. If the response is not received in the specified duration, it times out.
Else it returns proper response.
- Created the same function as above but put all in one function and took timeout from query
*/

type TimeoutHandler struct {
	dt   time.Duration
	hdlr func(w http.ResponseWriter, r *http.Request)
}

func (h *TimeoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx, cf := context.WithTimeout(r.Context(), h.dt)
	defer cf()

	done := make(chan struct{})

	go func() {
		h.hdlr(w, r)
		done <- struct{}{}
	}()

	select {
	case <-done:
		return
	case <-ctx.Done():
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusRequestTimeout)
		en := json.NewEncoder(w)
		en.Encode(map[string]string{"err": "request timeout"})
	}
}

func Timeout(dt time.Duration, hdlr func(w http.ResponseWriter, r *http.Request)) TimeoutHandler {
	return TimeoutHandler{dt, hdlr}
}

func SimpleTimeout(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	dt, err := time.ParseDuration(q.Get("timeout"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cl := context.WithTimeout(r.Context(), dt)
	defer cl()

	done := make(chan struct{})

	go func() {
		res, _ := http.Get("https://asdfg.free.beeceptor.com/delay")
		w.Header().Set("Content-Type", "application/json")
		x, _ := io.ReadAll(res.Body)
		w.Write(x)
		close(done)
	}()

	select {
	case <-done:
		return
	case <-ctx.Done():
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusRequestTimeout)
		en := json.NewEncoder(w)
		en.Encode(map[string]string{"err": "request timeout"})
	}

}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		en := json.NewEncoder(w)
		en.Encode(map[string]string{"success": "1"})
	})

	delayFunc := func(w http.ResponseWriter, r *http.Request) {
		res, _ := http.Get("https://asdfg.free.beeceptor.com/delay")
		w.Header().Set("Content-Type", "application/json")
		x, _ := io.ReadAll(res.Body)
		w.Write(x)
	}

	timeoutFunc := Timeout(1*time.Second, delayFunc)

	http.HandleFunc("/timeout", timeoutFunc.ServeHTTP)
	http.HandleFunc("/delay", SimpleTimeout)

	if err := http.ListenAndServe(":8888", nil); err != nil {
		panic(err)
	}
}
