package tosid

import (
	"errors"
	"regexp"
	"strings"
)

// Validator provides validation utilities for TOSID codes
type Validator struct {
	classifier *TaxonomyClassifier
}

// NewValidator creates a new TOSID validator
func NewValidator() *Validator {
	return &Validator{
		classifier: NewTaxonomyClassifier(),
	}
}

// ValidateFormat validates the basic format of a TOSID code
func (v *Validator) ValidateFormat(code string) error {
	pattern := regexp.MustCompile(`^(\d{2})([A-Z])(\d?[A-Z]{2}-[A-Z]{3}-[A-Z]{3})(:[A-Z0-9]{3}-[A-Z0-9]{3}-[A-Z0-9]{3}-[A-Z0-9]{3})?$`)
	
	if !pattern.MatchString(code) {
		return errors.New("invalid TOSID format")
	}
	
	return nil
}

// ValidateComponents validates the components of a TOSID
func (v *Validator) ValidateComponents(taxonomyCode, netmaskIndicator, identifier string) error {
	// Validate taxonomy code
	if err := v.ValidateTaxonomyCode(taxonomyCode); err != nil {
		return err
	}
	
	// Validate netmask indicator
	if err := v.ValidateNetmaskIndicator(taxonomyCode, netmaskIndicator); err != nil {
		return err
	}
	
	// Validate identifier
	if err := v.ValidateIdentifier(identifier); err != nil {
		return err
	}
	
	return nil
}

// ValidateTaxonomyCode validates a taxonomy code
func (v *Validator) ValidateTaxonomyCode(taxonomyCode string) error {
	if len(taxonomyCode) != 2 {
		return errors.New("taxonomy code must be exactly 2 characters")
	}
	
	if !strings.ContainsAny(string(taxonomyCode[0]), "01") {
		return errors.New("first taxonomy digit must be 0 or 1")
	}
	
	if !strings.ContainsAny(string(taxonomyCode[1]), "01") {
		return errors.New("second taxonomy digit must be 0 or 1")
	}
	
	if !v.classifier.IsValidTaxonomyCode(taxonomyCode) {
		return errors.New("invalid taxonomy code")
	}
	
	return nil
}

// ValidateNetmaskIndicator validates a netmask indicator
func (v *Validator) ValidateNetmaskIndicator(taxonomyCode, netmaskIndicator string) error {
	if len(netmaskIndicator) != 1 {
		return errors.New("netmask indicator must be exactly 1 character")
	}
	
	if netmaskIndicator < "A" || netmaskIndicator > "Z" {
		return errors.New("netmask indicator must be a letter A-Z")
	}
	
	if !v.classifier.IsValidNetmaskIndicator(taxonomyCode, netmaskIndicator) {
		return errors.New("invalid netmask indicator for this taxonomy code")
	}
	
	return nil
}

// ValidateIdentifier validates an identifier
func (v *Validator) ValidateIdentifier(identifier string) error {
	if identifier == "" {
		return errors.New("identifier cannot be empty")
	}
	
	// Basic validation of identifier structure
	pattern := regexp.MustCompile(`^[A-Z0-9]{3}-[A-Z0-9]{3}-[A-Z0-9]{3}(:[A-Z0-9]{3}-[A-Z0-9]{3}-[A-Z0-9]{3}-[A-Z0-9]{3})?$`)
	if !pattern.MatchString(identifier) {
		return errors.New("identifier format is invalid")
	}
	
	return nil
}

// ValidateSemanticConsistency performs semantic consistency checks
func (v *Validator) ValidateSemanticConsistency(tosid *TOSID) []string {
	var warnings []string
	
	// Check for common semantic inconsistencies
	if strings.Contains(tosid.Identifier, "ART") && tosid.TaxonomyCode[:1] == "0" {
		warnings = append(warnings, "identifier suggests artificial entity but taxonomy indicates natural")
	}
	
	if strings.Contains(tosid.Identifier, "NAT") && tosid.TaxonomyCode[:1] == "1" {
		warnings = append(warnings, "identifier suggests natural entity but taxonomy indicates artificial")
	}
	
	// Check scale consistency
	if tosid.NetmaskIndicator == "F" && strings.Contains(tosid.Identifier, "GAL") {
		warnings = append(warnings, "microscopic scale inconsistent with galactic identifier")
	}
	
	if tosid.NetmaskIndicator == "A" && strings.Contains(tosid.Identifier, "MOL") {
		warnings = append(warnings, "cosmic scale inconsistent with molecular identifier")
	}
	
	return warnings
}

// IsWellFormed checks if a TOSID is well-formed according to all rules
func (v *Validator) IsWellFormed(tosid *TOSID) (bool, []string) {
	var errors []string
	
	// Basic component validation
	if err := v.ValidateComponents(tosid.TaxonomyCode, tosid.NetmaskIndicator, tosid.Identifier); err != nil {
		errors = append(errors, err.Error())
	}
	
	// Semantic consistency warnings
	warnings := v.ValidateSemanticConsistency(tosid)
	errors = append(errors, warnings...)
	
	return len(errors) == 0, errors
}