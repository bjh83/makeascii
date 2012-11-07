package asciizer

import(
	"testing"
	. "makeascii/grayscaler"
)

func TestMakeAscii(t *testing.T) {
	picture := New(2, 2)
	picture.Set(0, 0, 0x00)
	picture.Set(1, 0, 0x60)
	picture.Set(0, 1, 0x80)
	picture.Set(1, 1, 0xff)
	newString := MakeAscii(picture)
	if newString != " +\n=#\n" {
		t.Error("MakeAscii() does not generate the text as expected")
	}
}

