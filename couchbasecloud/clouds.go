package couchbasecloud

import (
	"net/http"
	"net/url"
	"strconv"
)

type CloudsList struct {
	Cursor Cursor  `json:"cursor"`
	Data   []Cloud `json:"data"`
}

type Cloud struct {
	Id                 string `json:"id"`
	Name               string `json:"name"`
	Provider           string `json:"provider"`
	Region             string `json:"region"`
	Status             string `json:"status"`
	VirtualNetworkCIDR string `json:"virtualNetworkCidr"`
	VirtualNetworkID   string `json:"virtualNetworkId"`
}

type ListCloudsOptions struct {
	Page    int     `json:"page"`
	PerPage int     `json:"perPage"`
	SortBy  *string `json:"sortBy"`
}

const listCloudsUrl = "/clouds"

func (client *CouchbaseCloudClient) ListClouds(options *ListCloudsOptions) (*CloudsList, error) {
	cloudsUrl := client.BaseURL + client.getApiEndpoint(listCloudsUrl)

	if options != nil {
		setListCloudsParams(&cloudsUrl, *options)
	}

	req, err := http.NewRequest(http.MethodGet, cloudsUrl, nil)
	if err != nil {
		return nil, err
	}

	res := CloudsList{}

	if err := client.sendRequest(req, &res, true); err != nil {
		return nil, err
	}

	return &res, nil
}

func setListCloudsParams(urlStr *string, options ListCloudsOptions) {
	params := url.Values{}

	if options.SortBy != nil {
		params.Add("sortBy", *options.SortBy)
	}

	if options.Page != 0 {
		params.Add("page", strconv.Itoa(options.Page))
	}

	if options.PerPage != 0 {
		params.Add("perPage", strconv.Itoa(options.PerPage))
	}

	if urlParams := params.Encode(); urlParams != "" {
		*urlStr += "?" + urlParams
	}
}
