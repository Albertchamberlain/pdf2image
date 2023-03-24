package main

import (
	"fmt"
	"testing"

	"github.com/agiledragon/gomonkey"
	"github.com/gographics/imagick/imagick"
	"github.com/stretchr/testify/assert"
)

func TestConvertPDFToImage(t *testing.T) {
	// 模拟一个返回错误的ReadImage方法
	patch := gomonkey.ApplyFunc((*imagick.MagickWand).ReadImage, func(_ *imagick.MagickWand, _ string) error {
		return fmt.Errorf("error reading image")
	})
	defer patch.Reset()

	err := convertPDFToImage("test.pdf", "output/")
	assert.NotNil(t, err)
	assert.Equal(t, "error reading image", err.Error())

	// 模拟返回总页数为3的GetNumberImages方法
	patch2 := gomonkey.ApplyFunc((*imagick.MagickWand).GetNumberImages, func(_ *imagick.MagickWand) uint {
		return 3
	})
	defer patch2.Reset()

	err = convertPDFToImage("test.pdf", "output/")
	assert.Nil(t, err)
}
