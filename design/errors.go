package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var UploadErrorMedia = MediaType("application/vnd.upload.error", func() {
	Description("Upload error info")
	TypeName("UploadErrorMedia")
	Attributes(func() {
		Attribute("filename", String, "Uploaded file name")
		Attribute("uploaded_at", DateTime, "Upload timestamp")
		Attribute("description", String, "Error description")
		Required("filename", "uploaded_at", "description")
	})
	View("default", func() {
		Attribute("filename")
		Attribute("uploaded_at")
		Attribute("description")
	})
})
