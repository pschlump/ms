package ms

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

/// func init() {
/// 	trueValues = make(map[string]bool)
/// 	trueValues["t"] = true
/// 	trueValues["T"] = true
/// 	trueValues["yes"] = true
/// 	trueValues["Yes"] = true
/// 	trueValues["YES"] = true
/// 	trueValues["1"] = true
/// 	trueValues["true"] = true
/// 	trueValues["True"] = true
/// 	trueValues["TRUE"] = true
/// 	trueValues["on"] = true
/// 	trueValues["On"] = true
/// 	trueValues["ON"] = true
/// }

func ParseBool(s string) (b bool) {
	_, b = trueValues[s]
	return
}
