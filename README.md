# cinemate
Реализация API v2 сайта сinemate.cc на языке Go.

Установка

go get github.com/serbe/cinemate

Использование:

Инициализация:

import (
	"fmt"
	"github.com/serbe/cinemate"
)

c := cinemate.Init("ваш ключ API")

Получить подробную информацию о персоне:

person, _ := c.GetPerson(68675)
fmt.Println(person[0].Name)

-> Джейк Джилленхол

Получить подробную информацию о фильме:

movie, _ := c.GetMovie(68675)
fmt.Println(movie[0].TitleRussian)
fmt.Println(movie[0].Imdb.Rating)

-> Криминальная фишка от Генри
-> 6.0

Подробные примеры находятся в каталоге [examples](https://github.com/serbe/cinemate/examples).
