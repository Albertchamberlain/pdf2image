package pdf2image

import (
	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:   "pdf2image",
		Short: "Convert pdf to image",
	}
)

func init() {
	RootCmd.AddCommand(pdf2image.NewPdf2ImageCmd())
}

func main() {
	RootCmd.Execute()
}
