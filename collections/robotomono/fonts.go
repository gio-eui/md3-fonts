package robotomono

import (
	"fmt"
	"gioui.org/font"
	"gioui.org/font/opentype"
	"gioui.org/text"
	"github.com/gio-eui/md3-fonts/fonts/robotomono/robotomonobold"
	"github.com/gio-eui/md3-fonts/fonts/robotomono/robotomonobolditalic"
	"github.com/gio-eui/md3-fonts/fonts/robotomono/robotomonoextralight"
	"github.com/gio-eui/md3-fonts/fonts/robotomono/robotomonoextralightitalic"
	"github.com/gio-eui/md3-fonts/fonts/robotomono/robotomonolight"
	"github.com/gio-eui/md3-fonts/fonts/robotomono/robotomonolightitalic"
	"github.com/gio-eui/md3-fonts/fonts/robotomono/robotomonomedium"
	"github.com/gio-eui/md3-fonts/fonts/robotomono/robotomonomediumitalic"
	"github.com/gio-eui/md3-fonts/fonts/robotomono/robotomonoregular"
	"github.com/gio-eui/md3-fonts/fonts/robotomono/robotomonoregularitalic"
	"github.com/gio-eui/md3-fonts/fonts/robotomono/robotomonosemibold"
	"github.com/gio-eui/md3-fonts/fonts/robotomono/robotomonosemibolditalic"
	"github.com/gio-eui/md3-fonts/fonts/robotomono/robotomonothin"
	"github.com/gio-eui/md3-fonts/fonts/robotomono/robotomonothinitalic"
	"sync"
)

var (
	once       sync.Once
	collection []text.FontFace
)

func Collection() []font.FontFace {
	once.Do(func() {
		register(font.Font{Weight: font.Thin, Style: font.Regular}, robotomonothin.TypeFace, robotomonothin.Data)
		register(font.Font{Weight: font.Thin, Style: font.Italic}, robotomonothinitalic.TypeFace, robotomonothinitalic.Data)
		register(font.Font{Weight: font.ExtraLight, Style: font.Regular}, robotomonoextralight.TypeFace, robotomonoextralight.Data)
		register(font.Font{Weight: font.ExtraLight, Style: font.Italic}, robotomonoextralightitalic.TypeFace, robotomonoextralightitalic.Data)
		register(font.Font{Weight: font.Light, Style: font.Regular}, robotomonolight.TypeFace, robotomonolight.Data)
		register(font.Font{Weight: font.Light, Style: font.Italic}, robotomonolightitalic.TypeFace, robotomonolightitalic.Data)
		register(font.Font{Weight: font.Normal, Style: font.Regular}, robotomonoregular.TypeFace, robotomonoregular.Data)
		register(font.Font{Weight: font.Normal, Style: font.Italic}, robotomonoregularitalic.TypeFace, robotomonoregularitalic.Data)
		register(font.Font{Weight: font.Medium, Style: font.Regular}, robotomonomedium.TypeFace, robotomonomedium.Data)
		register(font.Font{Weight: font.Medium, Style: font.Italic}, robotomonomediumitalic.TypeFace, robotomonomediumitalic.Data)
		register(font.Font{Weight: font.SemiBold, Style: font.Regular}, robotomonosemibold.TypeFace, robotomonosemibold.Data)
		register(font.Font{Weight: font.SemiBold, Style: font.Italic}, robotomonosemibolditalic.TypeFace, robotomonosemibolditalic.Data)
		register(font.Font{Weight: font.Bold, Style: font.Regular}, robotomonobold.TypeFace, robotomonobold.Data)
		register(font.Font{Weight: font.Bold, Style: font.Italic}, robotomonobolditalic.TypeFace, robotomonobolditalic.Data)
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
