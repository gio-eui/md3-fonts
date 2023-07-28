package ubuntu

import (
	"fmt"
	"gioui.org/font"
	"gioui.org/font/opentype"
	"gioui.org/text"
	"github.com/gio-eui/md3-fonts/fonts/ubuntu/ubuntubold"
	"github.com/gio-eui/md3-fonts/fonts/ubuntu/ubuntubolditalic"
	"github.com/gio-eui/md3-fonts/fonts/ubuntu/ubuntulight"
	"github.com/gio-eui/md3-fonts/fonts/ubuntu/ubuntulightitalic"
	"github.com/gio-eui/md3-fonts/fonts/ubuntu/ubuntumedium"
	"github.com/gio-eui/md3-fonts/fonts/ubuntu/ubuntumediumitalic"
	"github.com/gio-eui/md3-fonts/fonts/ubuntu/ubunturegular"
	"github.com/gio-eui/md3-fonts/fonts/ubuntu/ubunturegularitalic"
	"sync"
)

var (
	once       sync.Once
	collection []text.FontFace
)

func Collection() []font.FontFace {
	once.Do(func() {
		register(font.Font{Weight: font.Light, Style: font.Regular}, ubuntulight.TypeFace, ubuntulight.Data)
		register(font.Font{Weight: font.Light, Style: font.Italic}, ubuntulightitalic.TypeFace, ubuntulightitalic.Data)
		register(font.Font{Weight: font.Normal, Style: font.Regular}, ubunturegular.TypeFace, ubunturegular.Data)
		register(font.Font{Weight: font.Normal, Style: font.Italic}, ubunturegularitalic.TypeFace, ubunturegularitalic.Data)
		register(font.Font{Weight: font.Medium, Style: font.Regular}, ubuntumedium.TypeFace, ubuntumedium.Data)
		register(font.Font{Weight: font.Medium, Style: font.Italic}, ubuntumediumitalic.TypeFace, ubuntumediumitalic.Data)
		register(font.Font{Weight: font.Bold, Style: font.Regular}, ubuntubold.TypeFace, ubuntubold.Data)
		register(font.Font{Weight: font.Bold, Style: font.Italic}, ubuntubolditalic.TypeFace, ubuntubolditalic.Data)
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
