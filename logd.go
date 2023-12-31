package logd

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

type logger struct{}

var (
	out   = fmt.Println
	build = fmt.Sprintf

	blue   = color.New(color.FgBlue).SprintfFunc()
	red    = color.New(color.FgRed).SprintfFunc()
	green  = color.New(color.FgHiGreen).SprintfFunc()
	yellow = color.New(color.FgYellow).SprintfFunc()
	violet = color.New(color.FgHiMagenta).SprintfFunc()
)

func New() *logger {
	return &logger{}
}

// FatalF and ErrorF are the same expect that FatalF uses os.Exit(1)
func (l *logger) Fatalf(format string, args ...interface{}) {
	messageToLog := red(format, args)

	fatal := color.New(color.Bold, color.BgHiRed, color.FgWhite).Sprint("FATAL")

	out(buildMessage(fatal + " " + messageToLog))
	os.Exit(1)
}

// see  difference between Fatalf and Errorf above
func (l *logger) Errorf(format string, args ...interface{}) {
	messageToLog := red(format, args)

	out(buildMessage(messageToLog))
}

func (l *logger) Infof(format string, args ...interface{}) {
	messageToLog := blue(format, args)

	out(buildMessage(messageToLog))
}

func (l *logger) Warnf(format string, args ...interface{}) {
	messageToLog := violet(format, args)

	out(buildMessage(messageToLog))
}

func (l *logger) Debugf(format string, args ...interface{}) {
	messageToLog := green(format, args)

	out(buildMessage(messageToLog))
}

func buildMessage(messageToPrint string) string {
	t := timeNow()
	return build("%s - %s ", t, messageToPrint)
}

func timeNow() string {
	t := time.Now().Format("02.01.2006 15:04:05")

	return yellow("%s", t)
}
