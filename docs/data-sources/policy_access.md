---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "prosimo_policy_access Data Source - terraform-provider-prosimo"
subcategory: ""
description: |-
  Use this data source to get information on existing access policies.
---

# prosimo_policy_access (Data Source)

Use this data source to get information on existing access policies.



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `filter` (String) Custom filters to scope specific results. Usage: filter = app_access_type==agent

### Read-Only

- `id` (String) The ID of this resource.
- `policy` (List of Object) (see [below for nested schema](#nestedatt--policy))

<a id="nestedatt--policy"></a>
### Nested Schema for `policy`

Read-Only:

- `app_access_type` (String)
- `createdtime` (String)
- `details` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details))
- `device_posture_configured` (Boolean)
- `id` (String)
- `name` (String)
- `teamid` (String)
- `type` (String)
- `updatedtime` (String)

<a id="nestedobjatt--policy--details"></a>
### Nested Schema for `policy.details`

Read-Only:

- `actions` (List of String)
- `apps` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--apps))
- `matches` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--matches))
- `networks` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--networks))

<a id="nestedobjatt--policy--details--apps"></a>
### Nested Schema for `policy.details.apps`

Read-Only:

- `selecteditems` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--apps--selecteditems))

<a id="nestedobjatt--policy--details--apps--selecteditems"></a>
### Nested Schema for `policy.details.apps.selecteditems`

Read-Only:

- `id` (String)
- `name` (String)



<a id="nestedobjatt--policy--details--matches"></a>
### Nested Schema for `policy.details.matches`

Read-Only:

- `advanced` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--matches--advanced))
- `devicepostureprofiles` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--matches--devicepostureprofiles))
- `devices` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--matches--devices))
- `fqdn` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--matches--fqdn))
- `idp` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--matches--idp))
- `networks` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--matches--networks))
- `time` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--matches--time))
- `url` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--matches--url))
- `users` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--matches--users))

<a id="nestedobjatt--policy--details--matches--advanced"></a>
### Nested Schema for `policy.details.matches.users`

Read-Only:

- `operations` (String)
- `property` (String)
- `values` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--matches--users--values))

<a id="nestedobjatt--policy--details--matches--users--values"></a>
### Nested Schema for `policy.details.matches.users.values`

Read-Only:

- `inputitems` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--matches--users--values--inputitems))
- `selecteditems` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--matches--users--values--selecteditems))

<a id="nestedobjatt--policy--details--matches--users--values--inputitems"></a>
### Nested Schema for `policy.details.matches.users.values.selecteditems`

Read-Only:

- `id` (String)
- `name` (String)


<a id="nestedobjatt--policy--details--matches--users--values--selecteditems"></a>
### Nested Schema for `policy.details.matches.users.values.selecteditems`

Read-Only:

- `id` (String)
- `name` (String)




<a id="nestedobjatt--policy--details--matches--devicepostureprofiles"></a>
### Nested Schema for `policy.details.matches.users`

Read-Only:

- `operations` (String)
- `property` (String)
- `values` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--matches--users--values))

<a id="nestedobjatt--policy--details--matches--users--values"></a>
### Nested Schema for `policy.details.matches.users.values`

Read-Only:

- `inputitems` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--matches--users--values--inputitems))
- `selecteditems` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--matches--users--values--selecteditems))

<a id="nestedobjatt--policy--details--matches--users--values--inputitems"></a>
### Nested Schema for `policy.details.matches.users.values.selecteditems`

Read-Only:

- `id` (String)
- `name` (String)


<a id="nestedobjatt--policy--details--matches--users--values--selecteditems"></a>
### Nested Schema for `policy.details.matches.users.values.selecteditems`

Read-Only:

- `id` (String)
- `name` (String)




<a id="nestedobjatt--policy--details--matches--devices"></a>
### Nested Schema for `policy.details.matches.users`

Read-Only:

- `operations` (String)
- `property` (String)
- `values` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--matches--users--values))

<a id="nestedobjatt--policy--details--matches--users--values"></a>
### Nested Schema for `policy.details.matches.users.values`

Read-Only:

- `inputitems` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--matches--users--values--inputitems))
- `selecteditems` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--matches--users--values--selecteditems))

<a id="nestedobjatt--policy--details--matches--users--values--inputitems"></a>
### Nested Schema for `policy.details.matches.users.values.selecteditems`

Read-Only:

- `id` (String)
- `name` (String)


<a id="nestedobjatt--policy--details--matches--users--values--selecteditems"></a>
### Nested Schema for `policy.details.matches.users.values.selecteditems`

Read-Only:

- `id` (String)
- `name` (String)




<a id="nestedobjatt--policy--details--matches--fqdn"></a>
### Nested Schema for `policy.details.matches.users`

Read-Only:

- `operations` (String)
- `property` (String)
- `values` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--matches--users--values))

<a id="nestedobjatt--policy--details--matches--users--values"></a>
### Nested Schema for `policy.details.matches.users.values`

Read-Only:

- `inputitems` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--matches--users--values--inputitems))
- `selecteditems` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--matches--users--values--selecteditems))

<a id="nestedobjatt--policy--details--matches--users--values--inputitems"></a>
### Nested Schema for `policy.details.matches.users.values.selecteditems`

Read-Only:

- `id` (String)
- `name` (String)


<a id="nestedobjatt--policy--details--matches--users--values--selecteditems"></a>
### Nested Schema for `policy.details.matches.users.values.selecteditems`

Read-Only:

- `id` (String)
- `name` (String)




<a id="nestedobjatt--policy--details--matches--idp"></a>
### Nested Schema for `policy.details.matches.users`

Read-Only:

- `operations` (String)
- `property` (String)
- `values` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--matches--users--values))

<a id="nestedobjatt--policy--details--matches--users--values"></a>
### Nested Schema for `policy.details.matches.users.values`

Read-Only:

- `inputitems` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--matches--users--values--inputitems))
- `selecteditems` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--matches--users--values--selecteditems))

<a id="nestedobjatt--policy--details--matches--users--values--inputitems"></a>
### Nested Schema for `policy.details.matches.users.values.selecteditems`

Read-Only:

- `id` (String)
- `name` (String)


<a id="nestedobjatt--policy--details--matches--users--values--selecteditems"></a>
### Nested Schema for `policy.details.matches.users.values.selecteditems`

Read-Only:

- `id` (String)
- `name` (String)




<a id="nestedobjatt--policy--details--matches--networks"></a>
### Nested Schema for `policy.details.matches.users`

Read-Only:

- `operations` (String)
- `property` (String)
- `values` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--matches--users--values))

<a id="nestedobjatt--policy--details--matches--users--values"></a>
### Nested Schema for `policy.details.matches.users.values`

Read-Only:

- `inputitems` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--matches--users--values--inputitems))
- `selecteditems` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--matches--users--values--selecteditems))

<a id="nestedobjatt--policy--details--matches--users--values--inputitems"></a>
### Nested Schema for `policy.details.matches.users.values.selecteditems`

Read-Only:

- `id` (String)
- `name` (String)


<a id="nestedobjatt--policy--details--matches--users--values--selecteditems"></a>
### Nested Schema for `policy.details.matches.users.values.selecteditems`

Read-Only:

- `id` (String)
- `name` (String)




<a id="nestedobjatt--policy--details--matches--time"></a>
### Nested Schema for `policy.details.matches.users`

Read-Only:

- `operations` (String)
- `property` (String)
- `values` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--matches--users--values))

<a id="nestedobjatt--policy--details--matches--users--values"></a>
### Nested Schema for `policy.details.matches.users.values`

Read-Only:

- `inputitems` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--matches--users--values--inputitems))
- `selecteditems` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--matches--users--values--selecteditems))

<a id="nestedobjatt--policy--details--matches--users--values--inputitems"></a>
### Nested Schema for `policy.details.matches.users.values.selecteditems`

Read-Only:

- `id` (String)
- `name` (String)


<a id="nestedobjatt--policy--details--matches--users--values--selecteditems"></a>
### Nested Schema for `policy.details.matches.users.values.selecteditems`

Read-Only:

- `id` (String)
- `name` (String)




<a id="nestedobjatt--policy--details--matches--url"></a>
### Nested Schema for `policy.details.matches.users`

Read-Only:

- `operations` (String)
- `property` (String)
- `values` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--matches--users--values))

<a id="nestedobjatt--policy--details--matches--users--values"></a>
### Nested Schema for `policy.details.matches.users.values`

Read-Only:

- `inputitems` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--matches--users--values--inputitems))
- `selecteditems` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--matches--users--values--selecteditems))

<a id="nestedobjatt--policy--details--matches--users--values--inputitems"></a>
### Nested Schema for `policy.details.matches.users.values.selecteditems`

Read-Only:

- `id` (String)
- `name` (String)


<a id="nestedobjatt--policy--details--matches--users--values--selecteditems"></a>
### Nested Schema for `policy.details.matches.users.values.selecteditems`

Read-Only:

- `id` (String)
- `name` (String)




<a id="nestedobjatt--policy--details--matches--users"></a>
### Nested Schema for `policy.details.matches.users`

Read-Only:

- `operations` (String)
- `property` (String)
- `values` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--matches--users--values))

<a id="nestedobjatt--policy--details--matches--users--values"></a>
### Nested Schema for `policy.details.matches.users.values`

Read-Only:

- `inputitems` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--matches--users--values--inputitems))
- `selecteditems` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--matches--users--values--selecteditems))

<a id="nestedobjatt--policy--details--matches--users--values--inputitems"></a>
### Nested Schema for `policy.details.matches.users.values.selecteditems`

Read-Only:

- `id` (String)
- `name` (String)


<a id="nestedobjatt--policy--details--matches--users--values--selecteditems"></a>
### Nested Schema for `policy.details.matches.users.values.selecteditems`

Read-Only:

- `id` (String)
- `name` (String)





<a id="nestedobjatt--policy--details--networks"></a>
### Nested Schema for `policy.details.networks`

Read-Only:

- `selecteditems` (Set of Object) (see [below for nested schema](#nestedobjatt--policy--details--networks--selecteditems))

<a id="nestedobjatt--policy--details--networks--selecteditems"></a>
### Nested Schema for `policy.details.networks.selecteditems`

Read-Only:

- `id` (String)
- `name` (String)

