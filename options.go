package gospotify

import (
	"net/url"
	"strconv"
)

// queryParameters is a global structure for all query parameters.
type queryParameters struct {
	// timeRange is over what time frame the affinities are computed.
	timeRange *TimeRange
	// limit is the maximum number of items to return.
	limit *int
	// offset is the index of the first item to return. Use with limit to get the next set of items.
	offset *int
	// after is the last item ID retrieved from the previous request.
	after *string
}

// applyQueryParameters applies the options to the queryParameters.
func applyQueryParameters(opts ...QueryOption) queryParameters {
	p := queryParameters{}

	for _, opt := range opts {
		opt(&p)
	}

	return p
}

// toQuery converts the parameters to a query string.
func (p queryParameters) toQuery() string {
	query := url.Values{}
	if p.timeRange != nil {
		query.Set("time_range", string(*p.timeRange))
	}
	if p.limit != nil {
		query.Set("limit", strconv.Itoa(*p.limit))
	}
	if p.offset != nil {
		query.Set("offset", strconv.Itoa(*p.offset))
	}
	if p.after != nil {
		query.Set("after", *p.after)
	}
	return query.Encode()
}

// QueryOption is an option for query parameters.
type QueryOption func(*queryParameters)

// WithTimeRange sets the time range for the affinities.
func WithTimeRange(timeRange TimeRange) QueryOption {
	return func(p *queryParameters) {
		p.timeRange = &timeRange
	}
}

// WithLimit sets the maximum number of items to return.
func WithLimit(limit int) QueryOption {
	return func(p *queryParameters) {
		if limit < 1 {
			limit = 1
		}
		if limit > 50 {
			limit = 50
		}
		p.limit = &limit
	}
}

// WithOffset sets the index of the first item to return.
func WithOffset(offset int) QueryOption {
	return func(p *queryParameters) {
		if offset < 0 {
			offset = 0
		}
		p.offset = &offset
	}
}

// WithAfter sets the last item ID retrieved from the previous request.
func WithAfter(after string) QueryOption {
	return func(p *queryParameters) {
		p.after = &after
	}
}
