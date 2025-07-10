package assets

import (
	"embed"
)

var (
	//go:embed sprites/runner_idle.png
	RunnerIdleImage []byte

	//go:embed sprites/runner_run.png
	RunnerRunImage []byte

	//go:embed sprites/runner_jump.png
	RunnerJumpImage []byte
)

var (
	//go:embed background/background.png
	BackgroundImage []byte
)

var (
	//go:embed tiles/*
	TileFiles embed.FS
)
