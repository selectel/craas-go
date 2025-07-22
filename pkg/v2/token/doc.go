/*
Package `tokenv2` provides a set of functions for working with with CRaaS tokens API.

Example of creating a token:

		var expiresAt, _ = time.Parse("2006-01-02T15:04:05Z", "2030-01-01T00:00:00Z")
		create := &tokenV2.TokenV2{
		Expiration: struct {
			IsSet     bool      "json:\"isSet\""
			ExpiresAt time.Time "json:\"expiresAt,omitempty\""
		}{
			IsSet:     true,
			ExpiresAt: expiresAt,
		},
		Scope: struct {
			ModeRW        bool
			AllRegistries bool
			RegistryIds   []string
		}{
			ModeRW:        true,
			AllRegistries: false,
			RegistryIds: []string{
				"888af692-c646-4b76-a234-81ca9b5bcafe",
				"6303699d-c2cd-40b1-8428-9dcd6cc3d00d",
			},
		},
	}

	craasToken, err := tokenV2.Create(ctx, client, create)
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Printf("CRaaS token: %+v", craasToken)

Example of getting a token by its ID:

	craasToken, _, err := tokenV2.GetByID(ctx, client, tokenID)
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Printf("CRaaS token: %+v", craasToken)

Example of revoking a token by its ID:

	_, err := tokenV2.Revoke(ctx, client, tokenID)
	if err != nil {
	    log.Fatal(err)
	}

Example of getting list of a tokens:

	lim := new(int)
	*lim = 1
	opts := tokenV2.Opts{
		Limit: lim,
	}
	_, _, err := tokenV2.List(ctx, client, opts)
	if err != nil {
	    log.Fatal(err)
	}

Example of refreshing a token by its ID:

	var expiresAt, _ = time.Parse("2006-01-02T15:04:05Z", "2030-01-01T00:00:00Z")
	var exp = tokenV2.Exp{
		IsSet:     true,
		ExpiresAt: expiresAt,
	}

	_, err := tokenV2.Refresh(ctx, client, tokenID, exp)
	if err != nil {
	    log.Fatal(err)
	}

Example of regenerate a token by its ID:

	var expiresAt, _ = time.Parse("2006-01-02T15:04:05Z", "2030-01-01T00:00:00Z")
	var exp = tokenV2.Exp{
		IsSet:     true,
		ExpiresAt: expiresAt,
	}

	_, err := tokenV2.Regenerate(ctx, client, tokenID, exp)
	if err != nil {
	    log.Fatal(err)
	}

Example of delete a token by its ID:

	_, err := tokenV2.Delete(ctx, client, tokenID)
	if err != nil {
	    log.Fatal(err)
	}

Example of patch a token by its ID:

		name := "token"
		scope := tokenV2.Scope{
			ModeRW:        true,
			AllRegistries: false,
			RegistryIds: []string{
				"888af692-c646-4b76-a234-81ca9b5bcafe",
				"6303699d-c2cd-40b1-8428-9dcd6cc3d00d",
		},
	}

	_, err := tokenV2.Patch(ctx, client, tokenID, name, scope)
	if err != nil {
	    log.Fatal(err)
	}
*/
package tokenv2
