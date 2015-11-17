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

// Response from api.cinemate.cc
type APIResponse struct {
	Movies  []Movie  `xml:"movie,omitempty"`
	Persons []Person `xml:"person,omitempty"`
}

// Movie is responce movie api from server. Now parse only xml
// id                  ID фильма
// type                Тип фильма
// title_russian       русскоязычное название фильма
// title_original      название фильма в оригинале
// title_english       англоязычное название фильма
// year                год выхода фильма
// poster              включает в себя 3 тега со ссылками на постеры разных размеров
// description         описание фильма
// runtime             длительность фильма в минутах
// release_date_world  дата выхода фильма в мире в ISO формате
// release_date_russia дата выхода фильма в России в ISO формате
// imdb	рейтинг фильма по версии imdb.com со следующими атрибутами:
// rating              рейтинг фильма по 10-балльной шкале
// votes               число голосов за фильм
// kinopoisk           рейтинг фильма по версии kinopoisk.ru со следующими атрибутами:
// rating              рейтинг фильма по 10-балльной шкале
// votes               число голосов за фильм
// country             список стран-создателей фильма, представленный списком тегов name с русским названием стран
// genre               список жанров фильма, представленный списком тегов name с русским названием жанров
// director            список режиссеров фильма, представленный списком тегов name с русским именами режиссеров и ID персоны
// cast                список актеров фильма, представленный списком тегов name с русским именами актеров и ID персоны
// url                 ссылка на страницу фильма
type Movie struct {
	ID                int64    `xml:"id,omitempty"`
	Type              string   `xml:"type,omitempty"`
	TitleRussian      string   `xml:"title_russian"`
	TitleOriginal     string   `xml:"title_original"`
	TitleEnglish      string   `xml:"title_english,omitempty"`
	Year              int64    `xml:"year,omitempty"`
	Runtime           int64    `xml:"runtime,omitempty"`
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
	Votes  int64   `xml:"votes,attr"`
}

type country struct {
	Name []string `xml:"name,omitempty"`
}

type genre struct {
	Name []string `xml:"name,omitempty"`
}

// Person is responce person api from server. Now parse only xml
// id            ID персоны
// name          русскоязычное имя персоны
// name_original имя персоны в оригинале
// photo         включает в себя 3 тега со ссылками на фотографии разных размеров
// url           ссылка на страницу персоны
type Person struct {
	ID           int64        `xml:"id,omitempty"`
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

// Account with passkey for access to account
type Account struct {
	passkey string
}

// AccountResponse is response account api from server
// username                логин пользователя
// reputation              репутация пользователя
// review_count            количество отзывов
// gold_badges             число золотых наград
// silver_badges           число серебряных наград
// bronze_badges           число бронзовых наград
// unread_pm_count         число непрочитанных личных сообщений
// unread_forum_count      число новых сообщений и/или тем на форуме в отслеживаемых темах и разделах
// unread_updatelist_count число новых записей в ленте обновлений
// subscription_count      общее число подписок в ленте обновлений
type AccountProfile struct {
	Username              string `xml:"username"`
	Reputation            int64  `xml:"reputation"`
	ReviewCount           int64  `xml:"review_count"`
	GoldBadges            int64  `xml:"gold_badges"`
	SilverBadges          int64  `xml:"silver_badges"`
	BronzeBadges          int64  `xml:"bronze_badges"`
	UnreadPmCount         int64  `xml:"unread_pm_count"`
	UnreadForumCount      int64  `xml:"unread_forum_count"`
	UnreadUpdatelistCount int64  `xml:"unread_updatelist_count"`
	SubscriptionCount     int64  `xml:"subscription_count"`
}

// UpdateList Записи ленты обновлений пользователя
// count число всех записей в ленте обновлений (новый)
// item  запись ленты обновлений
type UpdateList struct {
	Count int64            `xml:"count"`
	Items []updateListItem `xml:"item"`
}

// ListItem запись ленты обновлений
// date	дата и время добавления записи в ленту обновлений пользователя в ISO формате
// description	текстовое описание обновления
// url	ссылка на обновление; переход по ссылке отмечает запись в ленте прочитанной и производит редирект на страницу с обновлением
// new	флаг прочитанного обновления (1 - непрочтенное, 0 - прочтенное)
// for_object	список объектов object, список объектов movie, person или comment, к которым привязано обновление
type updateListItem struct {
	Date        string           `xml:"date"`
	Description string           `xml:"description"`
	URL         string           `xml:"url"`
	New         int64            `xml:"new"`
	ForObject   updateListObject `xml:"for_object"`
}

type updateListObject struct {
	Movie   updateListItemObject `xml:"movie,omitempty"`
	Person  updateListItemObject `xml:"person,omitempty"`
	Comment updateListItemObject `xml:"comment,omitempty"`
}

// title	строковое представление объекта
// url	ссылка на объект обновления
type updateListItemObject struct {
	ID    int64  `xml:"id"`
	Title string `xml:"title"`
}

// WatchList список объектов слежения пользователя
// Каждый узел представляет собой объект слежения одного из типов: movie, person или comment
type WatchList struct {
	Comments []watchListObject `xml:"comment"`
	Persons  []watchListObject `xml:"person"`
	Movies   []watchListObject `xml:"movie"`
}

// date	дата и время добавления объекта в список слежения в ISO формате
// name	строковое представление объекта слежения
// description	описание подписки на объект
// url	ссылка на объект слежения
type watchListObject struct {
	Date        string `xml:"date"`
	Name        string `xml:"name"`
	Description string `xml:"description"`
	URL         string `xml:"url"`
}

// Stats статистика сайта за последние сутки
// users_count    число новых пользователей
// reviews_count  число новых отзывов
// comments_count число новых комментариев к отзывам
// movies_count   число новых фильмов
type Stats struct {
	UsersCount    int64 `xml:"users_count"`
	ReviewsCount  int64 `xml:"reviews_count"`
	CommentsCount int64 `xml:"comments_count"`
	MoviesCount   int64 `xml:"movies_count"`
}

// Init CinemaCC to set API value
func Init(apiKey string) *API {
	return &API{apikey: apiKey}
}

// InitAccount CinemaCC to set API value
func InitAccount(passKey string) *Account {
	return &Account{passkey: passKey}
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
