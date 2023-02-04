package digiutils

// AddBinary takes two parameters a and b and returns
// their sum as a binary string.
func AddBinary(a string, b string) string {
	n, m := len(a), len(b)
	carry := 0
	res := []byte{}
	for i, j := n-1, m-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		x := 0
		if i >= 0 {
			x = int(a[i] - '0')
		}
		y := 0
		if j >= 0 {
			y = int(b[j] - '0')
		}
		sum := x + y + carry
		res = append([]byte{byte(sum%2) + '0'}, res...)
		carry = sum / 2
	}
	if carry > 0 {
		res = append([]byte{'1'}, res...)
	}
	return string(res)
}