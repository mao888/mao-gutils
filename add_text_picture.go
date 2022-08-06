package utils

import (
	"Project/global"
	"Project/utils/freetype-go/freetype"
	"github.com/fogleman/gg"
	"github.com/shamsher31/goimgtype"
	"go-utils/md5"
	"go.uber.org/zap"
	"golang.org/x/image/draw"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

// file "sss.jpg/sss.png"
// color.RGBA{255, 0, 0, 0}
// 生成图片的名称, resImage := "resImage.png"
// fontName := "fdbsjw.ttf"

var (
	fontName string  = "uploads/front/fdbsjw.ttf"
	text     string  = "uploads/front/text.png"
	fontSize float64 = 60
)

// ImageProcess 参数传入
func ImageProcess(filePath string, textLabel string) (url string, err error) {
	nums := strings.Split(filePath, "/")
	file := nums[len(nums)-1]

	src, err := os.Open(filePath)
	if err != nil {
		zap.L().Error("Open params failed", zap.Error(err))
		return url, err
	}
	datatype, err := imgtype.Get(filePath)
	if err != nil {
		zap.L().Error("imgType params failed", zap.Error(err))
		return url, err
	} else {
		// 根据文件类型执行响应的操作
		switch datatype {
		case `image/jpeg`:
			img, err := jpeg.Decode(src)
			if err != nil {
				zap.L().Error("jpegDecode params failed", zap.Error(err))
				return url, err
			}

			c := img.Bounds()

			// 字体大小
			if len(strconv.Itoa(c.Dx())) == 4 || len(strconv.Itoa(c.Dy())) == 4 {
				fontSize = 80
			} else if len(strconv.Itoa(c.Dx())) < 4 && len(strconv.Itoa(c.Dy())) < 4 {
				fontSize = 15
			} else if len(strconv.Itoa(c.Dx())) > 4 && len(strconv.Itoa(c.Dy())) > 4 {
				fontSize = 300
			}

			// 颜色反转
			x := int(float64(c.Dx()) - float64(c.Dx())*0.50) // 获取要添加字体的x位置
			y := int(float64(c.Dy()) - float64(c.Dy())*0.15) // 获取要添加字体的y位置
			colorRgb := c.At(x, y)
			r, g, b, a := colorRgb.RGBA()
			r_uint8 := uint8(r >> 8)
			g_uint8 := uint8(g >> 8)
			b_uint8 := uint8(b >> 8)
			a_uint8 := uint8(a >> 8)
			r_uint8 = 255 - r_uint8
			g_uint8 = 255 - g_uint8
			b_uint8 = 255 - b_uint8

			col := color.RGBA{R: r_uint8, G: g_uint8, B: b_uint8, A: a_uint8}

			res, err := addLabel(img, textLabel, x, y, col, fontSize, fontName)

			if err != nil {
				zap.L().Error("jpegAddLabel params failed", zap.Error(err))
				return url, err
			}

			fileName := strings.Split(file, ".")
			// 读取文件名并加密
			name := strings.TrimSuffix(fileName[0], ".jpg")
			name = md5.MD5V([]byte(name))
			// 拼接新文件名
			filename := name + "_" + time.Now().Format("20060102150405") + ".jpg"
			f, err := os.Create(global.KBB_CONFIG.Local.Path + "/" + filename)
			if err != nil {
				zap.L().Error("jpegCreate params failed", zap.Error(err))
				return url, err
			}
			defer f.Close()

			err = jpeg.Encode(f, res, &jpeg.Options{Quality: 100}) // 100表示质量最好的
			if err != nil {
				zap.L().Error("jpegCreate params failed", zap.Error(err))
				return url, err
			}

			url = global.KBB_CONFIG.Local.Path + "/" + filename

		case `image/png`:
			img, err := png.Decode(src)
			if err != nil {
				zap.L().Error("pngDecode params failed", zap.Error(err))
				return url, err
			}

			c := img.Bounds()

			// 字体大小
			if len(strconv.Itoa(c.Dx())) == 4 || len(strconv.Itoa(c.Dy())) == 4 {
				fontSize = 80
			} else if len(strconv.Itoa(c.Dx())) < 4 && len(strconv.Itoa(c.Dy())) < 4 {
				fontSize = 15
			} else if len(strconv.Itoa(c.Dx())) > 4 && len(strconv.Itoa(c.Dy())) > 4 {
				fontSize = 300
			}

			// 颜色反转
			x := int(float64(c.Dx()) - float64(c.Dx())*0.50) // 获取要添加字体的x位置
			y := int(float64(c.Dy()) - float64(c.Dy())*0.15) // 获取要添加字体的y位置
			colorRgb := c.At(x, y)
			r, g, b, a := colorRgb.RGBA()
			r_uint8 := uint8(r >> 8)
			g_uint8 := uint8(g >> 8)
			b_uint8 := uint8(b >> 8)
			a_uint8 := uint8(a >> 8)
			r_uint8 = 255 - r_uint8
			g_uint8 = 255 - g_uint8
			b_uint8 = 255 - b_uint8

			col := color.RGBA{R: r_uint8, G: g_uint8, B: b_uint8, A: a_uint8}

			res, err := addLabel(img, textLabel, x, y, col, fontSize, fontName)

			if err != nil {
				zap.L().Error("pngDecode params failed", zap.Error(err))
				return url, err
			}

			fileName := strings.Split(file, ".")
			// 读取文件名并加密
			name := strings.TrimSuffix(fileName[0], ".png")
			name = md5.MD5V([]byte(name))
			// 拼接新文件名
			filename := name + "_" + time.Now().Format("20060102150405") + ".png"

			f, err := os.Create(global.KBB_CONFIG.Local.Path + "/" + filename)
			if err != nil {
				zap.L().Error("pngCreate params failed", zap.Error(err))
				return url, err
			}
			defer f.Close()

			err = png.Encode(f, res)
			if err != nil {
				zap.L().Error("pngEncode params failed", zap.Error(err))
				log.Println(err)
				return url, err
			}

			url = global.KBB_CONFIG.Local.Path + "/" + filename
		}
	}
	return url, err
}

// addLabel 图片生成
func addLabel(img image.Image, label string, x, y int, fontColor color.Color, size float64, fontPath string) (image.Image, error) {
	bound := img.Bounds()
	// 创建一个新的图片
	rgba := image.NewRGBA(image.Rect(0, 0, bound.Dx(), bound.Dy()))
	// 读取字体
	fontBytes, err := ioutil.ReadFile(fontPath)
	if err != nil {
		zap.L().Error("ReadFile params failed", zap.Error(err))
		return rgba, err
	}

	myFont, err := freetype.ParseFont(fontBytes)
	if err != nil {
		zap.L().Error("ParseFont params failed", zap.Error(err))
		return rgba, err
	}

	draw.Draw(rgba, rgba.Bounds(), img, bound.Min, draw.Src)
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(myFont)
	c.SetFontSize(size)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	uni := image.NewUniform(fontColor)
	c.SetSrc(uni)
	c.SetHinting(freetype.NoHinting)

	// 在指定的位置显示
	pt := freetype.Pt(x, y+int(c.PointToFix32(size)>>6))
	if _, err := c.DrawString(label, pt); err != nil {
		zap.L().Error("DrawString params failed", zap.Error(err))
		return rgba, err
	}

	return rgba, nil
}

// Watermark 图片水印白底
func Watermark(filePath string, textLabel string) (url string, err error) {
	buyBox := "uploads/buy-box/" // + 盒子名称 注意：需要购买后的路径储存自己盒柜中
	nums := strings.Split(filePath, "/")
	file := nums[len(nums)-1]
	// 读取文件后缀
	ext := path.Ext(filePath)
	// 读取文件名并加密
	name := strings.TrimSuffix(file, ext)
	name = md5.MD5V([]byte(name + time.Now().Format("20060102150405")))
	// 拼接新文件名
	filename := name + "_" + time.Now().Format("20060102150405") + ext

	src, err := os.Open(filePath)
	if err != nil {
		zap.L().Error("Open params failed", zap.Error(err))
		return url, err
	}
	datatype, err := imgtype.Get(filePath)
	if err != nil {
		zap.L().Error("imgType params failed", zap.Error(err))
		return url, err
	} else {
		// 根据文件类型执行响应的操作
		switch datatype {
		case `image/jpeg`:
			img, err := jpeg.Decode(src)
			if err != nil {
				zap.L().Error("jpegEncode params failed", zap.Error(err))
				log.Println(err)
				return url, err
			}
			c := img.Bounds()

			dc := gg.NewContext(0, 0)                // 设置画笔颜色为黑色
			sWidth, _ := dc.MeasureString(textLabel) // 测算字符串将在画布中占用的宽与长
			dc = gg.NewContext(c.Dx(), 90)
			dc.SetRGB255(255, 255, 255)
			dc.Clear()
			if err := dc.LoadFontFace(fontName, fontSize); err != nil { // 从本地加载字体文件
				log.Println(err)
				return filePath, err
			}
			dc.SetRGB(0, 0, 0)
			dc.DrawString(textLabel, (float64(c.Dy())-sWidth)/2, 60) // 直接将文字贴入画布中
			err = dc.SavePNG(text)
			if err != nil {
				log.Println("SavePNG", err)
				return filePath, err
			} // 保存水印图片
			wmb, err := os.Open(text) // 打开水印图片
			if err != nil {
				zap.L().Error("Open params failed", zap.Error(err))
				return url, err
			}
			watermark, err := png.Decode(wmb) // 解析水印图片
			if err != nil {
				zap.L().Error("pngEncode params failed", zap.Error(err))
				log.Println(err)
				return url, err
			}
			defer wmb.Close()
			//把水印写到右下角，并向0坐标各偏移10个像素
			offset := image.Pt(img.Bounds().Dx()-watermark.Bounds().Dx(), img.Bounds().Dy()-watermark.Bounds().Dy())
			b := img.Bounds()
			m := image.NewNRGBA(b)
			draw.Draw(m, b, img, image.ZP, draw.Src)
			draw.Draw(m, watermark.Bounds().Add(offset), watermark, image.ZP, draw.Over)
			buyBox += filename // 拼接新的路径
			imgw, err := os.Create(buyBox)
			if err != nil {
				log.Println("Create::", err)
				return filePath, err
			}
			err = jpeg.Encode(imgw, m, &jpeg.Options{Quality: 100})
			if err != nil {
				log.Println("Encode::", err)
				return filePath, err
			}
			defer imgw.Close()
		case `image/png`:
			img, err := png.Decode(src)
			if err != nil {
				zap.L().Error("pngEncode params failed", zap.Error(err))
				log.Println(err)
				return url, err
			}
			c := img.Bounds()

			dc := gg.NewContext(0, 0)                // 设置画笔颜色为黑色
			sWidth, _ := dc.MeasureString(textLabel) // 测算字符串将在画布中占用的宽与长
			dc = gg.NewContext(c.Dx(), 90)

			dc.SetRGB255(255, 255, 255)
			dc.Clear()
			if err := dc.LoadFontFace(fontName, fontSize); err != nil { // 从本地加载字体文件
				log.Println(err)
				return filePath, err
			}
			dc.SetRGB(0, 0, 0)
			dc.DrawString(textLabel, (float64(c.Dy())-sWidth)/2, 60) // 直接将文字贴入画布中
			err = dc.SavePNG(text)
			if err != nil {
				log.Println("SavePNG", err)
				return filePath, err
			} // 保存水印图片
			wmb, err := os.Open(text) // 打开水印图片
			if err != nil {
				zap.L().Error("Open params failed", zap.Error(err))
				return url, err
			}
			watermark, err := png.Decode(wmb) // 解析水印图片
			if err != nil {
				zap.L().Error("pngEncode params failed", zap.Error(err))
				log.Println(err)
				return url, err
			}
			defer wmb.Close()
			//把水印写到右下角，并向0坐标各偏移10个像素
			offset := image.Pt(img.Bounds().Dx()-watermark.Bounds().Dx(), img.Bounds().Dy()-watermark.Bounds().Dy())
			b := img.Bounds()
			m := image.NewNRGBA(b)
			draw.Draw(m, b, img, image.ZP, draw.Src)
			draw.Draw(m, watermark.Bounds().Add(offset), watermark, image.ZP, draw.Over)

			// png转换成jpg
			newImg := image.NewRGBA(img.Bounds())
			draw.Draw(newImg, newImg.Bounds(), &image.Uniform{C: color.White}, image.Point{}, draw.Src)
			draw.Draw(newImg, newImg.Bounds(), img, img.Bounds().Min, draw.Over)

			buyBox += filename // 拼接新的路径
			imgw, err := os.Create(buyBox)
			if err != nil {
				log.Println("Create::", err)
				return filePath, err
			}
			err = jpeg.Encode(imgw, m, &jpeg.Options{Quality: 100})
			if err != nil {
				log.Println("Encode::", err)
				return filePath, err
			}
			defer imgw.Close()
		}
	}
	defer src.Close()
	// 只有执行成功后才会返回新的路径
	return buyBox, err
}

func CompressJpg(imagePath string) error {
	fileInfo, err := os.Stat(imagePath)
	if err != nil {
		log.Println("os.Stat", err)
		return err
	}
	size := float64(fileInfo.Size()) / 1048576
	imgQuality := 0
	if size < 2 {
		imgQuality = 90
	} else if size >= 2 && size <= 3.5 {
		imgQuality = 75
	} else if size > 3.5 && size <= 5 {
		imgQuality = 60
	} else if size > 5 && size <= 10 {
		imgQuality = 30
	} else {
		imgQuality = 10
	}
	// 需要压缩
	imgFile, err := os.Open(imagePath)
	if err != nil {
		log.Println("os.Open::", err)
		return err
	}
	defer imgFile.Close()

	// 更具图片大小选择压缩质量
	jpgImg, err := jpeg.Decode(imgFile)
	if err != nil {
		log.Println("jpeg.Decode::", err)
		return err
	}

	//保存到新文件中
	newfile, err := os.Create(imagePath)
	if err != nil {
		log.Println("Create::", err)
		return err
	}
	defer newfile.Close()

	// &jpeg.Options{Quality: 10} 图片压缩质量
	err = jpeg.Encode(newfile, jpgImg, &jpeg.Options{Quality: imgQuality})
	if err != nil {
		log.Println("Encode::", err)
		return err
	}
	return nil
}

func CompressPng(imagePath string) error {
	fileInfo, err := os.Stat(imagePath)
	if err != nil {
		log.Println("os.Stat", err)
		return err
	}
	size := float64(fileInfo.Size()) / 1048576
	imgQuality := 0
	if size < 2 {
		imgQuality = 90
	} else if size >= 2 && size <= 3.5 {
		imgQuality = 75
	} else if size > 3.5 && size <= 5 {
		imgQuality = 60
	} else if size > 5 && size <= 10 {
		imgQuality = 30
	} else {
		imgQuality = 10
	}
	// 需要压缩
	imgFile, err := os.Open(imagePath)
	if err != nil {
		log.Println("os.Open::", err)
		return err
	}
	defer imgFile.Close()

	// 更具图片大小选择压缩质量
	pngImg, err := png.Decode(imgFile)
	if err != nil {
		log.Println("png.Decode::", err)
		return err
	}
	// png转换成jpg
	newImg := image.NewRGBA(pngImg.Bounds())
	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{C: color.White}, image.Point{}, draw.Src)
	draw.Draw(newImg, newImg.Bounds(), pngImg, pngImg.Bounds().Min, draw.Over)

	//保存到新文件中
	newfile, err := os.Create(imagePath)
	if err != nil {
		log.Println("Create::", err)
		return err
	}
	defer newfile.Close()

	// &jpeg.Options{Quality: 10} 图片压缩质量
	err = jpeg.Encode(newfile, pngImg, &jpeg.Options{Quality: imgQuality})
	if err != nil {
		log.Println("Encode::", err)
		return err
	}
	return nil
}
