package main

import (
	"log/slog"
	"os"
)

// ref: https://go.dev/blog/slog

type Home struct {
	Name    string
	Country string
}

type HomeV2 struct {
	Name    string
	Country string
}

func (r HomeV2) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("diff_name", r.Name),
		slog.String("diff_country", r.Country),
	)
}

func main() {
	slog.Info("Hello World", "name", "aash", "surname", "dhariya")
	slog.Info("Hello World", "name", "aash", "surname") // doesn't work

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("Hello World", "aash", "dhariya") // prints everyting in the form of key value pairs

	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("Hello World", "aash", "dhariya") // prints everyting in the form of key value pairs and JSON

	// to have less memory allocations and increase performance
	logger.Info("Print a lot of vars", slog.Int("count", 10), slog.Bool("heavy", true), slog.Any("extra", map[string]string{"name": "surname"}))
	logger.Debug("Debug Print") // doesn't print

	logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug, AddSource: true})) // sets the level and adds source of the log as well.
	logger.Debug("Debug Print Again")

	logger.With("new", "args").Debug("Call With")
	logger.With(slog.Int("count", 100), slog.String("value", "new-val")).Debug("Call with again")

	h := Home{"Meadows", "Ahmedabad"}
	hv2 := HomeV2{"Meadows", "Ahmedabad"}
	logger.Info("Home Value", "Home", h)
	logger.Info("Home v2 Value", "HomeV2", hv2)

}
