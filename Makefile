# TAG = $(shell git tag -l --points-at HEAD)
# COMMIT = $(shell git rev-parse --short HEAD)
# BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
# BUILD_TIME = $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

# CORE_MOD_NAME = github.com/gbh007/hgraber-next-tools
# LDFLAGS = -ldflags "-X '$(CORE_MOD_NAME)/version.Version=$(TAG)' -X '$(CORE_MOD_NAME)/version.Commit=$(COMMIT)' -X '$(CORE_MOD_NAME)/version.BuildAt=$(BUILD_TIME)' -X '$(CORE_MOD_NAME)/version.Branch=$(BRANCH)'"


.PHONY: update-dep
update-dep:
	go get -u github.com/gbh007/hgraber-next-agent-core@master
	go get -u github.com/gbh007/hgraber-next@master
	go mod tidy