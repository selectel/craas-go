package testing

import (
	"time"

	"github.com/selectel/craas-go/pkg/v1/repository"
)

const (
	testRegistryID     = "9f3b5b5e-1b5a-4b5c-9b5a-5b5c1b5a4b5c"
	testRepositoryName = "test-repository"
	testImageDigest    = "sha256:0c2777301ee83e106586099533312b684b3782760d59d303865b64b90330a3e4"
)

const testListRepositoriesResponseRaw = `[
    {
        "name": "nginx",
        "size": 53620078,
        "updatedAt": "2022-09-22T10:15:40.362702Z"
    }
]`

var expectedListRepositoriesResponse = []*repository.Repository{
	{
		Name:      "nginx",
		Size:      53620078,
		UpdatedAt: time.Date(2022, 9, 22, 10, 15, 40, 362702000, time.UTC),
	},
}

const testGetRepositoryResponseRaw = `{
    "name": "test-repository",
    "updatedAt": "2022-10-25T14:21:29.321Z",
    "size": 34634720
}`

var expectedGetRepositoryResponse = &repository.Repository{
	Name:      "test-repository",
	UpdatedAt: time.Date(2022, 10, 25, 14, 21, 29, 321000000, time.UTC),
	Size:      34634720,
}

const testListImagesResponseRaw = `[
    {
        "createdAt": "2022-05-17T22:37:17.011072851Z",
        "digest": "sha256:a76df3b4f1478766631c794de7ff466aca466f995fd5bb216bb9643a3dd2a6bb",
        "layers": [
            {
                "digest": "sha256:214ca5fb90323fe769c63a12af092f2572bf1c6b300263e09883909fc865d260",
                "size": 31379476
            },
            {
                "digest": "sha256:50836501937ff210a4ee8eedcb17b49b3b7627c5b7104397b2a6198c569d9231",
                "size": 25338790
            }
        ],
        "size": 56718266,
        "tags": [
            "latest"
        ]
    }
]
`

var expectedListImagesResponse = []*repository.Image{
	{
		Digest:    "sha256:a76df3b4f1478766631c794de7ff466aca466f995fd5bb216bb9643a3dd2a6bb",
		CreatedAt: time.Date(2022, 5, 17, 22, 37, 17, 11072851, time.UTC),
		Layers: []repository.Layer{
			{
				Digest: "sha256:214ca5fb90323fe769c63a12af092f2572bf1c6b300263e09883909fc865d260",
				Size:   31379476,
			},
			{
				Digest: "sha256:50836501937ff210a4ee8eedcb17b49b3b7627c5b7104397b2a6198c569d9231",
				Size:   25338790,
			},
		},
		Size: 56718266,
		Tags: []string{"latest"},
	},
}

const testListTagsResponseRaw = `[
    "latest"
]
`

var expectedListTagsResponse = []string{"latest"}

const testListImageLayersResponseRaw = `[
    {
        "digest": "sha256:df9b9388f04ad6279a7410b85cedfdcb2208c0a003da7ab5613af71079148139",
        "size": 2814559
    },
    {
        "digest": "sha256:07cb125bb766321ae78c2a558eb61358ea0c322871a963cc343f07b19a5dfe9f",
        "size": 667045
    },
    {
        "digest": "sha256:a2388678e8f6795f5d3b643238e9a9f1e322727c82faffcf6c90cdd066dc4082",
        "size": 27429525
    },
    {
        "digest": "sha256:157b5e2e4346c2f22c275fae26a9b3bf83fbad24d1fbd18b1edc3a81bfcc622c",
        "size": 369
    }
]
`

var expectedListImageLayersResponse = []*repository.Layer{
	{
		Digest: "sha256:df9b9388f04ad6279a7410b85cedfdcb2208c0a003da7ab5613af71079148139",
		Size:   2814559,
	},
	{
		Digest: "sha256:07cb125bb766321ae78c2a558eb61358ea0c322871a963cc343f07b19a5dfe9f",
		Size:   667045,
	},
	{
		Digest: "sha256:a2388678e8f6795f5d3b643238e9a9f1e322727c82faffcf6c90cdd066dc4082",
		Size:   27429525,
	},
	{
		Digest: "sha256:157b5e2e4346c2f22c275fae26a9b3bf83fbad24d1fbd18b1edc3a81bfcc622c",
		Size:   369,
	},
}
