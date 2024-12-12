package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"
)

func (handler *Handler) AdminAddBlogHandlerGet(res http.ResponseWriter, req *http.Request) {
	tmp, err := template.ParseFiles("../web/html/base.html", "../web/html/admin/addBlog.html")
	if err != nil {
		handler.app.ServerError(res, err)
	}

	err = tmp.ExecuteTemplate(res, "base", nil)
	if err != nil {
		handler.app.ServerError(res, err)
	}
}

func (handler *Handler) UploadImage(res http.ResponseWriter , req *http.Request){
	err := req.ParseMultipartForm(2048)
	if err != nil {
		handler.log.Error("Failed to parse request body")
		handler.app.ServerError(res,err)
		return
	}

	srcFile,fileHeader,err := req.FormFile("image")
	if err != nil {
		handler.log.Error("Failed to get file from body")
		handler.app.ServerError(res,err)
		return
	}


	err = os.MkdirAll("../web/static/images/",0777)
	if err !=nil {
		handler.log.Error("Create Folder Image failed")
		handler.app.ServerError(res,err)
		return
	}
	
	destinationFile,err := os.Create(fmt.Sprintf("../web/static/images/%s",fileHeader.Filename))
	if err != nil {
		handler.log.Error("Create Image failed")
		handler.app.ServerError(res,err)
		return
	}
	defer destinationFile.Close()

	_,err = io.Copy(destinationFile,srcFile)
	if err != nil {
		handler.log.Error("Copy Image failed")
		handler.app.ServerError(res,err)
		return
	}

	response := ResponseBody {
		Data : struct {
			ImagePath string `json:"imagePath"`
		}{
			ImagePath : fmt.Sprintf("/static/images/%s",fileHeader.Filename),
		},
		Code : 0,
		Message: "success",
	}

	responseText,err := json.Marshal(response)
	if err != nil {
		handler.log.Error("Parse json failed")
		handler.app.ServerError(res,err)
		return
	}

	res.Write(responseText)
}
