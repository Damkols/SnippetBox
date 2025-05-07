package main

import (
    "net/http"
    "flag"
    "log/slog"
    "os"
    "database/sql"

    "snippetbox.usmkols.net/internal/models"

    _"github.com/go-sql-driver/mysql"
)

type application struct{ //--> this struct is a blueprint for what our application struct should look like, we will use it for dependency injection
    logger *slog.Logger
    snippets *models.SnippetModel
}

func main() {

    addr := flag.String("addr", ":4000", "HTTP network address") //--> command line flags

    dsn:= flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name") //--> cmd flag for MYSQL DSN string

    flag.Parse()

    logger:= slog.New(slog.NewTextHandler(os.Stdout, nil)) //--> initializing a structured logger

    db, err := openDB(*dsn)
    if err != nil {
        logger.Error(err.Error())
        os.Exit(1)
    }

    defer db.Close() //--> defer call to db.Close(), this allows connection pool to close before main() func exits

    app := &application{ //--> creates a new struct using the application blueprint, get the memory address and store it in app
        logger: logger, //--> stores the memory address of our initialized structured logger
    }

    logger.Info("starting server", "addr", *addr) //--> log starting server on port :4000 to the terminal

    err:= http.ListenAndServe(*addr, app.routes()) //--> check for errors

    logger.Error(err.Error()) //--> if there is an error log it to the terminal

    os.Exit(1)
}


func openDB(dsn string) (*sql.DB, error) {  //--> openDB func wraps sql.Open() and returns a sql.DB connection pool
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil.err
    }

    err = db.Ping()

    if err != nil {
        db.Close()
        return nil, err
    }

    return db, nil
}