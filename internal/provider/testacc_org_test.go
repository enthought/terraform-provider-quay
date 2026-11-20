package provider

import (
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
)

// testAccOrg returns the organization name used by an acceptance test, the HCL expression that
// references it, and the HCL that declares it.
//
// By default the test creates (and destroys) an organization called defaultName. When the
// QUAY_TEST_ORG environment variable is set, the test uses that existing organization instead and
// declares nothing, so nothing in the organization is created or deleted other than the test's own
// uniquely named resources (see testAccName).
func testAccOrg(defaultName, resourceName string) (name, ref, hcl string) {
	if existing := os.Getenv("QUAY_TEST_ORG"); existing != "" {
		return existing, fmt.Sprintf("%q", existing), ""
	}

	hcl = fmt.Sprintf(`
resource "quay_organization" %[1]q {
  name = %[2]q
  email = "quay+%[2]s@example.com"
}
`, resourceName, defaultName)
	return defaultName, fmt.Sprintf("quay_organization.%s.name", resourceName), hcl
}

// testAccName returns a unique, lowercase, Quay-safe name for robots and teams created by a test.
// The "tfacc" prefix makes leftovers easy to find and safe to sweep.
func testAccName(suffix string) string {
	return "tfacc" + acctest.RandStringFromCharSet(6, "abcdefghijklmnopqrstuvwxyz") + suffix
}
