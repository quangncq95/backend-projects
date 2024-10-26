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
