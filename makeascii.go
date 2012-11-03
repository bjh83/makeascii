package main

import(
	"image/jpeg"
	"image/png"
	"image"
	"os"
	"fmt"
	"regexp"
	. "./grayscaler"
	. "./asciizer"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please select one picture to makeascii...")
		return
	}
	fileName := os.Args[1]
	isJpeg, _ := regexp.MatchString(".+\\.jpg", fileName)
	isPng, _ := regexp.MatchString(".+\\.png", fileName)
	var picture image.Image
	var imageErr error
	fileIn, errIn := os.Open(fileName)
	if errIn != nil {
		fmt.Println(errIn.Error())
		return
	}
	if isJpeg {
		picture, imageErr = jpeg.Decode(fileIn)
	} else if isPng {
		picture, imageErr = png.Decode(fileIn)
	} else {
		fmt.Println("File type is not supported...")
		return
	}
	if imageErr != nil {
		fmt.Println(imageErr.Error())
		return
	}
	fmt.Print(MakeAscii(GetImage(picture)))
}

