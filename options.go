package gospotify

import (
	"net/url"
	"strconv"
	"strings"
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
	// market is an ISO 3166-1 alpha-2 country code.
	// If a country code is specified, only content that is available in that market will be returned.
	market *string
	// includeGroups is a list of keywords that will be used to filter the response.
	// If not supplied, all album types will be returned. Valid values are album, single, appears_on, compilation.
	includeGroups []string
	// includeExternal if include_external=audio is specified it signals that the client can play externally hosted audio content, and marks the content as playable in the response.
	// By default externally hosted audio content is marked as unplayable in the response.
	includeExternal *string
	// fields filters for the query: a comma-separated list of the fields to return.
	// If omitted, all fields are returned.
	// A dot separator can be used to specify non-reoccurring fields, while parentheses can be used to specify reoccurring fields within objects.
	// Use multiple parentheses to drill down into nested objects.
	// Fields can be excluded by prefixing them with an exclamation mark.
	fields *string
	// additionalTypes is a comma-separated list of item types that your client supports besides the default track type.
	// Valid types are: track and episode.
	additionalTypes []string
	// uris is a list of Spotify URIs.
	uris []string
	// position is the position to insert the items, a zero-based index.
	position *int
	// deviceID The id of the device this command is targeting.
	deviceID *string
	// afterMs is a Unix timestamp in milliseconds. Returns all items after (but not including) this cursor position.
	afterMs *int
	// beforeMs is a Unix timestamp in milliseconds. Returns all items before (but not including) this cursor position.
	beforeMs *int
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
	if p.market != nil {
		query.Set("market", *p.market)
	}
	if len(p.includeGroups) > 0 {
		query.Set("include_groups", strings.Join(p.includeGroups, ","))
	}
	if p.includeExternal != nil {
		query.Set("include_external", *p.includeExternal)
	}
	if p.fields != nil {
		query.Set("fields", *p.fields)
	}
	if len(p.additionalTypes) > 0 {
		query.Set("additional_types", strings.Join(p.additionalTypes, ","))
	}
	if len(p.uris) > 0 {
		query.Set("uris", strings.Join(p.uris, ","))
	}
	if p.position != nil {
		query.Set("position", strconv.Itoa(*p.position))
	}
	if p.deviceID != nil {
		query.Set("device_id", *p.deviceID)
	}
	if p.afterMs != nil {
		query.Set("after", strconv.Itoa(*p.afterMs))
	}
	if p.beforeMs != nil {
		query.Set("before", strconv.Itoa(*p.beforeMs))
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

// WithMarket sets the market for the content to be returned.
func WithMarket(market string) QueryOption {
	return func(p *queryParameters) {
		p.market = &market
	}
}

// WithIncludeGroups sets the include groups for the content to be returned.
func WithIncludeGroups(includeGroups ...string) QueryOption {
	return func(p *queryParameters) {
		p.includeGroups = includeGroups
	}
}

// WithIncludeExternal sets the include external for the content to be returned.
func WithIncludeExternal(includeExternal string) QueryOption {
	return func(p *queryParameters) {
		p.includeExternal = &includeExternal
	}
}

// WithFields sets the fields for the content to be returned.
func WithFields(fields string) QueryOption {
	return func(p *queryParameters) {
		p.fields = &fields
	}
}

// WithAdditionalTypes sets the additional types for the content to be returned.
func WithAdditionalTypes(additionalTypes ...string) QueryOption {
	return func(p *queryParameters) {
		p.additionalTypes = additionalTypes
	}
}

// WithURIs sets the specified Spotify URIs.
func WithURIs(uris ...string) QueryOption {
	return func(p *queryParameters) {
		p.uris = uris
	}
}

// WithPosition sets the position to insert the items, a zero-based index.
func WithPosition(position int) QueryOption {
	return func(p *queryParameters) {
		p.position = &position
	}
}

// WithDeviceID sets the device id the command is targeting.
func WithDeviceID(deviceID string) QueryOption {
	return func(p *queryParameters) {
		p.deviceID = &deviceID
	}
}

// WithAfterMs sets the cursor position to return all items after (but not including) it.
func WithAfterMs(afterMs int) QueryOption {
	return func(p *queryParameters) {
		p.afterMs = &afterMs
	}
}

// WithBeforeMs sets the cursor position to return all items before (but not including) it.
func WithBeforeMs(beforeMs int) QueryOption {
	return func(p *queryParameters) {
		p.beforeMs = &beforeMs
	}
}
