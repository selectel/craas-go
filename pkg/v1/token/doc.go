/*
Package `token` provides a set of functions for working with with CRaaS tokens API.

Example of creating a token:

	createOpts := &token.CreateOpts{
	    TokenTTL: token.TTL1Year,
	}
	craasToken, _, err := token.Create(ctx, client, createOpts)
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Printf("CRaaS token: %+v", craasToken)

Example of getting a token by its ID:

	craasToken, _, err := token.Get(ctx, client, tokenID)
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Printf("CRaaS token: %+v", craasToken)

Example of revoking a token by its ID:

	_, err := token.Revoke(ctx, client, tokenID)
	if err != nil {
	    log.Fatal(err)
	}

Example of refreshing a token by its ID:

	token, _, err := token.Refresh(ctx, client, tokenID)
	if err != nil {
	    log.Fatal(err)
	}
*/
package token
