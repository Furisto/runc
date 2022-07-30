package libcontainer

import (
	"github.com/Furisto/runc/libcontainer/cgroups"
	"github.com/Furisto/runc/libcontainer/intelrdt"
	"github.com/Furisto/runc/types"
)

type Stats struct {
	Interfaces    []*types.NetworkInterface
	CgroupStats   *cgroups.Stats
	IntelRdtStats *intelrdt.Stats
}
