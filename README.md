# cinemate
Реализация API v2 сайта сinemate.cc на языке Go.

## Установка

``` sh
go get github.com/serbe/cinemate
```

## Использование:

**Инициализация:**

``` go
import (
	"fmt"
	"github.com/serbe/cinemate"
)

c := cinemate.Init("ваш ключ API")
```

**Получить подробную информацию о персоне (актере/режиссере):**

```go
person, _ := c.GetPerson(68675)
fmt.Println(person.Name)

> Джейк Джилленхол
```

**Получить подробную информацию о фильме:**

``` go
movies, _ := c.GetMovie(68675)
fmt.Println(movies[0].TitleRussian)
fmt.Println(movies[0].Imdb.Rating)

> Криминальная фишка от Генри
> 6.0

```

**Получить статистику сайта за последние сутки:**

``` go
stats, _ := cinemate.GetStatsNew()
fmt.Println(stats.UsersCount)

> 4

```
