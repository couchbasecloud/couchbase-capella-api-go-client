package couchbasecloud

import (
	"net/http"
	"net/url"
	"strconv"
)

type ClustersList struct {
	Cursor Cursor    `json:"cursor"`
	Data   []Cluster `json:"data"`
}

type Cluster struct {
	Id        string   `json:"id"`
	CloudId   string   `json:"cloudId"`
	Name      string   `json:"name"`
	Nodes     int      `json:"nodes"`
	ProjectId string   `json:"projectId"`
	Services  []string `json:"services"`
	TenantId  string   `json:"tenantId"`
}

type ListClustersOptions struct {
	Page      int     `json:"page"`
	PerPage   int     `json:"perPage"`
	SortBy    *string `json:"sortBy"`
	CloudId   *string `json:"cloudId"`
	ProjectId *string `json:"projectId"`
}

const listClustersUrl = "/clusters"

func (client *CouchbaseCloudClient) ListClusters(options *ListClustersOptions) (*ClustersList, error) {
	cloudsUrl := client.BaseURL + client.getApiEndpoint(listClustersUrl)

	if options != nil {
		setListClustersParams(&cloudsUrl, *options)
	}

	req, err := http.NewRequest(http.MethodGet, cloudsUrl, nil)
	if err != nil {
		return nil, err
	}

	res := ClustersList{}

	if err := client.sendRequest(req, &res, true); err != nil {
		return nil, err
	}

	return &res, nil
}

func setListClustersParams(clustersUrl *string, options ListClustersOptions) {
	params := url.Values{}

	if options.Page != 0 {
		params.Add("page", strconv.Itoa(options.Page))
	}

	if options.PerPage != 0 {
		params.Add("perPage", strconv.Itoa(options.PerPage))
	}

	if options.SortBy != nil {
		params.Add("sortBy", *options.SortBy)
	}

	if options.CloudId != nil {
		params.Add("cloudId", *options.CloudId)
	}

	if options.ProjectId != nil {
		params.Add("projectId", *options.ProjectId)
	}

	if urlParams := params.Encode(); urlParams != "" {
		*clustersUrl += "?" + urlParams
	}
}
