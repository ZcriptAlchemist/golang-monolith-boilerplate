package service

import (
	"context"
	"sqlc-demo/internal/pkg"
	"sqlc-demo/internal/pkg/db/sqlc"
)

// ------------------------------
// Creates an admin record in db
// ------------------------------

func CreateAdmin(adminParams sqlc.CreateAdminParams) error {
	err := pkg.DB.CreateAdmin(context.Background(), adminParams)
	if err != nil {
		return err
	}

	return nil
}

// --------------------------
// Fetches all admin details
// --------------------------
func FetchAdmins() ([]sqlc.Admin, error) {
	admins, err := pkg.DB.ListAdmins(context.Background())
	if err != nil {
		return nil, err
	}

	return admins, nil
}
