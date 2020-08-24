package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)
import log "github.com/sirupsen/logrus"
import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)
import "github.com/robfig/cron/v3"

var token = os.Getenv("TOKEN")
var b, _ = tb.NewBot(tb.Settings{
	Token:  token,
	Poller: &tb.LongPoller{Timeout: 10 * time.Second},
})

var cache = make(map[int]string)

func main() {

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
▘ ▘▝▀▘▝▀▘▌  ▘ ▘▝▀▘▝▘▘ ▘▝▀▘▘ ▘ %s@%s by %s
`, version, hash, "BennyThink")

	fmt.Printf("\n %c[1;32m%s%c[0m\n\n", 0x1B, banner, 0x1B)

	c := cron.New()
	//scheduler()
	//_, _ = c.AddFunc("1 1 */3 * *", scheduler)
	_, _ = c.AddFunc("* * * * *", scheduler)
	c.Start()

	b.Handle("/add", add)
	b.Handle(tb.OnCallback, on)
	b.Handle(tb.OnText, text)
	b.Handle("/start", start)
	log.Infoln("I'm running...")
	b.Start()
}
