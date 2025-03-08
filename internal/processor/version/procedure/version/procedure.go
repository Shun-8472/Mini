package version

import (
	"mini/internal/processor/version/substance/version"
)

type Procedure interface {
	GetDemoInfo() (version.DemoInfo, error)
}
