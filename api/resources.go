package api

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "reflect"
)

// The official Unleashed API resource map
var API_RESOURCE_MAP ResourceMap = make(ResourceMap)

type Resource struct {
    Route     string
    Paginated bool
}

func (r *Resource) GetQuery(id string) string {
    return fmt.Sprintf("%s/%s/%s", SchemePrefix, r.Route, id)
}

func (r *Resource) ListQuery() string {
    return fmt.Sprintf("%s/%s/", SchemePrefix, r.Route)
}

// A type to describe a paginated resource response
type paginatedResource struct {
    Pagination Pagination
    Items      interface{}
}

func newPaginatedResource(items interface{}) *paginatedResource {
    return &paginatedResource{
        Items: items,
    }
}

type ResourceMap map[reflect.Type](*Resource)

func (m ResourceMap) RegisterResource(i interface{}, resource *Resource) error {
    if resource == nil {
        return fmt.Errorf("Nil Resource cannot be registered")
    }

    t := reflect.TypeOf(i)
    if t.Kind() != reflect.Ptr {
        return fmt.Errorf("Cannot register non-pointer type '%s'", t)
    }

    m[t.Elem()] = resource
    return nil
}

type ResourceMapper struct {
    routeMap  ResourceMap
    client   *http.Client
}

// Creates a new default resource mapper with all the standard
// API resources defined.
func NewResourceMapper(client *http.Client) *ResourceMapper {
    mapper := &ResourceMapper{
        routeMap: API_RESOURCE_MAP,
        client: client,
    }
    return mapper
}

func (m *ResourceMapper) lookupResource(o interface{}) (r *Resource, err error) {
    var t reflect.Type = reflect.TypeOf(o)

    if t.Kind() == reflect.Ptr {
        t = t.Elem()
    }
    if t.Kind() == reflect.Slice {
        t = t.Elem()
    }

    r, ok := m.routeMap[t]
    if !ok {
        err = fmt.Errorf("ResourceMapper: Could not find resource which maps to type %s", t)
    }
    return
}

// Gets the resource identified by the Guid
// The paramater `o` should be a pointer to an instance of the resource type.
func (m *ResourceMapper) Get(o interface{}, id *Guid) (err error) {
    if resource, err := m.lookupResource(o); err == nil {
        query := resource.GetQuery(id.String())
        err = m.get(o, query)
    }
    return
}

// Lists the entire collection of the given resource.
// The parameter `o` should be a pointer to an initialised slice
// of the resource type.
func (m *ResourceMapper) List(o interface{}) (err error) {
    if resource, err := m.lookupResource(o); err == nil {
        query := resource.ListQuery()
        if resource.Paginated {
            // TODO: Keep fetching until pagination exhausted
            page := newPaginatedResource(o)
            err = m.get(page, query)
        } else {
            err = m.get(o, query)
        }
    }
    return
}

func (m *ResourceMapper) get(o interface{}, query string) (err error) {
    if res, err := m.client.Get(query); err == nil {
        if body, err := ioutil.ReadAll(res.Body); err == nil {
            defer res.Body.Close()
            err = json.Unmarshal(body, o)
        }
    }
    return
}
