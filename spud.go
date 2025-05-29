// this is the core service for the spud application
// we use this to spin up the server, read the config, etc etc
package spud

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/orsinium-labs/enum"
	"golang.org/x/sync/errgroup"
)

type mode enum.Member[string]

var (
	Testing    = mode{"testing"}
	Production = mode{"production"}
)

type App struct {
	AdminPath string
	AppRouter *chi.Mux
	// SQLite path
	dbType  string
	DbPath  string
	port    int
	host    string
	appMode mode
}

func NewApp(adminPath string, appRouter *chi.Mux, options ...func(*App)) *App {
	app := &App{
		AdminPath: adminPath,
		AppRouter: appRouter,

		// setup sane defaults
		appMode: Production,
		port:    8080,
		host:    "0.0.0.0",
	}
	for _, option := range options {
		option(app)
	}
	// call initialize to prep stuff
	// we will use to this set it up aside from running, which
	// means we can reuse this for testing
	return app
}

func WithMode(mode mode) func(*App) {
	return func(a *App) {
		a.appMode = mode
	}
}

// functional options pattern on the app to control port + host
func WithPort(port int) func(*App) {
	return func(a *App) {
		a.Port = port
	}
}

func WithHost(host string) func(*App) {
	return func(a *App) {
		a.host = host
	}
}

func WithSQLite(dbPath string) func(*App) {
	return func(a *App) {
		a.dbType = "sqlite"
		a.DbPath = dbPath
	}
}
func (a *App) initialize() {
	// initialize the app
	// setup the db
	// logger
}

func (a *App) Run(ctx context.Context) error {
	// start http server within an err group so we can stop the db after

	sctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	slog.Info("Starting server", "port", a.port, "host", a.host)
	eg, egctx := errgroup.WithContext(sctx)
	eg.Go(func() error {
		router := chi.NewMux()

		router.Use(
			middleware.Logger,
			middleware.Recoverer,
		)

		// router.Handle("/static/*", static())
		// setup admin routes
		if err := routes.SetupRoutes(egctx, router); err != nil {
			return fmt.Errorf("error setting up routes: %w", err)
		}

		srv := &http.Server{
			Addr:     "0.0.0.0:" + getPort(),
			Handler:  router,
			ErrorLog: slog.NewLogLogger(slog.Default().Handler(), slog.LevelError),
		}

		go func() {
			<-egctx.Done()
			if err := srv.Shutdown(context.Background()); err != nil {
				log.Fatalf("error during shutdown: %v", err)
			}
		}()

		return srv.ListenAndServe()
	})

	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}
