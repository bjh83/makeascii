package asciizer

import(
	. "../grayscaler"
)

func MakeAscii(picture *Picture) string {
	out := ""
	for y := 0; y < picture.Height; y++ {
		for x := 0; x < picture.Width; x++ {
			colorVal := picture.Get(x, y)
			if colorVal >= 192 {
				out += " "
			} else if colorVal >= 128 {
				out += "."
			} else if colorVal >= 64 {
				out += "o"
			} else {
				out += "#"
			}
		}
		out += "\n"
	}
	return out
}

