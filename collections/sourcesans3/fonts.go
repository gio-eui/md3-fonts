package sourcesans3

import (
	"fmt"
	"gioui.org/font"
	"gioui.org/font/opentype"
	"gioui.org/text"
	"github.com/gio-eui/md3-fonts/fonts/sourcesans3/sourcesans3black"
	"github.com/gio-eui/md3-fonts/fonts/sourcesans3/sourcesans3blackitalic"
	"github.com/gio-eui/md3-fonts/fonts/sourcesans3/sourcesans3bold"
	"github.com/gio-eui/md3-fonts/fonts/sourcesans3/sourcesans3bolditalic"
	"github.com/gio-eui/md3-fonts/fonts/sourcesans3/sourcesans3extrabold"
	"github.com/gio-eui/md3-fonts/fonts/sourcesans3/sourcesans3extrabolditalic"
	"github.com/gio-eui/md3-fonts/fonts/sourcesans3/sourcesans3extralight"
	"github.com/gio-eui/md3-fonts/fonts/sourcesans3/sourcesans3extralightitalic"
	"github.com/gio-eui/md3-fonts/fonts/sourcesans3/sourcesans3light"
	"github.com/gio-eui/md3-fonts/fonts/sourcesans3/sourcesans3lightitalic"
	"github.com/gio-eui/md3-fonts/fonts/sourcesans3/sourcesans3medium"
	"github.com/gio-eui/md3-fonts/fonts/sourcesans3/sourcesans3mediumitalic"
	"github.com/gio-eui/md3-fonts/fonts/sourcesans3/sourcesans3regular"
	"github.com/gio-eui/md3-fonts/fonts/sourcesans3/sourcesans3regularitalic"
	"github.com/gio-eui/md3-fonts/fonts/sourcesans3/sourcesans3semibold"
	"github.com/gio-eui/md3-fonts/fonts/sourcesans3/sourcesans3semibolditalic"
	"sync"
)

var (
	once       sync.Once
	collection []text.FontFace
)

func Collection() []font.FontFace {
	once.Do(func() {
		register(font.Font{Weight: font.ExtraLight, Style: font.Regular}, sourcesans3extralight.TypeFace, sourcesans3extralight.Data)
		register(font.Font{Weight: font.ExtraLight, Style: font.Italic}, sourcesans3extralightitalic.TypeFace, sourcesans3extralightitalic.Data)
		register(font.Font{Weight: font.Light, Style: font.Regular}, sourcesans3light.TypeFace, sourcesans3light.Data)
		register(font.Font{Weight: font.Light, Style: font.Italic}, sourcesans3lightitalic.TypeFace, sourcesans3lightitalic.Data)
		register(font.Font{Weight: font.Normal, Style: font.Regular}, sourcesans3regular.TypeFace, sourcesans3regular.Data)
		register(font.Font{Weight: font.Normal, Style: font.Italic}, sourcesans3regularitalic.TypeFace, sourcesans3regularitalic.Data)
		register(font.Font{Weight: font.Medium, Style: font.Regular}, sourcesans3medium.TypeFace, sourcesans3medium.Data)
		register(font.Font{Weight: font.Medium, Style: font.Italic}, sourcesans3mediumitalic.TypeFace, sourcesans3mediumitalic.Data)
		register(font.Font{Weight: font.SemiBold, Style: font.Regular}, sourcesans3semibold.TypeFace, sourcesans3semibold.Data)
		register(font.Font{Weight: font.SemiBold, Style: font.Italic}, sourcesans3semibolditalic.TypeFace, sourcesans3semibolditalic.Data)
		register(font.Font{Weight: font.Bold, Style: font.Regular}, sourcesans3bold.TypeFace, sourcesans3bold.Data)
		register(font.Font{Weight: font.Bold, Style: font.Italic}, sourcesans3bolditalic.TypeFace, sourcesans3bolditalic.Data)
		register(font.Font{Weight: font.Bold, Style: font.Regular}, sourcesans3extrabold.TypeFace, sourcesans3extrabold.Data)
		register(font.Font{Weight: font.Bold, Style: font.Italic}, sourcesans3extrabolditalic.TypeFace, sourcesans3extrabolditalic.Data)
		register(font.Font{Weight: font.Black, Style: font.Regular}, sourcesans3black.TypeFace, sourcesans3black.Data)
		register(font.Font{Weight: font.Black, Style: font.Italic}, sourcesans3blackitalic.TypeFace, sourcesans3blackitalic.Data)
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
