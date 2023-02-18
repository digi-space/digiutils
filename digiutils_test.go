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

func TestContainsDuplicate(t *testing.T) {
	got := []bool{ContainsDuplicate([]int{1, 2, 3, 1}), ContainsDuplicate([]int{1, 2, 3, 4}), ContainsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2})}
	expected := []bool{true, false, true}

	for i := 0; i < len(got); i++ {
		if got[i] != expected[i] {
			t.Errorf("got %t, expected %t", got[i], expected[i])
		}
	}
}

func TestIsPowerOfTwo(t *testing.T) {
	got := []bool{IsPowerOfTwo(1), IsPowerOfTwo(16), IsPowerOfTwo(3)}
	expected := []bool{true, true, false}

	for i := 0; i < len(got); i++ {
		if got[i] != expected[i] {
			t.Errorf("got %t, expected %t", got[i], expected[i])
		}
	}
}

func TestAddDigits(t *testing.T) {
	got := []int{AddDigits(38), AddDigits(0)}
	expected := []int{2, 0}

	for i := 0; i < len(got); i++ {
		if got[i] != expected[i] {
			t.Errorf("got %d, expected %d", got[i], expected[i])
		}
	}
}

func TestGetMiddle(t *testing.T) {
	got := []int{GetMiddle([]int{1, 2, 3}), GetMiddle([]int{5, 4, 8, 9})}
	expected := []int{1, -1}

	for i := 0; i < len(got); i++ {
		if got[i] != expected[i] {
			t.Errorf("got %d, expected %d", got[i], expected[i])
		}
	}
}

func TestRoundToNext5(t *testing.T) {
	got := []int{RoundToNext5(0), RoundToNext5(3), RoundToNext5(21)}
	expected := []int{0, 5, 25}

	for i := 0; i < len(got); i++ {
		if got[i] != expected[i] {
			t.Errorf("got %d, expected %d", got[i], expected[i])
		}
	}
}

func TestIsTriangle(t *testing.T) {
	got := []bool{IsTriangle(5, 1, 2), IsTriangle(1, 2, 5), IsTriangle(4, 2, 3)}
	expected := []bool{false, false, true}

	for i := 0; i < len(got); i++ {
		if got[i] != expected[i] {
			t.Errorf("got %t, expected %t", got[i], expected[i])
		}
	}
}

func TestTwoOldestAges(t *testing.T) {
	got := [][2]int{TwoOldestAges([]int{6, 5, 83, 5, 3, 18}), TwoOldestAges([]int{1, 5, 87, 45, 8, 8})}
	expected := [][2]int{{18, 83}, {45, 87}}

	for i := 0; i < len(got); i++ {
		if got[i] != expected[i] {
			t.Errorf("got %v, expected %v", got[i], expected[i])
		}
	}
}

func TestGetCount(t *testing.T) {
	got := []int{GetCount("abracadabra")}
	expected := []int{5}

	for i := 0; i < len(got); i++ {
		if got[i] != expected[i] {
			t.Errorf("got %d, expected %d", got[i], expected[i])
		}
	}
}
