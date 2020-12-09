package backend

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRateHandler(t *testing.T) {
	var rateNGNTtoEURT = "/rate?sourceCode=NGNT&sourceIssuer=GAWODAROMJ33V5YDFY3NPYTHVYQG7MJXVJ2ND3AOGIHYRWINES6ACCPD&destCode=EURT&destIssuer=GAP5LETOV6YIE62YAM56STDANPRDO7ZFDBGSNHJQIYGGKSMOZAHOOS2S"
	var rateEURTtoNGNT = "/rate?sourceCode=EURT&sourceIssuer=GAP5LETOV6YIE62YAM56STDANPRDO7ZFDBGSNHJQIYGGKSMOZAHOOS2S&destCode=NGNT&destIssuer=GAWODAROMJ33V5YDFY3NPYTHVYQG7MJXVJ2ND3AOGIHYRWINES6ACCPD"
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
			r:              httptest.NewRequest("GET", rateNGNTtoEURT+"&start=1603315973&end=1603320883", nil),
			w:              httptest.NewRecorder(),
			expectedStatus: http.StatusOK,
			golden:         "NGNT_EURT_limited.golden",
			handler:        RateHandler(),
		},
		{
			name:           "reversed query (values should be reciprocal of above test)",
			r:              httptest.NewRequest("GET", rateEURTtoNGNT+"&start=1603315973&end=1603320883", nil),
			w:              httptest.NewRecorder(),
			expectedStatus: http.StatusOK,
			golden:         "EURT_NGNT_limited.golden",
			handler:        RateHandler(),
		},
	}

	for _, test := range tests {
		runTest(t, test, "../testdata/rate")
	}
}
