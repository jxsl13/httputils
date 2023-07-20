package pickle

const (
	ContentTypeBinary = "application/octet-stream"
	ContentTypeText   = "text/plain"
	ContentTypeJSON   = "application/json"
	ContentTypeXML    = "application/xml"
)

var ContentTypes = []string{
	ContentTypeBinary,
	ContentTypeText,
	ContentTypeJSON,
	ContentTypeXML,
}
