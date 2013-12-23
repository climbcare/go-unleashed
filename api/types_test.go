package api

import (
    . "launchpad.net/gocheck"

    "encoding/json"
    "log"
)

var _ = log.Println


type ProductSuite struct{}
var _ = Suite(&ProductSuite{})

func (s *ProductSuite) TestGuid(c *C) {
    var guid Guid
    err := json.Unmarshal([]byte(`"76b4e5d4-ff42-4785-93c5-a69a2980752d"`), &guid)

    c.Assert(err, IsNil)
}

func (s *ProductSuite) TestDateTime(c *C) {
    var dt DateTime
    err := json.Unmarshal([]byte(`"/Date(1378686019420)/"`), &dt)

    c.Assert(err, IsNil)
    c.Assert(dt.ms, Equals, int64(1378686019420))
}

func (s *ProductSuite) TestDateTimeEscaped(c *C) {
    var dt DateTime
    err := json.Unmarshal([]byte(`"\/Date(1378686019420)\/"`), &dt)

    c.Assert(err, IsNil)
    c.Assert(dt.ms, Equals, int64(1378686019420))
}

func (s *ProductSuite) TestProduct(c *C) {
    var p Product
    err := json.Unmarshal(productExample(), &p)

    c.Assert(err, IsNil)

    c.Assert(p.ProductCode, Equals, "_TEST01")
    c.Assert(p.SellPriceTier1.Name, Equals, "RRP")
    c.Assert(p.SellPriceTier1.Value, Equals, "15.0000")
}


func productExample() []byte {
    p := []byte(`{
        "ProductCode": "_TEST01",
        "ProductDescription": "Test Item 01",
        "Barcode": null,
        "PackSize": null,
        "Width": null,
        "Height": null,
        "Depth": null,
        "Weight": null,
        "MinStockAlertLevel": null,
        "MaxStockAlertLevel": null,
        "ReOrderPoint": null,
        "UnitOfMeasure": null,
        "NeverDiminishing": false,
        "LastCost": null,
        "DefaultPurchasePrice": 5,
        "DefaultSellPrice": 10,
        "AverageLandPrice": null,
        "Obsolete": false,
        "Notes": null,
        "SellPriceTier1": {
            "Name": "RRP",
            "Value": "15.0000"
        },
        "SellPriceTier2": {
            "Name": "",
            "Value": "0.0000"
        },
        "SellPriceTier3": {
            "Name": "Sell Price Tier 3",
            "Value": "0.0000"
        },
        "SellPriceTier4": {
            "Name": "Sell Price Tier 4",
            "Value": "0.0000"
        },
        "SellPriceTier5": {
            "Name": "Sell Price Tier 5",
            "Value": "0.0000"
        },
        "SellPriceTier6": {
            "Name": "Sell Price Tier 6",
            "Value": "0.0000"
        },
        "SellPriceTier7": {
            "Name": "Sell Price Tier 7",
            "Value": "0.0000"
        },
        "SellPriceTier8": {
            "Name": "Sell Price Tier 8",
            "Value": "0.0000"
        },
        "SellPriceTier9": {
            "Name": "Sell Price Tier 9",
            "Value": "0.0000"
        },
        "SellPriceTier10": {
            "Name": "Sell Price Tier 10",
            "Value": "0.0000"
        },
        "Taxable": false,
        "XeroTaxCode": "EXEMPTEXPENSES",
        "XeroTaxRate": 0,
        "TaxableSales": true,
        "XeroSalesTaxCode": "OUTPUT",
        "XeroSalesTaxRate": 0.1,
        "IsComponent": false,
        "IsAssembledProduct": false,
        "CanAutoAssemble": false,
        "ProductGroup": {
            "Guid": "e76c9496-1571-4ee7-bf75-334ccbf45d29",
            "GroupName": "_Test"
        },
        "XeroSalesAccount": null,
        "BinLocation": null,
        "Supplier": null,
        "SourceId": null,
        "CreatedBy": "admin@example.com",
        "SourceVariantParentId": null,
        "Guid": "76b4e5d4-ff42-4785-93c5-a69a2980752d",
        "LastModifiedOn": "\/Date(1378686019420)\/"
    }`)

    return p
}
