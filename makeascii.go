package main

import(
	"image/jpeg"
	"os"
	"fmt"
	. "./grayscaler"
	. "./asciizer"
)

func main() {
	fileIn, errIn := os.Open(os.Args[1])
	if errIn != nil {
		fmt.Println(errIn.Error())
		return
	}
	pic, imageErr := jpeg.Decode(fileIn)
	if imageErr != nil {
		fmt.Println(imageErr.Error())
	}
	fmt.Print(MakeAscii(GetImage(pic)))
}

