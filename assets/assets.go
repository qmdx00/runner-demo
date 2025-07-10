package assets

import (
	_ "embed"
)

var (
	//go:embed sprites/runner_idle.png
	RunnerIdleImage []byte

	//go:embed sprites/runner_run.png
	RunnerRunImage []byte

	//go:embed sprites/runner_jump.png
	RunnerJumpImage []byte

	//go:embed tiles/background.png
	BackgroundImage []byte
)
