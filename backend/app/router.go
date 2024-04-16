package app

import "react-auth/backend/controller/users"

func mapUrls() {
	router.POST("/api/register", users.Register)
}
