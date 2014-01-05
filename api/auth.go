package api

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/base64"
)

type Credentials struct {
    ApiId  string
    ApiKey string
}

/*
Use the Credentials to sign some bytes as per the Unleashed Authenticaion scheme.

See https://api.unleashedsoftware.com/AuthenticationHelp
*/
func (c *Credentials) Sign(data []byte) string {
    mac := hmac.New(sha256.New, []byte(c.ApiKey))
    mac.Write(data)
    return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
