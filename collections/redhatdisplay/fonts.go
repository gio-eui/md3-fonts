package redhatdisplay

import (
	"fmt"
	"gioui.org/font"
	"gioui.org/font/opentype"
	"gioui.org/text"
	"github.com/gio-eui/md3-fonts/fonts/redhatdisplay/redhatdisplayblack"
	"github.com/gio-eui/md3-fonts/fonts/redhatdisplay/redhatdisplayblackitalic"
	"github.com/gio-eui/md3-fonts/fonts/redhatdisplay/redhatdisplaybold"
	"github.com/gio-eui/md3-fonts/fonts/redhatdisplay/redhatdisplaybolditalic"
	"github.com/gio-eui/md3-fonts/fonts/redhatdisplay/redhatdisplayextrabold"
	"github.com/gio-eui/md3-fonts/fonts/redhatdisplay/redhatdisplayextrabolditalic"
	"github.com/gio-eui/md3-fonts/fonts/redhatdisplay/redhatdisplaylight"
	"github.com/gio-eui/md3-fonts/fonts/redhatdisplay/redhatdisplaylightitalic"
	"github.com/gio-eui/md3-fonts/fonts/redhatdisplay/redhatdisplaymedium"
	"github.com/gio-eui/md3-fonts/fonts/redhatdisplay/redhatdisplaymediumitalic"
	"github.com/gio-eui/md3-fonts/fonts/redhatdisplay/redhatdisplayregular"
	"github.com/gio-eui/md3-fonts/fonts/redhatdisplay/redhatdisplayregularitalic"
	"github.com/gio-eui/md3-fonts/fonts/redhatdisplay/redhatdisplaysemibold"
	"github.com/gio-eui/md3-fonts/fonts/redhatdisplay/redhatdisplaysemibolditalic"
	"sync"
)

var (
	once       sync.Once
	collection []text.FontFace
)

func Collection() []font.FontFace {
	once.Do(func() {
		register(font.Font{Weight: font.Light, Style: font.Regular}, redhatdisplaylight.TypeFace, redhatdisplaylight.Data)
		register(font.Font{Weight: font.Light, Style: font.Italic}, redhatdisplaylightitalic.TypeFace, redhatdisplaylightitalic.Data)
		register(font.Font{Weight: font.Normal, Style: font.Regular}, redhatdisplayregular.TypeFace, redhatdisplayregular.Data)
		register(font.Font{Weight: font.Normal, Style: font.Italic}, redhatdisplayregularitalic.TypeFace, redhatdisplayregularitalic.Data)
		register(font.Font{Weight: font.Medium, Style: font.Regular}, redhatdisplaymedium.TypeFace, redhatdisplaymedium.Data)
		register(font.Font{Weight: font.Medium, Style: font.Italic}, redhatdisplaymediumitalic.TypeFace, redhatdisplaymediumitalic.Data)
		register(font.Font{Weight: font.SemiBold, Style: font.Regular}, redhatdisplaysemibold.TypeFace, redhatdisplaysemibold.Data)
		register(font.Font{Weight: font.SemiBold, Style: font.Italic}, redhatdisplaysemibolditalic.TypeFace, redhatdisplaysemibolditalic.Data)
		register(font.Font{Weight: font.Bold, Style: font.Regular}, redhatdisplaybold.TypeFace, redhatdisplaybold.Data)
		register(font.Font{Weight: font.Bold, Style: font.Italic}, redhatdisplaybolditalic.TypeFace, redhatdisplaybolditalic.Data)
		register(font.Font{Weight: font.Bold, Style: font.Regular}, redhatdisplayextrabold.TypeFace, redhatdisplayextrabold.Data)
		register(font.Font{Weight: font.Bold, Style: font.Italic}, redhatdisplayextrabolditalic.TypeFace, redhatdisplayextrabolditalic.Data)
		register(font.Font{Weight: font.Black, Style: font.Regular}, redhatdisplayblack.TypeFace, redhatdisplayblack.Data)
		register(font.Font{Weight: font.Black, Style: font.Italic}, redhatdisplayblackitalic.TypeFace, redhatdisplayblackitalic.Data)
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
