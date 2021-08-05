package main

type Color string

// https://en.wikipedia.org/wiki/ANSI_escape_code#Colors
const (
	ColorBlack         = "\u001b[30m"
	ColorRed           = "\u001b[31m"
	ColorGreen         = "\u001b[32m"
	ColorYellow        = "\u001b[33m"
	ColorBlue          = "\u001b[34m"
	ColorMagneta       = "\u001b[35m"
	ColorCyan          = "\u001b[36m"
	ColorWhite         = "\u001b[37m"
	ColorGray          = "\u001b[90m"
	ColorBrightRed     = "\u001b[91m"
	ColorBrightGreen   = "\u001b[92m"
	ColorBrightYellow  = "\u001b[93m"
	ColorBrightBlue    = "\u001b[94m"
	ColorBrightMagneta = "\u001b[95m"
	ColorBrightCyan    = "\u001b[96m"
	ColorBrightWhite   = "\u001b[97m"
	ColorReset         = "\u001b[0m"
)

const (
	CORRECT_POINT = 3
	PASS_POINT    = 1
	WRONG_POINT   = -1
)

type QuizRecord struct {
	ID int
	Q  string
	A  string
}
