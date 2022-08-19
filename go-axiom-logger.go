package goaxiomlogger

import (
	log "github.com/sirupsen/logrus"

	adapter "github.com/axiomhq/axiom-go/adapters/logrus"
)

func init() {
	hook, err := adapter.New()
	if err != nil {
		log.Fatal(err)
	} else {
		log.RegisterExitHandler(hook.Close)
		log.AddHook(hook)
	}
}

// func main() {
// 	log.WithField("mood", "hyped").Info("This is awesome!")
// 	log.WithField("mood", "worried").Warn("This is no that awesome...")
// 	log.WithField("mood", "depressed").Error("This is rather bad.")
// 	log.Infof("Exiting...")
// 	log.Exit(0)
// }