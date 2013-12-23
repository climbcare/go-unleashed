package api

import (
    "crypto/rand"
    "log"
    "net/http"

    . "launchpad.net/gocheck"
    "github.com/nu7hatch/gouuid"
)

var _ = log.Println


type AuthSuite struct{
    ApiId string
    ApiKey string
}


func (s *AuthSuite) SetUpSuite(c *C) {
    // Make apiId
    fakeId, err := uuid.NewV4()
    c.Assert(err, IsNil)
    s.ApiId = fakeId.String()

    // Make apiKey
    b := make([]byte, 96)
    _, err = rand.Read(b)
    c.Assert(err, IsNil)
    s.ApiKey = string(b)
}

func (s *AuthSuite) TestAuth(c *C) {
    signer := NewRequestSigner(s.ApiId, s.ApiKey)

    url := "https://example.com/Products?productCode=ACME&includeObsolete=true"
    req, err := http.NewRequest("GET", url, nil)
    c.Assert(err, IsNil)

    req = signer(req)

    c.Assert(req, NotNil)

    // Ensure API headers were added
    c.Assert(req.Header.Get("accept"), Equals, "application/json")
    c.Assert(req.Header.Get("content-type"), Equals, "application/json")
    c.Assert(req.Header.Get("api-auth-id"), Equals, s.ApiId)

    c.Assert(req.Header.Get("api-auth-signature"), Not(Equals), "")
}

var _ = Suite(&AuthSuite{})
