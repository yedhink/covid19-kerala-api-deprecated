package date

import (
	"time"

	. "github.com/yedhink/covid19-kerala-api/internal/storage"
)

const (
	//              mm-dd-yyyy
	hyphenLayout = "02-01-2006"
	slashLayout  = "02/01/2006"
	nilLayout    = "00-00-0000"
)

// returns the date in zulu date format
// eg :- 2020-03-27T00:00:00Z
func ZuluDateFormat(path string) string {
	date, _ := time.Parse(hyphenLayout, path)
	return date.Format(time.RFC3339)
}

// checks if user param date contains a valid formatted date
// 02-04-20 or 02/04/20 etc are invalid
func IsDate(d string) bool {
	date := ""
	switch d[:1] {
	case "<", ">":
		date = d[1:]
	case "l":
		return true
	default:
		date = d
	}
	if _, err := time.Parse(hyphenLayout, date); err != nil {
		_, err = time.Parse(slashLayout, date)
		if err != nil {
			return false
		}
	}
	return true
}

// returns time objects for the strings k and d
func parseDate(k string, d string) (time.Time, time.Time) {
	userDate, err := time.Parse(hyphenLayout, d)
	if err != nil {
		userDate, _ = time.Parse(slashLayout, d)
	}
	userDate.Format(time.RFC3339)
	keyDate, _ := time.Parse(time.RFC3339, k)
	return keyDate, userDate
}

// validates the date for different cases like '<','>','latest' and eqauls
// k is the timestamp from the json file, d is the user specified date
func ValidDate(k string, d string, st *Storage) bool {
	switch d[:1] {
	case "<":
		kD, uD := parseDate(k, d[1:])
		return kD.Before(uD)
	case ">":
		kD, uD := parseDate(k, d[1:])
		return kD.After(uD)
	case "l":
		// latest data pdf
		date := GetLocalPdfDate(st.BasePath)
		kD, uD := parseDate(k, date)
		return kD.Equal(uD)
	default:
		kD, uD := parseDate(k, d)
		return kD.Equal(uD)
	}
}
