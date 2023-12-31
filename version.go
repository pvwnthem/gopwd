package main

import "github.com/blang/semver/v4"

func GetVersion() semver.Version {
	version := semver.Version{
		Major: 1,
		Minor: 2,
		Patch: 7,
		Pre: []semver.PRVersion{
			{VersionStr: "git"},
		},
		Build: []string{"HEAD"},
	}
	return version
}
