package backend

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var volumeToNGNT = "/volume?code=NGNT&issuer=GAWODAROMJ33V5YDFY3NPYTHVYQG7MJXVJ2ND3AOGIHYRWINES6ACCPD"
var volumeFromCENTUS = "/volume?code=CENTUS&issuer=GAKMVPHBET4T7DPN32ODVSI4AA3YEZX2GHGNNSBGFNRQ6QEVKFO4MNDZ&volumeFrom=true"

func TestVolumeHandler(t *testing.T) {
	tests := []queryTest{
		{
			name:           "volumeTo full history",
			r:              httptest.NewRequest("GET", volumeToNGNT, nil),
			w:              httptest.NewRecorder(),
			expectedStatus: http.StatusOK,
			golden:         "NGNT_all.golden",
			handler:        VolumeHandler(),
		},
		{
			name:           "volumeTo limited range",
			r:              httptest.NewRequest("GET", volumeToNGNT+"&start=20304878&end=20492609", nil),
			w:              httptest.NewRecorder(),
			expectedStatus: http.StatusOK,
			golden:         "NGNT_limited.golden",
			handler:        VolumeHandler(),
		},
		{
			name:           "volumeFrom full history",
			r:              httptest.NewRequest("GET", volumeFromCENTUS, nil),
			w:              httptest.NewRecorder(),
			expectedStatus: http.StatusOK,
			golden:         "CENTUS_all.golden",
			handler:        VolumeHandler(),
		},
		{
			name:           "volumeFrom limited range",
			r:              httptest.NewRequest("GET", volumeFromCENTUS+"&start=24334355&end=24799948", nil),
			w:              httptest.NewRecorder(),
			expectedStatus: http.StatusOK,
			golden:         "CENTUS_limited.golden",
			handler:        VolumeHandler(),
		},
	}

	for _, test := range tests {
		runTest(t, test, "../testdata/volume")
	}
}
