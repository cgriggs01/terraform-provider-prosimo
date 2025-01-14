---
page_title: "prosimo_ip_addresses Resource - terraform-provider-prosimo"
subcategory: ""
description: |-
  Use this resource to create/modify ip ranges.
---

# prosimo_ip_addresses (Resource)

Use this resource to create/modify ip ranges.

This resource is usually used along with `terraform-provider-prosimo`.



## Example Usage

```terraform
resource "prosimo_ip_addresses" "aws_ip_address" {
  cidr        = "172.16.0.0/16"
  cloud_type = "AWS"
}

output "aws_ip_address" {
  description = "AWS IP Address"
  value       = prosimo_ip_addresses.aws_ip_address
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cidr` (String) CIDR range
- `cloud_type` (String) Cloud Service Provider, e.g: AWS, AZURE, GCP

### Read-Only

- `id` (String) The ID of this resource.
- `name` (String) Name of cloud account
- `subnets_in_use` (Number) Subnets in use
- `total_subnets` (Number) Total Number of Generated Subnets

