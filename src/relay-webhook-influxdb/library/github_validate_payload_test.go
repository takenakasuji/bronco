package library

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGithubValidatePayload1(t *testing.T) {
	const (
		body            = `{"br":true}`
		signatureHeader = "X-Hub-Signature-256"
	)

	tests := []struct {
		name      string
		secret    string
		signature string
		want      bool
	}{
		{
			name:      "success",
			secret:    "SUCCESS_SECRET",
			signature: "sha256=3283ddf9448334fe28ad466832446836af7e48f0094f53469b60188399ca0773",
			want:      true,
		},
		{
			name:      "fail_no_prefix",
			secret:    "FAIL_SECRET",
			signature: "b86634e4703a184ae5bd12e515db0523a4c12a4ae3909e53c501fd7371cd7820",
			want:      false,
		},
		{
			name:      "fail_unmatched_secret",
			secret:    "FAIL_SECRET",
			signature: "sha256=0123456789abc",
			want:      false,
		},
	}
	for _, tt := range tests {
		buf := bytes.NewBufferString(body)
		req, _ := http.NewRequest("POST", "http://localhost", buf)
		req.Header.Set(signatureHeader, tt.signature)
		t.Run(tt.name, func(t *testing.T) {
			if got := GithubValidatePayload(req, tt.secret); got != tt.want {
				t.Errorf("GithubValidatePayload() = %v, want %v", got, tt.want)
			}
		})
	}
}
