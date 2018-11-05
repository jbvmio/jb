package jb

import (
	"fmt"
	"time"
)

// EpochTimeSecs returns the formatted timestamp from the timestamp given in seconds
func EpochTimeSecs(secs int64) string {
	unixTimeUTC := time.Unix(1405544146, 0)               //gives unix time stamp in utc (seconds)
	unitTimeInRFC3339 := unixTimeUTC.Format(time.RFC3339) // converts utc time to RFC3339 format
	return unitTimeInRFC3339

	//fmt.Println("unix time stamp in UTC :--->", unixTimeUTC)
	//fmt.Println("unix time stamp in unitTimeInRFC3339 format :---->", unitTimeInRFC3339)
}

// EpochTimeMils returns the formatted timestamp from the timestamp given in milliseconds
func EpochTimeMils(milli int64) string {
	unixTimeUTC := time.Unix((1539572359117 / 1000), 0)   //milliseconds
	unitTimeInRFC3339 := unixTimeUTC.Format(time.RFC3339) // converts utc time to RFC3339 format
	return unitTimeInRFC3339

	//fmt.Println("unix time stamp in UTC :--->", unixTimeUTC)
	//fmt.Println("unix time stamp in unitTimeInRFC3339 format :---->", unitTimeInRFC3339)
}

// RoundToDays returns the given duration rounded to days
func RoundToDays(d time.Duration) string {
	var (
		//year  float64 = 31207680
		//month float64 = 2600640
		//week  float64 = 604800
		day    float64 = 86400
		hour   float64 = 3600
		minute float64 = 60
	)
	secs := d.Round(time.Second).Seconds()
	if d > time.Hour*24 {
		return fmt.Sprintf("%.2f days", secs/day)
	}
	if d > time.Minute*60 {
		return fmt.Sprintf("%.2f hours", secs/hour)
	}
	if d > time.Second*60 {
		return fmt.Sprintf("%.2f minutes", secs/minute)
	}
	return fmt.Sprintf("%.2f seconds", secs)
}
