package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/realmallaury/teltech/cmd/handler"
	"github.com/realmallaury/teltech/internal/cache"

	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Config stores app related configuration data.
type Config struct {
	Host            string
	ShutdownTimeout time.Duration
	CacheSize       int
	CacheTTL        time.Duration
}

func main() {
	if err := run(); err != nil {
		log.Println("shutting down", "error:", err)
		os.Exit(1)
	}
}

func run() error {
	logger := log.New(os.Stdout, "Teltech : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	config := Config{
		Host:            "0.0.0.0:8080",
		ShutdownTimeout: 5 * time.Second,
		CacheSize:       1000,
		CacheTTL:        1 * time.Minute,
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	f := pflag.NewFlagSet(os.Args[0], pflag.ContinueOnError)
	f.String("host", config.Host, "the host and port of the CMS")
	f.Duration("shutdown-timeout", config.ShutdownTimeout, "server shutdown timeout")
	f.Int("cache-size", config.CacheSize, "maximum cache size")
	f.Duration("cache-ttl", config.CacheTTL, "cache ttl duration")

	if err := f.Parse(os.Args[1:]); err != nil {
		return err
	}

	if err := viper.BindPFlags(f); err != nil {
		return err
	}

	config.Host = viper.GetString("host")
	config.ShutdownTimeout = viper.GetDuration("shutdown-timeout")

	logger.Printf("Config: %+v", config)

	store := cache.NewStore(config.CacheSize, config.CacheTTL)

	api := &http.Server{
		Addr:    config.Host,
		Handler: handler.Router(ctx, logger, store),
	}

	serverErrors := make(chan error, 1)

	go func() {
		logger.Printf("API Listening on %s", config.Host)
		serverErrors <- api.ListenAndServe()
	}()

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		return errors.Wrap(err, "starting server")

	case <-osSignals:
		logger.Println("Start shutdown...")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), config.ShutdownTimeout)
		defer cancel()

		err := api.Shutdown(shutdownCtx)
		if err != nil {
			log.Printf("main : Graceful shutdown did not complete in %v : %v", config.ShutdownTimeout, err)
			err = api.Close()
		}

		if err != nil {
			return errors.Wrap(err, "could not stop server gracefully")
		}
	}

	return nil
}
