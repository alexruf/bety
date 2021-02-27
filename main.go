package main

import (
	"context"
	"database/sql"
	"github.com/alexruf/bety/api"
	"github.com/alexruf/bety/server"
	"github.com/alexruf/bety/ui"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	//configure logger
	log.SetLevel(log.DebugLevel)

	//init DB connection
	db, err := connectDatabase("sqlite3", "db.sqlite")
	if err != nil {
		log.WithError(err).Fatal("Error connecting to database")
	}
	defer func() {
		log.Debug("Closing database connection...")
		if err := db.Close(); err != nil {
			log.WithError(err).Error("Error closing database connection")
		}
	}()

	//DB migration
	if err := migrateDatabase(db, "file://./migrations", "sqlite3"); err != nil {
		log.WithError(err).Fatal("Error migrating database")
	}

	//initialize application server
	appServer := server.New(db)
	appServer.RegisterRoutes()
	//initialize api
	appApi := api.New(appServer)
	appApi.RegisterRoutes()
	//initialize web ui
	webUi := ui.New(appServer)
	webUi.RegisterRoutes()

	//configure HTTP server
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      appServer.Router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	//notify on system interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sig := <-c
		log.Debugf("Received system signal: %+v", sig)
		cancel()
	}()

	//start HTTP server
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.WithError(err).Fatal("Unhandled error")
		}
	}()

	log.Infof("Start listening on: '%s'", srv.Addr)

	//wait for system interrupt
	<-ctx.Done()

	log.Infof("Stopped listening on: '%s'", srv.Addr)

	//create shutdown context with timeout
	ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	//shutdown application server
	if err := appServer.Shutdown(ctxShutdown); err != nil {
		log.WithError(err).Error("Error during application server shutdown")
	}
	//shutdown HTTP server
	if err := srv.Shutdown(ctxShutdown); err != nil {
		log.WithError(err).Error("Error during HTTP server shutdown")
	}
	log.Info("Exited properly")
}

func connectDatabase(driverName string, dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func migrateDatabase(db *sql.DB, sourceURL string, databaseName string) error {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{MigrationsTable: "migration_table"})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(sourceURL, databaseName, driver)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
