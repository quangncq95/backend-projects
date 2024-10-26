package converter

// This table save ratio to convert  an unit to mg.
var weightConverterTable = map[string]float64{
	"mg": 1,
	"g":  1000,
	"kg": 1000000,
	"oz": 28350,
	"lb": 453600,
}

func convertWeight(input *ConversionInput) *ConversionResult {
	outValue := input.InputValue * weightConverterTable[input.FromUnit] / weightConverterTable[input.ToUnit]
	return &ConversionResult{
		Input:  input,
		Output: outValue,
	}
}
