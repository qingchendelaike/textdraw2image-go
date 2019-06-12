package src

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"io/ioutil"

	"github.com/disintegration/imaging"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

func main() {
	HandleUserImage()
}

// HandleUserImage paste user image onto background
func HandleUserImage() (string, error) {
	m, err := imaging.Open("test.png")
	if err != nil {
		fmt.Printf("open file failed")
	}

	bm, err := imaging.Open("bg.png")
	if err != nil {
		fmt.Printf("open file failed")
	}

	// 图片按比例缩放
	dst := imaging.Resize(m, 200, 200, imaging.Lanczos)
	// 将图片粘贴到背景图的固定位置
	result := imaging.Overlay(bm, dst, image.Pt(120, 140), 1)
	writeOnImage(result)

	fileName := fmt.Sprintf("%d.jpg", 1234)
	err = imaging.Save(result, fileName)
	if err != nil {
		return "", err
	}

	return fileName, nil
}

var dpi = flag.Float64("dpi", 256, "screen resolution")

func writeOnImage(target *image.NRGBA) {
	c := freetype.NewContext()

	c.SetDPI(*dpi)
	c.SetClip(target.Bounds())
	c.SetDst(target)
	c.SetHinting(font.HintingFull)

	// 设置文字颜色、字体、字大小
	c.SetSrc(image.NewUniform(color.RGBA{R: 240, G: 240, B: 245, A: 180}))
	c.SetFontSize(32)
	fontFam, err := getFontFamily()
	if err != nil {
		fmt.Println("get font family error")
	}
	c.SetFont(fontFam)

	pt := freetype.Pt(0, 200)

	_, err = c.DrawString("我是水印", pt)
	if err != nil {
		fmt.Printf("draw error: %v \n", err)
	}

}

func getFontFamily() (*truetype.Font, error) {
	// 这里需要读取中文字体，否则中文文字会变成方格
	fontBytes, err := ioutil.ReadFile("./fonts/FZYanSJW_Xian.ttf")
	if err != nil {
		fmt.Println("read file error:", err)
		return &truetype.Font{}, err
	}

	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		fmt.Println("parse font error:", err)
		return &truetype.Font{}, err
	}

	return f, err
}
