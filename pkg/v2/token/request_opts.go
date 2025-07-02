package tokenv2

import (
	"net/url"
	"strconv"
)

type Opts struct {
	Limit     *int
	Offset    *int
	SortField string
	SortType  string
	Search    string
	ScopeMode string
}

func makeQueryString(o Opts) string {
	val := url.Values{}
	if o.Limit != nil {
		limit := strconv.Itoa(*o.Limit)
		val.Add("limit", limit)
	}
	if o.Offset != nil {
		offset := strconv.Itoa(*o.Offset)
		val.Add("offset", offset)
	}
	if o.SortField != "" {
		val.Add("sort_field", o.SortField)
	}
	if o.SortType != "" {
		val.Add("sort_type", o.SortType)
	}
	if o.Search != "" {
		val.Add("search", o.Search)
	}
	if o.ScopeMode != "" {
		val.Add("scope_mode", o.ScopeMode)
	}
	query := val.Encode()

	return query
}
