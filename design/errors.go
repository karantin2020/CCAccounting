package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"

	"github.com/karantin2020/CCAccounting/types"
)

var UploadErrorMedia = MediaType("application/vnd.upload.error", func() {
	Description("Upload error info")
	TypeName("UploadErrorMedia")
	Attributes(func() {
		Attribute("filename", String, "Uploaded file name")
		Attribute("uploaded_at", DateTime, "Upload timestamp")
		Attribute("description", String, "Error description")
		Required("filename", "rows", "uploaded_at")
	})
})
