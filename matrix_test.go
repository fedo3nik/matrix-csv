package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isSquare(t *testing.T) {
	squareMatrix := [][]string{{"1", "2"}, {"3", "4"}}
	notSquareMatrix := [][]string{{"1", "2"}, {"3", "4", "5"}}

	tt := []struct {
		name           string
		providedMatrix [][]string
		expectedResult bool
	}{
		{
			name:           "fail: matrix is not square",
			providedMatrix: notSquareMatrix,
			expectedResult: false,
		},
		{
			name:           "success: matrix is square",
			providedMatrix: squareMatrix,
			expectedResult: true,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			res := isSquare(tc.providedMatrix)
			assert.Equal(t, tc.expectedResult, res)
		})
	}
}

func Test_convertToMatrixString(t *testing.T) {
	matrix := [][]string{{"1", "2"}, {"3", "4"}}

	tt := []struct {
		name           string
		providedMatrix [][]string
		expectedResult string
	}{
		{
			name:           "success: converted to matrix view string",
			providedMatrix: matrix,
			expectedResult: "1,2\n3,4\n",
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			res := convertToMatrixString(tc.providedMatrix)
			assert.Equal(t, tc.expectedResult, res)
		})
	}
}

func Test_convertToPlainString(t *testing.T) {
	matrix := [][]string{{"1", "2"}, {"3", "4"}}

	tt := []struct {
		name           string
		providedMatrix [][]string
		expectedResult string
	}{
		{
			name:           "success: converted to plain view string",
			providedMatrix: matrix,
			expectedResult: "1,2,3,4\n",
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			res := convertToPlainString(tc.providedMatrix)
			assert.Equal(t, tc.expectedResult, res)
		})
	}
}

func Test_invertMatrix(t *testing.T) {
	matrix := [][]string{{"1", "2"}, {"3", "4"}}
	inverted := [][]string{{"1", "3"}, {"2", "4"}}

	tt := []struct {
		name           string
		providedMatrix [][]string
		expectedResult [][]string
	}{
		{
			name:           "success: replaced rows with columns",
			providedMatrix: matrix,
			expectedResult: inverted,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			res := invertMatrix(tc.providedMatrix)
			assert.Equal(t, tc.expectedResult, res)
		})
	}
}

func Test_matrixToInt(t *testing.T) {
	validMatrix := [][]string{{"1", "2"}, {"3", "4"}}
	invalidMatrix := [][]string{{"1", "b"}, {"3", "4"}}

	tt := []struct {
		name           string
		providedMatrix [][]string
		expectedResult [][]int
		expectedErr    error
	}{
		{
			name:           "fail: matrix with non-int values",
			providedMatrix: invalidMatrix,
			expectedResult: nil,
			expectedErr:    errNotIntValue,
		},
		{
			name:           "success: valid matrix",
			providedMatrix: validMatrix,
			expectedResult: [][]int{{1, 2}, {3, 4}},
			expectedErr:    nil,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			res, err := matrixToInt(tc.providedMatrix)
			assert.Equal(t, tc.expectedResult, res)
			assert.ErrorIs(t, err, tc.expectedErr)
		})
	}
}

func Test_sumMatrix(t *testing.T) {
	validMatrix := [][]string{{"1", "2"}, {"3", "4"}}
	invalidMatrix := [][]string{{"1", "b"}, {"3", "4"}}

	tt := []struct {
		name           string
		providedMatrix [][]string
		expectedResult int
		expectedErr    error
	}{
		{
			name:           "fail: matrix with non-int values",
			providedMatrix: invalidMatrix,
			expectedResult: 0,
			expectedErr:    errNotIntValue,
		},
		{
			name:           "success: calculated sum of elements",
			providedMatrix: validMatrix,
			expectedResult: 10,
			expectedErr:    nil,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			res, err := sumMatrix(tc.providedMatrix)
			assert.Equal(t, tc.expectedResult, res)
			assert.ErrorIs(t, err, tc.expectedErr)
		})
	}
}

func Test_multiplyMatrix(t *testing.T) {
	validMatrix := [][]string{{"1", "2"}, {"3", "4"}}
	invalidMatrix := [][]string{{"1", "b"}, {"3", "4"}}

	tt := []struct {
		name           string
		providedMatrix [][]string
		expectedResult int
		expectedErr    error
	}{
		{
			name:           "fail: matrix with non-int values",
			providedMatrix: invalidMatrix,
			expectedResult: 0,
			expectedErr:    errNotIntValue,
		},
		{
			name:           "success: calculated multiplying result of elements",
			providedMatrix: validMatrix,
			expectedResult: 24,
			expectedErr:    nil,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			res, err := multiplyMatrix(tc.providedMatrix)
			assert.Equal(t, tc.expectedResult, res)
			assert.ErrorIs(t, err, tc.expectedErr)
		})
	}
}
