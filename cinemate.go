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
	Movie  []Movie  `xml:"movie,omitempty"`
	Person []Person `xml:"person,omitempty"`
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
	ID                int      `xml:"id,omitempty"`
	Type              string   `xml:"type,omitempty"`
	TitleRussian      string   `xml:"title_russian"`
	TitleOriginal     string   `xml:"title_original"`
	TitleEnglish      string   `xml:"title_english,omitempty"`
	Year              int      `xml:"year,omitempty"`
	Runtime           int      `xml:"runtime,omitempty"`
	Poster            image    `xml:"poster,omitempty"`
	URL               string   `xml:"url,omitempty"`
	Imdb              rating   `xml:"imdb,omitempty"`
	Kinopoisk         rating   `xml:"kinopoisk,omitempty"`
	Country           country  `xml:"country"`
	Genre             genre    `xml:"genre"`
	Description       string   `xml:"description,omitempty"`
	Trailer           string   `xml:"trailer,omitempty"`
	ReleaseDateWorld  string   `xml:"release_date_world,omitempty"`
	ReleaseDateRussia string   `xml:"release_date_russia,omitempty"`
	Director          Person   `xml:"director>person,omitempty"`
	Cast              []Person `xml:"cast>person,omitempty"`
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
	Name []string `xml:"name,omitempty"`
}

type genre struct {
	Name []string `xml:"name,omitempty"`
}

// Person is responce person api from server. Now parse only xml
// id	ID персоны
// name	русскоязычное имя персоны
// name_original	имя персоны в оригинале
// photo	включает в себя 3 тега со ссылками на фотографии разных размеров
// url	ссылка на страницу персоны
type Person struct {
	ID           int          `xml:"id,omitempty"`
	Name         string       `xml:"name,omitempty"`
	NameOriginal string       `xml:"name_original,omitempty"`
	Photo        image        `xml:"photo,omitempty"`
	URL          string       `xml:"url,omitempty"`
	Movies       personMovies `xml:"movies,omitempty"`
}

type personMovies struct {
	Director []Movie `xml:"director>movie,omitempty"`
	Actor    []Movie `xml:"actor>movie,omitempty"`
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
