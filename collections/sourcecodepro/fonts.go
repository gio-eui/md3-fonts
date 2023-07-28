package sourcecodepro

import (
	"fmt"
	"gioui.org/font"
	"gioui.org/font/opentype"
	"gioui.org/text"
	"github.com/gio-eui/md3-fonts/fonts/sourcecodepro/sourcecodeproblack"
	"github.com/gio-eui/md3-fonts/fonts/sourcecodepro/sourcecodeproblackitalic"
	"github.com/gio-eui/md3-fonts/fonts/sourcecodepro/sourcecodeprobold"
	"github.com/gio-eui/md3-fonts/fonts/sourcecodepro/sourcecodeprobolditalic"
	"github.com/gio-eui/md3-fonts/fonts/sourcecodepro/sourcecodeproextrabold"
	"github.com/gio-eui/md3-fonts/fonts/sourcecodepro/sourcecodeproextrabolditalic"
	"github.com/gio-eui/md3-fonts/fonts/sourcecodepro/sourcecodeproextralight"
	"github.com/gio-eui/md3-fonts/fonts/sourcecodepro/sourcecodeproextralightitalic"
	"github.com/gio-eui/md3-fonts/fonts/sourcecodepro/sourcecodeprolight"
	"github.com/gio-eui/md3-fonts/fonts/sourcecodepro/sourcecodeprolightitalic"
	"github.com/gio-eui/md3-fonts/fonts/sourcecodepro/sourcecodepromedium"
	"github.com/gio-eui/md3-fonts/fonts/sourcecodepro/sourcecodepromediumitalic"
	"github.com/gio-eui/md3-fonts/fonts/sourcecodepro/sourcecodeproregular"
	"github.com/gio-eui/md3-fonts/fonts/sourcecodepro/sourcecodeproregularitalic"
	"github.com/gio-eui/md3-fonts/fonts/sourcecodepro/sourcecodeprosemibold"
	"github.com/gio-eui/md3-fonts/fonts/sourcecodepro/sourcecodeprosemibolditalic"
	"sync"
)

var (
	once       sync.Once
	collection []text.FontFace
)

func Collection() []font.FontFace {
	once.Do(func() {
		register(font.Font{Weight: font.ExtraLight, Style: font.Regular}, sourcecodeproextralight.TypeFace, sourcecodeproextralight.Data)
		register(font.Font{Weight: font.ExtraLight, Style: font.Italic}, sourcecodeproextralightitalic.TypeFace, sourcecodeproextralightitalic.Data)
		register(font.Font{Weight: font.Light, Style: font.Regular}, sourcecodeprolight.TypeFace, sourcecodeprolight.Data)
		register(font.Font{Weight: font.Light, Style: font.Italic}, sourcecodeprolightitalic.TypeFace, sourcecodeprolightitalic.Data)
		register(font.Font{Weight: font.Normal, Style: font.Regular}, sourcecodeproregular.TypeFace, sourcecodeproregular.Data)
		register(font.Font{Weight: font.Normal, Style: font.Italic}, sourcecodeproregularitalic.TypeFace, sourcecodeproregularitalic.Data)
		register(font.Font{Weight: font.Medium, Style: font.Regular}, sourcecodepromedium.TypeFace, sourcecodepromedium.Data)
		register(font.Font{Weight: font.Medium, Style: font.Italic}, sourcecodepromediumitalic.TypeFace, sourcecodepromediumitalic.Data)
		register(font.Font{Weight: font.SemiBold, Style: font.Regular}, sourcecodeprosemibold.TypeFace, sourcecodeprosemibold.Data)
		register(font.Font{Weight: font.SemiBold, Style: font.Italic}, sourcecodeprosemibolditalic.TypeFace, sourcecodeprosemibolditalic.Data)
		register(font.Font{Weight: font.Bold, Style: font.Regular}, sourcecodeprobold.TypeFace, sourcecodeprobold.Data)
		register(font.Font{Weight: font.Bold, Style: font.Italic}, sourcecodeprobolditalic.TypeFace, sourcecodeprobolditalic.Data)
		register(font.Font{Weight: font.Bold, Style: font.Regular}, sourcecodeproextrabold.TypeFace, sourcecodeproextrabold.Data)
		register(font.Font{Weight: font.Bold, Style: font.Italic}, sourcecodeproextrabolditalic.TypeFace, sourcecodeproextrabolditalic.Data)
		register(font.Font{Weight: font.Black, Style: font.Regular}, sourcecodeproblack.TypeFace, sourcecodeproblack.Data)
		register(font.Font{Weight: font.Black, Style: font.Italic}, sourcecodeproblackitalic.TypeFace, sourcecodeproblackitalic.Data)
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
