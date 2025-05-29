// this is the core service for the spud application
// we use this to spin up the server, read the config, etc etc
package spud

import "github.com/go-chi/chi/v5"

type App struct {
	AdminPath string
	AppRouter *chi.Mux
	// SQLite path
	DbPath string
	Port   int
	Host   string
}

func NewApp(adminPath string, dbPath string) *App {
	return &App{
		AdminPath: adminPath,
		AppRouter: chi.NewRouter(),
		DbPath:    dbPath,
	}
}

func (a *App) Run() error {
	return nil
}
