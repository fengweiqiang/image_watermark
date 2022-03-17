package image_watermark

import (
	"bytes"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
)

type Watermark struct {
	//watermark text content
	Text string
	//font library,Default Microsoft Yahei Regular
	fontBytes *truetype.Font
	//watermark location
	X, Y int
	//font size
	FontSize float64
	//SetDPI sets the screen resolution in dots per inch.
	DPI float64
	//font color
	FontColor color.Color
	// Options are the encoding parameters.
	// Quality ranges from 1 to 100 inclusive, higher is better.
	Quality int
}

func NewWatermark(text string) *Watermark {
	return &Watermark{
		Text:      text,
		FontSize:  20,
		DPI:       75,
		FontColor: color.Black,
		Quality:   50,
	}
}

//load font library
func (w *Watermark) LoadFontBytes(fontLibrary []byte) error {
	fontLibraryBytes, err := freetype.ParseFont(fontLibrary)
	if err != nil {
		return err
	}
	w.fontBytes = fontLibraryBytes
	return nil
}

//write bytes
func (w *Watermark) TextWatermark(inFilePath string) ([]byte, error) {
	imagebytes, err := ioutil.ReadFile(inFilePath)
	if err != nil {
		return nil, err
	}
	filePathImg, imgtype, err := image.Decode(bytes.NewReader(imagebytes))
	if err != nil {
		return nil, err
	}
	filePathImgRGBA := image.NewRGBA(filePathImg.Bounds())
	if w.fontBytes == nil {
		font, _ := freetype.ParseFont(FontBytes)
		w.fontBytes = font
	}
	f := freetype.NewContext()
	f.SetDPI(w.DPI)           //设置DPI
	f.SetFont(w.fontBytes)    //设置字体
	f.SetFontSize(w.FontSize) //设置字号
	f.SetClip(filePathImg.Bounds())
	f.SetDst(filePathImgRGBA)
	f.SetSrc(image.NewUniform(w.FontColor)) //设置颜色

	draw.Draw(filePathImgRGBA, filePathImg.Bounds(), filePathImg, image.Point{}, draw.Src)
	_, err = f.DrawString(w.Text, freetype.Pt(w.X, w.Y))
	if err != nil {
		return nil, err
	}
	buf := bytes.Buffer{}
	if imgtype == "jpeg" {
		err := jpeg.Encode(&buf, filePathImgRGBA, &jpeg.Options{w.Quality})
		if err != nil {
			return nil, err
		}
	} else if imgtype == "png" {
		err := png.Encode(&buf, filePathImgRGBA)
		if err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}

//write to file
func (w *Watermark) TextWatermarkToFile(inFilePath, outFilePath string) error {
	watermarkBytes, err := w.TextWatermark(inFilePath)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(outFilePath, watermarkBytes, 0666)
}
