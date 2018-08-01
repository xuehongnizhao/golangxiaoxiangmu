package v1

import (
	"testing"
)

func TestPushImageToRespository(t *testing.T) {
	img := PushImageInfo{}
	img.FilePath = "C:\\Users\\gaoyong\\Desktop\\command-v1.tar.gz"
	img.ImageName = "command:v1"

	c := NewClient("192.168.10.127")

	err := c.PushImageToRespository(img)
	if err != nil {
		t.Fatal(err)
	}

}
