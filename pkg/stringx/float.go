package stringx

import (
	"fmt"
	"math"
)

// FloatFormatByDecimalPlace 根据指定小数位数保留小数，当无小数时不返回小数位
func FloatFormatByDecimalPlace(value float64, decimalPlaces int) string {
	multiplier := math.Pow(10, float64(decimalPlaces))
	roundedAmount := math.Round(value*multiplier) / multiplier
	formatSpecifier := fmt.Sprintf("%%.%df", decimalPlaces)
	if decimalPlaces == 0 {
		formatSpecifier = "%.0f"
	}
	if roundedAmount == float64(int(roundedAmount)) {
		return fmt.Sprintf("%.0f", roundedAmount)
	}
	return fmt.Sprintf(formatSpecifier, roundedAmount)
}
