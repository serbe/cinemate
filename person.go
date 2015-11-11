package cinemate

import (
	"encoding/xml"
	"net/url"
	"strconv"
	"time"
)

// GetPerson Основная информация о персоне
// example: http://api.cinemate.cc/person?id=3971&apikey=APIKEY&format=xml
// apikey	ключ разработчика
// id	ID персоны
// format	необязательный параметр формата возвращаемых сервером данных: xml (по умолчанию) или json
func (ccc *API) GetPerson(id int) ([]Person, error) {
	time.Sleep(1 * time.Second)
	var result Response
	var u url.URL
	u.Scheme = "http"
	u.Host = "api.cinemate.cc"
	u.Path = "/person"
	q := u.Query()
	q.Set("apikey", ccc.apikey)
	q.Set("id", strconv.Itoa(id))
	q.Set("format", "xml")
	u.RawQuery = q.Encode()
	xmlBody, err := getXML(u.String())
	if err != nil {
		return []Person{}, err
	}
	err = xml.Unmarshal(xmlBody, &result)

	return result.Person, err
}

// GetPersonMovies Информация о персоне, включая фильмы, в съемке которых персона 
// принимала участие в качестве актера или режиссера. Основная информация о 
// персоне идентична команде person.
// example: http://api.cinemate.cc/person.movies?id=3971&apikey=APIKEY&format=xml
// apikey	ключ разработчика
// id	ID персоны
// format	необязательный параметр формата возвращаемых сервером данных: xml (по умолчанию) или json
func (ccc *API) GetPersonMovies(id int) ([]Person, error) {
	time.Sleep(1 * time.Second)
	var result Response
	var u url.URL
	u.Scheme = "http"
	u.Host = "api.cinemate.cc"
	u.Path = "/person.movies"
	q := u.Query()
	q.Set("apikey", ccc.apikey)
	q.Set("id", strconv.Itoa(id))
	q.Set("format", "xml")
	u.RawQuery = q.Encode()
	xmlBody, err := getXML(u.String())
	if err != nil {
		return []Person{}, err
	}
	err = xml.Unmarshal(xmlBody, &result)

	return result.Person, err
}

// GetPersonSearch Метод возвращает первые 10 результатов поиска по базе персон
// example: http://api.cinemate.cc/person.search?apikey=APIKEY&term=гиленхол&format=xml
// apikey	ключ разработчика
// term	искомая строка; поддерживается уточняющий поиск по году выхода фильма (год должен быть указан в конце искомой строки, например, "Пираты кариб 2003") и коррекцию ошибок при печати
// format	необязательный параметр формата возвращаемых сервером данных: xml (по умолчанию) или json
func (ccc *API) GetPersonSearch(term string) ([]Person, error) {
	time.Sleep(1 * time.Second)
	var result Response
	var u url.URL
	u.Scheme = "http"
	u.Host = "api.cinemate.cc"
	u.Path = "/movie.search"
	q := u.Query()
	q.Set("apikey", ccc.apikey)
	q.Set("term", term)
	q.Set("format", "xml")
	u.RawQuery = q.Encode()
	xmlBody, err := getXML(u.String())
	if err != nil {
		return []Person{}, err
	}
	err = xml.Unmarshal(xmlBody, &result)
	return result.Person, err
}
