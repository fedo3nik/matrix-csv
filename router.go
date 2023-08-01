package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"go.uber.org/zap"
	http "net/http"
	"path/filepath"
)

const (
	// end-points paths
	echo     = "/echo"
	invert   = "/invert"
	flatten  = "/flatten"
	sum      = "/sum"
	multiply = "/multiply"

	csvExt  = ".csv"
	fileKey = "file"
)

var (
	errFileExtension   = errors.New("invalid file extension, should be \"*.csv\"")
	errEmptyFile       = errors.New("there are no data in file")
	errMatrixNotSquare = errors.New("matrix should be square, number of rows are equal to the number of columns")
)

type Router struct {
	*http.ServeMux
	log *zap.Logger
}

func NewRouter(log *zap.Logger) *Router {
	return &Router{
		ServeMux: http.NewServeMux(),
		log:      log,
	}
}

func (rout *Router) InitRoutes() {
	rout.HandleFunc(echo, rout.Echo)
	rout.HandleFunc(invert, rout.Invert)
	rout.HandleFunc(flatten, rout.Flatten)
	rout.HandleFunc(sum, rout.Sum)
	rout.HandleFunc(multiply, rout.Multiply)
}

func (rout *Router) Echo(w http.ResponseWriter, r *http.Request) {
	matrix, err := extractData(r, fileKey)
	if err != nil {
		rout.log.Error("extracting data from .csv file failed", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := convertToMatrixString(matrix)
	rout.log.Info("Echo command called")
	w.WriteHeader(http.StatusOK)

	_, err = fmt.Fprint(w, response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (rout *Router) Invert(w http.ResponseWriter, r *http.Request) {
	matrix, err := extractData(r, fileKey)
	if err != nil {
		rout.log.Error("extracting data from .csv file failed", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	inverted := invertMatrix(matrix)
	response := convertToMatrixString(inverted)
	rout.log.Info("Invert command called")
	w.WriteHeader(http.StatusOK)

	_, err = fmt.Fprint(w, response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (rout *Router) Flatten(w http.ResponseWriter, r *http.Request) {
	matrix, err := extractData(r, fileKey)
	if err != nil {
		rout.log.Error("extracting data from .csv file failed", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := convertToPlainString(matrix)
	rout.log.Info("Flatten command called")
	w.WriteHeader(http.StatusOK)

	_, err = fmt.Fprint(w, response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (rout *Router) Sum(w http.ResponseWriter, r *http.Request) {
	matrix, err := extractData(r, fileKey)
	if err != nil {
		rout.log.Error("extracting data from .csv file failed", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := sumMatrix(matrix)
	rout.log.Info("Sum command called")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := fmt.Sprintf("%d\n", res)
	w.WriteHeader(http.StatusOK)

	_, err = fmt.Fprint(w, response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (rout *Router) Multiply(w http.ResponseWriter, r *http.Request) {
	matrix, err := extractData(r, fileKey)
	if err != nil {
		rout.log.Error("extracting data from .csv file failed", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := multiplyMatrix(matrix)
	rout.log.Info("Multiply command called")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := fmt.Sprintf("%d\n", res)
	w.WriteHeader(http.StatusOK)

	_, err = fmt.Fprint(w, response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func extractData(r *http.Request, key string) ([][]string, error) {
	file, header, err := r.FormFile(key)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = file.Close()
	}()

	extension := filepath.Ext(header.Filename)
	if extension != csvExt {
		return nil, errFileExtension
	}

	matrix, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}

	if matrix == nil {
		return nil, errEmptyFile
	}

	if !isSquare(matrix) {
		return nil, errMatrixNotSquare
	}

	return matrix, err
}
