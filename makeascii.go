package main

import(
	"image/jpeg"
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
	result, _ := regexp.MatchString(".+\\.jpg", fileName)
	if !result {
		fmt.Println("Format of input is not supported...")
		return
	}
	fileIn, errIn := os.Open(fileName)
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

