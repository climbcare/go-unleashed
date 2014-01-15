package api

func init() {
    // Here we map all the Unleashed resources exposed through the API
    // into the resource map used for all new instances of the API client.

    API_RESOURCE_MAP.RegisterResource(&Product{}, &Resource{
        Route: "/Products",
        Paginated: true,
    })
    API_RESOURCE_MAP.RegisterResource(&ProductGroup{}, &Resource{
        Route: "/ProductGroups",
        Paginated: false,
    })
}
