---
page_title: "prosimo_ip_reputation Resource - terraform-provider-prosimo"
subcategory: ""
description: |-
  Use this resource to create/modify ip reputation settngs.
---

# prosimo_ip_reputation (Resource)

Use this resource to create/modify ip reputation settngs.

This resource is usually used along with `terraform-provider-prosimo`.



## Example Usage

```terraform
resource "prosimo_ip_reputation" "ip_reputation" {
  enabled        = true
  allowlist = ["1.1.1.3/16", "1.1.1.2/16", "1.1.1.4/16"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `enabled` (Boolean) Defaults to false, set it true to enable IP Reputation

### Optional

- `allowlist` (List of String) List of CIDRs to allow

### Read-Only

- `id` (String) The ID of this resource.
