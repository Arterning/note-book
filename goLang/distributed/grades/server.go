package grades

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func RegisterHandlers() {

	handler := new(studentsHandler)
	http.Handle("/student", handler)
	http.Handle("/student/", handler)
}

type studentsHandler struct{}

func (sh studentsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	pathSegments := strings.Split(r.URL.Path, "/")
	switch len(pathSegments) {
	case 2:
		sh.getAll(w, r)
	case 3:
		id, err := strconv.Atoi(pathSegments[2])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		sh.getOne(w, r, id)
	case 4:
		id, err := strconv.Atoi(pathSegments[2])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		sh.addGrade(w, r, id)
	default:
		w.WriteHeader(http.StatusNotFound)
	}

}

func (sh studentsHandler) getAll(w http.ResponseWriter, r *http.Request) {

	studentsMutex.Lock()
	defer studentsMutex.Unlock()
	data, err := sh.toJSON(students)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func (sh studentsHandler) toJSON(obj interface{}) ([]byte, error) {

	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	err := enc.Encode(obj)
	if err != nil {
		return nil, fmt.Errorf("failed to json")
	}
	return b.Bytes(), nil
}

func (sh studentsHandler) getOne(w http.ResponseWriter, r *http.Request, id int) {
	studentsMutex.Lock()
	defer studentsMutex.Unlock()
	student, err := students.GetById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	data, err := sh.toJSON(student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func (sh studentsHandler) addGrade(w http.ResponseWriter, r *http.Request, id int) {
	studentsMutex.Lock()
	defer studentsMutex.Unlock()
	student, err := students.GetById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	var g Grade
	dec := json.NewDecoder(r.Body)
	err = dec.Decode(&g)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	student.Grades = append(student.Grades, g)
	w.WriteHeader(http.StatusCreated)
	data, err := sh.toJSON(g)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}
