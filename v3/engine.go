package main

import (
	"database/sql"
	"fmt"

	"github.com/hhhhkkk/mini-blog/v3/dialect"
	"github.com/hhhhkkk/mini-blog/v3/log"
	"github.com/hhhhkkk/mini-blog/v3/session"

	_ "github.com/mattn/go-sqlite3"
)

type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

func NewEngine(driver, source string) (*Engine, error) {
	d, ok := dialect.GetDialect(driver)
	if !ok {
		return nil, fmt.Errorf("dialect %q is not registered", driver)
	}

	var db, err = sql.Open(driver, source)

	if err != nil {
		log.Error(err)
		return nil, err
	}
	// Send a ping to make sure the database connection is alive.
	if err = db.Ping(); err != nil {
		log.Error(err)
		return nil, err
	}
	e := &Engine{
		db:      db,
		dialect: d,
	}
	log.Info("Connect database success")
	return e, nil
}

func (engine *Engine) Close() {
	if engine == nil || engine.db == nil {
		return
	}
	if err := engine.db.Close(); err != nil {
		log.Error("Failed to close database")
	}
	log.Info("Close database success")
}

func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db, engine.dialect)
}
