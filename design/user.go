package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("authentication", func() {

	BasePath("/")

	Action("login", func() {
		Routing(
			POST("/login"),
		)
		Description("Log in with account credentials")
		Payload(func() {
			Member("name", String, "Accound name or email field")
			Member("password", String, "Accound password field")
			Required("name", "password")
		})
		Response(OK)
		Response(Unauthorized)
		Response(BadRequest, ErrorMedia)
	})

	Action("signup", func() {
		Routing(
			POST("/signup"),
		)
		Description("Signup with account credentials")
		Payload(func() {
			Member("name", String, "Accound name field")
			Member("password", String, "Accound password field")
			Member("password_repeat", String, "Accound password repeat field")
			Member("email", String, "Accound email field")
			Required("name", "password", "password_repeat", "email")
		})
		Response(Created)
		Response(BadRequest, ErrorMedia)
	})

	Action("logout", func() {
		Routing(
			GET("/logout"),
		)
		Description("Log out with account name")
		Response(OK)
	})
})

var _ = Resource("user", func() {

	DefaultMedia(User)
	BasePath("/users")

	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("Retrieve all users")
		Response(OK, CollectionOf(User))
	})

	Action("show", func() {
		Routing(
			GET("/:userID"),
		)
		Description("Retrieve user with given id")
		Params(func() {
			Param("userID", String, "User ID", func() {
				MinLength(3)
			})
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Create new user")
		Payload(func() {
			Member("name", String, "Accound name field", func() {
				MinLength(3)
			})
			Member("password", String, "Accound password field", func() {
				MinLength(8)
			})
			Member("email", String, "Accound email field", func() {
				Format("email")
			})
			Required("name", "password", "email")
		})
		Response(Created, "/users/[a-zA-Z0-9]+")
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Routing(
			PUT("/:userID"),
		)
		Description("Change user name")
		Params(func() {
			Param("userID", String, "User ID")
		})
		Payload(func() {
			Member("name", String, "Accound name field", func() {
				MinLength(3)
			})
			Member("password", String, "Accound password field", func() {
				MinLength(8)
			})
			Member("email", String, "Accound email field", func() {
				Format("email")
			})
			Required("name")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Routing(
			DELETE("/:userID"),
		)
		Params(func() {
			Param("userID", String, "User ID")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})
