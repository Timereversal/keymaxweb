package models

type CarKeyInfo struct {
	Brand               string
	Model               string
	Year                int
	VIN                 string
	Transponder         string
	BladeKey            string
	Frequency           int
	PartNumber          string
	RemotePhotoId       int
	RemoteFormatPhotoID int
	MatchingProcedure   string
	KeyCuttingInfo      string
}
