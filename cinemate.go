package cinemate

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	apiURL = "http://http://api.cinemate.cc"
)

// API with apikey for use api.cinemate.cc
type API struct {
	apikey string
}

// Response from  api.cinemate.cc
type Response struct {
	Movie  []Movie  `xml:"movie"`
	Person []Person `xml:"person"`
}

// Movie is responce movie api from server. Now parse only xml
// id	ID фильма
// type	Тип фильма
// title_russian	русскоязычное название фильма
// title_original	название фильма в оригинале
// title_english	англоязычное название фильма
// year	год выхода фильма
// poster	включает в себя 3 тега со ссылками на постеры разных размеров
// description	описание фильма
// runtime	длительность фильма в минутах
// release_date_world	дата выхода фильма в мире в ISO формате
// release_date_russia	дата выхода фильма в России в ISO формате
// imdb	рейтинг фильма по версии imdb.com со следующими атрибутами:
// rating	рейтинг фильма по 10-балльной шкале
// votes	число голосов за фильм
// kinopoisk	рейтинг фильма по версии kinopoisk.ru со следующими атрибутами:
// rating	рейтинг фильма по 10-балльной шкале
// votes	число голосов за фильм
// country	список стран-создателей фильма, представленный списком тегов name с русским названием стран
// genre	список жанров фильма, представленный списком тегов name с русским названием жанров
// director	список режиссеров фильма, представленный списком тегов name с русским именами режиссеров и ID персоны
// cast	список актеров фильма, представленный списком тегов name с русским именами актеров и ID персоны
// url	ссылка на страницу фильма
type Movie struct {
	ID                int      `xml:"id"`
	Type              string   `xml:"type"`
	TitleRussian      string   `xml:"title_russian"`
	TitleOriginal     string   `xml:"title_original"`
	TitleEnglish      string   `xml:"title_english"`
	Year              int      `xml:"year"`
	Runtime           int      `xml:"runtime"`
	Poster            image    `xml:"poster"`
	URL               string   `xml:"url"`
	Imdb              rating   `xml:"imdb"`
	Kinopoisk         rating   `xml:"kinopoisk"`
	Country           country  `xml:"country"`
	Genre             genre    `xml:"genre"`
	Description       string   `xml:"description"`
	Trailer           string   `xml:"trailer"`
	ReleaseDateWorld  string   `xml:"release_date_world"`
	ReleaseDateRussia string   `xml:"release_date_russia"`
	Director          Person   `xml:"director>person"`
	Cast              []Person `xml:"cast>person"`
}

type image struct {
	Small  urlStruct `xml:"small"`
	Big    urlStruct `xml:"big"`
	Medium urlStruct `xml:"medium"`
}

type urlStruct struct {
	URL string `xml:"url,attr"`
}

type rating struct {
	Rating float64 `xml:"rating,attr"`
	Votes  int     `xml:"votes,attr"`
}

type country struct {
	Name []string `xml:"name"`
}

type genre struct {
	Name []string `xml:"name"`
}

// Person is responce person api from server. Now parse only xml
// id	ID персоны
// name	русскоязычное имя персоны
// name_original	имя персоны в оригинале
// photo	включает в себя 3 тега со ссылками на фотографии разных размеров
// url	ссылка на страницу персоны
type Person struct {
	ID           int    `xml:"id"`
	Name         string `xml:"name"`
	NameOriginal string `xml:"name_original"`
	Photo        image  `xml:"photo"`
	URL          string `xml:"url"`
	/// 
}

// CCRequest struct for make search request
type CCRequest struct {
	ID      int64
	Type    string
	State   string
	Mode    string
	Year    int64
	Genre   string
	Country string
	OrderBy string
	Order   string
	From    string
	To      string
	Page    int64
	PerPage int64
	Format  string
}

// Init CinemaCC to set API value
func Init(apiKey string) *API {
	return &API{apikey: apiKey}
}

func getXML(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	if resp.StatusCode != 200 {
		log.Println(resp.Header)
		return []byte{}, fmt.Errorf("Status Code %d received from cinemate.cc", resp.StatusCode)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, err
}
