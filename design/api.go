package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("CCAccounting", func() {
	Description("Service to store and query criminal cases info")
	Version("1.0") // API version being described
	BasePath("/api/v1")
	Host("localhost:3451")
	Scheme("http")
})
