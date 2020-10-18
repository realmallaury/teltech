package arithmetic

import (
	"fmt"
	"math"
	"testing"

	"github.com/realmallaury/teltech/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	assert := assert.New(t)

	tables := []struct {
		x       string
		y       string
		xVal    float64
		yVal    float64
		errFlag bool
	}{
		{"1", "1", float64(1), float64(1), false},
		{"1.5", "1.5", float64(1.5), float64(1.5), false},
		{"1.55555555", "1.55555555", float64(1.55555555), float64(1.55555555), false},
		{"1.", "1..55555555", 0, 0, true},
		{
			fmt.Sprintf("%f", math.MaxFloat64),
			fmt.Sprintf("%f", -math.MaxFloat64),
			math.MaxFloat64,
			-math.MaxFloat64,
			false,
		},
	}

	for _, table := range tables {
		xRes, yRes, err := utils.Convert(table.x, table.y)

		assert.Equal(table.xVal, xRes, "Values should be the same")
		assert.Equal(table.yVal, yRes, "Values should be the same")

		if table.errFlag {
			assert.Error(err, "Should be error")
		} else {
			assert.NoError(err, "Error should be nil")
		}
	}
}

func TestAdd(t *testing.T) {
	assert := assert.New(t)

	tables := []struct {
		x       string
		y       string
		res     *Result
		errFlag bool
	}{
		{"1", "1", &Result{Action: AddConst, X: 1, Y: 1, Answer: "2", Cached: false}, false},
		{"1", "-1", &Result{Action: AddConst, X: 1, Y: -1, Answer: "0", Cached: false}, false},
		{"0.1", "0.1", &Result{Action: AddConst, X: 0.1, Y: 0.1, Answer: "0.2", Cached: false}, false},
		{"0.1", "0..1", nil, true},
		{
			fmt.Sprintf("%f", math.MaxFloat64),
			fmt.Sprintf("%f", math.MaxFloat64),
			&Result{
				Action: AddConst,
				X:      1.7976931348623157e+308,
				Y:      1.7976931348623157e+308,
				Answer: "+Inf",
				Cached: false,
			},
			false,
		},
	}

	for _, table := range tables {
		res, err := Add(table.x, table.y)

		assert.Equal(table.res, res, "Values should be the same")

		if table.errFlag {
			assert.Error(err, "Should be error")
		} else {
			assert.NoError(err, "Error should be nil")
		}
	}
}

func TestSubtract(t *testing.T) {
	assert := assert.New(t)

	tables := []struct {
		x       string
		y       string
		res     *Result
		errFlag bool
	}{
		{"1", "1", &Result{Action: SubtractConst, X: 1, Y: 1, Answer: "0", Cached: false}, false},
		{"1", "-1", &Result{Action: SubtractConst, X: 1, Y: -1, Answer: "2", Cached: false}, false},
		{"0.5", "0.1", &Result{Action: SubtractConst, X: 0.5, Y: 0.1, Answer: "0.4", Cached: false}, false},
		{"0.1", "0..1", nil, true},
		{
			fmt.Sprintf("%f", -math.MaxFloat64),
			fmt.Sprintf("%f", math.MaxFloat64),
			&Result{
				Action: SubtractConst,
				X:      -1.7976931348623157e+308,
				Y:      1.7976931348623157e+308,
				Answer: "-Inf",
				Cached: false,
			},
			false,
		},
	}

	for _, table := range tables {
		res, err := Subtract(table.x, table.y)

		assert.Equal(table.res, res, "Values should be the same")

		if table.errFlag {
			assert.Error(err, "Should be error")
		} else {
			assert.NoError(err, "Error should be nil")
		}
	}
}

func TestMultiply(t *testing.T) {
	assert := assert.New(t)

	tables := []struct {
		x       string
		y       string
		res     *Result
		errFlag bool
	}{
		{"1", "1", &Result{Action: MultiplyConst, X: 1, Y: 1, Answer: "1", Cached: false}, false},
		{"1", "-1", &Result{Action: MultiplyConst, X: 1, Y: -1, Answer: "-1", Cached: false}, false},
		{"0.5", "0.1", &Result{Action: MultiplyConst, X: 0.5, Y: 0.1, Answer: "0.05", Cached: false}, false},
		{"0.1", "0..1", nil, true},
		{
			fmt.Sprintf("%f", math.MaxFloat64),
			fmt.Sprintf("%f", math.MaxFloat64),
			&Result{
				Action: MultiplyConst,
				X:      1.7976931348623157e+308,
				Y:      1.7976931348623157e+308,
				Answer: "+Inf",
				Cached: false,
			},
			false,
		},
	}

	for _, table := range tables {
		res, err := Multiply(table.x, table.y)

		assert.Equal(table.res, res, "Values should be the same")

		if table.errFlag {
			assert.Error(err, "Should be error")
		} else {
			assert.NoError(err, "Error should be nil")
		}
	}
}

func TestDivide(t *testing.T) {
	assert := assert.New(t)

	tables := []struct {
		x       string
		y       string
		res     *Result
		errFlag bool
	}{
		{"1", "1", &Result{Action: DivideConst, X: 1, Y: 1, Answer: "1", Cached: false}, false},
		{"1", "-1", &Result{Action: DivideConst, X: 1, Y: -1, Answer: "-1", Cached: false}, false},
		{"0.5", "0.1", &Result{Action: DivideConst, X: 0.5, Y: 0.1, Answer: "5", Cached: false}, false},
		{"0.1", "0..1", nil, true},
		{
			fmt.Sprintf("%f", math.MaxFloat64),
			fmt.Sprintf("%f", 1.797e+300),
			&Result{
				Action: DivideConst,
				X:      1.7976931348623157e+308,
				Y:      1.797e+300,
				Answer: "1.0003857177864861e+08",
				Cached: false,
			},
			false,
		},
		{
			fmt.Sprintf("%f", -math.MaxFloat64),
			fmt.Sprintf("%f", 5.435e+30),
			&Result{
				Action: DivideConst,
				X:      -1.7976931348623157e+308,
				Y:      5.435e+30,
				Answer: "-3.3076230632241315e+277",
				Cached: false,
			},
			false,
		},
		{
			"5", "0", &Result{Action: DivideConst, X: 5, Y: 0, Answer: "+Inf", Cached: false}, false,
		},
	}

	for _, table := range tables {
		res, err := Divide(table.x, table.y)

		assert.Equal(table.res, res, "Values should be the same")

		if table.errFlag {
			assert.Error(err, "Should be error")
		} else {
			assert.NoError(err, "Error should be nil")
		}
	}
}
