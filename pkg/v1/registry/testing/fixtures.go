package testing

import (
	"time"

	"github.com/selectel/craas-go/pkg/v1/registry"
)

const testRegistryID = "9f3b5b5e-1b5a-4b5c-9b5a-5b5c1b5a4b5c"

const testCreateRegistryRequestRaw = `{
    "name": "test-registry"
}`

const testCreateRegistryResponseRaw = `{
    "id": "9f3b5b5e-1b5a-4b5c-9b5a-5b5c1b5a4b5c",
    "name": "test-registry",
    "createdAt": "2022-10-25T10:25:22.556Z",
    "status": "CREATING",
    "size": 0,
    "sizeLimit": 21474836480,
    "used": 0
}`

var expectedCreateRegistryResponse = &registry.Registry{
	ID:        "9f3b5b5e-1b5a-4b5c-9b5a-5b5c1b5a4b5c",
	Name:      "test-registry",
	CreatedAt: time.Date(2022, 10, 25, 10, 25, 22, 556000000, time.UTC),
	Status:    "CREATING",
	Size:      0,
	SizeLimit: 21474836480,
	Used:      0,
}

const testGetRegistryResponseRaw = `{
    "id": "9f3b5b5e-1b5a-4b5c-9b5a-5b5c1b5a4b5c",
    "name": "test-registry",
    "createdAt": "2022-10-25T10:25:22.556Z",
    "status": "ACTIVE",
    "size": 500000000,
    "sizeLimit": 21474836480,
      "used": 2.33
}`

var expectedGetRegistryResponse = &registry.Registry{
	ID:        "9f3b5b5e-1b5a-4b5c-9b5a-5b5c1b5a4b5c",
	Name:      "test-registry",
	CreatedAt: time.Date(2022, 10, 25, 10, 25, 22, 556000000, time.UTC),
	Status:    "ACTIVE",
	Size:      500000000,
	SizeLimit: 21474836480,
	Used:      2.33,
}

const testGetRegistryWithUnknownStatusResponseRaw = `{
    "id": "9f3b5b5e-1b5a-4b5c-9b5a-5b5c1b5a4b5c",
    "name": "test-registry",
    "createdAt": "2022-10-25T10:25:22.556Z",
    "status": "UNEXPECTED",
    "size": 500000000,
    "sizeLimit": 21474836480,
      "used": 2.33
}`

var expectedGetRegistryWithUnknownStatusResponse = &registry.Registry{
	ID:        "9f3b5b5e-1b5a-4b5c-9b5a-5b5c1b5a4b5c",
	Name:      "test-registry",
	CreatedAt: time.Date(2022, 10, 25, 10, 25, 22, 556000000, time.UTC),
	Status:    "UNKNOWN",
	Size:      500000000,
	SizeLimit: 21474836480,
	Used:      2.33,
}

const testListRegistriesResponseRaw = `[
    {
        "createdAt": "2022-06-22T10:13:45.895721Z",
        "id": "fc43e322-b084-4b3c-a04a-1ab2a28cd860",
        "name": "test-registry1",
        "size": 1811104096,
        "sizeLimit": 23622320128,
        "status": "ACTIVE",
        "used": 8.43
    },
    {
        "createdAt": "2022-08-11T08:27:06.446114Z",
        "id": "a71ec018-4317-4c2a-afaf-e5cbdc017e9c",
        "name": "test-registry2",
        "size": 984830374,
        "sizeLimit": 21474836480,
        "status": "ACTIVE",
        "used": 4.59
    }
]`

var expectedListRegistriesResponse = []*registry.Registry{
	{
		ID:        "fc43e322-b084-4b3c-a04a-1ab2a28cd860",
		Name:      "test-registry1",
		CreatedAt: time.Date(2022, 6, 22, 10, 13, 45, 895721000, time.UTC),
		Status:    "ACTIVE",
		Size:      1811104096,
		SizeLimit: 23622320128,
		Used:      8.43,
	},
	{
		ID:        "a71ec018-4317-4c2a-afaf-e5cbdc017e9c",
		Name:      "test-registry2",
		CreatedAt: time.Date(2022, 8, 11, 8, 27, 6, 446114000, time.UTC),
		Status:    "ACTIVE",
		Size:      984830374,
		SizeLimit: 21474836480,
		Used:      4.59,
	},
}
