package testing

import "github.com/selectel/craas-go/pkg/v1/gc"

const testRegistryID = "fc43e322-b084-4b3c-a04a-1ab2a28cd860"

const testGetGarbageSizeResponseRaw = `{
  "sizeNonReferenced": 56723502,
  "sizeUntagged": 30915818,
  "sizeSummary": 87639320
}`

var expectedGetGarbageSizeResponse = &gc.GarbageSize{
	NonReferenced: 56723502,
	Untagged:      30915818,
	Summary:       87639320,
}
