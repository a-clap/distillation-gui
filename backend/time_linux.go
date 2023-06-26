package backend

import (
	"fmt"
	"os/exec"
	"regexp"
	"time"

	"github.com/a-clap/logging"
)

const dateExecutable = "timedatectl"

var (
	ntpServiceRE = regexp.MustCompile(`NTP service:\s?(.+)\n`)
)

func (b *Backend) NTPGet() (bool, error) {
	out, err := exec.Command(dateExecutable, "status").Output()
	if err != nil {
		logger.Error("failed to read ntp", logging.String("error", err.Error()))
		return false, fmt.Errorf("%w: NTPGet", err)
	}

	got := ntpServiceRE.FindStringSubmatch(string(out))
	if len(got) != 2 {
		logger.Error("failed to parse output", logging.Strings("out", got))
		return false, fmt.Errorf("unexpected out: %v", got)
	}

	switch got[1] {
	case "inactive":
		return false, nil
	case "active":
		return true, nil
	}

	logger.Error("unknown output", logging.String("out", got[1]))

	return false, fmt.Errorf("unknown output %v", got[1])
}

func (b *Backend) NTPSet(enable bool) error {
	args := []string{"set-ntp", "false"}
	if enable {
		args[1] = "true"
	}

	logger.Debug("NTPSet")

	if err := exec.Command(dateExecutable, args...).Run(); err != nil {
		logger.Error("failed to set ntp", logging.Bool("ntp", enable), logging.String("error", err.Error()))
		b.eventEmitter.OnError(ErrSetTime)

		return fmt.Errorf("%w: NTPSet %v", err, dateExecutable)
	}

	return nil
}

func (b *Backend) TimeSet(ts int64) error {
	logger.Debug("TimeSet")
	dateString := time.UnixMilli(ts).Format("2006-01-02 15:04:05")
	args := []string{"set-time", dateString}

	if err := exec.Command(dateExecutable, args...).Run(); err != nil {
		logger.Error("failed to set time ", logging.String("error", err.Error()))
		b.eventEmitter.OnError(ErrSetTime)

		return fmt.Errorf("%w: TimeSet %v", err, dateExecutable)
	}

	return nil
}
