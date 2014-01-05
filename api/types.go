package api

import (
    "errors"
    "fmt"
    "regexp"
    "strconv"

    "github.com/nu7hatch/gouuid"
)

type Guid struct {
    guid *uuid.UUID
}

func (g *Guid) UnmarshalJSON(value []byte) (err error) {
    buffer := string(value[1 : len(value) - 1])  // value is quoted
    g.guid, err = uuid.ParseHex(buffer)
    return
}

func (g *Guid) String() string {
    if g.guid != nil {
        return g.guid.String()
    } else {
        return "nil"
    }
}

type DateTime struct {
    ms int64
}

func (dt *DateTime) UnmarshalJSON(value []byte) error {
    s := string(value[1 : len(value) - 1])

    re := regexp.MustCompile(`^\\?\/Date\((-?[0-9]+)\)\\?\/`)
    md := re.FindStringSubmatch(s)
    if md == nil {
        return errors.New("Invalid DateTime string")
    }

    if ms, err := strconv.ParseInt(md[1], 10, 0); err != nil {
        return errors.New("Could not Parse DateTime string")
    } else {
        dt.ms = ms
    }

    return nil
}

type Product struct {
      AverageLandPrice float32
      Barcode string
      BinLocation string
      CanAutoAssemble bool
      CreatedBy string
      DefaultPurchasePrice float32
      DefaultSellPrice float32
      Depth float32
      Guid Guid
      Height float32
      IsAssembledProduct bool
      IsComponent bool
      LastCost float32
      LastModifiedOn DateTime
      MaxStockAlertLevel float32
      MinStockAlertLevel float32
      NeverDiminishing bool
      Notes string
      Obsolete bool
      PackSize float32
      ProductCode string
      ProductDescription string
      ProductGroup ProductGroup
      ReOrderPoint float32
      SellPriceTier1 SellPriceTier
      SellPriceTier2 SellPriceTier
      SellPriceTier3 SellPriceTier
      SellPriceTier4 SellPriceTier
      SellPriceTier5 SellPriceTier
      SellPriceTier6 SellPriceTier
      SellPriceTier7 SellPriceTier
      SellPriceTier8 SellPriceTier
      SellPriceTier9 SellPriceTier
      SellPriceTier10 SellPriceTier
      SourceId string
      SourceVariantParentId string
      Supplier Supplier
      Taxable bool
      TaxableSales bool
      UnitOfMeasure UnitOfMeasure
      Weight float32
      Width float32
      XeroSalesAccount string
      XeroSalesTaxCode string
      XeroSalesTaxRate float32
      XeroTaxCode string
      XeroTaxRate float32
}

type Pagination struct {
    NumberOfItems int
    PageSize      int
    PageNumber    int
    NumberOfPages int
}

type ProductPagination struct {
    Pagination Pagination
    Items      []Product
}

type ProductGroup struct {
    Guid      *Guid
    GroupName string
}

func (pg ProductGroup) String() string {
    return fmt.Sprintf("{ ProductGroup: %s %s }", pg.Guid, pg.GroupName)
}

type SellPriceTier struct {
    Name  string
    Value string
}

type Supplier struct {
    Guid         *Guid
    SupplierCode string
    SupplierName string
}

func (s Supplier) String() string {
    return fmt.Sprintf("{ Supplier: %s %s %s }", s.Guid, s.SupplierCode, s.SupplierName)
}

type UnitOfMeasure struct {
    Guid *uuid.UUID
    Name string
}

func (um UnitOfMeasure) String() string {
    return fmt.Sprintf("{ UnitOfMeasure: %s %s }", um.Guid, um.Name)
}
