package grayscaler

import(
	"image"
	"image/color"
	. "../lindatrast"
)

const(
	charHeight = 10
	charWidth = 5
)

type Picture struct {
	Data []uint8
	Width, Height int
}

func New(width, height int) *Picture {
	return &Picture{make([]byte, width * height), width, height}
}

func (pic *Picture) Get(x, y int) uint8 {
	return pic.Data[x + y * pic.Width]
}

func (pic *Picture) Set(x, y int, color uint8) {
	pic.Data[x + y * pic.Width] = color
}

func GetImage(pic image.Image) *Picture {
	bounds := pic.Bounds()
	gray := image.NewGray(bounds)
	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			color := pic.At(x, y)
			gray.Set(x, y, color)
		}
	}
	return average(Contrast(gray, 30))
}

func average(oldPic *image.Gray) *Picture {
	bounds := oldPic.Bounds()
	newWidth := charWidth * (bounds.Dx() / charWidth)
	newHeight := charHeight * (bounds.Dy() / charHeight)
	newPic := New(newWidth / charWidth, newHeight / charHeight)
	for x := 0; x < newWidth - charWidth; x += charWidth {
		for y := 0; y < newHeight - charHeight; y += charHeight {
			sum := 0
			for i := x; i < x + charWidth; i++ {
				for j := y; j < y + charHeight; j++ {
					sum += int(oldPic.At(i, j).(color.Gray).Y)
				}
			}
			newPic.Set(x / charWidth, y / charHeight, uint8(sum / (charWidth * charHeight)))
		}
	}
	return newPic
}

