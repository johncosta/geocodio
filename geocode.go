package geocodio

import (
	"errors"
	"strconv"
	"strings"
)

// Geocode single address
// See: http://geocod.io/docs/#toc_4
func (g *Geocodio) Geocode(address string) (GeocodeResult, error) {
	if address == "" {
		return GeocodeResult{}, errors.New("address must not be empty")
	}

	results, err := g.Call("/geocode", map[string]string{"q": address})
	if err != nil {
		return GeocodeResult{}, err
	}

	return results, nil
}

func (g *Geocodio) GeocodeByComponents(street, city, state, postalCode, country string, limit int) (GeocodeResult, error) {
	cs, err := CombineComponents(street, city, state, postalCode, country, limit)
	if err != nil {
		return GeocodeResult{}, err
	}
	results, err := g.Call("/geocode", cs)
	if err != nil {
		return GeocodeResult{}, err
	}

	return results, nil
}

func CombineComponents(street, city, state, postalCode, country string, limit int) (map[string]string, error) {
	if street == "" {
		return nil, errors.New("street must not be empty")
	}
	if city == "" {
		return nil, errors.New("city must not be empty")
	}
	if state == "" {
		return nil, errors.New("state must not be empty")
	}
	if postalCode == "" {
		return nil, errors.New("postal code must not be empty")
	}
	if country == "" {
		country = "USA"
	}

	cs := map[string]string{
		"street":      street,
		"city":        city,
		"state":       state,
		"postal_code": postalCode,
		"country":     country,
		"limit":       strconv.Itoa(limit),
	}

	if strconv.Itoa(limit) == "0" {
		delete(cs, "limit")
	}

	return cs, nil
}

// GeocodeAndReturnTimezone will geocode and include Timezone in the fields response
func (g *Geocodio) GeocodeAndReturnTimezone(address string) (GeocodeResult, error) {
	return g.GeocodeReturnFields(address, "timezone")
}

// GeocodeAndReturnTimezone will geocode and include Timezone in the fields response
func (g *Geocodio) GeocodeByComponentsAndReturnTimezone(street, city, state, postalCode, country string, limit int) (GeocodeResult, error) {
	cs, err := CombineComponents(street, city, state, postalCode, country, limit)
	if err != nil {
		return GeocodeResult{}, err
	}
	cs["fields"] = "timezone"
	return g.Call("/geocode", cs)
}

// GeocodeAndReturnCongressionalDistrict will geocode and include Congressional District in the fields response
func (g *Geocodio) GeocodeAndReturnCongressionalDistrict(address string) (GeocodeResult, error) {
	return g.GeocodeReturnFields(address, "cd")
}

// GeocodeAndReturnStateLegislativeDistricts will geocode and include State Legislative Districts in the fields response
func (g *Geocodio) GeocodeAndReturnStateLegislativeDistricts(address string) (GeocodeResult, error) {
	return g.GeocodeReturnFields(address, "stateleg")
}

// TODO: School District (school)

// GeocodeReturnFields will geocode and includes additional fields in response
/*
 	See: http://geocod.io/docs/#toc_22
	Note:
		Each field counts as an additional lookup each
*/
func (g *Geocodio) GeocodeReturnFields(address string, fields ...string) (GeocodeResult, error) {
	if address == "" {
		return GeocodeResult{}, errors.New("address can not be empty")
	}

	fieldsCommaSeparated := strings.Join(fields, ",")

	results, err := g.Call("/geocode",
		map[string]string{
			"q":      address,
			"fields": fieldsCommaSeparated,
		})
	if err != nil {
		return GeocodeResult{}, err
	}

	return results, nil
}
