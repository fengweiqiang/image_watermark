# image_watermark
##Image watermark.Add text watermark to pictures

#Quick Start

## Install

```bash
go get github.com/fengweiqiang/image_watermark
```

## Usage and Examples
```go
w := &image_watermark.Watermark{
    Text:      "Hello, 世界",
    X:         20,
    Y:         30,
    FontSize:  10,
    DPI:       75,
    FontColor: color.Black,
    Quality:   50,
}
or
w := image_watermark.NewWatermark("Hello, 世界")
w.X= 25
w.Y= 30
```

###load font library
Available from C:\Windows\Fonts
```go
fontBytes, _ := ioutil.ReadFile(inFilePath)
w.LoadFontBytes(fontBytes)
```


###write to file
```go
w.TextWatermarkToFile("Hello.png","世界.png")
```

#comparison

Original image                     | Dst image 
-----------------------------------|----------------------------------------
![srcImage](Hello.png) | ![dstImage](世界.png) 
