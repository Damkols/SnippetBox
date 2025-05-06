package main

import (
    "net/http"
    "flag"
    "log/slog"
    "os"
)

type application struct{ //--> this struct is a blueprint for what our application struct should look like, we will use it for dependency injection
    logger *slog.Logger
}

func main() {

    addr := flag.String("addr", ":4000", "HTTP network address") //--> command line flags

    flag.Parse()

    logger:= slog.New(slog.NewTextHandler(os.Stdout, nil)) //--> initializing a structured logger

    app := &application{ //--> creates a new struct using the application blueprint, get the memory address and store it in app
        logger: logger, //--> stores the memory address of our initialized structured logger
    }

    logger.Info("starting server", "addr", *addr) //--> log starting server on port :4000 to the terminal

    err:= http.ListenAndServe(*addr, app.routes()) //--> check for errors

    logger.Error(err.Error()) //--> if there is an error log it to the terminal

    os.Exit(1)
}
