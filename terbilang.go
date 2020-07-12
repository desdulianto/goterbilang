package goterbilang

import (
	"fmt"
	"math"
	"strings"
)

var satuan = [...]string{"nol", "satu", "dua", "tiga", "empat",
	"lima", "enam", "tujuh", "delapan", "sembilan", "sepuluh", "sebelas"}

type terbilangSuffix struct {
	nilai int64
	label string
}

var suffixes = [...]terbilangSuffix{
	{1000000000000, "triliun"},
	{1000000000, "milyar"},
	{1000000, "juta"},
	{1000, "ribu"},
	{100, "ratus"},
	{10, "puluh"},
}

func findFirst(n int) (terbilangSuffix, error) {
	for _, suffix := range suffixes {
		if int64(n) >= suffix.nilai {
			return suffix, nil
		}
	}
	return terbilangSuffix{}, fmt.Errorf("not found")
}

func FromInt(n int) string {
	switch {
	case n < 0:
		return fmt.Sprintf("negatif %s", FromInt(int(math.Abs(float64(n)))))
	case n < 12:
		return satuan[n]
	case n >= 12 && n <= 19:
		return fmt.Sprintf("%s belas", satuan[n%10])
	default:
		suffix, err := findFirst(n)
		if err != nil {
			return ""
		}
		trail := ""
		if int64(n)%suffix.nilai > 0 {
			trail = FromInt(int(int64(n) % suffix.nilai))
		}
		return strings.TrimSpace(strings.ReplaceAll(
			strings.ReplaceAll(
				fmt.Sprintf("%s %s %s", FromInt(int(int64(n)/suffix.nilai)),
					suffix.label, trail),
				"satu ratus", "seratus"),
			"satu ribu", "seribu"))
	}
}
