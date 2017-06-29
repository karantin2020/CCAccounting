package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("journal", func() {
	BasePath("/journal")

	Action("upload", func() {
		Routing(POST("/"))
		Description("Upload journal")
		Response(OK, UploadJournalMedia)
		Response(BadRequest, UploadErrorMedia)
	})

	Action("show", func() {
		Routing(GET("/"))
		Description("Show all stored info")
		Response(OK, CollectionOf(CancBookV1))
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
	View("default", func() {
		Attribute("filename")
		Attribute("rows")
		Attribute("uploaded_at")
	})
})

var CancBookV1 = MediaType("application/vnd.journal.cancbookv1", func() {
	Description("CancBookV1 struct")
	TypeName("CancBookV1")
	Attributes(func() {
		Attribute("ID", String, "ID of field")
		Required("ID")
	})
	View("default", func() {
		Attribute("ID")
	})
})
