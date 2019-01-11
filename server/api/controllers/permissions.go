package controllers

// Permission ...
type Permission string

// Permissions ...
type Permissions []Permission

const (
	// PermissionSession ...
	PermissionSession Permission = "Account"

	// PermissionCreateUser ...
	PermissionCreateUser Permission = "CreateUser"
	// PermissionRetrieveUser ...
	PermissionRetrieveUser Permission = "RetrieveUser"
	// PermissionUpdateUser ...
	PermissionUpdateUser Permission = "UpdateUser"
	// PermissionDeleteUser ...
	PermissionDeleteUser Permission = "DeleteUser"

	// PermissionCreateGroup ...
	PermissionCreateGroup Permission = "CreateGroup"
	// PermissionRetrieveGroup ...
	PermissionRetrieveGroup Permission = "RetrieveGroup"
	// PermissionUpdateGroup ...
	PermissionUpdateGroup Permission = "UpdateGroup"
	// PermissionDeleteGroup ...
	PermissionDeleteGroup Permission = "DeleteGroup"

	// PermissionCreateGroupAdmin ...
	PermissionCreateGroupAdmin Permission = "CreateGroupAdmin"
	// PermissionRetrieveGroupAdmin ...
	PermissionRetrieveGroupAdmin Permission = "RetrieveGroupAdmin"
	// PermissionUpdateGroupAdmin ...
	PermissionUpdateGroupAdmin Permission = "UpdateGroupAdmin"
	// PermissionDeleteGroupAdmin ...
	PermissionDeleteGroupAdmin Permission = "DeleteGroupAdmin"

	// PermissionCreateGroupADUser ...
	PermissionCreateGroupADUser Permission = "CreateGroupADUser"
	// PermissionRetrieveGroupADUser ...
	PermissionRetrieveGroupADUser Permission = "RetrieveGroupADUser"
	// PermissionUpdateGroupADUser ...
	PermissionUpdateGroupADUser Permission = "UpdateGroupADUser"
	// PermissionDeleteGroupADUser ...
	PermissionDeleteGroupADUser Permission = "DeleteGroupADUser"

	// PermissionRetrieveBTU ...
	PermissionRetrieveBTU Permission = "PermissionRetrieveBTU"
)

func addSessionPerms(ps *Permissions) {
	*ps = append(*ps, PermissionSession)
}
func addUserPerms(ps *Permissions) {
	*ps = append(*ps, PermissionCreateUser)
	*ps = append(*ps, PermissionRetrieveUser)
	*ps = append(*ps, PermissionUpdateUser)
	*ps = append(*ps, PermissionDeleteUser)
}
func addGroupPerms(ps *Permissions) {
	*ps = append(*ps, PermissionCreateGroup)
	*ps = append(*ps, PermissionRetrieveGroup)
	*ps = append(*ps, PermissionUpdateGroup)
	*ps = append(*ps, PermissionDeleteGroup)
}
func addGroupAdminPerms(ps *Permissions) {
	*ps = append(*ps, PermissionCreateGroupAdmin)
	*ps = append(*ps, PermissionRetrieveGroupAdmin)
	*ps = append(*ps, PermissionUpdateGroupAdmin)
	*ps = append(*ps, PermissionDeleteGroupAdmin)
}
func addGroupADUserPerms(ps *Permissions) {
	*ps = append(*ps, PermissionCreateGroupADUser)
	*ps = append(*ps, PermissionRetrieveGroupADUser)
	*ps = append(*ps, PermissionUpdateGroupADUser)
	*ps = append(*ps, PermissionDeleteGroupADUser)
}
func addBTUPerms(ps *Permissions) {
	*ps = append(*ps, PermissionRetrieveBTU)
}
func readOnly(ps *Permissions) {
	*ps = append(*ps, PermissionRetrieveGroup)
	*ps = append(*ps, PermissionRetrieveGroupADUser)
	// *ps = append(*ps, PermissionRetrieveUser)
	// *ps = append(*ps, PermissionRetrieveGroupAdmin)
}
func all(ps *Permissions) {
	addSessionPerms(ps)
	addUserPerms(ps)
	addGroupPerms(ps)
	addGroupAdminPerms(ps)
	addGroupADUserPerms(ps)
	addBTUPerms(ps)
}

// Rol ...
type Rol string

const (
	// RolAdmin ...
	RolAdmin Rol = "Admin"
	// RolUser ...
	RolUser Rol = "User"
	// RolReadOnly ...
	RolReadOnly Rol = "ReadOnly"
)

// GetPermissions ...
func GetPermissions(rol Rol) Permissions {
	ps := &Permissions{}
	switch rol {
	case RolAdmin:
		all(ps)
	case RolUser:
		addSessionPerms(ps)
		addGroupADUserPerms(ps)
		addBTUPerms(ps)
		*ps = append(*ps, PermissionRetrieveGroup)
	case RolReadOnly:
		readOnly(ps)
	}
	return *ps
}
