package colors

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/common-nighthawk/go-figure"
)

func FigletLogo(text string) {
	type Combo struct {
		Text  string
		Font  string
		Color string
	}
	var possibleCombos = []Combo{
		{text, "cricket", "blue"},
		{text, "smisome1", "red"},
		{"        " + text, "contessa", "yellow"},
		{text, "binary", "green"},
		{text, "larry3d", "cyan"},
	}

	randomCombo := possibleCombos[rand.Intn(len(possibleCombos))]
	if randomCombo.Font == "binary" {
		fmt.Println()
	}
	myFigure := figure.NewColorFigure(randomCombo.Text, randomCombo.Font, randomCombo.Color, true)
	myFigure.Print()
	fmt.Println()
}

var defaultFont = "cybermedium"

func FigletPrint(text string) {
	myFigure := figure.NewColorFigure(text, defaultFont, "green", true)
	myFigure.Print()
}

func FigletPrintWithFont(text string, font string) {
	myFigure := figure.NewColorFigure(text, font, "green", true)
	myFigure.Print()
}

func FigletPrintWithFontAndColor(text string, font string, color string) {
	myFigure := figure.NewColorFigure(text, font, color, true)
	myFigure.Print()
}

func FigletPrintError(text string) {
	figure.NewColorFigure("ERROR: "+text, defaultFont, "red", true).Print()
	fmt.Println()
}

func FigletPrintErrorWithFont(text string, font string) {
	myFigure := figure.NewColorFigure("ERROR: "+text, font, "red", true)
	myFigure.Print()
	fmt.Println()
}

func ShowAllFigletFonts() {
	fonts, err := figure.AssetDir("fonts")
	if err != nil {
		fmt.Println("error getting fonts:", err)
	}
	for _, font := range fonts {
		font = strings.TrimSuffix(font, ".flf")
		fmt.Println(font)
		myFigure := figure.NewColorFigure(font, font, "green", true)
		myFigure.Print()
	}
}
