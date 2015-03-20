package config

import (
    "github.com/Sirupsen/logrus"
)


func GetLogEntry(ns string) (*logrus.Entry) {
    e := logrus.WithField("where", ns)
    return e
}
