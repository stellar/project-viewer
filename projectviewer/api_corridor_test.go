package projectviewer

import (
	"flag"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

var update = flag.Bool("update", false, "update the golden files of this test")
var NGNTtoEURTparams = "/?sourceCode=NGNT&sourceIssuer=GAWODAROMJ33V5YDFY3NPYTHVYQG7MJXVJ2ND3AOGIHYRWINES6ACCPD&destCode=EURT&destIssuer=GAP5LETOV6YIE62YAM56STDANPRDO7ZFDBGSNHJQIYGGKSMOZAHOOS2S"

type queryTest struct {
	name           string
	r              *http.Request
	w              *httptest.ResponseRecorder
	expectedStatus int
	golden         string
}

func TestCorridorHandler(t *testing.T) {
	tests := []queryTest{
		{
			name:           "full history query",
			r:              httptest.NewRequest("GET", NGNTtoEURTparams, nil),
			w:              httptest.NewRecorder(),
			expectedStatus: http.StatusOK,
			golden:         "NGNT_EURT_all.golden",
		},
		{
			name:           "limited range query",
			r:              httptest.NewRequest("GET", NGNTtoEURTparams+"&start=27839022&end=27839022", nil),
			w:              httptest.NewRecorder(),
			expectedStatus: http.StatusOK,
			golden:         "NGNT_EURT_limited.golden",
		},
	}

	for _, test := range tests {
		runTest(t, test, "../testdata/")
	}
}

func runTest(t *testing.T, test queryTest, goldenFolder string) {
	t.Run(test.name, func(t *testing.T) {
		handler := CorridorHandler()
		handler.ServeHTTP(test.w, test.r)

		assert.Equal(t, test.expectedStatus, test.w.Code)
		actualString := test.w.Body.String()
		wantString, err := getGolden(t, path.Join(goldenFolder, test.golden), actualString, *update)
		assert.NoError(t, err)
		assert.Equal(t, wantString, actualString)
	})
}

func getGolden(t *testing.T, goldenFile string, actual string, update bool) (string, error) {
	t.Helper()
	f, err := os.OpenFile(goldenFile, os.O_RDWR, 0644)
	defer f.Close()
	if err != nil {
		return "", err
	}

	// If the update flag is true, clear the current contents of the golden file and write the actual output
	// This is useful for when new tests or added or functionality changes that breaks current tests
	if update {
		err := os.Truncate(goldenFile, 0)
		if err != nil {
			return "", err
		}

		_, err = f.WriteString(actual)
		if err != nil {
			return "", err
		}

		return actual, nil
	}

	wantOutput, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}

	return string(wantOutput), nil
}
