// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccBeyondcorpAppConnection_beyondcorpAppConnectionBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBeyondcorpAppConnectionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBeyondcorpAppConnection_beyondcorpAppConnectionBasicExample(context),
			},
			{
				ResourceName:            "google_beyondcorp_app_connection.app_connection",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "region"},
			},
		},
	})
}

func testAccBeyondcorpAppConnection_beyondcorpAppConnectionBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_service_account" "service_account" {
  account_id   = "tf-test-my-account%{random_suffix}"
  display_name = "Test Service Account"
}

resource "google_beyondcorp_app_connector" "app_connector" {
  name = "tf-test-my-app-connector%{random_suffix}"
  principal_info {
    service_account {
     email = google_service_account.service_account.email
    }
  }
}

resource "google_beyondcorp_app_connection" "app_connection" {
  name = "tf-test-my-app-connection%{random_suffix}"
  type = "TCP_PROXY"
  application_endpoint {
    host = "foo-host"
    port = 8080
  }
  connectors = [google_beyondcorp_app_connector.app_connector.id]
}
`, context)
}

func TestAccBeyondcorpAppConnection_beyondcorpAppConnectionFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBeyondcorpAppConnectionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBeyondcorpAppConnection_beyondcorpAppConnectionFullExample(context),
			},
			{
				ResourceName:            "google_beyondcorp_app_connection.app_connection",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "region"},
			},
		},
	})
}

func testAccBeyondcorpAppConnection_beyondcorpAppConnectionFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_service_account" "service_account" {
  account_id   = "tf-test-my-account%{random_suffix}"
  display_name = "Test Service Account"
}

resource "google_beyondcorp_app_gateway" "app_gateway" {
  name = "tf-test-my-app-gateway%{random_suffix}"
  type = "TCP_PROXY"
  host_type = "GCP_REGIONAL_MIG"
}

resource "google_beyondcorp_app_connector" "app_connector" {
  name = "tf-test-my-app-connector%{random_suffix}"
  principal_info {
    service_account {
     email = google_service_account.service_account.email
    }
  }
}

resource "google_beyondcorp_app_connection" "app_connection" {
  name = "tf-test-my-app-connection%{random_suffix}"
  type = "TCP_PROXY"
  display_name = "some display name%{random_suffix}"
  application_endpoint {
    host = "foo-host"
    port = 8080
  }
  connectors = [google_beyondcorp_app_connector.app_connector.id]
  gateway {
    app_gateway = google_beyondcorp_app_gateway.app_gateway.id
  }
  labels = {
    foo = "bar"
    bar = "baz"
  }
}
`, context)
}

func testAccCheckBeyondcorpAppConnectionDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_beyondcorp_app_connection" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{BeyondcorpBasePath}}projects/{{project}}/locations/{{region}}/appConnections/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("BeyondcorpAppConnection still exists at %s", url)
			}
		}

		return nil
	}
}