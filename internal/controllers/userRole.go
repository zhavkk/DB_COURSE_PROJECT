package controllers

import (
	"dbproject/internal/db"
	"dbproject/internal/utils"
	"net/http"
)

func GetRolesHandler(w http.ResponseWriter, r *http.Request) {
	roles, err := db.GetAllRoles()
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "failed to fetch roles")
		return
	}
	utils.ResponseWithJson(w, http.StatusOK, roles)
}
