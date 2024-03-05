package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"text/template"

	"github.com/KasimKaizer/copyanon/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	logger        *slog.Logger
	gists         *models.GistModel
	templateCache map[string]*template.Template
}

func main() {
	addr := flag.String("addr", "localhost:4000", "HTTP Network Address")
	debugFlag := flag.Bool("debug", false, "Enable debug mode")

	dataSource := flag.String("dataSourceName", "", "MySQL database source name")
	flag.StringVar(dataSource, "d", *dataSource, "Short hand for dataSourceName")

	flag.Parse()
	logHandlerOps := slog.HandlerOptions{}

	if *debugFlag {
		logHandlerOps.Level = slog.LevelDebug
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &logHandlerOps))

	db, err := openDB(*dataSource)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	templateCache, err := NewTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := &application{
		logger:        logger,
		gists:         &models.GistModel{DB: db},
		templateCache: templateCache,
	}

	logger.Info("server started", slog.String("address", *addr))

	err = http.ListenAndServe(*addr, app.router())
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(name string) (*sql.DB, error) {
	db, err := sql.Open("mysql", name)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
