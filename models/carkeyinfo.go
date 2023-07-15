package models

import (
	"database/sql"
	"fmt"
	"time"
)

type CarKeyInfo struct {
	ID                  int
	CreatedAt           time.Time
	Brand               string
	Model               string
	Year                int
	VIN                 string
	Transponder         string
	BladeKey            string
	Frequency           int
	PartNumber          string
	RemotePhotoId       string
	RemoteFormatPhotoID string
	MatchingProcedure   string
	KeyCuttingInfo      string
}

type CarKeyInfoService struct {
	DB *sql.DB
}

func (ckinfoservice *CarKeyInfoService) Create(ckinfo *CarKeyInfo) error {

	query := `
		INSERT INTO carkeyinfo (brand, model, year, vin, transponder, blade, frequency, partnumber,
		                        keypicture, remoteformat, programmingprocedure, cutinfo)
		                        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		                        RETURNING id, created_at`
	fmt.Println(ckinfo.Brand)
	fmt.Println(ckinfo.Model)
	fmt.Println(ckinfoservice, &ckinfoservice)

	row := ckinfoservice.DB.QueryRow(query, ckinfo.Brand, ckinfo.Model, ckinfo.Year, ckinfo.VIN,
		ckinfo.Transponder, ckinfo.BladeKey, ckinfo.Frequency, ckinfo.PartNumber, ckinfo.RemotePhotoId,
		ckinfo.RemoteFormatPhotoID, ckinfo.MatchingProcedure, ckinfo.KeyCuttingInfo)
	fmt.Println("1 line after row")
	err := row.Scan(&ckinfo.ID, &ckinfo.CreatedAt)
	if err != nil {
		return fmt.Errorf("create carkeyinfo : %w", err)
	}

	fmt.Println(ckinfo)
	fmt.Printf(" struct: %+v", ckinfo)

	return nil

}
