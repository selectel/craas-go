/*
Package `gc` provides a set of functions for working with CRaaS garbage collection API.

Example of starting a garbage collection:

	opts := &gc.StartGCOpts{
	    DeleteUntagged: true,
	}
	_, err := gc.StartGarbageCollection(ctx, client, registryID, opts)
	if err != nil {
	    log.Fatal(err)
	}

Example of getting a garbage size:

	gcSize, _, err := gc.GetGarbageSize(ctx, client, registryID)
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Printf("Garbage size: %+v", gcSize)
*/
package gc
