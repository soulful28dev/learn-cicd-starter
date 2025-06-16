package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		input http.Header
		want  string
	}{
		{
			input: http.Header{"Authorization": []string{"ApiKey ApiKeyValue1"}},
			want:  "ApiKeyValue1",
		},
		{
			input: http.Header{"Authorization": []string{"ApiKey ApiKeyValue2"}},
			want:  "ApiKeyValue2",
		},
	}

	for _, tc := range tests {
		got, err := GetAPIKey(tc.input)

		if err != nil {
			t.Fatalf("expected: %v, got err: %v", tc.want, err)
		}

		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
