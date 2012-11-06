package main

import(
	"runtime"
	"time"
	"fmt"
	"os"
	"image/jpeg"
	. "../grayscaler"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide one fileName...")
		return
	}
	fileIn, fileErr := os.Open(os.Args[1])
	if fileErr != nil {
		fmt.Println(fileErr.Error())
		return
	}
	image, imageErr := jpeg.Decode(fileIn)
	if imageErr != nil {
		fmt.Println(imageErr.Error())
	}
	runtime.GOMAXPROCS(4)
	start := time.Now()
	GetImage(image)
	fmt.Println(time.Since(start))
	start = time.Now()
	ConcurrentGetImage(image)
	fmt.Println(time.Since(start))
}

