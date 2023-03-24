package main

import (
	"fmt"
	"os"

	"github.com/gographics/imagick/imagick"
	"github.com/spf13/cobra"
)

func convertPDFToImage(inputPath, outputPath string) error {
	// 初始化 imagick 库和 magickWand 对象
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	pdf := imagick.NewMagickWand()
	defer pdf.Destroy()

	// 读取输入 PDF 文件并检查错误
	if err := pdf.ReadImage(inputPath); err != nil {
		return err
	}

	// 将 PDF 的每一页转换为图像
	totalPages := pdf.GetNumberImages()
	for i := 0; i < int(totalPages); i++ {
		if bool1 := pdf.SetIteratorIndex(i); bool1 != false {
			continue
		}

		// 将 PDF 页面转换为 PNG 图像
		if err := mw.ReadImageBlob(pdf.GetImageBlob()); err != nil {
			continue
		}

		// 将 PNG 图像文件保存到输出路径中
		filename := fmt.Sprintf("%s_%d.png", outputPath, i+1)
		if err := mw.WriteImage(filename); err != nil {
			continue
		}
		fmt.Printf("Page %d saved as %s\n", i+1, filename)
	}
	return nil
}

func run(cmd *cobra.Command, args []string) {
	fmt.Println("Converting PDF to image...")

	if len(args) != 2 {
		fmt.Println("requires an input PDF file and an output image directory")
		os.Exit(1)
	}

	inputPath := args[0]
	outputPath := args[1]

	err := convertPDFToImage(inputPath, outputPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println("Conversion complete.")
}
