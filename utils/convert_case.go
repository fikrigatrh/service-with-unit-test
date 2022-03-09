package utils

import (
	"strings"
)

func LowerCamelCase(v string) string {
	s := strings.Split(v, "")
	s[0] = strings.ToLower(s[0])

	res := strings.Join(s, "")

	return res
}

func ConvertMonth(str string) string {
	var ss string
	sTemp := strings.Split(str, "-")
	switch sTemp[1] {
	case "01":
		ss = "Jan"
	case "02":
		ss = "Feb"
	case "03":
		ss = "Mar"
	case "04":
		ss = "Apr"
	case "05":
		ss = "Mai"
	case "06":
		ss = "Jun"
	case "07":
		ss = "Jul"
	case "08":
		ss = "Agu"
	case "09":
		ss = "Sep"
	case "10":
		ss = "Okt"
	case "11":
		ss = "Nov"
	case "12":
		ss = "Des"
	}

	return ss
}
