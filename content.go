package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"qiniupkg.com/http/httputil.v2"
	xbytes "qiniupkg.com/x/bytes.v7"
)

var (
	srcPoweredBy = []byte(` by &#160;<a href="http://www.doxygen.org/index.html">`)
	destPoweredBy = []byte(`. <a href="javascript:document.getElementsByName('x-refresh')[0].submit();" title="Refresh this page from the source.">Refresh now</a>. <a href="./?tools">Tools</a> for package owners. Powered by <a href="http://qiniu.com/"><img class="footer" src="http://assets.qiniu.com/qiniu-white-97x51.png" alt="qiniu"/></a> <a href="http://www.doxygen.org/index.html">`)
)

var (
	srcFooter = []byte(`<hr class="footer"/>`)
	destFooterTempl = `<hr class="footer"/><form name="x-refresh" method="POST" action="/-/refresh"><input type="hidden" name="path" value="%s"></form>`
)

func serveContent(
	w http.ResponseWriter, req *http.Request,
	pkg, name string, modtime time.Time, content io.ReadSeeker) {

	if !strings.HasSuffix(name, ".html") {
		http.ServeContent(w, req, name, modtime, content)
		return
	}

	b, err := ioutil.ReadAll(content)
	if err != nil {
		httputil.Error(w, err)
		return
	}

	var destFooter bytes.Buffer
	fmt.Fprintf(&destFooter, destFooterTempl, pkg)

	b = xbytes.Replace(b, srcFooter, destFooter.Bytes(), 1)
	b = xbytes.Replace(b, srcPoweredBy, destPoweredBy, 1)

	http.ServeContent(w, req, name, modtime, bytes.NewReader(b))
}

