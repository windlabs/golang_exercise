package enum

type (
	PermissionIndex int
	PermissionValue uint64
)

const (
	Executable PermissionIndex = iota + 1
	Write
	Read
)

func (v PermissionValue) Index() (i PermissionIndex) {
	i = 1
	for {
		if v == 0 {
			return 0
		}
		if v == 1 {
			return i
		}
		v = v >> 1
		i++
	}
}

func (i PermissionIndex) Value() PermissionValue {
	return 1 << (i - 1)
}

func HasPermission(permission PermissionValue, i PermissionIndex) (ret bool) {
	return permission&(1<<(i-1)) > 0
}

func HasRead(permission PermissionValue) (ret bool) {
	return permission&(1<<(Read-1)) > 0
}

func HasWrite(permission PermissionValue) (ret bool) {
	return permission&(1<<(Write-1)) > 0
}

func HasExecutable(permission PermissionValue) (ret bool) {
	return permission&(1<<(Executable-1)) > 0
}

func Split(permission PermissionValue) (PermissionList []PermissionIndex) {
	var i PermissionIndex = 1
	for {
		if permission == 0 {
			return
		}
		if permission&1 == 1 {
			PermissionList = append(PermissionList, i)
		}
		permission = permission >> 1
		i++
	}
	return
}

func Join(PermissionList ...PermissionIndex) (permission PermissionValue) {
	for _, v := range PermissionList {
		permission |= 1 << v
	}
	return
}
