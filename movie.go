package cinemate

import (
	"encoding/xml"
	"net/url"
	"strconv"
	"time"
)

// GetMovie Информация о фильме
// example: http://api.cinemate.cc/movie?apikey=APIKEY&id=68675&format=xml
// apikey	ключ разработчика
// id	ID фильма
// format	необязательный параметр формата возвращаемых сервером данных: xml (по умолчанию) или json
func (ccc *API) GetMovie(id int) ([]Movie, error) {
	time.Sleep(1 * time.Second)
	var result Response
	var u url.URL
	u.Scheme = "http"
	u.Host = "api.cinemate.cc"
	u.Path = "/movie"
	q := u.Query()
	q.Set("apikey", ccc.apikey)
	q.Set("id", strconv.Itoa(id))
	q.Set("format", "xml")
	u.RawQuery = q.Encode()
	xmlBody, err := getXML(u.String())
	if err != nil {
		return []Movie{}, err
	}
	err = xml.Unmarshal(xmlBody, &result)

	return result.Movie, err
}

// GetMovieList Результаты поиска фильмов, используя заданные фильтры. Возвращается 10 первых фильмов.
// example: http://api.cinemate.cc/movie.list?apikey=APIKEY&year=2010&format=xml
// apikey	ключ разработчика
// type	тип фильмов. Возможные значения: movie, serial, short
// state	состояние фильма. Возможные значения: soon, cinema
// mode	специальный режим отображения лучших фильмов, отсортированных по рейтингу IMDB. Возможные значения: best
// year	год выпуска фильма или сериала
// genre	жанр. В качестве значения (используется slug со страницы http://cinemate.cc/movie/genre/, например, "sport")
// country	отбор по стране (используется slug со страницы http://cinemate.cc/movie/country/, например, "kazakhstan")
// order_by	критерий сортировки: create_date (по дате добавления на сайт), release_date (по дате выхода в мире), ru_release_date (по дате выхода в России, по умолчанию)
// order	порядок сортировки параметра order_by: desc (по убыванию, по умолчанию), asc (по возрастанию)
// from, to	значения среза параметра order_by в формате даты ДД.ММ.ГГГГ. Включительно.
// page, per_page	страница и количество записей в выборке. По умолчанию 0 и 10 соответственно. per_page не может быть более 25.
// format	необязательный параметр формата возвращаемых сервером данных: xml (по умолчанию) или json
func (ccc *API) GetMovieList(ccr CCRequest) ([]Movie, error) {
	time.Sleep(1 * time.Second)
	var result Response
	var u url.URL
	u.Scheme = "http"
	u.Host = "api.cinemate.cc"
	u.Path = "/movie.list"
	q := u.Query()
	q.Set("apikey", ccc.apikey)
	if ccr.Type != "" {
		q.Set("type", ccr.Type)
	}
	if ccr.State != "" {
		q.Set("state", ccr.State)
	}
	if ccr.Mode != "" {
		q.Set("mode", ccr.Mode)
	}
	if ccr.Year != 0 {
		q.Set("year", strconv.FormatInt(ccr.Year, 10))
	}
	if ccr.Genre != "" {
		q.Set("genre", ccr.Genre)
	}
	if ccr.Country != "" {
		q.Set("country", ccr.Country)
	}
	if ccr.OrderBy != "" {
		q.Set("order_by", ccr.OrderBy)
	}
	if ccr.Order != "" {
		q.Set("order", ccr.Order)
	}
	if ccr.From != "" {
		q.Set("from", ccr.From)
	}
	if ccr.To != "" {
		q.Set("to", ccr.To)
	}
	if ccr.Page != 0 {
		q.Set("page", strconv.FormatInt(ccr.Page, 10))
	}
	if ccr.PerPage != 0 {
		q.Set("per_page", strconv.FormatInt(ccr.PerPage, 10))
	}
	q.Set("format", "xml")
	u.RawQuery = q.Encode()
	xmlBody, err := getXML(u.String())
	if err != nil {
		return []Movie{}, err
	}
	err = xml.Unmarshal(xmlBody, &result)
	if err != nil {
		return []Movie{}, err
	}
	return result.Movie, err
}

// GetMovieSearch Поиск по заголовкам фильмов
// example: http://api.cinemate.cc/movie.search?apikey=APIKEY&term=Пираты%20кариб&format=xml
// apikey	ключ разработчика
// term	искомая строка; поддерживается уточняющий поиск по году выхода фильма (год должен быть указан в конце искомой строки, например, "Пираты кариб 2003") и коррекцию ошибок при печати
// format	необязательный параметр формата возвращаемых сервером данных: xml (по умолчанию) или json
func (ccc *API) GetMovieSearch(term string) ([]Movie, error) {
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
		return []Movie{}, err
	}
	err = xml.Unmarshal(xmlBody, &result)
	return result.Movie, err
}
