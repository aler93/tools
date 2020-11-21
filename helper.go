package helpers

import (
	"fmt"
	"math"
)

const SYMBOL = "R$"

func Money32(value float32, symbol bool) string {
	if symbol {
		return fmt.Sprintf(SYMBOL+" %.2f", value)
	}

	return fmt.Sprintf("%.2f", value)
}

func Money64(money float64, symbol bool) string {
	if symbol {
		return SYMBOL + " " + F64ts(money)
	}

	return F64ts(money)
}

func Round32(value float32) float32 {
	return F64to32(math.Round(F32to64(value)))
}

func RoundFloor32(value float32) float32 {
	return F64to32(math.Floor(F32to64(value)))
}

func RoundCeil32(value float32) float32 {
	return F64to32(math.Ceil(F32to64(value)))
}

func Round64(value float64) float64 {
	return math.Round(value)
}

func RoundFloor64(value float64) float64 {
	return (value)
}

func RoundCeil64(value float64) float64 {
	return math.Ceil(value)
}
