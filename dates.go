package golang_helpers

//Created by Aditya Raj
//Functions to handle dates

import (
	"time"
	"log"
)

func GetTodaysDate() string {
	loc, _ := time.LoadLocation("America/Los_Angeles")
	current_time := time.Now().In(loc)
	return current_time.Format("2006-01-02")
}

func GetTodaysDateTime() string {
	loc, _ := time.LoadLocation("America/Los_Angeles")
	current_time := time.Now().In(loc)
	return current_time.Format("2006-01-02 15:04:05")
}

func GetTodaysDateTimeFormatted() string {
	loc, _ := time.LoadLocation("America/Los_Angeles")
	current_time := time.Now().In(loc)
	return current_time.Format("Jan 2, 2006 at 3:04 PM")
}

func GetTimeStampFromDate(dtformat string) string {
	form := "Jan 2, 2006 at 3:04 PM"
	t2, _ := time.Parse(form, dtformat)
	return t2.Format("20060102150405")
}

func GetTimeStamp() string {
	loc, _ := time.LoadLocation("America/Los_Angeles")
	t := time.Now().In(loc)
	return t.Format("20060102150405")
}

func GetFormattedDate(dtval string) string {
	dtnew := LocaliseTime(dtval)
	return dtnew.Format("Jan 2, 2006") // at 15:04
}

func GetRawDate(dtval string) string {
	dtnew := LocaliseTime(dtval)
	return dtnew.Format("2006-01-02") // at 15:04
}

func GetChatCodeTimeStamp(timeval string) string {
	dtnew := LocaliseDate().Format("Jan 2, 2006")
	return GetTimeStampFromDate(dtnew + " at "+ timeval)
}

func GetChatCodeTimeStampNextDay(timeval string) string {
	nextday := LocaliseDate().Add(time.Hour * 24)
	dtnew := nextday.Format("Jan 2, 2006")
	return GetTimeStampFromDate(dtnew + " at "+ timeval)
}

func GetChatCodePrevDayTimeStamp(ts string) string {
	location, err := time.LoadLocation("America/Los_Angeles")
	CheckErr2(err)
	tm, _ := time.Parse("20060102150405",ts)
	return tm.In(location).Add(time.Hour * -24).Format("20060102150405")
}

func GetFormattedDateTime(dtval string) string {
	dtnew := LocaliseTime(dtval)
	return dtnew.Format("Jan 2, 2006 at 3:04 PM") // at 15:04
}

func GetRawFormattedDateTime(dtval string) string {
	dtnew := LocaliseTime(dtval)
	return dtnew.Format("2006-01-02 15:04") // at 15:04
}

func GetRawFormattedDateTimeFull(dtval string) string {
	dtnew := LocaliseTime(dtval)
	return dtnew.Format("2006-01-02 15:04:05") // at 15:04
}

//20060102150405
func LocaliseTime(dtval string) time.Time {
	location, err := time.LoadLocation("America/Los_Angeles")
	CheckErr2(err)
	dt,_ := time.Parse("2006-01-02 15:04:05", dtval)
	return dt.In(location)
}

func DateMaxTime(dtval string) time.Time {
	location, err := time.LoadLocation("America/Los_Angeles")
	CheckErr2(err)
	dt,_ := time.Parse("2006-01-02 15:04:05", dtval)
	return dt.In(location)
}


func LocaliseDate() time.Time {
	location, err := time.LoadLocation("America/Los_Angeles")
	CheckErr2(err)
	t := time.Now()
	return t.In(location)
}

func CheckErr2(err error) {
	if err != nil {
		log.Panic(err)
		return
	}
}
