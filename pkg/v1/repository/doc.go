/*
Package `repository` provides a set of functions for working with CRaaS registry repositories API.

Example of getting a list of registry repositories:

	repositories, _, err := repository.ListRepositories(ctx, client, registryID)
	if err != nil {
	    log.Fatal(err)
	}
	for _, repo := range repositories {
	    fmt.Printf("Repository: %+v", repo)
	}

Example of getting a repository by its name:

	repo, _, err := repository.GetRepository(ctx, client, registryID, repositoryName)
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Printf("Repository: %+v", repo)

Example of deleting a repository by its name:

	_, err := repository.DeleteRepository(ctx, client, registryID, repositoryName)
	if err != nil {
	    log.Fatal(err)
	}

Example of getting a list of repository images:

	images, _, err := repository.ListImages(ctx, client, registryID, repositoryName)
	if err != nil {
	    log.Fatal(err)
	}
	for _, image := range images {
	    fmt.Printf("Image: %+v", image)
	}

Example of getting a list of repository tags:

	tags, _, err := repository.ListTags(ctx, client, registryID, repositoryName)
	if err != nil {
	    log.Fatal(err)
	}
	for _, tag := range tags {
	    fmt.Printf("Tag: %+v", tag)
	}

Example of getting a list of image layers:

	layers, _, err := repository.ListImageLayers(ctx, client, registryID, repositoryName, imageDigest)
	if err != nil {
	    log.Fatal(err)
	}
	for _, layer := range layers {
	    fmt.Printf("Layer: %+v", layer)
	}

Example of deleting an image by its digest:

	_, err := repository.DeleteImageManifest(ctx, client, registryID, repositoryName, imageDigest)
	if err != nil {
	    log.Fatal(err)
	}
*/
package repository
