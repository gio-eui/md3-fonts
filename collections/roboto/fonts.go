package roboto

import (
	"fmt"
	"gioui.org/font"
	"gioui.org/font/opentype"
	"gioui.org/text"
	"github.com/gio-eui/md3-fonts/fonts/roboto/robotoblack"
	"github.com/gio-eui/md3-fonts/fonts/roboto/robotoblackitalic"
	"github.com/gio-eui/md3-fonts/fonts/roboto/robotobold"
	"github.com/gio-eui/md3-fonts/fonts/roboto/robotobolditalic"
	"github.com/gio-eui/md3-fonts/fonts/roboto/robotolight"
	"github.com/gio-eui/md3-fonts/fonts/roboto/robotolightitalic"
	"github.com/gio-eui/md3-fonts/fonts/roboto/robotomedium"
	"github.com/gio-eui/md3-fonts/fonts/roboto/robotomediumitalic"
	"github.com/gio-eui/md3-fonts/fonts/roboto/robotoregular"
	"github.com/gio-eui/md3-fonts/fonts/roboto/robotoregularitalic"
	"github.com/gio-eui/md3-fonts/fonts/roboto/robotothin"
	"github.com/gio-eui/md3-fonts/fonts/roboto/robotothinitalic"
	"sync"
)

var (
	once       sync.Once
	collection []text.FontFace
)

func Collection() []font.FontFace {
	once.Do(func() {
		register(font.Font{Weight: font.Thin, Style: font.Regular}, robotothin.TypeFace, robotothin.Data)
		register(font.Font{Weight: font.Thin, Style: font.Italic}, robotothinitalic.TypeFace, robotothinitalic.Data)
		register(font.Font{Weight: font.Light, Style: font.Regular}, robotolight.TypeFace, robotolight.Data)
		register(font.Font{Weight: font.Light, Style: font.Italic}, robotolightitalic.TypeFace, robotolightitalic.Data)
		register(font.Font{Weight: font.Normal, Style: font.Regular}, robotoregular.TypeFace, robotoregular.Data)
		register(font.Font{Weight: font.Normal, Style: font.Italic}, robotoregularitalic.TypeFace, robotoregularitalic.Data)
		register(font.Font{Weight: font.Medium, Style: font.Regular}, robotomedium.TypeFace, robotomedium.Data)
		register(font.Font{Weight: font.Medium, Style: font.Italic}, robotomediumitalic.TypeFace, robotomediumitalic.Data)
		register(font.Font{Weight: font.Bold, Style: font.Regular}, robotobold.TypeFace, robotobold.Data)
		register(font.Font{Weight: font.Bold, Style: font.Italic}, robotobolditalic.TypeFace, robotobolditalic.Data)
		register(font.Font{Weight: font.Black, Style: font.Regular}, robotoblack.TypeFace, robotoblack.Data)
		register(font.Font{Weight: font.Black, Style: font.Italic}, robotoblackitalic.TypeFace, robotoblackitalic.Data)
	})
	return collection
}

func register(fnt font.Font, Typeface string, fontByte []byte) {
	face, err := opentype.Parse(fontByte)
	if err != nil {
		panic(fmt.Errorf("failed to parse font: %v", err))
	}
	fnt.Typeface = font.Typeface(Typeface)
	collection = append(collection, font.FontFace{Font: fnt, Face: face})
}
