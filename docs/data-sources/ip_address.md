---
page_title: "prosimo_ip_address Data Source - terraform-provider-prosimo"
subcategory: ""
description: |-
  Use this data source to get information on available ip ranges.
---

# prosimo_ip_address (Data Source)

Use this data source to get information on available ip ranges.

This resource is usually used along with `terraform-provider-prosimo`.



## Example Usage

```terraform
data "prosimo_ip_address" "ip_details" {
  filter="cloudtype==AZURE"
  }

  output "ip_details" {
  description = "ip"
  value       = data.prosimo_ip_address.ip_details
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `filter` (String)

### Read-Only

- `id` (String) The ID of this resource.
- `ip_pools` (Set of Object) (see [below for nested schema](#nestedatt--ip_pools))

<a id="nestedatt--ip_pools"></a>
### Nested Schema for `ip_pools`

Read-Only:

- `cidr` (String)
- `cloudtype` (String)
- `id` (String)
- `name` (String)
- `subnetsinuse` (Number)
- `totalsubnets` (Number)


