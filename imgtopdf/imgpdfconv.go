package imgtopdf

import (
	// "fmt"
	// "golibcustom/gokyeuerun"
	"fmt"
	"image"
	"log"
	"os"

	"github.com/signintech/gopdf"
	// heap "golibcustom/priorityqueue"
)

func getImageBytes(filepath string) []byte {
	b, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	return b
}

func ImgToPdf() {
	homeDir, _ := os.UserHomeDir()
	tempFilePath1 := fmt.Sprintf("%v/%v", homeDir, "try1.png")
	tempFilePath2 := fmt.Sprintf("%v/%v", homeDir, "try2.png")
	file1, _ := os.Open(tempFilePath1)
	file2, _ := os.Open(tempFilePath2)
	image1, _, _ := image.DecodeConfig(file1)
	image2, _, _ := image.DecodeConfig(file2)
	file1.Close()
	file2.Close()

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
	pdf.AddPage()

	//use path
	pdf.Image(tempFilePath1, 0, 0, nil)
	pdf.SetY(float64(image1.Height))
	pdf.SetNewXY(0, 0, float64(image2.Height))
	//use image holder by []byte
	imgH1, err := gopdf.ImageHolderByBytes(getImageBytes(tempFilePath2))
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.ImageByHolder(imgH1, 0, float64(image1.Height), nil)
	fmt.Printf("next x %v y %v \n", pdf.GetX(), pdf.GetY())

	pdf.WritePdf("image2.pdf")
	fmt.Printf("final x %v y %v \n", pdf.GetX(), pdf.GetY())
}
