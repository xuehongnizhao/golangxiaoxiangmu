package v1

import (
	"common/base"
	"encoding/json"
	"fmt"

	"io/ioutil"

	"net/http"
)

type RepoTags struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

func (c *Client) GetRepos() ([]RepoTags, error) {
	url := fmt.Sprintf("http://%s/platform_repository/v1/repos", c.Server)

	req, err := http.NewRequest("GET", url, nil)
	response, err := http.DefaultClient.Do(req)
	defer response.Body.Close()
	if err != nil {
		return make([]RepoTags, 0), err
	}

	if response.StatusCode != 200 {
		return make([]RepoTags, 0), fmt.Errorf("获取注册服务失败，返回状态码为：[ %s ]", response.StatusCode)
	}

	result := base.ApiResult{}

	bys, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return make([]RepoTags, 0), err
	}

	err = json.Unmarshal(bys, &result)
	if err != nil {
		return make([]RepoTags, 0), err
	}

	if !result.Success {
		return make([]RepoTags, 0), fmt.Errorf(result.Error.Msg)
	}

	list := make([]RepoTags, 0)
	bys, _ = json.Marshal(result.Result["repos"])
	json.Unmarshal(bys, &list)
	return list, nil

}
