package functions

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
	"time"
)

func Base64(path string) string {
	imgFile, err := os.Open(path) // a QR code image

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer imgFile.Close()

	// create a new buffer base on file size
	fInfo, _ := imgFile.Stat()
	var size = fInfo.Size()
	buf := make([]byte, size)

	// read file content into buffer
	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)

	// convert the buffer bytes to base64 string - use buf.Bytes() for new image
	imgBase64Str := base64.StdEncoding.EncodeToString(buf)

	return imgBase64Str
}

func Base64ToImage(base64Image string, directory string) (imagePath, image string, err error) {
	idx := strings.Index(base64Image, ";base64,")

	if idx < 0 {
		panic("InvalidImage")
	}

	ImageType := base64Image[11:idx]

	unbased, err := base64.StdEncoding.DecodeString(base64Image[idx+8:])

	if err != nil {
		panic("Cannot decode b64")
	}

	r := bytes.NewReader(unbased)

	if _, err := os.Stat(directory); os.IsNotExist(err) {
		_ = os.Mkdir(directory, os.ModePerm)
	}

	filename := fmt.Sprintf("%d.%s", time.Now().UnixMicro(), ImageType)
	fileDirectory := fmt.Sprintf("%s/%s", directory, filename)

	switch ImageType {
	case "png":
		src, err := png.Decode(r)
		src.ColorModel().Convert(color.Black)
		if err != nil {
			panic("Bad png")
		}

		f, err := os.OpenFile(fileDirectory, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			panic("Cannot open file")
		}

		png.Encode(f, src)
	case "jpeg":
		im, err := jpeg.Decode(r)
		if err != nil {
			panic("Bad jpeg")
		}

		f, err := os.OpenFile(fileDirectory, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			panic("Cannot open file")
		}

		jpeg.Encode(f, im, nil)
	case "gif":
		im, err := gif.Decode(r)
		if err != nil {
			panic("Bad gif")
		}

		f, err := os.OpenFile(fileDirectory, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			panic("Cannot open file")
		}

		gif.Encode(f, im, nil)
	}

	image = filename
	imagePath = fileDirectory

	return
}
