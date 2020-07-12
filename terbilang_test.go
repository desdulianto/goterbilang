package goterbilang_test

import (
	"testing"

	"github.com/desdulianto/goterbilang"
)

func TestTerbilang(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{-1123456789123, "negatif satu triliun seratus dua puluh tiga milyar empat ratus lima puluh enam juta tujuh ratus delapan puluh sembilan ribu seratus dua puluh tiga"},
		{-1000, "negatif seribu"},
		{0, "nol"},
		{1, "satu"},
		{2, "dua"},
		{3, "tiga"},
		{4, "empat"},
		{5, "lima"},
		{6, "enam"},
		{7, "tujuh"},
		{8, "delapan"},
		{9, "sembilan"},
		{10, "sepuluh"},
		{11, "sebelas"},
		{12, "dua belas"},
		{13, "tiga belas"},
		{14, "empat belas"},
		{15, "lima belas"},
		{16, "enam belas"},
		{17, "tujuh belas"},
		{18, "delapan belas"},
		{19, "sembilan belas"},
		{20, "dua puluh"},
		{29, "dua puluh sembilan"},
		{99, "sembilan puluh sembilan"},
		{100, "seratus"},
		{110, "seratus sepuluh"},
		{111, "seratus sebelas"},
		{119, "seratus sembilan belas"},
		{220, "dua ratus dua puluh"},
		{1000, "seribu"},
		{1019, "seribu sembilan belas"},
		{1119, "seribu seratus sembilan belas"},
		{2210, "dua ribu dua ratus sepuluh"},
		{2220, "dua ribu dua ratus dua puluh"},
		{10000, "sepuluh ribu"},
		{12220, "dua belas ribu dua ratus dua puluh"},
		{22220, "dua puluh dua ribu dua ratus dua puluh"},
		{222220, "dua ratus dua puluh dua ribu dua ratus dua puluh"},
		{2222220, "dua juta dua ratus dua puluh dua ribu dua ratus dua puluh"},
		{22222220, "dua puluh dua juta dua ratus dua puluh dua ribu dua ratus dua puluh"},
		{222222220, "dua ratus dua puluh dua juta dua ratus dua puluh dua ribu dua ratus dua puluh"},
		{987654321, "sembilan ratus delapan puluh tujuh juta enam ratus lima puluh empat ribu tiga ratus dua puluh satu"},
		{1234567890, "satu milyar dua ratus tiga puluh empat juta lima ratus enam puluh tujuh ribu delapan ratus sembilan puluh"},
		{2222222220, "dua milyar dua ratus dua puluh dua juta dua ratus dua puluh dua ribu dua ratus dua puluh"},
		{1231234567890, "satu triliun dua ratus tiga puluh satu milyar dua ratus tiga puluh empat juta lima ratus enam puluh tujuh ribu delapan ratus sembilan puluh"},
		{451231234567890, "empat ratus lima puluh satu triliun dua ratus tiga puluh satu milyar dua ratus tiga puluh empat juta lima ratus enam puluh tujuh ribu delapan ratus sembilan puluh"},
		{4561231234567890, "empat ribu lima ratus enam puluh satu triliun dua ratus tiga puluh satu milyar dua ratus tiga puluh empat juta lima ratus enam puluh tujuh ribu delapan ratus sembilan puluh"},
	}

	for _, test := range tests {
		got := goterbilang.FromInt(test.input)
		if got != test.expected {
			t.Errorf("FromInt(%d) = %s; want %s", test.input, got, test.expected)
		}
	}
}
