package geocodio_test

import (
	"testing"

	"github.com/johncosta/geocodio"
)

func TestReverseGeocodeFullAddress(t *testing.T) {
	Geocodio, err := geocodio.NewGeocodio(APIKey())

	if err != nil {
		t.Error("Failed with API KEY set.", APIKey(), err)
	}

	result, err := Geocodio.ReverseGeocode(AddressTestOneLatitude, AddressTestOneLongitude)
	if err != nil {
		t.Error(err)
	}

	// t.Log(result.ResponseAsString())

	if len(result.Results) == 0 {
		t.Error("Results length is 0")
	}

	if len(result.Results) < 3 {
		t.Error("Results found length is less than 3", len(result.Results))
	}

	if len(result.Results) == 0 {
		t.Error("No results were found.")
		return
	}

	if result.Results[0].Formatted != AddressTestOneFull {
		t.Error("Location latitude does not match", result.Results[0].Formatted, AddressTestOneFull)
	}

}

func TestReverseGeocodeWithoutLatLng(t *testing.T) {
	Geocodio, err := geocodio.NewGeocodio(APIKey())
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}
	_, err = Geocodio.ReverseGeocode(0.0, 0.0)
	if err == nil {
		t.Error("Error should not be nil.")
	}
}
