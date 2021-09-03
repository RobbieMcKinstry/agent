package main

import (
  log "github.com/sirupsen/logrus"
  "github.com/RobbieMcKinstry/colorformatter"
  // • Use Viper to load environment variables.
  _ "github.com/RobbieMcKinstry/agent/cli"
)

// TODO: 
// • Use Codegangsta CLI to capture CLI flags and specify
// • Identify the primary abstractions:
//     • Have an interface for Postgres vs MySQL vs Aurora
//     • Have a Map containing | Version -> Knobs
//     • Use an interface to abstract this detail. We should be able
//       to query a backend and collect the knobs and metrics we
//       want to select without hard-coding them.
// • Provide a way to marshal Knob and Metric data into strong types?
//     • Do we use `go generate`? Reflection?
// • What do we use to query? Squirrel?

func main() {
  // • Set up Logrus to colorize output if its in a TTY
  var formatter = colorformatter.New()
  log.SetFormatter(formatter)
  log.Info("Foobar")
}
