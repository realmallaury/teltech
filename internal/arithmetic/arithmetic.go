package arithmetic

import (
	"github.com/pkg/errors"
	"github.com/realmallaury/teltech/internal/utils"
)

// Arithmetic constants.
const (
	AddConst      string = "add"
	SubtractConst string = "subtract"
	MultiplyConst string = "multiply"
	DivideConst   string = "divide"
)

// Result contains data asociated with arithmetic operation.
type Result struct {
	Action string  `json:"action"`
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Answer string  `json:"answer"`
	Cached bool    `json:"cached"`
}

// Add converts x and y to float and return their addition.
func Add(x, y string) (*Result, error) {
	xVal, yVal, err := utils.Convert(x, y)
	if err != nil {
		return nil, errors.Wrapf(err, "add values: %v and %v", x, y)
	}

	return &Result{
		Action: AddConst,
		X:      xVal,
		Y:      yVal,
		Answer: utils.FloatToString(xVal + yVal),
	}, nil
}

// Subtract converts x and y to float and return their subtraction.
func Subtract(x, y string) (*Result, error) {
	xVal, yVal, err := utils.Convert(x, y)
	if err != nil {
		return nil, errors.Wrapf(err, "subtract values: %v and %v", x, y)
	}

	return &Result{
		Action: SubtractConst,
		X:      xVal,
		Y:      yVal,
		Answer: utils.FloatToString(xVal - yVal),
	}, nil
}

// Multiply converts x and y to float and return their product.
func Multiply(x, y string) (*Result, error) {
	xVal, yVal, err := utils.Convert(x, y)
	if err != nil {
		return nil, errors.Wrapf(err, "subtract values: %v and %v", x, y)
	}

	return &Result{
		Action: MultiplyConst,
		X:      xVal,
		Y:      yVal,
		Answer: utils.FloatToString(xVal * yVal),
	}, nil
}

// Divide converts x and y to float and return their division.
func Divide(x, y string) (*Result, error) {
	xVal, yVal, err := utils.Convert(x, y)
	if err != nil {
		return nil, errors.Wrapf(err, "subtract values: %v and %v", x, y)
	}

	return &Result{
		Action: DivideConst,
		X:      xVal,
		Y:      yVal,
		Answer: utils.FloatToString(xVal / yVal),
	}, nil
}
