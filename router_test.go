package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

const (
	validPath     = "./data/matrix.csv"
	txtPath       = "./data/text.txt"
	emptyPath     = "./data/empty.csv"
	notSquarePath = "./data/notSquare.csv"

	testURL = "http://localhost:3000"
)

func setupRouter() (*Router, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	router := NewRouter(logger)
	router.InitRoutes()
	return router, nil
}

func createReq(filePath string, url string) (*http.Request, *multipart.Writer, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	formFile, err := writer.CreateFormFile(fileKey, filePath)
	if err != nil {
		return nil, nil, err
	}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}

	_, err = io.Copy(formFile, file)
	if err != nil {
		return nil, nil, err
	}

	// INFO: here we close files to avoid EOF error
	err = file.Close()
	if err != nil {
		return nil, nil, err
	}
	err = writer.Close()
	if err != nil {
		return nil, nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, nil, err
	}

	return req, writer, nil
}

func TestRouter_Echo(t *testing.T) {
	validBody := "1,2,3\n4,5,6\n7,8,9\n"
	router, err := setupRouter()
	assert.NoError(t, err)

	successReq, successW, err := createReq(validPath, testURL)
	assert.NoError(t, err)
	successReq.Header.Set("Content-Type", successW.FormDataContentType())

	txtReq, txtW, err := createReq(txtPath, testURL)
	assert.NoError(t, err)
	txtReq.Header.Set("Content-Type", txtW.FormDataContentType())

	tt := []struct {
		name         string
		providedReq  *http.Request
		expectedBody string
		expectedCode int
	}{
		{
			name:         "fail: extract data - BadRequest",
			providedReq:  txtReq,
			expectedBody: "invalid file extension, should be \"*.csv\"\n",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "success: valid data provided",
			providedReq:  successReq,
			expectedBody: validBody,
			expectedCode: http.StatusOK,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			router.Echo(w, tc.providedReq)
			assert.NoError(t, err)

			assert.Equal(t, tc.expectedCode, w.Result().StatusCode)
			assert.Equal(t, tc.expectedBody, w.Body.String())
		})
	}
}

func TestRouter_Invert(t *testing.T) {
	validBody := "1,4,7\n2,5,8\n3,6,9\n"
	router, err := setupRouter()
	assert.NoError(t, err)

	successReq, successW, err := createReq(validPath, testURL)
	assert.NoError(t, err)
	successReq.Header.Set("Content-Type", successW.FormDataContentType())

	txtReq, txtW, err := createReq(txtPath, testURL)
	assert.NoError(t, err)
	txtReq.Header.Set("Content-Type", txtW.FormDataContentType())

	tt := []struct {
		name         string
		providedReq  *http.Request
		expectedBody string
		expectedCode int
	}{
		{
			name:         "fail: extract data - BadRequest",
			providedReq:  txtReq,
			expectedBody: "invalid file extension, should be \"*.csv\"\n",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "success: inverted successful",
			providedReq:  successReq,
			expectedBody: validBody,
			expectedCode: http.StatusOK,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			router.Invert(w, tc.providedReq)
			assert.NoError(t, err)

			assert.Equal(t, tc.expectedCode, w.Result().StatusCode)
			assert.Equal(t, tc.expectedBody, w.Body.String())
		})
	}
}

func TestRouter_Flatten(t *testing.T) {
	validBody := "1,2,3,4,5,6,7,8,9\n"
	router, err := setupRouter()
	assert.NoError(t, err)

	successReq, successW, err := createReq(validPath, testURL)
	assert.NoError(t, err)
	successReq.Header.Set("Content-Type", successW.FormDataContentType())

	txtReq, txtW, err := createReq(txtPath, testURL)
	assert.NoError(t, err)
	txtReq.Header.Set("Content-Type", txtW.FormDataContentType())

	tt := []struct {
		name         string
		providedReq  *http.Request
		expectedBody string
		expectedCode int
	}{
		{
			name:         "fail: extract data - BadRequest",
			providedReq:  txtReq,
			expectedBody: "invalid file extension, should be \"*.csv\"\n",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "success: flatten string matrix view",
			providedReq:  successReq,
			expectedBody: validBody,
			expectedCode: http.StatusOK,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			router.Flatten(w, tc.providedReq)
			assert.NoError(t, err)

			assert.Equal(t, tc.expectedCode, w.Result().StatusCode)
			assert.Equal(t, tc.expectedBody, w.Body.String())
		})
	}
}

func TestRouter_Sum(t *testing.T) {
	router, err := setupRouter()
	assert.NoError(t, err)

	successReq, successW, err := createReq(validPath, testURL)
	assert.NoError(t, err)
	successReq.Header.Set("Content-Type", successW.FormDataContentType())

	txtReq, txtW, err := createReq(txtPath, testURL)
	assert.NoError(t, err)
	txtReq.Header.Set("Content-Type", txtW.FormDataContentType())

	tt := []struct {
		name         string
		providedReq  *http.Request
		expectedBody string
		expectedCode int
	}{
		{
			name:         "fail: extract data - BadRequest",
			providedReq:  txtReq,
			expectedBody: "invalid file extension, should be \"*.csv\"\n",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "success: sum calculated",
			providedReq:  successReq,
			expectedBody: "45\n",
			expectedCode: http.StatusOK,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			router.Sum(w, tc.providedReq)
			assert.NoError(t, err)

			assert.Equal(t, tc.expectedCode, w.Result().StatusCode)
			assert.Equal(t, tc.expectedBody, w.Body.String())
		})
	}
}

func TestRouter_Multiply(t *testing.T) {
	router, err := setupRouter()
	assert.NoError(t, err)

	successReq, successW, err := createReq(validPath, testURL)
	assert.NoError(t, err)
	successReq.Header.Set("Content-Type", successW.FormDataContentType())

	txtReq, txtW, err := createReq(txtPath, testURL)
	assert.NoError(t, err)
	txtReq.Header.Set("Content-Type", txtW.FormDataContentType())

	tt := []struct {
		name         string
		providedReq  *http.Request
		expectedBody string
		expectedCode int
	}{
		{
			name:         "fail: extract data - BadRequest",
			providedReq:  txtReq,
			expectedBody: "invalid file extension, should be \"*.csv\"\n",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "success: sum calculated",
			providedReq:  successReq,
			expectedBody: "362880\n",
			expectedCode: http.StatusOK,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			router.Multiply(w, tc.providedReq)
			assert.NoError(t, err)

			assert.Equal(t, tc.expectedCode, w.Result().StatusCode)
			assert.Equal(t, tc.expectedBody, w.Body.String())
		})
	}
}

func Test_extractData(t *testing.T) {
	successReq, successW, err := createReq(validPath, testURL)
	assert.NoError(t, err)
	successReq.Header.Set("Content-Type", successW.FormDataContentType())

	txtReq, txtW, err := createReq(txtPath, testURL)
	assert.NoError(t, err)
	txtReq.Header.Set("Content-Type", txtW.FormDataContentType())

	emptyReq, emptyW, err := createReq(emptyPath, testURL)
	assert.NoError(t, err)
	emptyReq.Header.Set("Content-Type", emptyW.FormDataContentType())

	notSquareReq, notSquareW, err := createReq(notSquarePath, testURL)
	assert.NoError(t, err)
	notSquareReq.Header.Set("Content-Type", notSquareW.FormDataContentType())

	tt := []struct {
		name           string
		providedReq    *http.Request
		expectedResult [][]string
		expectedErr    error
	}{

		{
			name:           "fail: wrong file extension",
			providedReq:    txtReq,
			expectedResult: nil,
			expectedErr:    errFileExtension,
		},
		{
			name:           "fail: empty .csv file",
			providedReq:    emptyReq,
			expectedResult: nil,
			expectedErr:    errEmptyFile,
		},
		{
			name:           "fail: not square matrix",
			providedReq:    notSquareReq,
			expectedResult: nil,
			expectedErr:    errMatrixNotSquare,
		},
		{
			name:           "success: valid file provided",
			providedReq:    successReq,
			expectedResult: [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}},
			expectedErr:    nil,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			res, err := extractData(tc.providedReq, fileKey)

			assert.Equal(t, tc.expectedResult, res)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
