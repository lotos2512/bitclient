package bitclient

import (
	"errors"
	"fmt"
)

var SSHKeyAlreadyExists = errors.New("This SSH key already provides access to this repository")

type SetRepositorySSHKeyRequest struct {
	Key        Key    `json:"key"`
	Permission string `json:"permission"`
}

type Key struct {
	Text string `json:"text"`
}

func (bc *BitClient) SetRepositorySSHKey(projectKey, repositorySlug string, params SetRepositorySSHKeyRequest) (err error) {
	rError := new(ErrorResponse)
	_, err = bc.sling.New().Post(fmt.Sprintf("/rest/keys/1.0/projects/%s/repos/%s/ssh", projectKey, repositorySlug)).BodyJSON(params).Receive(nil, rError)
	if err != nil {
		return
	}

	if len(rError.Errors) > 0 {
		err = errors.New(rError.Errors[0].Message)
	}
	
	return

}
