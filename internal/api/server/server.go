package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	mid "test_postgis/internal/api/middleware"
	pdb "test_postgis/internal/gis"

	"github.com/gorilla/mux"
)

// Data for handlers
type supportHTTP struct {
	db   *pdb.ParamDB
	host string
	port string
}

// New creates new server struct
func New(db *pdb.ParamDB, rhost string, rport string) *supportHTTP {
	return &supportHTTP{
		db:   db,
		host: rhost,
		port: rport,
	}
}

// Method create a server
func (s *supportHTTP) Start() error {
	s.db.Log.Println("Server REST init.")

	ll := &mid.LogData{
		Log:  s.db.Log,
		Name: "Server REST",
	}

	router := mux.NewRouter()
	router.HandleFunc("/cities", ll.MidLogger(s.GetCity())).Methods("GET")
	router.HandleFunc("/cities", ll.MidLogger(s.PostCity())).Methods("POST")
	router.HandleFunc("/cities/{id}", ll.MidLogger(s.DelCity())).Methods("DELETE")
	router.HandleFunc("/cities/{id}", ll.MidLogger(s.PutCity())).Methods("PUT")
	router.HandleFunc("/cities/{id}", ll.MidLogger(s.PatchCity())).Methods("PATCH")
	router.HandleFunc("/cities/find/{long}/{lat}", ll.MidLogger(s.FindNearestCity())).Methods("GET")

	s.db.Log.Printf("Starting HTTP server at %s:%s\n", s.host, s.port)

	err := http.ListenAndServe(s.host+":"+s.port, router)
	if err != nil {
		s.db.Log.Printf("err=%v\n", err)
	}

	return err
}

// Handler for get request
func (s *supportHTTP) GetCity() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var bb pdb.City
		city, err := bb.SelectCity(s.db)
		if err != nil {
			w.WriteHeader(500) // Internal Server Error
			return
		}

		res, err := json.Marshal(city)
		if err != nil {
			w.WriteHeader(500) // Internal Server Error
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200) // OK
		fmt.Fprintf(w, "%s", string(res))
	}
}

// Handler for post request
func (s *supportHTTP) PostCity() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(500) // Internal Server Error
			return
		}

		b := &pdb.City{}
		err = json.Unmarshal(data, b)
		if err != nil {
			w.WriteHeader(400) // Bad Request
			return
		}

		id, err := b.InsertCity(s.db)
		if err != nil {
			w.WriteHeader(500) // Internal Server Error
			return
		}

		b.Id = id
		// Maybe select not need
		city, err := b.SelectCity(s.db)
		if err != nil {
			w.WriteHeader(400) // Bad Request
			return
		}

		res, err := json.Marshal(city)
		if err != nil {
			w.WriteHeader(500) // Internal Server Error
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200) // OK
		fmt.Fprintf(w, "%s", string(res))
	}
}

// Handler for delete request
func (s *supportHTTP) DelCity() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			w.WriteHeader(400) // Bad Request
			return
		}
		b := &pdb.City{}
		b.Id = int64(id)

		err = b.DeleteCity(s.db)
		if err != nil {
			w.WriteHeader(500) // Internal Server Error
			return
		}

		w.WriteHeader(200) // OK
		//fmt.Fprintf(w, "City %v deleted from database.\n", id)
	}
}

// Handler for put request
func (s *supportHTTP) PutCity() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(400) // Bad Request
			return
		}

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			w.WriteHeader(400) // Bad Request
			return
		}

		b := &pdb.City{}
		err = json.Unmarshal(data, b)
		if err != nil {
			w.WriteHeader(400) // Bad Request
			return
		}

		b.Id = int64(id)
		if b.Coords == "" || b.Title == "" {
			//fmt.Fprintf(w, "Some parameters not set for PUT request.\n")
			w.WriteHeader(400) // Bad Request
			return
		}

		err = b.UpdateCity(s.db)
		if err != nil {
			w.WriteHeader(500) // Internal Server Error
			return
		}

		w.WriteHeader(200) // OK
		//fmt.Fprintf(w, "Update database record id=%v\n", id)
	}
}

// Handler for patch request
func (s *supportHTTP) PatchCity() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(400) // Bad Request
			return
		}

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			w.WriteHeader(400) // Bad Request
			return
		}
		b := &pdb.City{}
		b.Id = int64(id)

		err = json.Unmarshal(data, b)
		if err != nil {
			w.WriteHeader(400) // Bad Request
			return
		}

		b.Id = int64(id)
		err = b.UpdateCity(s.db)
		if err != nil {
			w.WriteHeader(500) // Internal Server Error
			return
		}

		w.WriteHeader(200) // OK
		//fmt.Fprintf(w, "Update database record id=%v\n", id)
	}
}

// Handler for find nearest request
func (s *supportHTTP) FindNearestCity() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(400) // Bad Request
			return
		}

		params := mux.Vars(r)

		long, err := strconv.ParseFloat(params["long"], 64)
		if err != nil {
			w.WriteHeader(400) // Bad Request
			return
		}
		lat, err := strconv.ParseFloat(params["lat"], 64)
		if err != nil {
			w.WriteHeader(400) // Bad Request
			return
		}

		b := &pdb.City{}
		b.Long = long
		b.Lat = lat

		city, err := b.FindNearestCity(s.db)
		if err != nil {
			w.WriteHeader(500) // Internal Server Error
			return
		}

		res, err := json.Marshal(city)
		if err != nil {
			w.WriteHeader(500) // Internal Server Error
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200) // OK
		fmt.Fprintf(w, "%s", string(res))
	}
}
