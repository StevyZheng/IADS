package zfs

import (
	"github.com/bicomsystems/go-libzfs"
	"github.com/pkg/errors"
	"path"
)

type VDevType string

type Pool struct {
	Disks    []string
	Level    VDevType
	PoolName string
}

func (p *Pool) AppendDisk(disk string) {
	_ = append(p.Disks, disk)
}

func (p *Pool) SetLevel(l VDevType) {
	p.Level = l
}

func (p *Pool) SetPoolName(poolName string) {
	p.PoolName = poolName
}

func (p Pool) Create() (err error) {
	if len(p.Disks) == 0 {
		err = errors.Wrap(err, "zpool disks is nil.")
		return
	}
	if p.PoolName == "" || p.Level == "" {
		err = errors.Wrap(err, "poolname or level is nil.")
		return
	}
	var vdev zfs.VDevTree
	var vdevs, mdevs []zfs.VDevTree
	for _, d := range p.Disks {
		mdevs = append(mdevs, zfs.VDevTree{Type: zfs.VDevTypeDisk, Path: d})
	}
	vdevs = []zfs.VDevTree{
		zfs.VDevTree{Type: zfs.VDevTypeRaidz, Devices: mdevs},
	}
	poolName := p.PoolName
	vdev.Devices = vdevs
	props := make(map[zfs.Prop]string)
	fsprops := make(map[zfs.Prop]string)
	features := make(map[string]string)
	fsprops[zfs.DatasetPropMountpoint] = path.Join("/mnt", poolName)
	features["async_destroy"] = "enabled"
	features["empty_bpobj"] = "enabled"
	features["lz4_compress"] = "enabled"
	pool, err := zfs.PoolCreate(poolName, vdev, features, props, fsprops)
	if err != nil {
		println("Error: ", err.Error())
		return
	}

	defer pool.Close()
	return err
}
