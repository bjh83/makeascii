package grayscaler

import(
	"image"
	. "makeascii/contrastor"
)

func ConcurrentGetImage(pic image.Image) *Picture {
	bounds := pic.Bounds()
	gray := image.NewGray(bounds)
	thread := func(channel chan int) {
		for y := range channel {
			for x := 0; x < bounds.Dx(); x++ {
				pixel := pic.At(x, y)
				gray.Set(x, y, pixel)
			}
		}
	}
	channel1 := make(chan int)
	channel2 := make(chan int)
	channel3 := make(chan int)
	go thread(channel1)
	go thread(channel2)
	go thread(channel3)
	for y := 0; y < bounds.Dy(); y++ {
		select {
		case channel1 <- y:
		case channel2 <- y:
		case channel3 <- y:
		}
	}
	close(channel1)
	close(channel2)
	close(channel3)
	return average(Contrast(gray, 30))
}

