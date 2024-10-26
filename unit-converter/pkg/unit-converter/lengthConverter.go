package converter

// This table save ratio to convert  an unit to mm.
var lengthConverterTable = map[string]float64{
	"mm":  1,
	"cm":  10,
	"dm":  100,
	"m":   1000,
	"km":  1000000,
	"in":  25.4,
	"ft":  304.8,
	"yds": 914.4,
	"mi":  1609344,
}

func convertLength(input *ConversionInput) *ConversionResult {
	outValue := input.InputValue * lengthConverterTable[input.FromUnit] / lengthConverterTable[input.ToUnit]
	return &ConversionResult{
		Input:  input,
		Output: outValue,
	}
}
