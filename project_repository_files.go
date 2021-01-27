package bitclient

import (
	"fmt"
	"io/ioutil"
)

func (bc *BitClient) GetFile(projectKey string, repositorySlug string, path string) (data []byte, err error) {
	resp, err := bc.Request(
		fmt.Sprintf("/projects/%s/repos/%s/raw/%s", projectKey, repositorySlug, path),
		nil,
	)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}
