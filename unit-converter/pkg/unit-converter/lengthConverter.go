package converter

// This table save ratio to convert  an unit to mm.
var converterTable = map[string]float64{
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

type LengthConverter struct {
	ConverterName string
}

func (converter *LengthConverter) Convert(input *ConversionInput) *ConversionResult {
	outValue := input.InputValue * converterTable[input.FromUnit] / converterTable[input.ToUnit]
	return &ConversionResult{
		Input:  input,
		Output: outValue,
	}
}
