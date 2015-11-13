package cinemate

import (
	"encoding/xml"
	"net/url"
	"time"
)

// GetStatsNew возвращает статистику сайта за последние сутки
// Пример запроса: http://api.cinemate.cc/stats.new?format=xml
// format	необязательный параметр формата возвращаемых сервером данных: xml (по умолчанию) или json
func GetStatsNew() (stats Stats, err error) {
	time.Sleep(1 * time.Second)
	var u url.URL
	u.Scheme = "http"
	u.Host = "api.cinemate.cc"
	u.Path = "/stats.new"
	q := u.Query()
	q.Set("format", "xml")
	u.RawQuery = q.Encode()
	xmlBody, err := getXML(u.String())
	if err != nil {
		return
	}
	err = xml.Unmarshal(xmlBody, &stats)
	return
}
