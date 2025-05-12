Қазақша | [English](README-en.md) | [Qazaq Grammar Latın](README-qg.md)

# Köreyik!

![Banner](img/small.png)

###### Қазақ аниме энциклопедиясы. The Kazakh anime encyclopedia. 

Бұл — кіші жоба. Қазіргі таңға, ол әрең-әрең жылжып дамып жатыр. Оның басты мақсаты — аниме туралы қазақша қорлар көбейту.

Бұл — жобаның **backend-side** репозиторийі. **Frontend-side** репозиторийі -> [merki86/koreyik-frontend](https://github.com/merki86/koreyik-frontend) <- сол жерде орналасқан.

## Бағдарлар
`api/anime/` — **GET** — Аниме жазбаларының барлығын дерекқордан алу.\
`api/anime/` — **POST** — Дерекқорға жаңа аниме жазбасы қою.

`api/anime/<id>` — **GET** — Көрсетілген ID дерекқордағы аниме жазбасын алу.\
`api/anime/<id>` — **PUT** — Көрсетілген ID дерекқордағы аниме жазбасын жаңарту.\
`api/anime/<id>` — **DELETE** — Көрсетілген ID дерекқордағы аниме жазбасын жою.

`api/anime/random` — **GET** — Кездейсоқ аниме жазбасы бар бағдарға жеткізілу.

## JSON'ды жариялау (POST request)
POST сұратымының денесінің (request body) үлгісі.
```json
{
  "ID": 1,
  "ThumbnailURL": "https://some-site.kz/some-thumbnail.png",
  "Description": "Бір қарағанда, Хори кәдімгі жасөспірім қыз...",
  "Rating": "PG-13",
  "TitleKk": "Хоримия",
  "TitleJp": "ホリミヤ",
  "TitleEn": "Horimiya",
  "Status": "Аяқталған",
  "StartedAiring": "2021-01-10T00:00:00Z",
  "FinishedAiring": "2021-04-04T00:00:00Z",
  "Genres": [
    "Романтика",
    "Өмір үзіндісі"
  ],
  "Themes": [
    "Мектеп"
  ],
  "Seasons": 2,
  "Episodes": 23,
  "Duration": 23,
  "Studios": [
    "CloverWorks"
  ],
  "Producers": [
    "Тошихиро Маеда",
    "Тайто Ито",
    "..."
  ]
}
```

## Конфигурация (пішіндеме) файлы
Көбірек ақпарат [осы жерде](/config/README.md) оқи аласыз.

## Жүргізу жолы
Taskfile арқылы серверді жүргізе аласыз
```
> task
or
> task run
or
> task r
```
Сондай-ақ, билддау (to build) үшін
```
> task build
or
> task b
```

## Қолданылатын технологиялар:

### Басты:
- Main programming language - [Go](https://go.dev/)
- Database - [PostgreSQL](https://www.postgresql.org/about/)
- ORM - [GORM](https://gorm.io/)

### Go тілінің кітапханалары:
- HTTP Server - [go-chi/chi](https://github.com/go-chi/chi)
- Dotenv support - [joho/godotenv](https://github.com/joho/godotenv)
- Config support - [ilyakaznacheev/cleanenv](https://github.com/ilyakaznacheev/cleanenv)
- Pretty slog logger - [greyxor/slogor](https://gitlab.com/greyxor/slogor)
