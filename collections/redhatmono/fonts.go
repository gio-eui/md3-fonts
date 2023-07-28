package redhatmono

import (
	"fmt"
	"gioui.org/font"
	"gioui.org/font/opentype"
	"gioui.org/text"
	"github.com/gio-eui/md3-fonts/fonts/redhatmono/redhatmonobold"
	"github.com/gio-eui/md3-fonts/fonts/redhatmono/redhatmonobolditalic"
	"github.com/gio-eui/md3-fonts/fonts/redhatmono/redhatmonolight"
	"github.com/gio-eui/md3-fonts/fonts/redhatmono/redhatmonolightitalic"
	"github.com/gio-eui/md3-fonts/fonts/redhatmono/redhatmonomedium"
	"github.com/gio-eui/md3-fonts/fonts/redhatmono/redhatmonomediumitalic"
	"github.com/gio-eui/md3-fonts/fonts/redhatmono/redhatmonoregular"
	"github.com/gio-eui/md3-fonts/fonts/redhatmono/redhatmonoregularitalic"
	"github.com/gio-eui/md3-fonts/fonts/redhatmono/redhatmonosemibold"
	"github.com/gio-eui/md3-fonts/fonts/redhatmono/redhatmonosemibolditalic"
	"sync"
)

var (
	once       sync.Once
	collection []text.FontFace
)

func Collection() []font.FontFace {
	once.Do(func() {
		register(font.Font{Weight: font.Light, Style: font.Regular}, redhatmonolight.TypeFace, redhatmonolight.Data)
		register(font.Font{Weight: font.Light, Style: font.Italic}, redhatmonolightitalic.TypeFace, redhatmonolightitalic.Data)
		register(font.Font{Weight: font.Normal, Style: font.Regular}, redhatmonoregular.TypeFace, redhatmonoregular.Data)
		register(font.Font{Weight: font.Normal, Style: font.Italic}, redhatmonoregularitalic.TypeFace, redhatmonoregularitalic.Data)
		register(font.Font{Weight: font.Medium, Style: font.Regular}, redhatmonomedium.TypeFace, redhatmonomedium.Data)
		register(font.Font{Weight: font.Medium, Style: font.Italic}, redhatmonomediumitalic.TypeFace, redhatmonomediumitalic.Data)
		register(font.Font{Weight: font.SemiBold, Style: font.Regular}, redhatmonosemibold.TypeFace, redhatmonosemibold.Data)
		register(font.Font{Weight: font.SemiBold, Style: font.Italic}, redhatmonosemibolditalic.TypeFace, redhatmonosemibolditalic.Data)
		register(font.Font{Weight: font.Bold, Style: font.Regular}, redhatmonobold.TypeFace, redhatmonobold.Data)
		register(font.Font{Weight: font.Bold, Style: font.Italic}, redhatmonobolditalic.TypeFace, redhatmonobolditalic.Data)
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
