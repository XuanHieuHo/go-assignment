package importers

import (
	"context"
	"strings"

	"github.com/XuanHieuHo/go-assignment/requests"
	userService "github.com/XuanHieuHo/go-assignment/services/user"
)




type UserCSVSaver struct {
	Request requests.CreateUserRequest
	Service userService.UserService
}

// ParseFromRow implements CSVModelImporter.
func (u *UserCSVSaver) ParseFromRow(row []string, headerCheck map[string]int) error {
	get := func(key string) string {
		return strings.TrimSpace(row[headerCheck[key]])
	}
	u.Request.Name = get("name")
	u.Request.Email = get("email")
	return nil
}

// Save implements CSVModelImporter.
func (u *UserCSVSaver) Save(ctx context.Context) error {
	_, err := u.Service.CreateUser(ctx, u.Request)
	return err
}

func NewUserCSVSaver(request requests.CreateUserRequest, service userService.UserService) CSVImporter {
	return &UserCSVSaver{Request: request, Service: service}
}
