// Code generated by qtc from "basepage.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// This is a base page template. All the other template pages implement this interface.
//

//line ../internal/templates/basepage.qtpl:3
package templates

//line ../internal/templates/basepage.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line ../internal/templates/basepage.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line ../internal/templates/basepage.qtpl:4
type Page interface {
//line ../internal/templates/basepage.qtpl:4
	Title() string
//line ../internal/templates/basepage.qtpl:4
	StreamTitle(qw422016 *qt422016.Writer)
//line ../internal/templates/basepage.qtpl:4
	WriteTitle(qq422016 qtio422016.Writer)
//line ../internal/templates/basepage.qtpl:4
	Body() string
//line ../internal/templates/basepage.qtpl:4
	StreamBody(qw422016 *qt422016.Writer)
//line ../internal/templates/basepage.qtpl:4
	WriteBody(qq422016 qtio422016.Writer)
//line ../internal/templates/basepage.qtpl:4
}

// Page prints a page implementing Page interface.

//line ../internal/templates/basepage.qtpl:12
func StreamPageTemplate(qw422016 *qt422016.Writer, p Page) {
//line ../internal/templates/basepage.qtpl:12
	qw422016.N().S(`
<html>
	<head>
		<title>`)
//line ../internal/templates/basepage.qtpl:15
	p.StreamTitle(qw422016)
//line ../internal/templates/basepage.qtpl:15
	qw422016.N().S(`</title>
	</head>
	<body>
	<script src="https://unpkg.com/leaflet@1.9.3/dist/leaflet.js" integrity="sha256-WBkoXOwTeyKclOHuWtc+i2uENFpDZ9YPdf5Hf+D7ewM=" crossorigin=""></script>
		`)
//line ../internal/templates/basepage.qtpl:19
	p.StreamBody(qw422016)
//line ../internal/templates/basepage.qtpl:19
	qw422016.N().S(`
	</body>
</html>
`)
//line ../internal/templates/basepage.qtpl:22
}

//line ../internal/templates/basepage.qtpl:22
func WritePageTemplate(qq422016 qtio422016.Writer, p Page) {
//line ../internal/templates/basepage.qtpl:22
	qw422016 := qt422016.AcquireWriter(qq422016)
//line ../internal/templates/basepage.qtpl:22
	StreamPageTemplate(qw422016, p)
//line ../internal/templates/basepage.qtpl:22
	qt422016.ReleaseWriter(qw422016)
//line ../internal/templates/basepage.qtpl:22
}

//line ../internal/templates/basepage.qtpl:22
func PageTemplate(p Page) string {
//line ../internal/templates/basepage.qtpl:22
	qb422016 := qt422016.AcquireByteBuffer()
//line ../internal/templates/basepage.qtpl:22
	WritePageTemplate(qb422016, p)
//line ../internal/templates/basepage.qtpl:22
	qs422016 := string(qb422016.B)
//line ../internal/templates/basepage.qtpl:22
	qt422016.ReleaseByteBuffer(qb422016)
//line ../internal/templates/basepage.qtpl:22
	return qs422016
//line ../internal/templates/basepage.qtpl:22
}

// Base page implementation. Other pages may inherit from it if they need
// overriding only certain Page methods

//line ../internal/templates/basepage.qtpl:27
type BasePage struct{}

//line ../internal/templates/basepage.qtpl:28
func (p *BasePage) StreamTitle(qw422016 *qt422016.Writer) {
//line ../internal/templates/basepage.qtpl:28
	qw422016.N().S(`This is a base title`)
//line ../internal/templates/basepage.qtpl:28
}

//line ../internal/templates/basepage.qtpl:28
func (p *BasePage) WriteTitle(qq422016 qtio422016.Writer) {
//line ../internal/templates/basepage.qtpl:28
	qw422016 := qt422016.AcquireWriter(qq422016)
//line ../internal/templates/basepage.qtpl:28
	p.StreamTitle(qw422016)
//line ../internal/templates/basepage.qtpl:28
	qt422016.ReleaseWriter(qw422016)
//line ../internal/templates/basepage.qtpl:28
}

//line ../internal/templates/basepage.qtpl:28
func (p *BasePage) Title() string {
//line ../internal/templates/basepage.qtpl:28
	qb422016 := qt422016.AcquireByteBuffer()
//line ../internal/templates/basepage.qtpl:28
	p.WriteTitle(qb422016)
//line ../internal/templates/basepage.qtpl:28
	qs422016 := string(qb422016.B)
//line ../internal/templates/basepage.qtpl:28
	qt422016.ReleaseByteBuffer(qb422016)
//line ../internal/templates/basepage.qtpl:28
	return qs422016
//line ../internal/templates/basepage.qtpl:28
}

//line ../internal/templates/basepage.qtpl:29
func (p *BasePage) StreamBody(qw422016 *qt422016.Writer) {
//line ../internal/templates/basepage.qtpl:29
	qw422016.N().S(`This is a base body`)
//line ../internal/templates/basepage.qtpl:29
}

//line ../internal/templates/basepage.qtpl:29
func (p *BasePage) WriteBody(qq422016 qtio422016.Writer) {
//line ../internal/templates/basepage.qtpl:29
	qw422016 := qt422016.AcquireWriter(qq422016)
//line ../internal/templates/basepage.qtpl:29
	p.StreamBody(qw422016)
//line ../internal/templates/basepage.qtpl:29
	qt422016.ReleaseWriter(qw422016)
//line ../internal/templates/basepage.qtpl:29
}

//line ../internal/templates/basepage.qtpl:29
func (p *BasePage) Body() string {
//line ../internal/templates/basepage.qtpl:29
	qb422016 := qt422016.AcquireByteBuffer()
//line ../internal/templates/basepage.qtpl:29
	p.WriteBody(qb422016)
//line ../internal/templates/basepage.qtpl:29
	qs422016 := string(qb422016.B)
//line ../internal/templates/basepage.qtpl:29
	qt422016.ReleaseByteBuffer(qb422016)
//line ../internal/templates/basepage.qtpl:29
	return qs422016
//line ../internal/templates/basepage.qtpl:29
}
