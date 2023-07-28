package sourcesanspro

import (
	"fmt"
	"gioui.org/font"
	"gioui.org/font/opentype"
	"gioui.org/text"
	"github.com/gio-eui/md3-fonts/fonts/sourcesanspro/sourcesansproblack"
	"github.com/gio-eui/md3-fonts/fonts/sourcesanspro/sourcesansproblackitalic"
	"github.com/gio-eui/md3-fonts/fonts/sourcesanspro/sourcesansprobold"
	"github.com/gio-eui/md3-fonts/fonts/sourcesanspro/sourcesansprobolditalic"
	"github.com/gio-eui/md3-fonts/fonts/sourcesanspro/sourcesansproextralight"
	"github.com/gio-eui/md3-fonts/fonts/sourcesanspro/sourcesansproextralightitalic"
	"github.com/gio-eui/md3-fonts/fonts/sourcesanspro/sourcesansprolight"
	"github.com/gio-eui/md3-fonts/fonts/sourcesanspro/sourcesansprolightitalic"
	"github.com/gio-eui/md3-fonts/fonts/sourcesanspro/sourcesansproregular"
	"github.com/gio-eui/md3-fonts/fonts/sourcesanspro/sourcesansproregularitalic"
	"github.com/gio-eui/md3-fonts/fonts/sourcesanspro/sourcesansprosemibold"
	"github.com/gio-eui/md3-fonts/fonts/sourcesanspro/sourcesansprosemibolditalic"
	"sync"
)

var (
	once       sync.Once
	collection []text.FontFace
)

func Collection() []font.FontFace {
	once.Do(func() {
		register(font.Font{Weight: font.ExtraLight, Style: font.Regular}, sourcesansproextralight.TypeFace, sourcesansproextralight.Data)
		register(font.Font{Weight: font.ExtraLight, Style: font.Italic}, sourcesansproextralightitalic.TypeFace, sourcesansproextralightitalic.Data)
		register(font.Font{Weight: font.Light, Style: font.Regular}, sourcesansprolight.TypeFace, sourcesansprolight.Data)
		register(font.Font{Weight: font.Light, Style: font.Italic}, sourcesansprolightitalic.TypeFace, sourcesansprolightitalic.Data)
		register(font.Font{Weight: font.Normal, Style: font.Regular}, sourcesansproregular.TypeFace, sourcesansproregular.Data)
		register(font.Font{Weight: font.Normal, Style: font.Italic}, sourcesansproregularitalic.TypeFace, sourcesansproregularitalic.Data)
		register(font.Font{Weight: font.SemiBold, Style: font.Regular}, sourcesansprosemibold.TypeFace, sourcesansprosemibold.Data)
		register(font.Font{Weight: font.SemiBold, Style: font.Italic}, sourcesansprosemibolditalic.TypeFace, sourcesansprosemibolditalic.Data)
		register(font.Font{Weight: font.Bold, Style: font.Regular}, sourcesansprobold.TypeFace, sourcesansprobold.Data)
		register(font.Font{Weight: font.Bold, Style: font.Italic}, sourcesansprobolditalic.TypeFace, sourcesansprobolditalic.Data)
		register(font.Font{Weight: font.Black, Style: font.Regular}, sourcesansproblack.TypeFace, sourcesansproblack.Data)
		register(font.Font{Weight: font.Black, Style: font.Italic}, sourcesansproblackitalic.TypeFace, sourcesansproblackitalic.Data)
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
