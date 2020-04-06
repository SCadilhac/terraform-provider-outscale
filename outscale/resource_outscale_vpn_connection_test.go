package outscale

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/antihax/optional"
	oscgo "github.com/marinsalinas/osc-sdk-go"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccOutscaleVPNConnection_basic(t *testing.T) {
	resourceName := "outscale_vpn_connection.foo"

	publicIP := fmt.Sprintf("172.0.0.%d", acctest.RandIntRange(1, 255))

	log.Printf("publicIP: %#+v\n", publicIP)
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		IDRefreshName: resourceName,
		Providers:     testAccProviders,
		CheckDestroy:  testAccOutscaleVPNConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccOutscaleVPNConnectionConfig(publicIP),
				Check: resource.ComposeTestCheckFunc(
					testAccOutscaleVPNConnectionExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "client_gateway_id"),
					resource.TestCheckResourceAttrSet(resourceName, "virtual_gateway_id"),
					resource.TestCheckResourceAttrSet(resourceName, "connection_type"),

					resource.TestCheckResourceAttr(resourceName, "connection_type", "ipsec.1"),
				),
			},
		},
	})
}

func TestAccOutscaleVPNConnection_withoutStaticRoutes(t *testing.T) {
	resourceName := "outscale_vpn_connection.foo"
	publicIP := fmt.Sprintf("172.0.0.%d", acctest.RandIntRange(0, 255))

	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		IDRefreshName: "outscale_vpn_connection.foo",
		Providers:     testAccProviders,
		CheckDestroy:  testAccOutscaleVPNConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccOutscaleVPNConnectionConfigWithoutStaticRoutes(publicIP),
				Check: resource.ComposeTestCheckFunc(
					testAccOutscaleVPNConnectionExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "client_gateway_id"),
					resource.TestCheckResourceAttrSet(resourceName, "virtual_gateway_id"),
					resource.TestCheckResourceAttrSet(resourceName, "connection_type"),

					resource.TestCheckResourceAttr(resourceName, "connection_type", "ipsec.1"),
				),
			},
		},
	})
}

func testAccOutscaleVPNConnectionExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		conn := testAccProvider.Meta().(*OutscaleClient).OSCAPI

		if rs.Primary.ID == "" {
			return fmt.Errorf("No VPN Connection ID is set")
		}

		filter := oscgo.ReadVpnConnectionsRequest{
			Filters: &oscgo.FiltersVpnConnection{
				VpnConnectionIds: &[]string{rs.Primary.ID},
			},
		}

		resp, _, err := conn.VpnConnectionApi.ReadVpnConnections(context.Background(), &oscgo.ReadVpnConnectionsOpts{
			ReadVpnConnectionsRequest: optional.NewInterface(filter),
		})
		if err != nil || len(resp.GetVpnConnections()) < 1 {
			return fmt.Errorf("Outscale VPN Connection not found (%s)", rs.Primary.ID)
		}
		return nil
	}
}

func testAccOutscaleVPNConnectionDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*OutscaleClient).OSCAPI
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "outscale_vpn_connection" {
			continue
		}

		filter := oscgo.ReadVpnConnectionsRequest{
			Filters: &oscgo.FiltersVpnConnection{
				VpnConnectionIds: &[]string{rs.Primary.ID},
			},
		}

		resp, _, err := conn.VpnConnectionApi.ReadVpnConnections(context.Background(), &oscgo.ReadVpnConnectionsOpts{
			ReadVpnConnectionsRequest: optional.NewInterface(filter),
		})
		if err != nil ||
			len(resp.GetVpnConnections()) > 0 && resp.GetVpnConnections()[0].GetState() != "deleted" {
			return fmt.Errorf("Outscale VPN Connection still exists (%s): %s", rs.Primary.ID, err)
		}
	}
	return nil
}

func testAccOutscaleVPNConnectionConfig(publicIP string) string {
	return fmt.Sprintf(`
		resource "outscale_virtual_gateway" "virtual_gateway" {
			connection_type = "ipsec.1"
		}

		resource "outscale_client_gateway" "customer_gateway" {
			bgp_asn         = 3
			public_ip       = "%s"
			connection_type = "ipsec.1"
		}

		resource "outscale_vpn_connection" "foo" {
			client_gateway_id  = "${outscale_client_gateway.customer_gateway.id}"
			virtual_gateway_id = "${outscale_virtual_gateway.virtual_gateway.id}"
			connection_type    = "ipsec.1"
			static_routes_only  = true
		}
	`, publicIP)
}

func testAccOutscaleVPNConnectionConfigWithoutStaticRoutes(publicIP string) string {
	return fmt.Sprintf(`
		resource "outscale_virtual_gateway" "virtual_gateway" {
			connection_type = "ipsec.1"
		}

		resource "outscale_client_gateway" "customer_gateway" {
			bgp_asn         = 3
			public_ip       = "%s"
			connection_type = "ipsec.1"
		}

		resource "outscale_vpn_connection" "foo" {
			client_gateway_id  = "${outscale_client_gateway.customer_gateway.id}"
			virtual_gateway_id = "${outscale_virtual_gateway.virtual_gateway.id}"
			connection_type    = "ipsec.1"
		}
	`, publicIP)
}
