package v1

import (
	"bytes"
	"common/base"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type PushImageInfo struct {
	FilePath  string
	ImageName string
}

//推送镜像至镜像仓库
func (c *Client) PushImageToRespository(img PushImageInfo) error {

	url := fmt.Sprintf("http://%s/platform_repository/v1/images", c.Server)

	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)

	fw, err := w.CreateFormFile("file", filepath.Base(img.FilePath))
	if err != nil {
		return err
	}

	fd, err := os.Open(img.FilePath)
	if err != nil {
		return err
	}
	defer fd.Close()

	// Write file field from file to upload
	_, err = io.Copy(fw, fd)
	if err != nil {
		return err
	}
	w.WriteField("imageName", img.ImageName)
	w.Close()

	req, err := http.NewRequest("POST", url, buf)

	req.Header.Set("Content-Type", w.FormDataContentType())
	response, err := http.DefaultClient.Do(req)
	defer response.Body.Close()
	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return fmt.Errorf("镜像推送失败，返回状态码为：[ %d ]", response.StatusCode)
	}

	result := base.ApiResult{}

	bys, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bys, &result)
	if err != nil {
		return err
	}

	if !result.Success {
		return fmt.Errorf(result.Error.Msg)
	}

	return nil
}
