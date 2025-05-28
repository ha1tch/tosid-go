package tosid

import (
	"errors"
	"regexp"
)

// Parser handles parsing of TOSID codes
type Parser struct {
	pattern *regexp.Regexp
}

// NewParser creates a new TOSID parser
func NewParser() *Parser {
	pattern := regexp.MustCompile(`^(\d{2})([A-Z])(\d?[A-Z]{2}-[A-Z]{3}-[A-Z]{3})(:[A-Z0-9]{3}-[A-Z0-9]{3}-[A-Z0-9]{3}-[A-Z0-9]{3})?$`)
	return &Parser{
		pattern: pattern,
	}
}

// Parse creates a TOSID from a string representation
func (p *Parser) Parse(code string) (*TOSID, error) {
	matches := p.pattern.FindStringSubmatch(code)

	if matches == nil {
		return nil, errors.New("invalid TOSID format")
	}

	taxonomyCode := matches[1]
	netmaskIndicator := matches[2]
	categoryIdentifier := matches[3]
	specificIdentifier := ""

	if len(matches) > 4 && matches[4] != "" {
		specificIdentifier = matches[4]
	}

	identifier := categoryIdentifier
	if specificIdentifier != "" {
		identifier += specificIdentifier
	}

	return &TOSID{
		TaxonomyCode:     taxonomyCode,
		NetmaskIndicator: netmaskIndicator,
		Identifier:       identifier,
	}, nil
}

// ParseBatch parses multiple TOSID codes
func (p *Parser) ParseBatch(codes []string) ([]*TOSID, []error) {
	tosids := make([]*TOSID, 0, len(codes))
	errors := make([]error, 0)

	for _, code := range codes {
		tosid, err := p.Parse(code)
		if err != nil {
			errors = append(errors, err)
		} else {
			tosids = append(tosids, tosid)
		}
	}

	return tosids, errors
}

// ValidateFormat checks if a string matches TOSID format without full parsing
func (p *Parser) ValidateFormat(code string) bool {
	return p.pattern.MatchString(code)
}

// ExtractComponents extracts the main components without creating a TOSID object
func (p *Parser) ExtractComponents(code string) (taxonomyCode, netmaskIndicator, identifier string, err error) {
	matches := p.pattern.FindStringSubmatch(code)

	if matches == nil {
		return "", "", "", errors.New("invalid TOSID format")
	}

	taxonomyCode = matches[1]
	netmaskIndicator = matches[2]
	categoryIdentifier := matches[3]
	specificIdentifier := ""

	if len(matches) > 4 && matches[4] != "" {
		specificIdentifier = matches[4]
	}

	identifier = categoryIdentifier
	if specificIdentifier != "" {
		identifier += specificIdentifier
	}

	return taxonomyCode, netmaskIndicator, identifier, nil
}
