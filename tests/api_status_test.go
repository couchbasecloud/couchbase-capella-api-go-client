package tests

import (
	"context"
	"fmt"
	"os"
	"testing"

	couchbasecloud "github.com/couchbaselabs/couchbase-cloud-go-client"
)

func TestStatus(t *testing.T) {
	ctx := context.WithValue(
		context.Background(),
		couchbasecloud.ContextAPIKeys,
		map[string]couchbasecloud.APIKey{
			"accessKey": {
				Key: os.Getenv("CBC_ACCESS_KEY"),
			},
			"secretKey": {
				Key: os.Getenv("CBC_SECRET_KEY"),
			},
		},
	)
	configuration := couchbasecloud.NewConfiguration()
	apiClient := couchbasecloud.NewAPIClient(configuration)
	resp, _, err := apiClient.StatusApi.StatusShow(ctx).Execute()
	if err != nil {
		t.Fatalf("Error when calling `StatusApi.StatusShow(ctx)`: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from `StatusApi.StatusShow(ctx)`: %v\n", resp)
}
