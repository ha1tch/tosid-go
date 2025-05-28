package tosid

import (
	"regexp"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	// Test successful parsing
	tosidCode := "00B2-SOL-STR-SUN:000-000-000-001"
	tosid, err := Parse(tosidCode)
	if err != nil {
		t.Errorf("Failed to parse valid TOSID code: %v", err)
	}

	if tosid.TaxonomyCode != "00" {
		t.Errorf("Expected taxonomy code 00, got %s", tosid.TaxonomyCode)
	}

	if tosid.NetmaskIndicator != "B" {
		t.Errorf("Expected netmask indicator B, got %s", tosid.NetmaskIndicator)
	}

	if !strings.Contains(tosid.Identifier, "SOL-STR-SUN") {
		t.Errorf("Expected identifier to contain SOL-STR-SUN, got %s", tosid.Identifier)
	}

	// Test parsing failure
	invalidCode := "invalid-tosid-code"
	_, err = Parse(invalidCode)
	if err == nil {
		t.Errorf("Expected error for invalid TOSID code, got nil")
	}
}

func TestCreate(t *testing.T) {
	// Test successful creation
	tosid, err := Create("00", "B", "SOL-SYS-ERT:000-000-000-001")
	if err != nil {
		t.Errorf("Failed to create valid TOSID: %v", err)
	}

	expectedString := "00B-SOL-SYS-ERT:000-000-000-001"
	if tosid.String() != expectedString {
		t.Errorf("Expected string %s, got %s", expectedString, tosid.String())
	}

	// Test creation failures
	testCases := []struct {
		taxonomyCode    string
		netmaskIndicator string
		identifier      string
		expectError     bool
		errorPattern    string
	}{
		{"0", "B", "SOL-SYS-ERT", true, "taxonomy code must be exactly 2 characters"},
		{"00", "BB", "SOL-SYS-ERT", true, "netmask indicator must be exactly 1 character"},
		{"00", "B", "INVALID", true, "identifier format is invalid"},
		{"00", "1", "SOL-SYS-ERT", true, "netmask indicator must be a letter A-Z"},
		{"2X", "B", "SOL-SYS-ERT", true, "first taxonomy digit must be 0 or 1"},
	}

	for _, tc := range testCases {
		_, err := Create(tc.taxonomyCode, tc.netmaskIndicator, tc.identifier)
		if tc.expectError && err == nil {
			t.Errorf("Expected error for %s, %s, %s but got nil", tc.taxonomyCode, tc.netmaskIndicator, tc.identifier)
		} else if tc.expectError && !regexp.MustCompile(tc.errorPattern).MatchString(err.Error()) {
			t.Errorf("Expected error matching '%s', got '%s'", tc.errorPattern, err.Error())
		} else if !tc.expectError && err != nil {
			t.Errorf("Expected no error for %s, %s, %s but got %v", tc.taxonomyCode, tc.netmaskIndicator, tc.identifier, err)
		}
	}
}

func TestClassificationDescription(t *testing.T) {
	tosidObj, err := Parse("00B2-SOL-STR-SUN:000-000-000-001")
	if err != nil {
		t.Fatalf("Failed to parse TOSID: %v", err)
	}

	description := tosidObj.ClassificationDescription()
	expected := "Celestial/Natural - Physical/Material - Stellar Scale"
	if description != expected {
		t.Errorf("Expected description '%s', got '%s'", expected, description)
	}

	// Test with different taxonomy/netmask
	tosidObj, err = Parse("11A3-SCI-PHY-EIN:THE-REL-100")
	if err != nil {
		t.Fatalf("Failed to parse TOSID: %v", err)
	}

	description = tosidObj.ClassificationDescription()
	expected = "Artificial/Intelligent - Conceptual/Abstract - Civilizational Systems"
	if description != expected {
		t.Errorf("Expected description '%s', got '%s'", expected, description)
	}
}

func TestIsCompatibleWith(t *testing.T) {
	sun, err := Parse("00B2-SOL-STR-SUN:000-000-000-001")
	if err != nil {
		t.Fatalf("Failed to parse Sun TOSID: %v", err)
	}

	earth, err := Parse("00B3-SOL-SYS-ERT:000-000-000-001")
	if err != nil {
		t.Fatalf("Failed to parse Earth TOSID: %v", err)
	}

	mars, err := Parse("00B3-SOL-SYS-MRS:000-000-000-001")
	if err != nil {
		t.Fatalf("Failed to parse Mars TOSID: %v", err)
	}

	// Earth and Mars should be compatible (same taxonomy and netmask)
	if !earth.IsCompatibleWith(mars) {
		t.Error("Expected Earth and Mars to be compatible")
	}

	// Sun and Earth should not be compatible (different netmasks)
	if sun.IsCompatibleWith(earth) {
		t.Error("Expected Sun and Earth to be incompatible")
	}
}

func TestMatchesPattern(t *testing.T) {
	sun, err := Parse("00B2-SOL-STR-SUN:000-000-000-001")
	if err != nil {
		t.Fatalf("Failed to parse Sun TOSID: %v", err)
	}

	patterns := []struct {
		pattern string
		matches bool
	}{
		{"00B", true},       // Domain and netmask match
		{"00B2", true},      // Domain, netmask and first level match
		{"00*", true},       // Domain match with wildcard
		{"*SOL*", true},     // Wildcard with substring match
		{"11*", false},      // Different domain
		{"00C*", false},     // Different netmask
		{"", true},          // Empty pattern matches everything
		{"00B2-SOL-STR-SUN:000-000-000-001", true}, // Exact match
	}

	for _, p := range patterns {
		result := sun.MatchesPattern(p.pattern)
		if result != p.matches {
			t.Errorf("Pattern '%s' expected match=%v, got %v", p.pattern, p.matches, result)
		}
	}
}
