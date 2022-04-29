package ms

import "fmt"

var trueValues = map[string]bool{
	"t":    true,
	"T":    true,
	"yes":  true,
	"Yes":  true,
	"YES":  true,
	"1":    true,
	"true": true,
	"True": true,
	"TRUE": true,
	"on":   true,
	"On":   true,
	"ON":   true,
}

func ParseBool(s string) (b bool) {
	_, b = trueValues[s]
	return
}

var trueValues2 = map[string]bool{
	"t":    true,
	"T":    true,
	"yes":  true,
	"Yes":  true,
	"YES":  true,
	"1":    true,
	"true": true,
	"True": true,
	"TRUE": true,
	"on":   true,
	"On":   true,
	"ON":   true,

	"f":     false,
	"F":     false,
	"no":    false,
	"No":    false,
	"NO":    false,
	"0":     false,
	"false": false,
	"False": false,
	"FALSE": false,
	"off":   false,
	"Off":   false,
	"OFF":   false,
}

func ParseBoolChecked(s string) (b bool, err error) {
	v, ok := trueValues2[s]
	if ok {
		b = v
		return
	}
	err = fmt.Errorf("Invalid Boolean Value")
	return
}

func InList(s string, v []string) bool {
	for _, q := range v {
		if s == q {
			return true
		}
	}
	return false
}
