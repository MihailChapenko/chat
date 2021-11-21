module github.com/MihailChapenko/chat

// +heroku goVersion go1.16
go 1.17

require (
	github.com/go-chi/chi/v5 v5.0.0
	github.com/lib/pq v1.2.0
	golang.org/x/crypto v0.0.0-20211108221036-ceb1ce70b4fa
)

require github.com/jmoiron/sqlx v1.3.4

require gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
