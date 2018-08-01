package v1

import (
	"common/base"
	"encoding/json"
	"fmt"

	"io/ioutil"

	"net/http"
)

type ImageSummary struct {

	// containers
	// Required: true
	Containers int64 `json:"Containers"`

	// created
	// Required: true
	Created int64 `json:"Created"`

	// Id
	// Required: true
	ID string `json:"Id"`

	// labels
	// Required: true
	Labels map[string]string `json:"Labels"`

	// parent Id
	// Required: true
	ParentID string `json:"ParentId"`

	// repo digests
	// Required: true
	RepoDigests []string `json:"RepoDigests"`

	// repo tags
	// Required: true
	RepoTags []string `json:"RepoTags"`

	// shared size
	// Required: true
	SharedSize int64 `json:"SharedSize"`

	// size
	// Required: true
	Size int64 `json:"Size"`

	// virtual size
	// Required: true
	VirtualSize int64 `json:"VirtualSize"`
}

func (c *Client) GetImagesLocal() ([]ImageSummary, error) {
	url := fmt.Sprintf("http://%s/platform_repository/v1/images", c.Server)

	req, err := http.NewRequest("GET", url, nil)
	response, err := http.DefaultClient.Do(req)
	defer response.Body.Close()
	if err != nil {
		return make([]ImageSummary, 0), err
	}

	if response.StatusCode != 200 {
		return make([]ImageSummary, 0), fmt.Errorf("获取注册服务失败，返回状态码为：[ %s ]", response.StatusCode)
	}

	result := base.ApiResult{}

	bys, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return make([]ImageSummary, 0), err
	}

	err = json.Unmarshal(bys, &result)
	if err != nil {
		return make([]ImageSummary, 0), err
	}

	if !result.Success {
		return make([]ImageSummary, 0), fmt.Errorf(result.Error.Msg)
	}

	list := make([]ImageSummary, 0)
	bys, _ = json.Marshal(result.Result["images"])
	json.Unmarshal(bys, &list)
	return list, nil

}
