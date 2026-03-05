package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/nfnt/resize"
)

func main() {
	// 1. 定义命令行参数
	inputPath := flag.String("i", "", "输入图片的路径 (必填)")
	outputWidth := flag.Int("w", 100, "输出宽度 (字符数)")
	threshold := flag.Int("t", 40000, "亮度阈值 (0-65535, 越小点越少)")
	invert := flag.Bool("v", false, "是否反转颜色 (点亮浅色区域)")

	flag.Parse()

	// 2. 检查必填项
	if *inputPath == "" {
		fmt.Println("用法: braille-tool -i <图片路径> [-w 宽度] [-t 阈值] [-v]")
		flag.PrintDefaults()
		return
	}

	// 3. 打开并解码
	file, err := os.Open(*inputPath)
	if err != nil {
		fmt.Printf("错误: 无法打开文件: %v\n", err)
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Printf("错误: 无法解码图片: %v\n", err)
		return
	}

	// 4. 缩放图片
	// 每个盲文符宽 2 像素
	pixelWidth := uint(*outputWidth * 2)
	img = resize.Resize(pixelWidth, 0, img, resize.Lanczos3)

	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y

	// 5. 盲文点位映射
	dots := []struct{ dx, dy, bit uint }{
		{0, 0, 0}, {0, 1, 1}, {0, 2, 2},
		{1, 0, 3}, {1, 1, 4}, {1, 2, 5},
		{0, 3, 6}, {1, 3, 7},
	}

	// 6. 遍历转换
	for y := 0; y < h; y += 4 {
		line := ""
		for x := 0; x < w; x += 2 {
			var mask uint = 0
			for _, d := range dots {
				px, py := x+int(d.dx), y+int(d.dy)
				if px < w && py < h {
					if shouldLightUp(img.At(px, py), uint32(*threshold), *invert) {
						mask |= (1 << d.bit)
					}
				}
			}
			line += string(rune(0x2800 + mask))
		}
		fmt.Println(line)
	}
}

// shouldLightUp 判断该点是否应该显示
func shouldLightUp(c color.Color, threshold uint32, invert bool) bool {
	r, g, b, _ := c.RGBA()
	lum := uint32(0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b))
	
	if invert {
		return lum > threshold
	}
	return lum < threshold
}
