package axiomLogger

import (
	"testing"
)

func TestLogging(t *testing.T) {
	logger := New()
	logger.WithField("mood", "hyped").Info("This is awesome!")
	logger.WithField("mood", "worried").Warn("This is no that awesome...")
	logger.WithField("mood", "depressed").Error("This is rather bad.")
	logger.Infof("Exiting...")
	
}
