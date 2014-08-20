package pretty

import "strconv"

// Periodify renders an int with period markers to make it more easily
// human-readable.
func Periodify(num int, marker byte) string {
	numStr := strconv.Itoa(num)
	numMarkers := (len(numStr)-1) / 3
	b := make([]byte, len(numStr) + numMarkers)
	for i := len(numStr); i > 0; i-- {
		if (len(numStr) - i) % 3 == 0 && i != len(numStr) {
			b[i+numMarkers-1] = marker
			numMarkers--
		}
		b[i+numMarkers-1] = numStr[i-1]
	}
	return string(b)
}

// PeriodifyUS returns a human-readable number with a comma character
// as the period marker. There's a cleaner way to do this, but it
// works for simple use cases.
func PeriodifyUS(num int) string {
  return Periodify(num, ',')
}
