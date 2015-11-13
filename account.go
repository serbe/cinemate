package cinemate

import (
	"encoding/xml"
	"net/url"
	"strconv"
	"time"
)

// GetAccountAuth Авторизация по логину и паролю.
// Пример запроса: http://api.cinemate.cc/account.auth?username=USERNAME&password=PASSWORD
// username логин пользователя
// password пароль пользователя
func GetAccountAuth(username string, password string) (passkey string, err error) {
	time.Sleep(1 * time.Second)
	var u url.URL
	var result Account
	u.Scheme = "http"
	u.Host = "api.cinemate.cc"
	u.Path = "/account.auth"
	q := u.Query()
	q.Set("username", username)
	q.Set("password", password)
	u.RawQuery = q.Encode()
	xmlBody, err := getXML(u.String())
	if err != nil {
		return
	}
	err = xml.Unmarshal(xmlBody, &result)
	return
}

// GetAccountProfile Данные и статистика пользовательского аккаунта
// Пример запроса: http://api.cinemate.cc/account.profile?passkey=PASSKEY&format=xml
// PASSKEY уникальное для каждого пользователя 40-значное 16-ричное число, получить которое можно на странице настроек
// format  необязательный параметр формата возвращаемых сервером данных: xml (по умолчанию) или json
func (acc *Account) GetAccountProfile() (profile AccountProfile, err error) {
	time.Sleep(1 * time.Second)
	var u url.URL
	u.Scheme = "http"
	u.Host = "api.cinemate.cc"
	u.Path = "/account.profile"
	q := u.Query()
	q.Set("passkey", acc.passkey)
	q.Set("format", "xml")
	u.RawQuery = q.Encode()
	xmlBody, err := getXML(u.String())
	if err != nil {
		return
	}
	err = xml.Unmarshal(xmlBody, &profile)
	return
}

// GetAccountUpdateList Метод возвращает записи ленты обновлений пользователя
// Пример запроса: http://api.cinemate.cc/account.updatelist?passkey=PASSKEY&newonly=1&format=xml
// PASSKEY уникальное для каждого пользователя 40-значное 16-ричное число, получить которое можно на странице настроек
// newonly если 1, то возвращается список только непрочитанных записей в ленте
// format  необязательный параметр формата возвращаемых сервером данных: xml (по умолчанию) или json
func (acc *Account) GetAccountUpdateList(newonly ...bool) (list UpdateList, err error) {
	time.Sleep(1 * time.Second)
	var u url.URL
	newOnlyInt := 1
	u.Scheme = "http"
	u.Host = "api.cinemate.cc"
	u.Path = "/account.updatelist"
	q := u.Query()
	q.Set("passkey", acc.passkey)
	if len(newonly) > 0 {
		if newonly[0] == false {
			newOnlyInt = 0
		}
	}
	q.Set("newonly", strconv.FormatInt(int64(newOnlyInt), 10))
	q.Set("format", "xml")
	u.RawQuery = q.Encode()
	xmlBody, err := getXML(u.String())
	if err != nil {
		return
	}
	err = xml.Unmarshal(xmlBody, &list)
	return
}

// GetAccountWatchlist Метод возвращает список объектов слежения пользователя
// Пример запроса: http://api.cinemate.cc/account.watchlist?passkey=PASSKEY&format=xml
// PASSKEY уникальное для каждого пользователя 40-значное 16-ричное число, получить которое можно на странице настроек
// format  необязательный параметр формата возвращаемых сервером данных: xml (по умолчанию) или json
func (acc *Account) GetAccountWatchlist() (list WatchList, err error) {
	time.Sleep(1 * time.Second)
	var u url.URL
	u.Scheme = "http"
	u.Host = "api.cinemate.cc"
	u.Path = "/account.watchlist"
	q := u.Query()
	q.Set("passkey", acc.passkey)
	q.Set("format", "xml")
	u.RawQuery = q.Encode()
	xmlBody, err := getXML(u.String())
	if err != nil {
		return
	}
	err = xml.Unmarshal(xmlBody, &list)
	return
}
