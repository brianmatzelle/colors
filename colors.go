package colors

import "fmt"

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func Colorize(text string, color string) string {
	return color + text + Reset
}

func UserF(text string) string {
	return Colorize(text, Blue)
}

func ModF(text string, color int) string {
	color = color + 1
	mod := color % 256
	return CustomColorF("["+text+"]", mod)
}

func ErrorF(text string) string {
	return Colorize(text, Red)
}

func WarningF(text string) string {
	return Colorize(text, Yellow)
}

func InfoF(text string) string {
	return Colorize(text, Cyan)
}

func SuccessF(text string) string {
	return Colorize(text, Green)
}

func PathF(text string) string {
	return Colorize(text, Cyan)
}

var ExtendedColor = func(colorCode int) string {
	return fmt.Sprintf("\033[38;5;%dm", colorCode)
}

var RGBColor = func(r, g, b int) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

func CustomColorF(text string, colorCode int) string {
	return Colorize(text, ExtendedColor(colorCode))
}

func CustomRGBF(text string, r, g, b int) string {
	return Colorize(text, RGBColor(r, g, b))
}

func ShowAllColors() {
	for i := 0; i < 256; i++ {
		fmt.Println(CustomColorF(fmt.Sprintf("%d", i), i))
	}
}
