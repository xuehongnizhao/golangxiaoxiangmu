package v1

import (
	"testing"
)

func TestGetImagesLocal(t *testing.T) {

	c := NewClient("192.168.10.127")

	images, err := c.GetImagesLocal()
	if err != nil {
		t.Fatal(err)
	}

	for _, img := range images {
		t.Log(img)
	}

}
