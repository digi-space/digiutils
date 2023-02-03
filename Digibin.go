package Digibin

import "strconv"

func AddBinary(a string, b string) (result string) {
	n, m := len(a), len(b)
	carry := 0
	res := []int{}
	for i, j := n-1, m-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		x := 0
		if i >= 0 {
			x, _ = strconv.Atoi(string(a[i]))
		}
		y := 0
		if j >= 0 {
			y, _ = strconv.Atoi(string(b[j]))
		}
		sum := x + y + carry
		res = append([]int{sum % 2}, res...)
		carry = sum / 2
	}
	if carry > 0 {
		res = append([]int{1}, res...)
	}

	for _, val := range res {
		result += strconv.Itoa(val)
	}
	return
}
