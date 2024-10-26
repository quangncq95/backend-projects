package handler

import (
	"html/template"
	converter "ncquang/unit-converter/pkg/unit-converter"
	"net/http"
	"strconv"
)

func LengthHandlerPost(res http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("../web/template/base.html", "../web/template/result.html")
	if err != nil {
		http.Error(res, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = req.ParseForm()
	if err != nil {
		http.Error(res, "Internal server error", http.StatusInternalServerError)
		return
	}

	lengthConvert := &converter.LengthConverter{
		ConverterName: "length",
	}

	inputValue, err := strconv.ParseFloat(req.PostForm.Get("inputValue"), 64)
	if err != nil {
		http.Error(res, "Bad request", http.StatusBadRequest)
		return
	}

	input := &converter.ConversionInput{
		InputValue: inputValue,
		FromUnit:   req.PostForm.Get("fromUnit"),
		ToUnit:     req.PostForm.Get("toUnit"),
	}
	output := lengthConvert.Convert(input)

	err = tmpl.ExecuteTemplate(res, "base", output)
	if err != nil {
		http.Error(res, "Internal server error", http.StatusInternalServerError)
		return
	}

}
