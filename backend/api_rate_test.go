package backend

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRateHandler(t *testing.T) {
	var rateNGNTtoEURT = "/rate?sourceCode=NGNT&sourceIssuer=GAWODAROMJ33V5YDFY3NPYTHVYQG7MJXVJ2ND3AOGIHYRWINES6ACCPD&destCode=EURT&destIssuer=GAP5LETOV6YIE62YAM56STDANPRDO7ZFDBGSNHJQIYGGKSMOZAHOOS2S"
	tests := []queryTest{
		{
			name:           "full history query",
			r:              httptest.NewRequest("GET", rateNGNTtoEURT, nil),
			w:              httptest.NewRecorder(),
			expectedStatus: http.StatusOK,
			golden:         "NGNT_EURT_all.golden",
			handler:        RateHandler(),
		},
		{
			name:           "limited range query",
			r:              httptest.NewRequest("GET", rateNGNTtoEURT+"&start=32214562&end=32215487", nil),
			w:              httptest.NewRecorder(),
			expectedStatus: http.StatusOK,
			golden:         "NGNT_EURT_limited.golden",
			handler:        RateHandler(),
		},
	}

	for _, test := range tests {
		runTest(t, test, "../testdata/rate")
	}
}
