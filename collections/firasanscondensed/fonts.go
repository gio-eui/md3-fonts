package firasanscondensed

import (
	"fmt"
	"gioui.org/font"
	"gioui.org/font/opentype"
	"gioui.org/text"
	"github.com/gio-eui/md3-fonts/fonts/firasanscondensed/firasanscondensedblack"
	"github.com/gio-eui/md3-fonts/fonts/firasanscondensed/firasanscondensedblackitalic"
	"github.com/gio-eui/md3-fonts/fonts/firasanscondensed/firasanscondensedbold"
	"github.com/gio-eui/md3-fonts/fonts/firasanscondensed/firasanscondensedbolditalic"
	"github.com/gio-eui/md3-fonts/fonts/firasanscondensed/firasanscondensedextrabold"
	"github.com/gio-eui/md3-fonts/fonts/firasanscondensed/firasanscondensedextrabolditalic"
	"github.com/gio-eui/md3-fonts/fonts/firasanscondensed/firasanscondensedextralight"
	"github.com/gio-eui/md3-fonts/fonts/firasanscondensed/firasanscondensedextralightitalic"
	"github.com/gio-eui/md3-fonts/fonts/firasanscondensed/firasanscondensedlight"
	"github.com/gio-eui/md3-fonts/fonts/firasanscondensed/firasanscondensedlightitalic"
	"github.com/gio-eui/md3-fonts/fonts/firasanscondensed/firasanscondensedmedium"
	"github.com/gio-eui/md3-fonts/fonts/firasanscondensed/firasanscondensedmediumitalic"
	"github.com/gio-eui/md3-fonts/fonts/firasanscondensed/firasanscondensedregular"
	"github.com/gio-eui/md3-fonts/fonts/firasanscondensed/firasanscondensedregularitalic"
	"github.com/gio-eui/md3-fonts/fonts/firasanscondensed/firasanscondensedsemibold"
	"github.com/gio-eui/md3-fonts/fonts/firasanscondensed/firasanscondensedsemibolditalic"
	"github.com/gio-eui/md3-fonts/fonts/firasanscondensed/firasanscondensedthin"
	"github.com/gio-eui/md3-fonts/fonts/firasanscondensed/firasanscondensedthinitalic"
	"sync"
)

var (
	once       sync.Once
	collection []text.FontFace
)

func Collection() []font.FontFace {
	once.Do(func() {
		register(font.Font{Weight: font.Thin, Style: font.Regular}, firasanscondensedthin.TypeFace, firasanscondensedthin.Data)
		register(font.Font{Weight: font.Thin, Style: font.Italic}, firasanscondensedthinitalic.TypeFace, firasanscondensedthinitalic.Data)
		register(font.Font{Weight: font.ExtraLight, Style: font.Regular}, firasanscondensedextralight.TypeFace, firasanscondensedextralight.Data)
		register(font.Font{Weight: font.ExtraLight, Style: font.Italic}, firasanscondensedextralightitalic.TypeFace, firasanscondensedextralightitalic.Data)
		register(font.Font{Weight: font.Light, Style: font.Regular}, firasanscondensedlight.TypeFace, firasanscondensedlight.Data)
		register(font.Font{Weight: font.Light, Style: font.Italic}, firasanscondensedlightitalic.TypeFace, firasanscondensedlightitalic.Data)
		register(font.Font{Weight: font.Normal, Style: font.Regular}, firasanscondensedregular.TypeFace, firasanscondensedregular.Data)
		register(font.Font{Weight: font.Normal, Style: font.Italic}, firasanscondensedregularitalic.TypeFace, firasanscondensedregularitalic.Data)
		register(font.Font{Weight: font.Medium, Style: font.Regular}, firasanscondensedmedium.TypeFace, firasanscondensedmedium.Data)
		register(font.Font{Weight: font.Medium, Style: font.Italic}, firasanscondensedmediumitalic.TypeFace, firasanscondensedmediumitalic.Data)
		register(font.Font{Weight: font.SemiBold, Style: font.Regular}, firasanscondensedsemibold.TypeFace, firasanscondensedsemibold.Data)
		register(font.Font{Weight: font.SemiBold, Style: font.Italic}, firasanscondensedsemibolditalic.TypeFace, firasanscondensedsemibolditalic.Data)
		register(font.Font{Weight: font.Bold, Style: font.Regular}, firasanscondensedbold.TypeFace, firasanscondensedbold.Data)
		register(font.Font{Weight: font.Bold, Style: font.Italic}, firasanscondensedbolditalic.TypeFace, firasanscondensedbolditalic.Data)
		register(font.Font{Weight: font.Bold, Style: font.Regular}, firasanscondensedextrabold.TypeFace, firasanscondensedextrabold.Data)
		register(font.Font{Weight: font.Bold, Style: font.Italic}, firasanscondensedextrabolditalic.TypeFace, firasanscondensedextrabolditalic.Data)
		register(font.Font{Weight: font.Black, Style: font.Regular}, firasanscondensedblack.TypeFace, firasanscondensedblack.Data)
		register(font.Font{Weight: font.Black, Style: font.Italic}, firasanscondensedblackitalic.TypeFace, firasanscondensedblackitalic.Data)
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
