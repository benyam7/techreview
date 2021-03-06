package repository

import (
"github.com/hellyab/techreview/entities"
"github.com/hellyab/techreview/user"
"github.com/jinzhu/gorm"
)

// RoleGormRepo implements the menu.RoleRepository interface
type RoleGormRepo struct {
	conn *gorm.DB
}

// NewRoleGormRepo returns a new a new object of RoleGormRepo
func NewRoleGormRepo(db *gorm.DB) user.RoleRepository {
	return &RoleGormRepo{conn: db}
}

// Roles returns all user roles stored in the database
func (roleRepo *RoleGormRepo) Roles() ([]entities.Role, []error) {
	roles := []entities.Role{}
	errs := roleRepo.conn.Find(&roles).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return roles, errs
}

// Role retrieves a role by its id from the database
func (roleRepo *RoleGormRepo) Role(id string) (*entities.Role, []error) {
	role := entities.Role{}
	errs := roleRepo.conn.First(&role, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &role, errs
}

// RoleByName retrieves a role by its name from the database
func (roleRepo *RoleGormRepo) RoleByName(name string) (*entities.Role, []error) {
	role := entities.Role{}
	errs := roleRepo.conn.Find(&role, "name=?", name).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &role, errs
}

// UpdateRole updates a given user role in the database
func (roleRepo *RoleGormRepo) UpdateRole(role *entities.Role) (*entities.Role, []error) {
	r := role
	errs := roleRepo.conn.Save(r).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return r, errs
}

// DeleteRole deletes a given user role from the database
func (roleRepo *RoleGormRepo) DeleteRole(id string) (*entities.Role, []error) {
	r, errs := roleRepo.Role(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = roleRepo.conn.Delete(r, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return r, errs
}

// StoreRole stores a given user role in the database
func (roleRepo *RoleGormRepo) StoreRole(role *entities.Role) (*entities.Role, []error) {
	r := role
	errs := roleRepo.conn.Create(r).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return r, errs
}
