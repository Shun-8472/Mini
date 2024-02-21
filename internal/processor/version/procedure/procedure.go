package procedure

import "demo/internal/processor/version/substance"

type Procedure interface {
	GetDemoInfo() (substance.DemoInfo, error)
}
