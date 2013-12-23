package api

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/base64"
    "fmt"
    "log"
    "net/http"
    "net/url"
    "strings"
)

const (
    schemeName = "unleashed"
    schemePrefix = "unleashed://"
)

type Credentials struct {
    ApiId string
    ApiKey string
}

type UnleashedTransport struct {
    http.Transport
    apiUrl      string
    credentials *Credentials
}

func NewUnleashedTransport(credentials *Credentials, apiUrl string) *UnleashedTransport {
    return &UnleashedTransport{credentials: credentials, apiUrl: apiUrl}
}

func (t *UnleashedTransport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
    req.URL = t.rewriteURL(req.URL)
    req = t.signRequest(req)
    return t.Transport.RoundTrip(req)
}

func (t *UnleashedTransport) rewriteURL(URL *url.URL) *url.URL {
    path := strings.TrimPrefix(URL.String(), schemePrefix)
    request := fmt.Sprintf("%s%s", t.apiUrl, path)
    newURL, err := url.Parse(request)
    if err != nil {
        panic(err)
    }
    return newURL
}

func (t *UnleashedTransport) signRequest(req *http.Request) *http.Request {
    req.Header.Set("api-auth-id", t.credentials.ApiId)
    req.Header.Set("content-type", "application/json")
    req.Header.Set("accept", "application/json")

    mac := hmac.New(sha256.New, []byte(t.credentials.ApiKey))
    mac.Write([]byte(req.URL.RawQuery))
    signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))
    req.Header.Set("api-auth-signature", signature)

    log.Println(req.URL.String())
    return req
}

func NewClient(transport *UnleashedTransport) *http.Client {
    t := new(http.Transport)
    t.RegisterProtocol(schemeName, transport)
    return &http.Client{Transport: t}
}
