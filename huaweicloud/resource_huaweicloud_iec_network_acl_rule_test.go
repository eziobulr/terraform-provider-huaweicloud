package huaweicloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/huaweicloud/golangsdk/openstack/iec/v1/firewalls"
)

func TestAccIecNetworkACLRuleResource_basic(t *testing.T) {
	aclResourceName := "huaweicloud_iec_network_acl.acl_test"
	aclRuleResourceName := "huaweicloud_iec_network_acl_rule.rule_test"
	var fwGroup firewalls.RespFirewallRulesEntity

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIecNetworkACLDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIecNetworkACLRule_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIecNetworkACLRuleExists(aclResourceName, "ingress", &fwGroup),
					resource.TestCheckResourceAttrPtr(aclRuleResourceName, "protocol", &fwGroup.Protocol),
					resource.TestCheckResourceAttrPtr(aclRuleResourceName, "action", &fwGroup.Action),
					resource.TestCheckResourceAttrPtr(aclRuleResourceName, "destination_port", &fwGroup.DstPort),
				),
			},
			{
				Config: testAccIecNetworkACLRule_basic_update(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIecNetworkACLRuleExists(aclResourceName, "ingress", &fwGroup),
					resource.TestCheckResourceAttrPtr(aclRuleResourceName, "protocol", &fwGroup.Protocol),
					resource.TestCheckResourceAttrPtr(aclRuleResourceName, "action", &fwGroup.Action),
					resource.TestCheckResourceAttrPtr(aclRuleResourceName, "destination_port", &fwGroup.DstPort),
				),
			},
		},
	})
}

func testAccCheckIecNetworkACLRuleDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	iecV1Client, err := config.IECV1Client(HW_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating Huaweicloud IEC client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "huaweicloud_iec_network_acl" {
			continue
		}

		_, err := firewalls.Get(iecV1Client, rs.Primary.ID).Extract()
		if err == nil {
			return fmt.Errorf("IEC network acl still exists")
		}
	}

	return nil
}

func testAccCheckIecNetworkACLRuleExists(n, direction string, resource *firewalls.RespFirewallRulesEntity) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		iecV1Client, err := config.IECV1Client(HW_REGION_NAME)
		if err != nil {
			return fmt.Errorf("Error creating Huaweicloud IEC client: %s", err)
		}

		found, err := firewalls.Get(iecV1Client, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if direction == "ingress" {
			*resource = found.IngressFWPolicy.FirewallRules[0]
		} else if direction == "egress" {
			*resource = found.EgressFWPolicy.FirewallRules[0]
		} else {
			return fmt.Errorf("IEC Network ACL Rule not found")
		}

		return nil
	}
}

func testAccIecNetworkACLRule_basic() string {
	return fmt.Sprintf(`
resource "huaweicloud_iec_network_acl" "acl_test" {
  name = "iec-acl-basic"
}

resource "huaweicloud_iec_network_acl_rule" "rule_test" {
  network_acl_id         = huaweicloud_iec_network_acl.acl_test.id
  direction              = "ingress"
  protocol               = "tcp"
  action                 = "allow"
  source_ip_address      = "0.0.0.0/0"
  destination_ip_address = "0.0.0.0/0"
  destination_port       = "445"
  enabled                = true
}
`)
}

func testAccIecNetworkACLRule_basic_update() string {
	return fmt.Sprintf(`
resource "huaweicloud_iec_network_acl" "acl_test" {
  name = "iec-acl-update"
}

resource "huaweicloud_iec_network_acl_rule" "rule_test" {
  network_acl_id         = huaweicloud_iec_network_acl.acl_test.id
  direction              = "ingress"
  protocol               = "udp"
  action                 = "deny"
  source_ip_address      = "0.0.0.0/0"
  destination_ip_address = "0.0.0.0/0"
  destination_port       = "23-30"
  enabled                = true
}
`)
}
