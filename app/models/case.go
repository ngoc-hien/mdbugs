package models

import "github.com/ottob/go-semver/semver"

type Case struct {
	ID           int64
	Message      string
	GuideVersion *semver.Version
}

func NewCase(message string, version string) *Case {
	v, err := semver.NewVersion(version)
	if (err != nil) {
		panic(err)
	}
	
	return &Case{Message: message, GuideVersion : v}
}
