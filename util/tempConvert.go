package tempConvert

// FToC convert temperature from celsius to fahrenheit
func FToC(f float64) float64 {
    return (f - 32) * 5 / 9
}

// CToF fahrenheit to celsius
func CToF(f float64) float64 {
    return (f * 1.8) + 32
}
