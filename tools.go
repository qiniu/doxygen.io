package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

func handleBadge(w http.ResponseWriter, req *http.Request) {

	http.ServeFile(w, req, "./static/badge.svg")
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


