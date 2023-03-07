/*
Package `registry` provides a set of functions for working with with CRaaS registries API.

Example of creating a registry:

	createOpts := &registry.CreateOpts{
	    Name: "test-registry",
	}
	createdRegistry, _, err := registry.Create(ctx, client, createOpts)
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Printf("Created registry: %+v", createdRegistry)

Example of getting a registry by its ID:

	gotRegistry, _, err := registry.Get(ctx, client, registryID)
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Printf("Registry: %+v", gotRegistry)

Example of listing registries:

	registries, _, err := registry.List(ctx, client)
	if err != nil {
	    log.Fatal(err)
	}
	for _, registry := range registries {
	    fmt.Printf("Registry: %+v", registry)
	}

Example of deleting a registry by its ID:

	_, err := registry.Delete(ctx, client, registryID)
	if err != nil {
	    log.Fatal(err)
	}
*/
package registry
