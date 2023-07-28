package roboto

import (
	"fmt"
	"gioui.org/font"
	"gioui.org/font/opentype"
	"gioui.org/text"
	"github.com/gio-eui/md3-fonts/fonts/firasans/firasansblack"
	"github.com/gio-eui/md3-fonts/fonts/firasans/firasansblackitalic"
	"github.com/gio-eui/md3-fonts/fonts/firasans/firasansbold"
	"github.com/gio-eui/md3-fonts/fonts/firasans/firasansbolditalic"
	"github.com/gio-eui/md3-fonts/fonts/firasans/firasansextrabold"
	"github.com/gio-eui/md3-fonts/fonts/firasans/firasansextrabolditalic"
	"github.com/gio-eui/md3-fonts/fonts/firasans/firasansextralight"
	"github.com/gio-eui/md3-fonts/fonts/firasans/firasansextralightitalic"
	"github.com/gio-eui/md3-fonts/fonts/firasans/firasanslight"
	"github.com/gio-eui/md3-fonts/fonts/firasans/firasanslightitalic"
	"github.com/gio-eui/md3-fonts/fonts/firasans/firasansmedium"
	"github.com/gio-eui/md3-fonts/fonts/firasans/firasansmediumitalic"
	"github.com/gio-eui/md3-fonts/fonts/firasans/firasansregular"
	"github.com/gio-eui/md3-fonts/fonts/firasans/firasansregularitalic"
	"github.com/gio-eui/md3-fonts/fonts/firasans/firasanssemibold"
	"github.com/gio-eui/md3-fonts/fonts/firasans/firasanssemibolditalic"
	"github.com/gio-eui/md3-fonts/fonts/firasans/firasansthin"
	"github.com/gio-eui/md3-fonts/fonts/firasans/firasansthinitalic"
	"sync"
)

var (
	once       sync.Once
	collection []text.FontFace
)

func Collection() []font.FontFace {
	once.Do(func() {
		register(font.Font{Weight: font.Thin, Style: font.Regular}, firasansthin.TypeFace, firasansthin.Data)
		register(font.Font{Weight: font.Thin, Style: font.Italic}, firasansthinitalic.TypeFace, firasansthinitalic.Data)
		register(font.Font{Weight: font.ExtraLight, Style: font.Regular}, firasansextralight.TypeFace, firasansextralight.Data)
		register(font.Font{Weight: font.ExtraLight, Style: font.Italic}, firasansextralightitalic.TypeFace, firasansextralightitalic.Data)
		register(font.Font{Weight: font.Light, Style: font.Regular}, firasanslight.TypeFace, firasanslight.Data)
		register(font.Font{Weight: font.Light, Style: font.Italic}, firasanslightitalic.TypeFace, firasanslightitalic.Data)
		register(font.Font{Weight: font.Normal, Style: font.Regular}, firasansregular.TypeFace, firasansregular.Data)
		register(font.Font{Weight: font.Normal, Style: font.Italic}, firasansregularitalic.TypeFace, firasansregularitalic.Data)
		register(font.Font{Weight: font.Medium, Style: font.Regular}, firasansmedium.TypeFace, firasansmedium.Data)
		register(font.Font{Weight: font.Medium, Style: font.Italic}, firasansmediumitalic.TypeFace, firasansmediumitalic.Data)
		register(font.Font{Weight: font.SemiBold, Style: font.Regular}, firasanssemibold.TypeFace, firasanssemibold.Data)
		register(font.Font{Weight: font.SemiBold, Style: font.Italic}, firasanssemibolditalic.TypeFace, firasanssemibolditalic.Data)
		register(font.Font{Weight: font.Bold, Style: font.Regular}, firasansbold.TypeFace, firasansbold.Data)
		register(font.Font{Weight: font.Bold, Style: font.Italic}, firasansbolditalic.TypeFace, firasansbolditalic.Data)
		register(font.Font{Weight: font.Bold, Style: font.Regular}, firasansextrabold.TypeFace, firasansextrabold.Data)
		register(font.Font{Weight: font.Bold, Style: font.Italic}, firasansextrabolditalic.TypeFace, firasansextrabolditalic.Data)
		register(font.Font{Weight: font.Black, Style: font.Regular}, firasansblack.TypeFace, firasansblack.Data)
		register(font.Font{Weight: font.Black, Style: font.Italic}, firasansblackitalic.TypeFace, firasansblackitalic.Data)
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
