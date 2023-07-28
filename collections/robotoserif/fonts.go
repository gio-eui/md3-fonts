package robotoserif

import (
	"fmt"
	"gioui.org/font"
	"gioui.org/font/opentype"
	"gioui.org/text"
	"github.com/gio-eui/md3-fonts/fonts/robotoserif/robotoserifblack"
	"github.com/gio-eui/md3-fonts/fonts/robotoserif/robotoserifblackitalic"
	"github.com/gio-eui/md3-fonts/fonts/robotoserif/robotoserifbold"
	"github.com/gio-eui/md3-fonts/fonts/robotoserif/robotoserifbolditalic"
	"github.com/gio-eui/md3-fonts/fonts/robotoserif/robotoserifextrabold"
	"github.com/gio-eui/md3-fonts/fonts/robotoserif/robotoserifextrabolditalic"
	"github.com/gio-eui/md3-fonts/fonts/robotoserif/robotoserifextralight"
	"github.com/gio-eui/md3-fonts/fonts/robotoserif/robotoserifextralightitalic"
	"github.com/gio-eui/md3-fonts/fonts/robotoserif/robotoseriflight"
	"github.com/gio-eui/md3-fonts/fonts/robotoserif/robotoseriflightitalic"
	"github.com/gio-eui/md3-fonts/fonts/robotoserif/robotoserifmedium"
	"github.com/gio-eui/md3-fonts/fonts/robotoserif/robotoserifmediumitalic"
	"github.com/gio-eui/md3-fonts/fonts/robotoserif/robotoserifregular"
	"github.com/gio-eui/md3-fonts/fonts/robotoserif/robotoserifregularitalic"
	"github.com/gio-eui/md3-fonts/fonts/robotoserif/robotoserifsemibold"
	"github.com/gio-eui/md3-fonts/fonts/robotoserif/robotoserifsemibolditalic"
	"github.com/gio-eui/md3-fonts/fonts/robotoserif/robotoserifthin"
	"github.com/gio-eui/md3-fonts/fonts/robotoserif/robotoserifthinitalic"
	"sync"
)

var (
	once       sync.Once
	collection []text.FontFace
)

func Collection() []font.FontFace {
	once.Do(func() {
		register(font.Font{Weight: font.Thin, Style: font.Regular}, robotoserifthin.TypeFace, robotoserifthin.Data)
		register(font.Font{Weight: font.Thin, Style: font.Italic}, robotoserifthinitalic.TypeFace, robotoserifthinitalic.Data)
		register(font.Font{Weight: font.ExtraLight, Style: font.Regular}, robotoserifextralight.TypeFace, robotoserifextralight.Data)
		register(font.Font{Weight: font.ExtraLight, Style: font.Italic}, robotoserifextralightitalic.TypeFace, robotoserifextralightitalic.Data)
		register(font.Font{Weight: font.Light, Style: font.Regular}, robotoseriflight.TypeFace, robotoseriflight.Data)
		register(font.Font{Weight: font.Light, Style: font.Italic}, robotoseriflightitalic.TypeFace, robotoseriflightitalic.Data)
		register(font.Font{Weight: font.Normal, Style: font.Regular}, robotoserifregular.TypeFace, robotoserifregular.Data)
		register(font.Font{Weight: font.Normal, Style: font.Italic}, robotoserifregularitalic.TypeFace, robotoserifregularitalic.Data)
		register(font.Font{Weight: font.Medium, Style: font.Regular}, robotoserifmedium.TypeFace, robotoserifmedium.Data)
		register(font.Font{Weight: font.Medium, Style: font.Italic}, robotoserifmediumitalic.TypeFace, robotoserifmediumitalic.Data)
		register(font.Font{Weight: font.SemiBold, Style: font.Regular}, robotoserifsemibold.TypeFace, robotoserifsemibold.Data)
		register(font.Font{Weight: font.SemiBold, Style: font.Italic}, robotoserifsemibolditalic.TypeFace, robotoserifsemibolditalic.Data)
		register(font.Font{Weight: font.Bold, Style: font.Regular}, robotoserifbold.TypeFace, robotoserifbold.Data)
		register(font.Font{Weight: font.Bold, Style: font.Italic}, robotoserifbolditalic.TypeFace, robotoserifbolditalic.Data)
		register(font.Font{Weight: font.Bold, Style: font.Regular}, robotoserifextrabold.TypeFace, robotoserifextrabold.Data)
		register(font.Font{Weight: font.Bold, Style: font.Italic}, robotoserifextrabolditalic.TypeFace, robotoserifextrabolditalic.Data)
		register(font.Font{Weight: font.Black, Style: font.Regular}, robotoserifblack.TypeFace, robotoserifblack.Data)
		register(font.Font{Weight: font.Black, Style: font.Italic}, robotoserifblackitalic.TypeFace, robotoserifblackitalic.Data)
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
