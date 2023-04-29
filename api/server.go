package api

import (
	"ascii-art-web/lib/utils"
	"io"
	"net/http"
)

func routes()  {
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
	})
	http.HandleFunc("/ascii-art",asciiArt)
}

func InitServer()  {
	routes()
	http.ListenAndServe(":8080",nil)
}

func asciiArt( w http.ResponseWriter, r *http.Request)  {
	body,_ := io.ReadAll(r.Body)
	*utils.Text = string(body)
}