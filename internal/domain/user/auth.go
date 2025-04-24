package user

func containsRequiredRoles(userPerms []int, required int) bool {
	for _, perm := range userPerms {
		if perm == required {
			return true
		}
	}
	return false
}
