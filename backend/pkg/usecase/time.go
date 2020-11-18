package usecase

import (
	"strconv"
	"time"
)

const location = "Asia/Tokyo"

func init() {
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc
}

func GetDeteInTokyo() int {
	dateTime, _ := strconv.Atoi(time.Now().Format("20060102"))
	return dateTime
}
