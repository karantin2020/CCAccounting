package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"

	"github.com/karantin2020/CCAccounting/types"
)

var _ = Resource("journal", func() {
	BasePath("/journal")

	Action("upload", func() {
		Routing(POST("/"))
		Description("Upload journal")
		Response(OK, UploadJournalMedia)
		Resource(BadRequest, UploadErrorMedia)
	})

	Action("show", func() {
		Routing(GET("/"))
		Description("Show all stored info")
		Response(OK, CollectionOf(types.CancBookV1))
		Response(NotFound)
	})
})

var UploadJournalMedia = MediaType("application/vnd.upload.journal", func() {
	Description("Upload info")
	TypeName("UploadJournalMedia")
	Attributes(func() {
		Attribute("filename", String, "Uploaded file name")
		Attribute("rows", Integer, "Number of inserted rows")
		Attribute("uploaded_at", DateTime, "Upload timestamp")
		Required("filename", "rows", "uploaded_at")
	})
})
