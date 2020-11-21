package helpers

import (
	"time"
)

// Now: Alias to time.Time
func Now() time.Time {
	return time.Now()
}

// Benchmark: Basically an alis to time.Since, printing the result; also return the result as a string and time.Duration
func Benchmark(s time.Time) (string, time.Duration) {
	since := time.Since(s)
	Print(since)

	return Dts(since), since
}

// NowUTC: Current datetime as UTC
func NowUTC() time.Time {
	return Now().UTC()
}

// NowTimestamp: Current Timestamp (unix, INT64)
func NowTimestamp() int64 {
	return Now().Unix()
}

// NowString: Current date as STRING
func NowString() string {
	return Now().String()
}

// Year: Current year
func Year() string {
	return Its(Now().Year())
}

// MonthPtBR: Current month in pr-BR (written long)
func MonthPtBR() string {
	key := month()

	month := make(map[string]string, 12)

	month["January"] = "Janeiro"
	month["February"] = "Fevereiro"
	month["March"] = "Março"
	month["April"] = "Abril"
	month["May"] = "Maio"
	month["June"] = "Junho"
	month["July"] = "Julho"
	month["August"] = "Agosto"
	month["September"] = "Setembro"
	month["October"] = "Outubro"
	month["November"] = "Novembro"
	month["December"] = "Dezembro"

	return month[key]
}

func month() string {
	return Now().Month().String()
}

// ShortMonthPtBR: Current month in pr-BR (written short)
func ShortMonthPtBR() string {
	key := month()

	month := make(map[string]string, 12)

	month["January"] = "Jan"
	month["February"] = "Fev"
	month["March"] = "Mar"
	month["April"] = "Abr"
	month["May"] = "Mai"
	month["June"] = "Jun"
	month["July"] = "Jul"
	month["August"] = "Ago"
	month["September"] = "Set"
	month["October"] = "Out"
	month["November"] = "Nov"
	month["December"] = "Dez"

	return month[key]
}

// MonthNum: Current month in numeric format (return as string)
func MonthNum() string {
	months := make(map[string]string, 12)

	months["January"] = "01"
	months["February"] = "02"
	months["March"] = "03"
	months["April"] = "04"
	months["May"] = "05"
	months["June"] = "06"
	months["July"] = "07"
	months["August"] = "08"
	months["September"] = "09"
	months["October"] = "10"
	months["November"] = "11"
	months["December"] = "12"

	return months[month()]
}

// Day: Current day as string
func Day() string {
	day := time.Now().Day()
	if day < 10 {
		return "0" + Its(day)
	}
	return Its(day)
}

// Hour: Current hour as string
func Hour() string {
	hour := Now().Hour()
	if hour < 10 {
		return "0" + Its(hour)
	}

	return Its(hour)
}

// Minute: Current minute as string
func Minute() string {
	minute := Now().Minute()
	if minute < 10 {
		return "0" + Its(Now().Minute())
	}

	return Its(minute)
}

// Second: Current second as string
func Second() string {
	second := Now().Second()
	if second < 10 {
		return "0" + Its(second)
	}

	return Its(second)
}

// Nanosecond: Current nanosecond as string
func Nanosecond() string {
	nanosecond := Now().Nanosecond()
	if nanosecond < 10 {
		return "0" + Its(nanosecond)
	}

	return Its(nanosecond)
}

// GetToday: Current full date; if database ? YYYY-MM-DD : DD/MM/YYYY
func GetToday(databaseFormat bool) string {
	if databaseFormat {
		return GetCustomTime("Y-M-D")
	}

	return GetCustomTime("D/M/Y")
}

// GetNow: Current full date with time; database ? YYYY-MM-DD HH:II:SS : DD/MM/YYYY HH:II:SS
func GetNow(databaseFormat bool) string {
	return GetToday(databaseFormat) + " " + GetCustomTime("H:I:S")
}

// GetCustomTime: User string expr to get datetime
// Y: year;
// M: month (numeric);
// L: month (written long);
// A: month (written short);
// D: day;
// H: hour;
// I: minute;
// S: second;
// N: Nanosecond;
func GetCustomTime(expr string) string {
	var date string
	for _, char := range expr {
		date += returnFormat(string(char))
	}

	return date
}

func returnFormat(char string) string {
	if char == "Y" {
		return Year()
	}
	if char == "M" {
		return MonthNum()
	}
	if char == "L" {
		return MonthPtBR()
	}
	if char == "A" {
		return ShortMonthPtBR()
	}
	if char == "D" {
		return Day()
	}
	if char == "H" {
		return Hour()
	}
	if char == "I" {
		return Minute()
	}
	if char == "S" {
		return Second()
	}
	if char == "N" {
		return Nanosecond()
	}

	return char
}

// AddDuration: Alias to time.Add()
func AddDuration(t time.Time, add time.Duration) time.Time {
	return t.Add(add)
}

// Wait: Alias to time.sleep()
func Wait(t time.Duration) {
	time.Sleep(t)
}
