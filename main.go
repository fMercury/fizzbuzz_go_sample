package main

import (
	"expvar"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	str1  = "fizz"
	str2  = "buzz"
	int1    = 3
	int2    = 5
	limit = 100
)

var requestCalls = expvar.NewMap("request_calls")
var requestMaxCalls = expvar.NewString("request_max_calls")
var requestMaxNumber = expvar.NewInt("request_max_number")

func parseRequest(req *http.Request) (q query, err error) {
	q.str1 = req.FormValue("str1")
	if q.str1 == "" {
		q.str1 = str1
	}
	q.str2 = req.FormValue("str2")
	if q.str2 == "" {
		q.str2 = str2
	}

	i1 := req.FormValue("int1")
	if i1 == "" {
		q.int1 = int1
	} else {
		q.int1, err = strconv.Atoi(i1)
		if err != nil {
			return q, fmt.Errorf("Unable to parse int1 value %s not an integer", i1)
		}
	}

	i2 := req.FormValue("int2")
	if i2 == "" {
		q.int2 = int2
	} else {
		q.int2, err = strconv.Atoi(req.FormValue("int2"))
		if err != nil {
			return q, fmt.Errorf("Unable to parse int2 value %s not an integer", i2)
		}
	}

	l := req.FormValue("limit")
	if l == "" {
		q.limit = limit
	} else {
		q.limit, err = strconv.Atoi(req.FormValue("limit"))
		if err != nil {
			return q, fmt.Errorf("Unable to parse limit value %s not an integer", l)
		}
	}
	return q, nil
}

func fizzBuzzHandler(w http.ResponseWriter, req *http.Request) {

	q, err := parseRequest(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	log.Printf("Serving query: %v", q)

	requestString := "int1:" + strconv.Itoa(q.int1) + " int2:" + strconv.Itoa(q.int2) + " limit:" + strconv.Itoa(q.limit) + " str1:" + q.str1 + " str2:" + q.str2   
	prs := requestCalls.Get(requestString)
	if prs == nil {
		requestCalls.Set(requestString, new(expvar.Int) )
	} else {
		requestCalls.Add(requestString, 1)
	}

	err = FizzBuzz(q, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func statistics(w http.ResponseWriter, req *http.Request) {
	fmt.Println("map:", requestCalls)
	requestMaxNumber.Set(0)
	requestCalls.Do( func(kv expvar.KeyValue) {
		fmt.Println("kv: ", kv)
		i, _ := strconv.ParseInt(kv.Value.String(), 10, 64)
		if requestMaxNumber.Value() == 0 { 
			requestMaxCalls.Set(kv.Key)
			requestMaxNumber.Set( i )
		} 
		if requestMaxNumber.Value() < i {
			requestMaxCalls.Set(kv.Key)
			requestMaxNumber.Set( i )
		}
	} )
    fmt.Fprintf(w, 	 requestMaxCalls.Value() +" "+ strconv.FormatInt(requestMaxNumber.Value(), 10)	 )
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/statistics", statistics)
	mux.HandleFunc("/fizzbuzz", fizzBuzzHandler)

	c := loadConfig()

	srv := &http.Server{
		Addr:         c.addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      mux,
	}

	log.Printf("Starting to serve on %s\n", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}


