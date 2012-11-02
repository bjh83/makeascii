package asciizer

import(
	. "../grayscaler"
)

func MakeAscii(picture *Picture) string {
	out := ""
	for y := 0; y < picture.Height; y++ {
		for x := 0; x < picture.Width; x++ {
			colorVal := picture.Get(x, y)
			if colorVal >= 213 {
				out += "#"
			} else if colorVal >= 171 {
				out += "="
			} else if colorVal >= 128 {
				out += "*"
			} else if colorVal >= 85 {
				out += ":"
			} else if colorVal >= 43 {
				out += "."
			} else {
				out += " "
			}
		}
		out += "\n"
	}
	return out
}

