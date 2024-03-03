package term

import (
	"time"

	"github.com/briandowns/spinner"
)

const withMissageMinDuration = 700 * time.Millisecond
const withoutMissageMinDuration = 400 * time.Millisecond

var s = spinner.New(spinner.CharSets[33], 100*time.Millisecond)
var startedAt time.Time

var lastMessage string
var active bool

func StartSpinner(msg string) {
	if active {
		if msg == lastMessage {
			return
		}

		s.Stop()
	}

	startedAt = time.Now()
	s.Prefix = msg + " "
	lastMessage = msg
	s.Start()
	active = true
}

func StopSpinner() {
	elapsed := time.Since(startedAt)

	if lastMessage != "" && elapsed < withMissageMinDuration {
		time.Sleep(withMissageMinDuration - elapsed)
	} else if elapsed < withoutMissageMinDuration {
		time.Sleep(withoutMissageMinDuration - elapsed)
	}

	s.Stop()
	ClearCurrentLine()

	active = false
}

func ResumeSpinner() {
	if !active {
		StartSpinner(lastMessage)
	}
}
