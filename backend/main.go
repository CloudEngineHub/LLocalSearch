package main

import (
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
	"github.com/tmc/langchaingo/memory"
)

type Sessions map[string]*memory.ConversationWindowBuffer

var sessions Sessions = make(Sessions)

func main() {
	w := os.Stderr
	slog.SetDefault(slog.New(
		tint.NewHandler(w, &tint.Options{
			AddSource: true,
			Level:     slog.LevelDebug,
		}),
	))
	// lschains.RunSourceChainExample()
	slog.Info("Starting the server")
	StartApiServer()
}
