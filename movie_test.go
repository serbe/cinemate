package cinemate

import (
	. "github.com/smartystreets/goconvey/convey"
	"net/url"
	"strconv"
	"testing"
)

func TestGetMovie(t *testing.T) {
	Convey("Given url value", t, func() {
		// var result APIResponse
		var u url.URL
		u.Scheme = "http"
		u.Host = "api.cinemate.cc"
		u.Path = "/movie"
		q := u.Query()
		q.Set("apikey", "APIKEY")
		q.Set("id", strconv.FormatInt(31, 10))
		q.Set("format", "xml")
		u.RawQuery = q.Encode()
		testUrl := "http://api.cinemate.cc/movie?apikey=APIKEY&format=xml&id=31"
		Convey("The values must be equal", func() {
			So(u.String(), ShouldEqual, testUrl)
		})
	})
	// xmlBody, err := getXML(u.String())
	// if err != nil {
	// 	return
	// }
	// err = xml.Unmarshal(xmlBody, &result)
	// movie = result.Movies[0]
	// if movie.ID == 0 {
	// 	err = fmt.Errorf("Movie not found")
	// }
	// return
}
