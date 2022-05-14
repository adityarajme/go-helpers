package golang_helpers

import (
	"github.com/disintegration/imaging"
	"image"
	"image/color"
	"io"
	"os"
)

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func GenerateFAQUploadThumbs(image_name string) {

	reader, err0 := os.Open("/file_location/" + image_name)
	CheckErr(err0)

	im, _, err111 := image.DecodeConfig(reader)
	CheckErr(err111)

	//Generating THumbs
	img, err22 := imaging.Open("/file_location/" + image_name)
	CheckErr(err22)

	//Image size to 50% - Normal Thumb
	float_imgwidth := float32(im.Width) * 0.5
	float_imgheight := float32(im.Height) * 0.5

	//Image size to 30%
	float_imgwidth30 := float32(im.Width) * 0.3
	float_imgheight30 := float32(im.Height) * 0.3

	//Generating 50%
	thumb := imaging.Thumbnail(img, int(float_imgwidth), int(float_imgheight), imaging.CatmullRom)
	dst := imaging.New(int(float_imgwidth), int(float_imgheight), color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, thumb, image.Pt(0, 0))
	errimg := imaging.Save(dst, "/file_location/thumb_"+image_name)
	CheckErr(errimg)

	//Generating 30%
	thumb = imaging.Thumbnail(img, int(float_imgwidth30), int(float_imgheight30), imaging.CatmullRom)
	dst = imaging.New(int(float_imgwidth30), int(float_imgheight30), color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, thumb, image.Pt(0, 0))
	errimg = imaging.Save(dst, "/file_location/thumb30_"+image_name)
	CheckErr(errimg)

	//Generating square 200 x 200
	thumb = imaging.Thumbnail(img, 200, 200, imaging.CatmullRom)
	dst = imaging.New(200, 200, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, thumb, image.Pt(0, 0))
	errimg = imaging.Save(dst, "/file_location/thumb_sq200_"+image_name)
	CheckErr(errimg)

	reader.Close()
}

func CopyFile(source string, desti string) {
	srcFile, err := os.Open(source)
	CheckErr(err)

	destFile, err := os.Create(desti) // creates if file doesn't exist
	CheckErr(err)

	_, err = io.Copy(destFile, srcFile) // check first var for number of bytes copied
	CheckErr(err)

	err = destFile.Sync()

	destFile.Close()
	srcFile.Close()
	CheckErr(err)
}
