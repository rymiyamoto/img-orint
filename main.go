package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/gographics/imagick.v3/imagick"
)

func main() {
	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()
	name := "orient_3.jpeg" // file名を適宜変える

	imageBytes, err := ioutil.ReadFile("in/" + name)
	if err != nil {
		fmt.Println("failed ReadFile")
		return
	}

	err = mw.ReadImageBlob(imageBytes)
	if err != nil {
		fmt.Println("failed ReadImageBlob")
		return
	}

	fmt.Printf("Orientation: %v\n", mw.GetImageOrientation())

	err = mw.AutoOrientImage()
	if err != nil {
		fmt.Println("failed AutoOrientImage")
		return
	}

	fmt.Printf("Orientation: %v\n", mw.GetImageOrientation())

	err = ioutil.WriteFile("out/"+name, mw.GetImagesBlob(), 0644)
	if err != nil {
		fmt.Println("failed WriteFile")
		return
	}

	fmt.Println("＼(^o^)／")
}
