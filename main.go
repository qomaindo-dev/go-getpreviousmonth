package main

import (
	"fmt"
	"time"
)

var (
  //Example of a collection of months
	tgl1 = 1717121780 //Mei
	tgl2 = 1722392180 //July
	tgl3 = 1725070580 //Augustus
	tgl4 = 1730340980 //October
	tgl5 = 1735611380 //December
	tgl6 = 1706667380 //January
)

func main() {
	timeLocation, _ := time.LoadLocation("Asia/Jakarta")
  
	timeOneDate := time.Unix(int64(tgl5), 0)
  
  //To validate if there are several months with a total of 31 days
	if timeOneDate.Day() == 31 {
		timeOneDate = time.Date(timeOneDate.Year(), timeOneDate.Month(), 29, timeOneDate.Hour(), timeOneDate.Minute(), timeOneDate.Second(), 0, timeLocation)
	}

	previousMonth1 := timeOneDate.AddDate(0, -1, 0)

	previousMonth2 := timeOneDate.AddDate(0, -2, 0)

	previousMonth3 := timeOneDate.AddDate(0, -3, 0)

	fmt.Println(previousMonth1.Day(), "-", previousMonth1.Month(), "-", previousMonth1.Year())

	firstDay1, lastDay1 := MonthInterval(previousMonth1.Year(), previousMonth1.Month())
	firstDay2, lastDay2 := MonthInterval(previousMonth2.Year(), previousMonth2.Month())
	firstDay3, lastDay3 := MonthInterval(previousMonth3.Year(), previousMonth3.Month())

	fmt.Println("Start Day-1", firstDay1.Format("02 Jan 2006"), "End Day-1: ", lastDay1.Format("02 Jan 2006"))
	fmt.Println("Start Day-2", firstDay2.Format("02 Jan 2006"), "End Day-2: ", lastDay2.Format("02 Jan 2006"))
	fmt.Println("Start Day-3", firstDay3.Format("02 Jan 2006"), "End Day-3: ", lastDay3.Format("02 Jan 2006"))
}

func MonthInterval(y int, m time.Month) (firstDay, lastDay time.Time) {
	timeLocation, _ := time.LoadLocation("Asia/Jakarta")

	firstDay = time.Date(y, m, 1, 0, 0, 0, 0, timeLocation)
	lastDay = time.Date(y, m+1, 1, 0, 0, 0, -1, timeLocation)
	return firstDay, lastDay
}
