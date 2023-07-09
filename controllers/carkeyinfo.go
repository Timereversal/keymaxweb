package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type CarKeyInfo struct {
}

func (c *CarKeyInfo) Create(w http.ResponseWriter, r *http.Request) {

	log.Println("path", r.URL.Path)
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		log.Println("error", err)
	}

	mForm := r.MultipartForm
	//log.Println(mForm)
	for k := range mForm.File {
		//file, handler, err := r.FormFile("keyPicture")
		f, handler, err := r.FormFile(k)
		if err != nil {
			log.Println("Error: ", err)
			return
		}
		defer f.Close()

		file, err := os.OpenFile(fmt.Sprintf("./images/%s", handler.Filename), os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Println("Error", err)
		}
		io.Copy(file, f)

	}

	log.Println(fmt.Sprintf("brand: %s, model: %s, year: %s, vin: %s, transponder: %s, blade: %s, frequency: %s, "+
		"partNumber: %s, programmingProcedure: %s, cutInfo: %s",
		r.FormValue("brand"),
		r.FormValue("model"),
		r.FormValue("year"),
		r.FormValue("vin"),
		r.FormValue("transponder"),
		r.FormValue("blade"),
		r.FormValue("frequency"),
		r.FormValue("partNumber"),
		r.FormValue("programmingProcedure"),
		r.FormValue("cutInfo"),
	))
	//log.Println("fromFile", r)
	log.Println(time.Now().Unix())

	//filename := handler.Filename
	//brand := r.FormValue("brand")
	//log.Println("filename: ", filename)
	//log.Println("brand : ", brand)

	w.Header().Set("Access-Control-Allow-Origin", "*")
}
