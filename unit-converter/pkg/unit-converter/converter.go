package converter

type ConversionInput struct {
	InputValue float64
	FromUnit   string
	ToUnit     string
}

type ConversionResult struct {
	Input  *ConversionInput
	Output float64
}

type Converter interface {
	Convert(input *ConversionInput) *ConversionResult
}

type ConverterFunc func(input *ConversionInput) *ConversionResult

func (f ConverterFunc) Convert(input *ConversionInput) *ConversionResult {
	return f(input)
}

func GetConverter(converterType string) Converter {
	switch converterType {
	case "length":
		return ConverterFunc(convertLength)
	case "weight":
		return ConverterFunc(convertWeight)
	case "temperature":
		return ConverterFunc(convertTemperature)
	default:
		return nil
	}
}
