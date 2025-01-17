---
layout: "outscale"
page_title: "OUTSCALE: outscale_product_type"
sidebar_current: "outscale-product-type"
description: |-
  [Provides information about a specific product type.]
---

# outscale_product_type Data Source

Provides information about a specific product type.
For more information on this resource, see the [User Guide](https://wiki.outscale.net/display/EN/Software+Licenses).
For more information on this resource actions, see the [API documentation](https://docs.outscale.com/api#3ds-outscale-api-producttype).

## Example Usage

```hcl
data "outscale_product_type" "product_type01" {
  filter {
    name   = "product_type_ids"
    values = ["0001"]
  }
}
```

## Argument Reference

The following arguments are supported:

* `filter` - (Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:
    * `product_type_ids` - (Optional) The IDs of the product types.

## Attribute Reference

The following attributes are exported:

* `description` - The description of the product type.
* `product_type_id` - The ID of the product type.
* `vendor` - The vendor of the product type.
