package converter

import "fmt"

// This table save ratio to convert  an unit to mm.
var temperatureConverterTable = map[string]func(input float64) float64{
	"c-f": func(input float64) float64 {
		return (input * float64(9) / float64(5)) + 32
	},
	"f-c": func(input float64) float64 {
		return (input - 32) * float64(5) / float64(9)
	},
	"c-k": func(input float64) float64 {
		return input + 273.15
	},
	"k-c": func(input float64) float64 {
		return input - 273.15
	},
	"f-k": func(input float64) float64 {
		return (input-32)*float64(5)/float64(9) + 273.15
	},
	"k-f": func(input float64) float64 {
		return ((input - 273.15) * float64(9) / float64(5)) + 32
	},
}

func convertTemperature(input *ConversionInput) *ConversionResult {
	unitConvertKey := fmt.Sprintf("%v-%v", input.FromUnit, input.ToUnit)
	fn := temperatureConverterTable[unitConvertKey]
	output := fn(input.InputValue)
	return &ConversionResult{
		Input:  input,
		Output: output,
	}
}
