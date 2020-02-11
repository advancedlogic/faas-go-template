package function

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func HandleTest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(Handle))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	expected, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	assert.Nil(t, err)
	assert.Equal(t, string(expected), "Hello world, input was: ")
}
