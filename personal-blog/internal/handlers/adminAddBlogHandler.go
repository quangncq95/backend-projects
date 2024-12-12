package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/gosimple/slug"
)

type Article struct {
	Title string `json:"title"`
	PublishDate string `json:"publishDate"`
	ContentPath string `json:"contentPath"`
	BannerImg string `json:"bannerImg"`
}

type ListArticle []Article

func (handler *Handler) AdminAddBlogHandlerGet(res http.ResponseWriter, req *http.Request) {
	tmp, err := template.ParseFiles("../web/html/base.html", "../web/html/admin/addBlog.html")
	if err != nil {
		handler.app.ServerError(res, err)
		return
	}

	err = tmp.ExecuteTemplate(res, "base", nil)
	if err != nil {
		handler.app.ServerError(res, err)
		return
	}
}

func (handler *Handler) AdminAddBlogHandlerPost(res http.ResponseWriter, req *http.Request) {
	err:=req.ParseMultipartForm(10 << 20)
	if err !=nil {
		handler.log.Error("Failed to parse request body")
		handler.app.ServerError(res,err)
		return
	}

	err = os.MkdirAll("data/contents",0777)
	if err != nil {
		handler.log.Error("Failed to create data directory")
		handler.app.ServerError(res,err)
		return
	}

	fileContentPath := fmt.Sprintf("data/contents/%s.txt",slug.Make(req.PostForm.Get("title")))

	fileContent,err := os.Create(fileContentPath)
	if err != nil {
		handler.log.Error("Failed to create file content article")
		handler.app.ServerError(res,err)
		return
	}

	defer fileContent.Close()

	_,err = io.WriteString(fileContent,req.PostForm.Get("content"))
	if err != nil {
		handler.log.Error("Failed to write file content")
		handler.app.ServerError(res,err)
		return
	}

	fileListContent,err := os.ReadFile("data/list.json")
	if err != nil && os.IsNotExist(err) {
		file,_ := os.Create("data/list.json")
		file.Close()
	}

	fileList,err := os.OpenFile("data/list.json",os.O_WRONLY,0777)
	if err !=nil {
		handler.log.Error("Failed to create file list article")
		handler.app.ServerError(res,err)
		return
	}

	defer fileList.Close()

	var listArticle []Article 
	if len(fileListContent) > 0 {
		err = json.Unmarshal(fileListContent,&listArticle)
		if err != nil {
			handler.log.Error("Failed to parse file content")
			handler.app.ServerError(res,err)
			return
		}
	}

	newArticle := Article{
		Title: req.PostForm.Get("title"),
		BannerImg: req.PostForm.Get("bannerImg"),
		ContentPath: fileContentPath,
		PublishDate: time.Now().Format(time.RFC3339),
	}


	listArticle = append(listArticle,newArticle)

	listArticleContent,err := json.Marshal(listArticle)

	if err != nil {
		handler.log.Error("Failed to parse file content")
		handler.app.ServerError(res,err)
		return
	}

	_,err = fileList.WriteString(string(listArticleContent)) 
	if err != nil {
		handler.log.Error("Failed to write to file list")
		handler.app.ServerError(res,err)
		return
	}

	res.Write([]byte("Success"))

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
