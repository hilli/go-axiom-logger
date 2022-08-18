package goaxiomlogger

import (
	log "github.com/sirupsen/logrus"

	adapter "github.com/axiomhq/axiom-go/adapters/logrus"
)

func init() {
	// Export `AXIOM_TOKEN`, `AXIOM_ORG_ID` (when using a personal token) and
	// `AXIOM_DATASET` for Axiom Cloud.
	// Export `AXIOM_URL`, `AXIOM_TOKEN` and `AXIOM_DATASET` for Axiom Selfhost.

	hook, err := adapter.New()
	if err != nil {
		log.Fatal(err)
	}
	log.RegisterExitHandler(hook.Close)
	log.AddHook(hook)
}

// func main() {
// 	log.WithField("mood", "hyped").Info("This is awesome!")
// 	log.WithField("mood", "worried").Warn("This is no that awesome...")
// 	log.WithField("mood", "depressed").Error("This is rather bad.")
// 	log.Infof("Exiting...")
// 	log.Exit(0)
// }