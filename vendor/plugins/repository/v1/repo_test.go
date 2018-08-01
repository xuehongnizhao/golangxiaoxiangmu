package v1

import (
	"testing"
)

func TestGetRepos(t *testing.T) {

	c := NewClient("192.168.10.127")

	repos, err := c.GetRepos()
	if err != nil {
		t.Fatal(err)
	}

	for _, repo := range repos {
		t.Log(repo)
	}

}
