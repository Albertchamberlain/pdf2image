package main

import (
	"fmt"
	"os"

	"github.com/gographics/imagick/imagick"
)

// convertPDFToImage 将 PDF 转换为图片文件
func convertPDFToImage(inputPath, outputPath string) error {
	imagick.Initialize()      // 初始化 imagick 库
	defer imagick.Terminate() // 程序结束时释放资源

	// 创建 PDF 物件
	pdf := imagick.NewMagickWand()
	defer pdf.Destroy()

	// 读取 PDF 文件
	if err := pdf.ReadImage(inputPath); err != nil {
		return err
	}

	// 获取 PDF 页面数量
	totalPages := pdf.GetNumberImages()

	// 循环遍历 PDF 页面并转换为图片文件
	for i := 0; i < int(totalPages); i++ {
		// 设置 PDF 页面迭代器位置
		if bool1 := pdf.SetIteratorIndex(i); bool1 != false {
			continue
		}

		// 创建图片文件物件
		image := imagick.NewMagickWand()
		defer image.Destroy()

		// 将 PDF 页面转换为图片文件
		if err := image.ReadImageBlob(pdf.GetImageBlob()); err != nil {
			continue
		}

		// 设置图片文件名称
		filename := fmt.Sprintf("%s/page%d.jpg", outputPath, i+1)

		// 将图片文件保存到输出目录
		if err := image.WriteImage(filename); err != nil {
			continue
		}

		fmt.Printf("Page %d saved as %s\n", i+1, filename)
	}

	return nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: pdf2image input.pdf output_directory/")
		return
	}

	inputPath := os.Args[1]
	outputPath := os.Args[2]

	err := convertPDFToImage(inputPath, outputPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Conversion failed:", err)
		os.Exit(1)
	}

	fmt.Println("Conversion complete.")
}
