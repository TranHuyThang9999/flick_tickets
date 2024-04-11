package middlewares

// func (m *MiddleWare) CheckPermission(keys ...string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		roleName := c.GetString("role_id")
// 		if roleName == "root" {
// 			c.Next()
// 			return
// 		}
// 		role, err := m.roleUseCase.GetRoleById(c, roleName)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, errors.NewSystemError("role dont exist"))
// 			c.Abort()
// 			return
// 		}
// 		if role.Status != true {
// 			c.JSON(http.StatusForbidden, errors.NewCustomHttpError(errors.UserStatusFalse, "status role is false"))
// 			c.Abort()
// 			return
// 		}
// 		exist := false
// 		for _, key := range keys {
// 			if helpers.InArray(role.Abilities, key) {
// 				exist = true
// 				break
// 			}
// 		}
// 		if !exist {
// 			c.JSON(http.StatusForbidden, errors.NewCustomHttpError(errors.Forbidden, "permission denied"))
// 			c.Abort()
// 			return
// 		}
// 		c.Next()
// 	}
// }
