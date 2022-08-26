package enum

type (
	PermissionType uint64
)

const (
	PermissionExecutable PermissionType = 1 << iota
	PermissionWrite
	PermissionRead
)

func (permission PermissionType) HasPermission(t PermissionType) bool {
	return permission&t == t
}

func (permission PermissionType) HasRead() bool {
	return permission&PermissionRead == PermissionRead
}

func (permission PermissionType) HasWrite() bool {
	return permission&PermissionWrite == PermissionWrite
}

func (permission PermissionType) HasExecutable() bool {
	return permission&PermissionExecutable == PermissionExecutable
}

func (permission PermissionType) Split() (PermissionList []PermissionType) {
	var i = 0
	for {
		if permission == 0 {
			return
		}
		if permission&1 == 1 {
			PermissionList = append(PermissionList, 1<<i)
		}
		permission = permission >> 1
		i++
	}
	return
}

func (permission PermissionType) Join(PermissionList ...PermissionType) PermissionType {
	for _, v := range PermissionList {
		permission |= v
	}
	return permission
}
