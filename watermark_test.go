package image_watermark_test

import (
	"github.com/fengweiqiang/image_watermark"
	"testing"
	"time"
)

func TestWatermark_TextWatermarkToFile(t *testing.T) {

	timeformat := time.Now().Format(time.RFC3339Nano)
	t.Log(timeformat)
	watermark := image_watermark.NewWatermark("当前时间：" + timeformat)

	//watermark.LoadFontBytes()
	watermark.X = 50
	watermark.Y = 30
	t.Log(watermark.TextWatermarkToFile("Hello.png", "watermark.png"))
}

func TestNewWatermark(t *testing.T) {
	w := image_watermark.NewWatermark("Hello, 世界")
	w.X = 25
	w.Y = 30
	t.Log(w.TextWatermarkToFile("Hello.png", "世界.png"))
}
