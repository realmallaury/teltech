package utils

import "strconv"

// Convert converts input string numbers to float.
func Convert(x, y string) (float64, float64, error) {
	xVal, err := strconv.ParseFloat(x, 64)
	if err != nil {
		return 0, 0, err
	}

	yVal, err := strconv.ParseFloat(y, 64)
	if err != nil {
		return 0, 0, err
	}

	return xVal, yVal, nil
}

// FloatToString converts float typte to string.
func FloatToString(f float64) string {
	return strconv.FormatFloat(float64(f), 'g', -1, 64)
}
