package api

import (
    "crypto/rand"
    "log"
    "net/http"
    "net/http/httptest"

    . "launchpad.net/gocheck"
    "github.com/nu7hatch/gouuid"
)

var _ = log.Println


type TransportSuite struct {
    credentials *Credentials
}

func (s *TransportSuite) SetUpSuite(c *C) {
    s.credentials = new(Credentials)

    // Make API ID
    fakeId, err := uuid.NewV4()
    c.Assert(err, IsNil)
    s.credentials.ApiId = fakeId.String()

    // Make API Key
    b := make([]byte, 96)
    _, err = rand.Read(b)
    c.Assert(err, IsNil)
    s.credentials.ApiKey = string(b)
}

func (s *TransportSuite) TestUnleashedTransport(c *C) {
    ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
        // Ensure API headers were added
        c.Assert(req.Header.Get("accept"), Equals, "application/json")
        c.Assert(req.Header.Get("content-type"), Equals, "application/json")
        c.Assert(req.Header.Get("api-auth-id"), Equals, s.credentials.ApiId)
        c.Assert(req.Header.Get("api-auth-signature"), Not(Equals), "")
    }))
    defer ts.Close()

    transport := NewUnleashedTransport(s.credentials, ts.URL)
    client := transport.NewClient()

    url := "unleashed:///Products?productCode=ACME&includeObsolete=true"
    _, err := client.Get(url)
    c.Assert(err, IsNil)
}

var _ = Suite(&TransportSuite{})
