package api

import (
    "net/http"
    "net/url"
    "strings"
)

const (
    SchemeName = "unleashed"
    SchemePrefix = "unleashed://"
)

type UnleashedTransport struct {
    http.Transport
    apiUrl      string
    credentials *Credentials
}

func NewUnleashedTransport(credentials *Credentials, apiUrl string) *UnleashedTransport {
    return &UnleashedTransport{credentials: credentials, apiUrl: apiUrl}
}

func (t *UnleashedTransport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
    t.rewriteURL(req)
    t.signRequest(req)
    return t.Transport.RoundTrip(req)
}

func (t *UnleashedTransport) rewriteURL(req *http.Request) {
    var err error
    request := strings.Replace(req.URL.String(), SchemePrefix, t.apiUrl, 1)
    req.URL, err = url.Parse(request)
    if err != nil {
        panic(err)
    }
}

func (t *UnleashedTransport) signRequest(req *http.Request) {
    req.Header.Set("api-auth-id", t.credentials.ApiId)
    req.Header.Set("content-type", "application/json")
    req.Header.Set("accept", "application/json")
    req.Header.Set("api-auth-signature", t.credentials.Sign([]byte(req.URL.RawQuery)))
}

func (t *UnleashedTransport) NewClient() *http.Client {
    transport := new(http.Transport)
    transport.RegisterProtocol(SchemeName, t)
    return &http.Client{Transport: transport}
}
