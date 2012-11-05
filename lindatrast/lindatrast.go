package lindatrast

import(
	"image"
	"image/color"
)

func getAverage(picture *image.Gray) uint8 {
	var sum uint64 = 0
	bounds := picture.Bounds()
	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			sum += uint64(picture.At(x, y).(color.Gray).Y)
		}
	}
	return uint8(sum / uint64(bounds.Dx() * bounds.Dy()))
}

func Contrast(picture *image.Gray, contrast uint8) *image.Gray {
	bounds := picture.Bounds()
	newPic := image.NewGray(bounds)
	average := getAverage(picture)
	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			pixel := picture.At(x, y).(color.Gray).Y
			if pixel > average {
				if pixel + contrast < pixel {
					pixel = 0xff
				} else {
					pixel += contrast
				}
			} else if pixel < average {
				if pixel - contrast > pixel {
					pixel = 0x00
				} else {
					pixel -= contrast
				}
			}
			newPic.Set(x, y, color.Gray{pixel})
		}
	}
	return newPic
}

