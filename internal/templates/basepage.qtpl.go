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
	 <link rel="stylesheet" href="https://unpkg.com/leaflet@1.9.4/dist/leaflet.css"
     integrity="sha256-p4NxAoJBhIIN+hmNHrzRCf9tD/miZyoHS5obTRR9BMY="
     crossorigin=""/>
 <!-- Make sure you put this AFTER Leaflet's CSS -->
 <script src="https://unpkg.com/leaflet@1.9.4/dist/leaflet.js"
     integrity="sha256-20nQCchB9co0qIjJZRGuk2/Z9VM+kNiyxNV1lvTlZBo="
     crossorigin=""></script>
		`)
//line ../internal/templates/basepage.qtpl:25
	p.StreamBody(qw422016)
//line ../internal/templates/basepage.qtpl:25
	qw422016.N().S(`
	</body>
</html>
`)
//line ../internal/templates/basepage.qtpl:28
}

//line ../internal/templates/basepage.qtpl:28
func WritePageTemplate(qq422016 qtio422016.Writer, p Page) {
//line ../internal/templates/basepage.qtpl:28
	qw422016 := qt422016.AcquireWriter(qq422016)
//line ../internal/templates/basepage.qtpl:28
	StreamPageTemplate(qw422016, p)
//line ../internal/templates/basepage.qtpl:28
	qt422016.ReleaseWriter(qw422016)
//line ../internal/templates/basepage.qtpl:28
}

//line ../internal/templates/basepage.qtpl:28
func PageTemplate(p Page) string {
//line ../internal/templates/basepage.qtpl:28
	qb422016 := qt422016.AcquireByteBuffer()
//line ../internal/templates/basepage.qtpl:28
	WritePageTemplate(qb422016, p)
//line ../internal/templates/basepage.qtpl:28
	qs422016 := string(qb422016.B)
//line ../internal/templates/basepage.qtpl:28
	qt422016.ReleaseByteBuffer(qb422016)
//line ../internal/templates/basepage.qtpl:28
	return qs422016
//line ../internal/templates/basepage.qtpl:28
}

// Base page implementation. Other pages may inherit from it if they need
// overriding only certain Page methods

//line ../internal/templates/basepage.qtpl:33
type BasePage struct{}

//line ../internal/templates/basepage.qtpl:34
func (p *BasePage) StreamTitle(qw422016 *qt422016.Writer) {
//line ../internal/templates/basepage.qtpl:34
	qw422016.N().S(`This is a base title`)
//line ../internal/templates/basepage.qtpl:34
}

//line ../internal/templates/basepage.qtpl:34
func (p *BasePage) WriteTitle(qq422016 qtio422016.Writer) {
//line ../internal/templates/basepage.qtpl:34
	qw422016 := qt422016.AcquireWriter(qq422016)
//line ../internal/templates/basepage.qtpl:34
	p.StreamTitle(qw422016)
//line ../internal/templates/basepage.qtpl:34
	qt422016.ReleaseWriter(qw422016)
//line ../internal/templates/basepage.qtpl:34
}

//line ../internal/templates/basepage.qtpl:34
func (p *BasePage) Title() string {
//line ../internal/templates/basepage.qtpl:34
	qb422016 := qt422016.AcquireByteBuffer()
//line ../internal/templates/basepage.qtpl:34
	p.WriteTitle(qb422016)
//line ../internal/templates/basepage.qtpl:34
	qs422016 := string(qb422016.B)
//line ../internal/templates/basepage.qtpl:34
	qt422016.ReleaseByteBuffer(qb422016)
//line ../internal/templates/basepage.qtpl:34
	return qs422016
//line ../internal/templates/basepage.qtpl:34
}

//line ../internal/templates/basepage.qtpl:35
func (p *BasePage) StreamBody(qw422016 *qt422016.Writer) {
//line ../internal/templates/basepage.qtpl:35
	qw422016.N().S(`This is a base body`)
//line ../internal/templates/basepage.qtpl:35
}

//line ../internal/templates/basepage.qtpl:35
func (p *BasePage) WriteBody(qq422016 qtio422016.Writer) {
//line ../internal/templates/basepage.qtpl:35
	qw422016 := qt422016.AcquireWriter(qq422016)
//line ../internal/templates/basepage.qtpl:35
	p.StreamBody(qw422016)
//line ../internal/templates/basepage.qtpl:35
	qt422016.ReleaseWriter(qw422016)
//line ../internal/templates/basepage.qtpl:35
}

//line ../internal/templates/basepage.qtpl:35
func (p *BasePage) Body() string {
//line ../internal/templates/basepage.qtpl:35
	qb422016 := qt422016.AcquireByteBuffer()
//line ../internal/templates/basepage.qtpl:35
	p.WriteBody(qb422016)
//line ../internal/templates/basepage.qtpl:35
	qs422016 := string(qb422016.B)
//line ../internal/templates/basepage.qtpl:35
	qt422016.ReleaseByteBuffer(qb422016)
//line ../internal/templates/basepage.qtpl:35
	return qs422016
//line ../internal/templates/basepage.qtpl:35
}
