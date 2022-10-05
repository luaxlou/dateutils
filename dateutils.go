package dateutils

import "time"

const STANARD_FORMAT = "2006-01-02 15:04:05"
const DATE_FORMAT = "2006-01-02"
const MONTH_FORMAT = "2006-01"

const FLAT_DATE = "20060102"
const FLAT_FORMAT = "20060102150405"

func Format(ts int64) string {
	return time.Unix(ts, 0).Format(STANARD_FORMAT)
}

func Parse(value string) (time.Time, error) {

	return time.ParseInLocation(STANARD_FORMAT, value, time.Local)

}
