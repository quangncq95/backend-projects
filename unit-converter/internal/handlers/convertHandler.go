package handler

import (
	"html/template"
	converterPkg "ncquang/unit-converter/pkg/unit-converter"
	"net/http"
	"strconv"
)

func ConvertHandlerPost(res http.ResponseWriter, req *http.Request) {
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

	inputValue, err := strconv.ParseFloat(req.PostForm.Get("inputValue"), 64)
	if err != nil {
		http.Error(res, "Bad request", http.StatusBadRequest)
		return
	}

	converter := converterPkg.GetConverter(req.PostForm.Get("type"))

	input := &converterPkg.ConversionInput{
		InputValue: inputValue,
		FromUnit:   req.PostForm.Get("fromUnit"),
		ToUnit:     req.PostForm.Get("toUnit"),
	}
	output := converter.Convert(input)

	err = tmpl.ExecuteTemplate(res, "base", output)
	if err != nil {
		http.Error(res, "Internal server error", http.StatusInternalServerError)
		return
	}

}
