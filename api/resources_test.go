package api

import (
    . "launchpad.net/gocheck"

    //"fmt"
    "log"
    "net/http"
)

var _ = log.Println

type ResourcesSuite struct{
    mapper *ResourceMapper
}

func (s *ResourcesSuite) SetUpSuite(c *C) {
    resourceMap := make(ResourceMap)
    err := resourceMap.RegisterResource(&Product{}, &Resource{
        Route: "/Products",
        Paginated: true,
    })
    c.Assert(err, IsNil)

    s.mapper = newCustomResourceMapper(resourceMap, nil)
}

func (s *ResourcesSuite) TestBadRegisterResource(c *C) {
    var (
        badType int
        err     error
    )

    resourceMap := make(ResourceMap)

    // Test registration of non-pointer type
    err = resourceMap.RegisterResource(badType, &Resource{
        Route: "/Products",
        Paginated: true,
    })
    c.Assert(err, NotNil)
    log.Println(err)

    // Test registration of nil Resource
    err = resourceMap.RegisterResource(&badType, nil)
    c.Assert(err, NotNil)
}

func (s *ResourcesSuite) TestMapper(c *C) {
    resource, err := s.mapper.lookupResource(&Product{})
    c.Assert(err, IsNil)
    c.Assert(resource.Route, Equals, "/Products")
}

func (s *ResourcesSuite) TestMapperFailedLookuyp(c *C) {
    var invalid int
    resource, err := s.mapper.lookupResource(&invalid)
    c.Assert(err, NotNil)
    c.Assert(resource, IsNil)
}

var _ = Suite(&ResourcesSuite{})

func newCustomResourceMapper(m ResourceMap, c *http.Client) *ResourceMapper {
    return &ResourceMapper{
        routeMap: m,
        client: c,
    }
}
