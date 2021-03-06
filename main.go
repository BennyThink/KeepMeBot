package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	log "github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/robfig/cron/v3"
)

var token = os.Getenv("TOKEN")
var b, err = tb.NewBot(tb.Settings{
	Token:  token,
	Poller: &tb.LongPoller{Timeout: 10 * time.Second},
})

func main() {
	deferInit()
	if err != nil {
		log.Panicf("Please check your network or TOKEN! %v", err)
	}
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)
	Formatter := &log.TextFormatter{
		EnvironmentOverrideColors: true,
		FullTimestamp:             true,
		TimestampFormat:           "2006-01-02 15:04:05",
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return fmt.Sprintf("[%s()]", f.Function), ""
		},
	}
	log.SetFormatter(Formatter)

	//  toilet  KeepMe.Run -f smblock
	banner := fmt.Sprintf(`
▌ ▌         ▙▗▌     ▛▀▖
▙▞ ▞▀▖▞▀▖▛▀▖▌▘▌▞▀▖  ▙▄▘▌ ▌▛▀▖
▌▝▖▛▀ ▛▀ ▙▄▘▌ ▌▛▀ ▗▖▌▚ ▌ ▌▌ ▌
▘ ▘▝▀▘▝▀▘▌  ▘ ▘▝▀▘▝▘▘ ▘▝▀▘▘ ▘ 
%s-%s at %s by %s
`, hash, branch, compileTime, "BennyThink")

	fmt.Printf("\n %c[1;32m%s%c[0m\n\n", 0x1B, banner, 0x1B)

	c := cron.New()
	switch os.Getenv("dev") {
	case "true":
		scheduler()
		_, _ = c.AddFunc("* * * * *", scheduler)
		log.SetLevel(log.DebugLevel)
	default:
		_, _ = c.AddFunc("*/15 * * * *", scheduler)
	}

	c.Start()

	b.Handle("/add", add)
	b.Handle(tb.OnText, onText)
	b.Handle(tb.OnCallback, onCallback)
	b.Handle("/start", start)
	b.Handle("/list", list)
	b.Handle("/history", history)
	b.Handle("/ping", ping)
	b.Handle(tb.OnEdited, edited)
	log.Infoln("I'm running...")
	b.Start()
}
