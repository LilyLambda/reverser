// +build !pure

package reverser

import (
	"io/ioutil"
	"net/http"
	"testing"
	"strings"
	"github.com/stretchr/testify/assert"
)

func TestGoogleRobots(t *testing.T) {
	resp, _ := http.Get("http://www.google.com/robots.txt")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	firstLine := strings.Split(string(body), "\n")[0]
	assert.Equal(t, firstLine, Reverse(Reverse(firstLine)), "google robots.txt first line is reversible")
}
