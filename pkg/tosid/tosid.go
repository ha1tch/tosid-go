package tosid

import (
	internal_tosid "github.com/ha1tch/tosid-go/internal/tosid"
)

// Re-export types from internal package
type TOSID = internal_tosid.TOSID

// Re-export maps and constants
var (
	TaxonomyDomains      = internal_tosid.TaxonomyDomains
	TaxonomyTypes        = internal_tosid.TaxonomyTypes
	NetmaskDescriptions  = internal_tosid.NetmaskDescriptions
)

// Parse creates a TOSID from a string representation
func Parse(code string) (*TOSID, error) {
	parser := internal_tosid.NewParser()
	return parser.Parse(code)
}

// Create creates a new TOSID with the specified components
func Create(taxonomyCode, netmaskIndicator, identifier string) (*TOSID, error) {
	validator := internal_tosid.NewValidator()
	if err := validator.ValidateComponents(taxonomyCode, netmaskIndicator, identifier); err != nil {
		return nil, err
	}
	
	return &TOSID{
		TaxonomyCode:     taxonomyCode,
		NetmaskIndicator: netmaskIndicator,
		Identifier:       identifier,
	}, nil
}

// ValidateFormat validates the basic format of a TOSID code
func ValidateFormat(code string) error {
	validator := internal_tosid.NewValidator()
	return validator.ValidateFormat(code)
}

// GetClassification returns the classification description for a TOSID
func GetClassification(taxonomyCode, netmaskIndicator string) string {
	classifier := internal_tosid.NewTaxonomyClassifier()
	return classifier.GetFullClassification(taxonomyCode, netmaskIndicator)
}