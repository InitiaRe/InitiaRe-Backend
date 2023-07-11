package conversion

import "time"

func FormatUnixToString(d int, format string) interface{} {
	if d != 0 {
		switch format {
		case "RFC822":
			return time.UnixMilli(int64(d)).Format(time.RFC822)
		case "Kitchen":
			return time.UnixMilli(int64(d)).Format(time.Kitchen)
		case "UnixDate":
			return time.UnixMilli(int64(d)).Format(time.UnixDate)
		case "YYYY-MM-DD":
			return time.UnixMilli(int64(d)).Format("2006-01-02")
		case "DD-MM-YYYY":
			return time.UnixMilli(int64(d)).Format("02-01-2006")
		case "YYYY-MM-DD HH:mm:ss":
			return time.UnixMilli(int64(d)).Format("2006-01-02 15:04:05")
		case "DD-MM-YYYY HH:mm:ss":
			return time.UnixMilli(int64(d)).Format("02-01-2006 15:04:05")
		default:
			return time.UnixMilli(int64(d)).Format("2006-01-02 15:04:05")
		}
	}
	return nil
}
