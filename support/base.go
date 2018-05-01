package support

import (
	"image/png"
	"os"
	"image"
	"bytes"
	"time"
	"github.com/tebeka/selenium"
	"fmt"
)

var driver selenium.WebDriver

//WDInit retorna uma instancia do WebDriver
func WDInit() selenium.WebDriver{
	var err error
		caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome"})
		driver, err = selenium.NewRemote(caps,"")

		if(err != nil){
			fmt.Println("Erro ao instanciar o driver.", err.Error())
		}

		driver.SetImplicitWaitTimeout(time.Second * 10)
		driver.ResizeWindow("note", 1280, 800)

		return driver
}

//SaveImage salva uma imagem 
func SaveImage(foto []byte, name string){
	
	img, _, _ := image.Decode(bytes.NewReader(foto))

	out, err := os.Create("./log/screenshots/" + name + ".png")

	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}

	err = png.Encode(out,img)
  
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}