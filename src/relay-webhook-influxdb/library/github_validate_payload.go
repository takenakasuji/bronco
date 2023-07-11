package library

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"net/http"
	"strings"
)

func GithubValidatePayload(r *http.Request, secret string) bool {
	h := strings.SplitN(r.Header.Get("X-Hub-Signature-256"), "=", 2)
	if h[0] != "sha256" {
		return false
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Cannot body close: %s\n", err)
		}
	}(r.Body)
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Cannot read the request body: %s\n", err)
		return false
	}

	hash := hmac.New(sha256.New, []byte(secret))
	if _, err := hash.Write(b); err != nil {
		log.Printf("Cannot hash write: %s\n", err)
		return false
	}

	expectedHash := hex.EncodeToString(hash.Sum(nil))
	return h[1] == expectedHash
}
