# couchbase-cloud-go-client
Work in progress Go wrapper for the Couchbase Cloud REST API

## Usage

`go get github.com/couchbaselabs/couchbase-cloud-go-client`

### Import
```go
import (
    "github.com/couchbaselabs/couchbase-cloud-go-client"
)
```

### Client

```go
client := couchbasecloudclient.NewClient(os.Getenv("CBC_ACCESS_KEY"), os.Getenv("CBC_SECRET_KEY"))
```

### List all clouds
Options can be specified with the `ListCloudOptions` type (some are optional)
```go
page := 1
lastPage := math.MaxInt16

for ok := true; ok; ok = page <= lastPage {
    listCloudsResponse, err := client.ListClouds(&couchbasecloudclient.ListCloudsOptions{Page: page, PerPage: 10})

    if err != nil {
        return nil, err
    }

    for _, cloud := range listCloudsResponse.Data {
        fmt.Printf(cloud.Name)
    }

    lastPage = listCloudsResponse.Cursor.Pages.Last
    page++
}
```

### List all clusters
Options can be specified with the `ListClustersOptions` type (some are optional)
```go
page := 1
lastPage := math.MaxInt16

for ok := true; ok; ok = page <= lastPage {
    listClustersResponse, err := client.ListClusters(&couchbasecloudclient.ListClustersOptions{Page: page, PerPage: 10})

    if err != nil {
        return nil, err
    }

    for _, cluster := range listClustersResponse.Data {
        fmt.Printf(cluster.Name)
    }

    lastPage = listClustersResponse.Cursor.Pages.Last
    page++
}
```

## TODO

- Build better (functional) pagination capabilities into the client. E.g.  iterator functions like aws has 
  https://docs.aws.amazon.com/sdk-for-go/api/service/s3/#S3.ListObjectsV2Pages
- Unit/integration testing
- Add documentation strings to all public types and functions so that documentation can be auto-generated from the code.
- Add linter https://github.com/golangci/golangci-lint to help maintain code quality.
- Auto-generate boilerplate code for the rest of the client using OpenAPI spec for the Couchbase Cloud API. This could 
  also be used to generate a server stub for unit testing or as an emulator for integration level tests.