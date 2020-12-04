package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/stellar/project-viewer/backend"
)

// ServeMux creates Mux to serve api
func ServeMux() http.Handler {
	mux := chi.NewMux()
	mux.Handle("/corridor", backend.CorridorHandler())
	mux.Handle("/volume", backend.VolumeHandler())
	return mux
}

func main() {
	fmt.Println("Starting server on port :8080")
	mux := ServeMux()
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("ListenAndServe:", err)
	}
}

/*
Some example URLs:

Corridors:
http://localhost:8080/corridor?sourceCode=NGNT&sourceIssuer=GAWODAROMJ33V5YDFY3NPYTHVYQG7MJXVJ2ND3AOGIHYRWINES6ACCPD&destCode=EURT&destIssuer=GAP5LETOV6YIE62YAM56STDANPRDO7ZFDBGSNHJQIYGGKSMOZAHOOS2S
http://localhost:8080/corridor?sourceCode=CENTUS&sourceIssuer=GAKMVPHBET4T7DPN32ODVSI4AA3YEZX2GHGNNSBGFNRQ6QEVKFO4MNDZ&destCode=USD&destIssuer=GB2O5PBQJDAFCNM2U2DIMVAEI7ISOYL4UJDTLN42JYYXAENKBWY6OBKZ

Volumes:
http://localhost:8080/volume?code=NGNT&issuer=GAWODAROMJ33V5YDFY3NPYTHVYQG7MJXVJ2ND3AOGIHYRWINES6ACCPD
http://localhost:8080/volume?code=CENTUS&issuer=GAKMVPHBET4T7DPN32ODVSI4AA3YEZX2GHGNNSBGFNRQ6QEVKFO4MNDZ&volumeFrom=true

*/
