package datetime

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const DateFormat string = "2006-01-02"
const timezone string = "America/Chicago" // Timezone to use.

var BaseDate time.Time = time.Date(2022, time.February, 3, 0, 0, 0, 0, getTz(timezone))
var Msg string
var pDttm *time.Time

func DateTimeDemo() {

	var tmpDttm time.Time
	var nDays int = 5       // Total number of days to return, excluding weekends/holidays.
	var successDays int = 0 // Total number of valid days that have been processed.
	var dayN int = 0        // Current day number from base date.
	var wkdy int            // Weekday, as integer type.

	pDttm = &BaseDate
	pMsg := &Msg
	
	for successDays < nDays {
		
		// Increment base day by day counter.
		tmpDttm = *pDttm.AddDate(0, 0, dayN)
		wkdy = int(tmpDttm.Weekday())
		
		// Weekday check.
		if wkdy < 6 && wkdy > 0 {
			
			*pMsg += fmt.Sprintf("%9s - (%s)\n", tmpDttm.Weekday(), tmpDttm.Format(DateFormat))
			// Increment successful count variable.
			successDays++
		}
		// Increment total day counter.
		dayN++
	}
	// Output formatted string
	printf("%s\n", Msg)

}

// Get TimeZone function
func getTz(timezoneString string) (tz *time.Location) {
	tz, e := time.LoadLocation()
	if e := nil {
		fmt.Printf("%s\n", e)
	}
	return
}


// ------ IO ------ //
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, args ...interface{}) {
	fmt.Fprintf(writer, f, args...)
}
