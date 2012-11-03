package asciizer

import(
	. "../grayscaler"
)

func MakeAscii(picture *Picture) string {
	out := ""
	for y := 0; y < picture.Height; y++ {
		for x := 0; x < picture.Width; x++ {
			colorVal := picture.Get(x, y)
			if colorVal >= 224 {
				out += "#"
			} else if colorVal >= 192 {
				out += "X"
			} else if colorVal >= 160 {
				out += "o"
			} else if colorVal >= 128 {
				out += "="
			} else if colorVal >= 96 {
				out += "+"
			} else if colorVal >= 64 {
				out += ":"
			} else if colorVal >= 32 {
				out += "-"
			} else {
				out += " "
			}
		}
		out += "\n"
	}
	return out
}

