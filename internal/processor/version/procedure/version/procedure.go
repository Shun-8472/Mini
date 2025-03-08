package version

import (
	"mini/internal/processor/version/substance/version"
)

type Procedure interface {
	GetVersionInfo() (version.DemoInfo, error)
}
