package main

import (
  "context"
  "fmt"
  "net/http"

  "github.com/jszwedko/go-circleci"
)

func doCircleCI() ([]*termui.Table, error) {
  // Check that the CircleCI API token is not empty.
  if len(circleciToken) <= 0 {
    logrus.Warn("CircleCI API token cannot be empty")
    logrus.Info("skipping CircleCI data")
    return nil, nil
  }

  tables := []*termui.Table{}

  // Initialize the CirclCI client
  circleciClient := &circleci.Client{Token: circleciToken}
}
