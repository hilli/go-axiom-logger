package axiomLogger

import (
	"os"

	"github.com/sirupsen/logrus"

	adapter "github.com/axiomhq/axiom-go/adapters/logrus"
)

var hostname string

func init() {

}

func New() *logrus.Entry {
	logger := logrus.New()
	hostname, _ = os.Hostname() // Best effort here

	hook, err := adapter.New()
	if err != nil {
		logrus.Warnf("Unable to configure Axiom: %v", err) // Still best effort here
	} else {
		logrus.RegisterExitHandler(hook.Close)
		logrus.AddHook(hook)
		// logrus.SetReportCaller(true)
	}
	entry := logger.WithFields(logrus.Fields{
		"hostname": hostname,
	})
	return entry
}
