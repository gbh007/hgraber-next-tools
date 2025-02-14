package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"strings"

	"github.com/gbh007/hgraber-next-tools/config"
	"github.com/gbh007/hgraber-next-tools/dataprovider/masterAPI"
)

func main() {
	cfg, err := parseConfig()
	if err != nil {
		panic(err)
	}

	logger := initLogger(cfg.MainConfig)

	client, err := masterAPI.New(cfg.MainConfig.MasterAPI.Addr, cfg.MainConfig.MasterAPI.Token)
	if err != nil {
		logger.Error(
			"fail init master api",
			slog.String("error", err.Error()),
		)

		os.Exit(1)
	}

	f, err := os.Open(cfg.Path)
	if err != nil {
		logger.Error(
			"fail open archive",
			slog.String("error", err.Error()),
		)

		os.Exit(1)
	}

	id, err := client.UploadArchive(context.Background(), f)

	_ = f.Close()

	if err != nil {
		logger.Error(
			"fail upload archive",
			slog.String("error", err.Error()),
		)

		os.Exit(1)
	}

	// Не самая лучшая сборка адреса, но в данном случае это не критично.
	u := strings.TrimRight(cfg.MainConfig.MasterAPI.Addr, "/") +
		"/#/book/" +
		id.String()

	logger.Info(
		"upload archive success",
		slog.String("id", id.String()),
		slog.String("url", u),
	)

}

type Config struct {
	MainConfig config.Config
	Path       string
}

func parseConfig() (Config, error) {
	configPath := flag.String("config", "config.yaml", "path to config")
	generateConfig := flag.String("generate-config", "", "generate example config")
	archivePath := flag.String("archive", "", "path to upload archive")
	useEnv := flag.Bool("use-env", false, "use env config")
	flag.Parse()

	if *generateConfig != "" {
		err := config.ExportToFile(config.DefaultConfig(), *generateConfig)
		if err != nil {
			panic(err)
		}

		os.Exit(0)
	}

	c, err := config.ImportConfig(*configPath, *useEnv)

	return Config{
		MainConfig: c,
		Path:       *archivePath,
	}, err
}

func initLogger(cfg config.Config) *slog.Logger {
	slogOpt := &slog.HandlerOptions{
		AddSource: cfg.Application.Debug,
		Level:     slog.LevelInfo,
	}

	if cfg.Application.Debug {
		slogOpt.Level = slog.LevelDebug
	}

	return slog.New(
		slog.NewJSONHandler(
			os.Stderr,
			slogOpt,
		),
	)
}
