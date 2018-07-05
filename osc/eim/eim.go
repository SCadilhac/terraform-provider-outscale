package eim

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/terraform-providers/terraform-provider-outscale/osc"
	"github.com/terraform-providers/terraform-provider-outscale/osc/handler"
)

//EIM the name of the api for url building
const EIM = "eim"

//Client manages the EIM API
type Client struct {
	client *osc.Client
	API    Service
}

// NewEIMClient return a client to operate EIM resources
func NewEIMClient(config osc.Config) (*Client, error) {
	s := &v4.Signer{
		Credentials: credentials.NewStaticCredentials(config.Credentials.AccessKey,
			config.Credentials.SecretKey, ""),
	}

	u, err := url.Parse(fmt.Sprintf(osc.DefaultBaseURL, EIM, config.Credentials.Region))
	if err != nil {
		return nil, err
	}

	config.Target = EIM
	config.BaseURL = u
	config.UserAgent = osc.UserAgent
	config.Client = &http.Client{}

	c := osc.Client{
		Config:                config,
		Signer:                s,
		MarshalHander:         handler.URLLBUEncodeMarshalHander,
		BuildRequestHandler:   handler.BuildURLEncodedRequest,
		UnmarshalHandler:      handler.UnmarshalXML,
		UnmarshalErrorHandler: handler.UnmarshalLBUErrorHandler,
		SetHeaders:            handler.SetHeaders,
	}

	f := &Client{client: &c,
		API: Operations{client: &c},
	}
	return f, nil
}
