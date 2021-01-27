package bitclient

import (
	"github.com/dghubble/sling"
	"net/http"
)

type BitClient struct {
	sling  *sling.Sling
	client *http.Client
}

type BitConf struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Url      string `yaml:"url"`
}

func NewBitClient(conf BitConf) *BitClient {
	client := &http.Client{}
	return &BitClient{
		sling: sling.New().Base(conf.Url).Set("Accept", "application/json").SetBasicAuth(conf.UserName, conf.Password).Client(client),
	}
}
