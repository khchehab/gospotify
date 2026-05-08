package gospotify

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"testing"
)

// Compile-time assertion: *ErrorResponse satisfies the error interface.
var _ error = (*ErrorResponse)(nil)

// ---- Error() formatting ----

func newErrResp(status int, message string) *ErrorResponse {
	e := &ErrorResponse{}
	e.ErrorObject.Status = status
	e.ErrorObject.Message = message
	return e
}

func TestErrorResponse_Error_Formatting(t *testing.T) {
	cases := []struct {
		name    string
		status  int
		message string
		want    string
	}{
		{"typical_404", 404, "Resource not found", "[404] Resource not found"},
		{"unauthorized_401", 401, "No token provided", "[401] No token provided"},
		{"server_error_500", 500, "Internal server error", "[500] Internal server error"},
		{"zero_status", 0, "something went wrong", "[0] something went wrong"},
		{"empty_message", 400, "", "[400] "},
		{"both_zero_values", 0, "", "[0] "},
		{"negative_status", -1, "invalid", "[-1] invalid"},
		{"special_chars", 403, "Access denied: user \"admin\" is not allowed\n", "[403] Access denied: user \"admin\" is not allowed\n"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := newErrResp(tc.status, tc.message).Error()
			if got != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}

func TestErrorResponse_Error_LongMessage(t *testing.T) {
	msg := strings.Repeat("a", 10_000)
	e := newErrResp(429, msg)
	got := e.Error()
	want := "[429] " + msg
	if got != want {
		t.Errorf("long message truncated or mangled: got len %d, want len %d", len(got), len(want))
	}
}

// ---- error interface compliance ----

func TestErrorResponse_ReturnedAsError(t *testing.T) {
	makeErr := func() error {
		e := &ErrorResponse{}
		e.ErrorObject.Status = 404
		e.ErrorObject.Message = "not found"
		return e
	}

	err := makeErr()
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if err.Error() != "[404] not found" {
		t.Errorf("got %q, want '[404] not found'", err.Error())
	}
	var target *ErrorResponse
	if !errors.As(err, &target) {
		t.Fatal("type assertion to *ErrorResponse failed")
	}
	if target.ErrorObject.Status != 404 {
		t.Errorf("got status %d, want 404", target.ErrorObject.Status)
	}
}

func TestErrorResponse_ErrorsAs(t *testing.T) {
	inner := &ErrorResponse{}
	inner.ErrorObject.Status = 503
	inner.ErrorObject.Message = "service unavailable"

	wrapped := fmt.Errorf("outer: %w", inner)

	var target *ErrorResponse
	if !errors.As(wrapped, &target) {
		t.Fatal("errors.As failed to find *ErrorResponse in chain")
	}
	if target.ErrorObject.Status != 503 {
		t.Errorf("got status %d, want 503", target.ErrorObject.Status)
	}
	if target.ErrorObject.Message != "service unavailable" {
		t.Errorf("got message %q, want 'service unavailable'", target.ErrorObject.Message)
	}
}

// ---- JSON unmarshalling ----

func unmarshalErrResp(t *testing.T, data string) (*ErrorResponse, error) {
	t.Helper()
	var e ErrorResponse
	err := json.Unmarshal([]byte(data), &e)
	return &e, err
}

func TestErrorResponse_Unmarshal_Valid(t *testing.T) {
	e, err := unmarshalErrResp(t, `{"error":{"status":404,"message":"Service not found"}}`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if e.ErrorObject.Status != 404 {
		t.Errorf("got status %d, want 404", e.ErrorObject.Status)
	}
	if e.ErrorObject.Message != "Service not found" {
		t.Errorf("got message %q, want 'Service not found'", e.ErrorObject.Message)
	}
}

func TestErrorResponse_Unmarshal_Unauthorized(t *testing.T) {
	e, err := unmarshalErrResp(t, `{"error":{"status":401,"message":"No token provided"}}`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if e.ErrorObject.Status != 401 || e.ErrorObject.Message != "No token provided" {
		t.Errorf("got status=%d message=%q", e.ErrorObject.Status, e.ErrorObject.Message)
	}
}

func TestErrorResponse_Unmarshal_EmptyMessage(t *testing.T) {
	e, err := unmarshalErrResp(t, `{"error":{"status":400,"message":""}}`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if e.ErrorObject.Status != 400 || e.ErrorObject.Message != "" {
		t.Errorf("got status=%d message=%q", e.ErrorObject.Status, e.ErrorObject.Message)
	}
}

func TestErrorResponse_Unmarshal_MissingMessage(t *testing.T) {
	e, err := unmarshalErrResp(t, `{"error":{"status":400}}`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if e.ErrorObject.Status != 400 || e.ErrorObject.Message != "" {
		t.Errorf("got status=%d message=%q", e.ErrorObject.Status, e.ErrorObject.Message)
	}
}

func TestErrorResponse_Unmarshal_MissingStatus(t *testing.T) {
	e, err := unmarshalErrResp(t, `{"error":{"message":"Bad request"}}`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if e.ErrorObject.Status != 0 || e.ErrorObject.Message != "Bad request" {
		t.Errorf("got status=%d message=%q", e.ErrorObject.Status, e.ErrorObject.Message)
	}
}

func TestErrorResponse_Unmarshal_EmptyErrorObject(t *testing.T) {
	e, err := unmarshalErrResp(t, `{"error":{}}`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if e.ErrorObject.Status != 0 || e.ErrorObject.Message != "" {
		t.Errorf("got status=%d message=%q", e.ErrorObject.Status, e.ErrorObject.Message)
	}
}

func TestErrorResponse_Unmarshal_NullErrorObject(t *testing.T) {
	e, err := unmarshalErrResp(t, `{"error":null}`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if e.ErrorObject.Status != 0 || e.ErrorObject.Message != "" {
		t.Errorf("got status=%d message=%q", e.ErrorObject.Status, e.ErrorObject.Message)
	}
}

func TestErrorResponse_Unmarshal_MissingTopLevelKey(t *testing.T) {
	e, err := unmarshalErrResp(t, `{}`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if e.ErrorObject.Status != 0 || e.ErrorObject.Message != "" {
		t.Errorf("got status=%d message=%q", e.ErrorObject.Status, e.ErrorObject.Message)
	}
}

func TestErrorResponse_Unmarshal_InvalidJSON(t *testing.T) {
	_, err := unmarshalErrResp(t, `not json at all`)
	if err == nil {
		t.Fatal("expected non-nil error for invalid JSON")
	}
	if _, ok := errors.AsType[*json.SyntaxError](err); !ok {
		t.Errorf("expected *json.SyntaxError, got %T: %v", err, err)
	}
}

func TestErrorResponse_Unmarshal_StatusWrongType(t *testing.T) {
	_, err := unmarshalErrResp(t, `{"error":{"status":"404","message":"not found"}}`)
	if err == nil {
		t.Fatal("expected non-nil error for wrong status type")
	}
	if _, ok := errors.AsType[*json.UnmarshalTypeError](err); !ok {
		t.Errorf("expected *json.UnmarshalTypeError, got %T: %v", err, err)
	}
}

func TestErrorResponse_Unmarshal_ExtraFields(t *testing.T) {
	e, err := unmarshalErrResp(t, `{"error":{"status":429,"message":"rate limited","reason":"PREMIUM_REQUIRED"}}`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if e.ErrorObject.Status != 429 || e.ErrorObject.Message != "rate limited" {
		t.Errorf("got status=%d message=%q", e.ErrorObject.Status, e.ErrorObject.Message)
	}
}

func TestErrorResponse_Unmarshal_UnicodeMessage(t *testing.T) {
	msg := "无效请求 — données invalides"
	e, err := unmarshalErrResp(t, `{"error":{"status":400,"message":"`+msg+`"}}`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if e.ErrorObject.Message != msg {
		t.Errorf("got message %q, want %q", e.ErrorObject.Message, msg)
	}
}
