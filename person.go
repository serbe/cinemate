package cinemate

import (
	"encoding/xml"
	"net/url"
	"strconv"
	"time"
)

// GetPerson Основная информация о персоне
// Пример запроса: http://api.cinemate.cc/person?id=3971&apikey=APIKEY&format=xml
// apikey ключ разработчика
// id     ID персоны
// format необязательный параметр формата возвращаемых сервером данных: xml (по умолчанию) или json
func (api *API) GetPerson(id int64) (person Person, err error) {
	time.Sleep(1 * time.Second)
	var result APIResponse
	var u url.URL
	u.Scheme = "http"
	u.Host = "api.cinemate.cc"
	u.Path = "/person"
	q := u.Query()
	q.Set("apikey", api.apikey)
	q.Set("id", strconv.FormatInt(id, 10))
	q.Set("format", "xml")
	u.RawQuery = q.Encode()
	xmlBody, err := getXML(u.String())
	if err != nil {
		return
	}
	err = xml.Unmarshal(xmlBody, &result)
	if len(result.Persons[0]) > 0 {
		person = result.Persons[0]
	}
	return
}

// GetPersonMovies Информация о персоне, включая фильмы, в съемке которых персона
// принимала участие в качестве актера или режиссера. Основная информация о
// персоне идентична команде person.
// Пример запроса: http://api.cinemate.cc/person.movies?id=3971&apikey=APIKEY&format=xml
// apikey ключ разработчика
// id     ID персоны
// format необязательный параметр формата возвращаемых сервером данных: xml (по умолчанию) или json
func (api *API) GetPersonMovies(id int64) (persons []Person, err error) {
	time.Sleep(1 * time.Second)
	var result APIResponse
	var u url.URL
	u.Scheme = "http"
	u.Host = "api.cinemate.cc"
	u.Path = "/person.movies"
	q := u.Query()
	q.Set("apikey", api.apikey)
	q.Set("id", strconv.FormatInt(id, 10))
	q.Set("format", "xml")
	u.RawQuery = q.Encode()
	xmlBody, err := getXML(u.String())
	if err != nil {
		return
	}
	err = xml.Unmarshal(xmlBody, &result)
	persons = result.Persons
	return
}

// GetPersonSearch Метод возвращает первые 10 результатов поиска по базе персон
// Пример запроса: http://api.cinemate.cc/person.search?apikey=APIKEY&term=гиленхол&format=xml
// apikey ключ разработчика
// term   искомая строка; поддерживается уточняющий поиск по году выхода фильма (год должен быть указан в конце искомой строки, например, "Пираты кариб 2003") и коррекцию ошибок при печати
// format необязательный параметр формата возвращаемых сервером данных: xml (по умолчанию) или json
func (api *API) GetPersonSearch(term string) (persons []Person, err error) {
	time.Sleep(1 * time.Second)
	var result APIResponse
	var u url.URL
	u.Scheme = "http"
	u.Host = "api.cinemate.cc"
	u.Path = "/person.search"
	q := u.Query()
	q.Set("apikey", api.apikey)
	q.Set("term", term)
	q.Set("format", "xml")
	u.RawQuery = q.Encode()
	xmlBody, err := getXML(u.String())
	if err != nil {
		return
	}
	err = xml.Unmarshal(xmlBody, &result)
	persons = result.Persons
	return
}
