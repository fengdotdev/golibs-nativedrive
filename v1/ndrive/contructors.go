package ndrive

import "github.com/fengdotdev/golibs-nativedrive/v1/nelement"

func NewNDrive(workingDir string) *NDrive {
	return &NDrive{
		workingDir: workingDir,
		self:       nelement.RootElement(),
	}
}
