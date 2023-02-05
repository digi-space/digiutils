package digiutils

import "testing"

func TestAddBinary(t *testing.T) {
	got := []string{AddBinary("11", "1"), AddBinary("1010", "1011")}
	expected := []string{"100", "10101"}

	for i := 0; i < len(got); i++ {
		if got[i] != expected[i] {
			t.Errorf("got %s, expected %s", got[i], expected[i])
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	got := []bool{IsPalindrome("A man, a plan, a canal: Panama"), IsPalindrome("race a car"), IsPalindrome(" ")}
	expected := []bool{true, false, true}

	for i := 0; i < len(got); i++ {
		if got[i] != expected[i] {
			t.Errorf("got %t, expected %t", got[i], expected[i])
		}
	}
}

func TestRomanToInt(t *testing.T) {
	got := []int{RomanToInt("III"), RomanToInt("LVIII"), RomanToInt("MCMXCIV"), RomanToInt("QWERWQR")}
	expected := []int{3, 58, 1994, 0}

	for i := 0; i < len(got); i++ {
		if got[i] != expected[i] {
			t.Errorf("got %d, expected %d", got[i], expected[i])
		}
	}
}
