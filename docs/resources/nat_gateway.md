---
subcategory: "NAT Gateway (NAT)"
---

# huaweicloud\_nat\_gateway

Manages a Nat gateway resource within HuaweiCloud Nat
This is an alternative to `huaweicloud_nat_gateway_v2`

## Example Usage

```hcl
resource "huaweicloud_nat_gateway" "nat_1" {
  name                = "Terraform"
  description         = "test for terraform"
  spec                = "3"
  router_id           = "2c1fe4bd-ebad-44ca-ae9d-e94e63847b75"
  internal_network_id = "dc8632e2-d9ff-41b1-aa0c-d455557314a0"
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the Nat gateway resource. If omitted, the provider-level region will be used. Changing this creates a new Nat gateway resource.

* `name` - (Required, String) The name of the nat gateway.

* `description` - (Optional, String) The description of the nat gateway.

* `spec` - (Required, String) The specification of the nat gateway, valid values are "1",
    "2", "3", "4".

* `tenant_id` - (Optional, String, ForceNew) The target tenant ID in which to allocate the nat
    gateway. Changing this creates a new nat gateway.

* `router_id` - (Required, String, ForceNew) ID of the router this nat gateway belongs to. Changing
    this creates a new nat gateway.

* `internal_network_id` - (Required, String, ForceNew) ID of the network this nat gateway connects to.
    Changing this creates a new nat gateway.

* `enterprise_project_id` - (Optional, String, ForceNew) The enterprise project id of the nat gateway. 
    Changing this creates a new nat gateway.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a resource ID in UUID format.

* `status` - The status of the nat gateway.

## Timeouts
This resource provides the following timeouts configuration options:
- `create` - Default is 10 minute.
- `delete` - Default is 10 minute.

## Import

Nat gateway can be imported using the following format:

```
$ terraform import huaweicloud_nat_gateway.nat_1 d126fb87-43ce-4867-a2ff-cf34af3765d9
```