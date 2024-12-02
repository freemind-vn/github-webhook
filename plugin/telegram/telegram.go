package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"text/template"

	"freemind.com/webhook/plugin"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Config struct {
	BotToken string `yaml:"bot_token"`
	Template string
}

var (
	opts   *Config
	tpl    *template.Template
	ctx    context.Context
	cancel context.CancelFunc
	b      *bot.Bot
)

// -----------------------------------------------------------------------------
// Plugin implementation
// -----------------------------------------------------------------------------

func Start(config string) error {
	var err error

	if err = plugin.ReadConfig(config, &opts); err != nil {
		return err
	}

	// opts = Default

	if tpl, err = template.New("telegram").Parse(opts.Template); err != nil {
		return err
	}

	if b, err = bot.New(opts.BotToken); err != nil {
		return err
	}

	go func() {
		ctx, cancel = signal.NotifyContext(context.Background(), os.Interrupt)
		defer cancel()
		b.Start(ctx)
	}()

	return err
}

func Get(req *http.Request) ([]byte, error) {
	return []byte("Hello, world!"), nil
}

func Post(req *http.Request) ([]byte, error) {
	chatID := req.URL.Query().Get("chat_id")
	if chatID == "" {
		slog.Info("no chat_id")
		return nil, nil
	}

	payload, err := plugin.ReadBodyJson(req)
	payload["event"] = req.Header.Get("X-GitHub-Event")

	var buf bytes.Buffer
	err = tpl.Execute(&buf, payload)
	if err != nil {
		return nil, err
	}

	msg := strings.TrimSpace(buf.String())
	if msg == "" {
		return nil, errors.New("chat_id must be present in the query params")
	}

	slog.Info("send message", "body", msg)

	m, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		// Text:      bot.EscapeMarkdown(msg),
		// ParseMode: models.ParseModeMarkdown,
		Text:      msg,
		ParseMode: models.ParseModeHTML,
	})
	if err != nil {
		return nil, err
	}

	return json.Marshal(&m)
}
