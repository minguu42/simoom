package usecase

import (
	"runtime/debug"
	"slices"
)

type MonitoringUsecase struct{}

type CheckHealthOutput struct {
	Revision string
}

func (u MonitoringUsecase) CheckHealth() CheckHealthOutput {
	revision := "xxxxxxx"
	if info, ok := debug.ReadBuildInfo(); ok {
		if i := slices.IndexFunc(info.Settings, func(s debug.BuildSetting) bool {
			return s.Key == "vcs.revision"
		}); i != -1 {
			revision = info.Settings[i].Value[:len(revision)]
		}
	}
	return CheckHealthOutput{Revision: revision}
}
