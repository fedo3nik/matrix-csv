package main

import (
	"errors"
	"strconv"
	"strings"
)

var (
	errNotIntValue = errors.New("only Integer value is allowed")
)

// INFO: check is matrix square. Compare number of rows and number of notes in each row (columns).
func isSquare(matrix [][]string) bool {
	rowsNumber := len(matrix)
	for _, row := range matrix {
		if len(row) != rowsNumber {
			return false
		}
	}
	return true
}

// INFO: converts data from .csv into string format. Use builder here because of reduced time spent in runtime
// and memory optimality.
func convertToMatrixString(matrix [][]string) string {
	var b strings.Builder
	for i := range matrix {
		row := matrix[i]
		b.WriteString(strings.Join(row, ","))
		b.WriteString("\n")
	}

	return b.String()
}

func convertToPlainString(matrix [][]string) string {
	var rows []string
	for i := range matrix {
		row := matrix[i]
		rows = append(rows, strings.Join(row, ","))
	}
	return strings.Join(rows, ",") + "\n"
}

// INFO: inverting matrix by replacing rows with columns.
func invertMatrix(matrix [][]string) [][]string {
	rowsNumber := len(matrix)
	columnsNumber := len(matrix[0])
	inverted := make([][]string, columnsNumber)
	for i := range inverted {
		inverted[i] = make([]string, rowsNumber)
	}

	for i := 0; i < rowsNumber; i++ {
		for j := 0; j < columnsNumber; j++ {
			inverted[j][i] = matrix[i][j]
		}
	}

	return inverted
}

// INFO: converts each element from string to int. Returns errNotIntValue in case of wrong data type.
func matrixToInt(matrix [][]string) ([][]int, error) {
	rowsNumber := len(matrix)
	columnsNumber := len(matrix[0])
	res := make([][]int, columnsNumber)
	for i := range res {
		res[i] = make([]int, rowsNumber)
	}

	for i := 0; i < rowsNumber; i++ {
		for j := 0; j < columnsNumber; j++ {
			elem, err := strconv.Atoi(matrix[i][j])
			if err != nil {
				return nil, errNotIntValue
			}
			res[i][j] = elem
		}
	}

	return res, nil
}

func sumMatrix(matrix [][]string) (int, error) {
	var sum int
	intMatrix, err := matrixToInt(matrix)
	if err != nil {
		return 0, err
	}

	for i := range intMatrix {
		for j := range intMatrix[i] {
			sum += intMatrix[i][j]
		}
	}

	return sum, nil
}

func multiplyMatrix(matrix [][]string) (int, error) {
	result := 1
	intMatrix, err := matrixToInt(matrix)
	if err != nil {
		return 0, err
	}

	for i := range intMatrix {
		for j := range intMatrix[i] {
			result *= intMatrix[i][j]
		}
	}

	return result, nil
}
