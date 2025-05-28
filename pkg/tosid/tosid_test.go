package tosid

import (
	"testing"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
		taxonomy string
		netmask  string
	}{
		{"00B2-SOL-STR-SUN:000-000-000-001", true, "00", "B"},
		{"11A3-SCI-PHY-EIN:THE-REL-100", true, "11", "A"},
		{"invalid-tosid", false, "", ""},
		{"00X2-SOL-STR-SUN:000-000-000-001", false, "", ""},
	}

	for _, tc := range testCases {
		tosid, err := Parse(tc.input)
		if tc.expected && err != nil {
			t.Errorf("Expected to parse %s successfully, got error: %v", tc.input, err)
		}
		if !tc.expected && err == nil {
			t.Errorf("Expected parsing %s to fail, but it succeeded", tc.input)
		}
		if tc.expected && tosid != nil {
			if tosid.TaxonomyCode != tc.taxonomy {
				t.Errorf("Expected taxonomy %s, got %s", tc.taxonomy, tosid.TaxonomyCode)
			}
			if tosid.NetmaskIndicator != tc.netmask {
				t.Errorf("Expected netmask %s, got %s", tc.netmask, tosid.NetmaskIndicator)
			}
		}
	}
}

func TestCreate(t *testing.T) {
	tosid, err := Create("00", "B", "SOL-STR-SUN:000-000-000-001")
	if err != nil {
		t.Fatalf("Failed to create TOSID: %v", err)
	}

	expected := "00B-SOL-STR-SUN:000-000-000-001"
	if tosid.String() != expected {
		t.Errorf("Expected %s, got %s", expected, tosid.String())
	}
}

func BenchmarkParse(b *testing.B) {
	tosidCode := "00B2-SOL-STR-SUN:000-000-000-001"
	for i := 0; i < b.N; i++ {
		_, err := Parse(tosidCode)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCreate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Create("00", "B", "SOL-STR-SUN:000-000-000-001")
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkPatternMatch(b *testing.B) {
	tosid, _ := Parse("00B2-SOL-STR-SUN:000-000-000-001")
	pattern := "00B*"
	
	for i := 0; i < b.N; i++ {
		tosid.MatchesPattern(pattern)
	}
}