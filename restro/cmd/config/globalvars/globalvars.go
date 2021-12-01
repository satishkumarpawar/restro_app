package globalvars

type Glbv struct {
	UserID       uint64
	RoleID       uint64
	UserRoleName string
}

var Glbvs Glbv

func SetVars(gv Glbv) {
	Glbvs = gv
}

func GetVars() Glbv {
	return Glbvs
}
