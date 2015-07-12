package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func handleBadge(w http.ResponseWriter, req *http.Request) {

	http.ServeContent(w, req, "badge.svg", toolsModtime, strings.NewReader(badgeSvg))
}

func handleTools(w http.ResponseWriter, req *http.Request, pkg string) {

	var buf bytes.Buffer
	fmt.Fprintf(&buf, toolsHtmlTempl, pkg, pkg, pkg, pkg, pkg, pkg, pkg)
	http.ServeContent(w, req, "tools.html", toolsModtime, bytes.NewReader(buf.Bytes()))
}

var (
	toolsModtime = time.Now() // Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
)

const (
	toolsHtmlTempl = `
<html>
<body>
  <h2>Tools for %s</h2>

  <h3>Badge</h3>

  <p><a href="/%s/"><img src="/%s/?status.svg" alt="doxygen.io"></a>

  <p>Use one of the snippets below to add a link to doxygen.io from your project
  website or README file:</a>

  <h5>HTML</h5>
  <input type="text" value='<a href="http://doxygen.io/%s/"><img src="http://doxygen.io/%s/?status.svg" alt="doxygen.io"></a>' size=100 class="click-select form-control">

  <h5>Markdown</h5>
  <input type="text" value="[![doxygen.io](http://doxygen.io/%s/?status.svg)](http://doxygen.io/%s/)" size=100 class="click-select form-control">

</body>
</html>
`
)

const (
	badgeSvg = `<svg xmlns="http://www.w3.org/2000/svg" width="79" height="20"><linearGradient id="b" x2="0" y2="100%"><stop offset="0" stop-color="#bbb" stop-opacity=".1"/><stop offset="1" stop-opacity=".1"/></linearGradient><mask id="a"><rect width="79" height="20" rx="3" fill="#fff"/></mask><g mask="url(#a)"><path fill="#555" d="M0 0h59v20H0z"/><path fill="#007ec6" d="M59 0h20v20H59z"/><path fill="url(#b)" d="M0 0h79v20H0z"/></g><g fill="#fff" text-anchor="middle" font-family="DejaVu Sans,Verdana,Geneva,sans-serif" font-size="11"><text x="29.5" y="15" fill="#010101" fill-opacity=".3">doxygen</text><text x="29.5" y="14">doxygen</text><text x="68" y="15" fill="#010101" fill-opacity=".3">io</text><text x="68" y="14">io</text></g></svg>`
)

