package controllers

import (
	"fmt"
	"github.com/timereversal/keymaxweb/models"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type CarKeyInfo struct {
	CarInfoService *models.CarKeyInfoService
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

	year, err := strconv.Atoi(r.FormValue("year"))
	if err != nil {
		// send response Year parameter is not a valid year
		fmt.Println(err)
		fmt.Sprintf("error conversion year to string: %s", err)
		return
	}
	freq, err := strconv.Atoi(r.FormValue("frequency"))
	if err != nil {
		// send response Year parameter is not a valid year
		fmt.Println(err)
		fmt.Sprintf("error conversion frequency to string: %s", err)
		return
	}
	fmt.Println("pre cInfo")
	fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	cInfo := models.CarKeyInfo{
		ID:                0,
		CreatedAt:         time.Now(),
		Brand:             r.FormValue("brand"),
		Model:             r.FormValue("model"),
		Year:              year,
		VIN:               r.FormValue("vin"),
		Transponder:       r.FormValue("transponder"),
		BladeKey:          r.FormValue("blade"),
		Frequency:         freq,
		PartNumber:        r.FormValue("partNumber"),
		MatchingProcedure: r.FormValue("programmingProcedure"),
		KeyCuttingInfo:    r.FormValue("cutInfo"),
	}
	fmt.Println("uuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuu")
	fmt.Println(cInfo)
	//fmt.Sprintf("%+v", cInfo)
	// ID and CreatedAt will be filled inside Create function
	err = c.CarInfoService.Create(&cInfo)
	if err != nil {
		//send response message to frontend
		fmt.Println(err)
		return
	}
	fmt.Sprintf("%+v", cInfo)
	//
	//log.Println(fmt.Sprintf("brand: %s, model: %s, year: %s, vin: %s, transponder: %s, blade: %s, frequency: %s, "+
	//	"partNumber: %s, programmingProcedure: %s, cutInfo: %s",
	//	r.FormValue("brand"),
	//	r.FormValue("model"),
	//	r.FormValue("year"),
	//	r.FormValue("vin"),
	//	r.FormValue("transponder"),
	//	r.FormValue("blade"),
	//	r.FormValue("frequency"),
	//	r.FormValue("partNumber"),
	//	r.FormValue("programmingProcedure"),
	//	r.FormValue("cutInfo"),
	//))
	//log.Println("fromFile", r)
	log.Println(time.Now().Unix())

	//filename := handler.Filename
	//brand := r.FormValue("brand")
	//log.Println("filename: ", filename)
	//log.Println("brand : ", brand)
	result := fmt.Sprintf(`{"ID": %d, "Created_at": %q}`, cInfo.ID, cInfo.CreatedAt)

	w.Header().Set("content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//fmt.Fprintf(w)
	w.Write([]byte(result))
}
