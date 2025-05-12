[Қазақша](README.md) | English | [Qazaq Grammar Latın](README-qg.md)

# Köreyik!

![Banner](img/small.png)

###### Қазақ аниме энциклопедиясы. The Kazakh anime encyclopedia. 

This is a small project. It is developing very slowly at the moment. It's main goal is to increase resources about anime in Kazakh.

It is the repository of **backend-side** of the project. The **frontend-side** repository is located here -> [merki86/koreyik-frontend](https://github.com/merki86/koreyik-frontend).

## Routes
`api/anime/` — **GET** — Get all records of anime in database.\
`api/anime/` — **POST** — Create a new record of anime in database.

`api/anime/<id>` — **GET** — Get a record of anime in database with specified ID.\
`api/anime/<id>` — **PUT** — Update a record of anime in database with specified ID.\
`api/anime/<id>` — **DELETE** — Delete a record of anime in database with specified ID.

`api/anime/random` — **GET** — Redirect on a route with random anime record.

## Post JSON
Example body of POST request.
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

## Configuration file
See more info [there](/config/README.md).

## How to run
You can run the server via Taskfile command:
```
> task
or
> task run
or
> task r
```
And also you can build the project via
```
> task build
or
> task b
```

## Used technologies:

### General:
- Main programming language - [Go](https://go.dev/)
- Database - [PostgreSQL](https://www.postgresql.org/about/)
- ORM - [GORM](https://gorm.io/)

### Go language external packages/libraries:
- HTTP Server - [go-chi/chi](https://github.com/go-chi/chi)
- Dotenv support - [joho/godotenv](https://github.com/joho/godotenv)
- Config support - [ilyakaznacheev/cleanenv](https://github.com/ilyakaznacheev/cleanenv)
- Pretty slog logger - [greyxor/slogor](https://gitlab.com/greyxor/slogor)
