package api

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/base64"
    "net/http"
)

type RequestSignerFunc func(*http.Request) (*http.Request)

func NewRequestSigner(apiId, apiKey string) RequestSignerFunc {
    return func(req *http.Request) (*http.Request) {
        req.Header.Set("api-auth-id", apiId)
        req.Header.Set("content-type", "application/json")
        req.Header.Set("accept", "application/json")

        mac := hmac.New(sha256.New, []byte(apiKey))
        mac.Write([]byte(req.URL.RawQuery))
        signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))
        req.Header.Set("api-auth-signature", signature)

        return req
    }
}
