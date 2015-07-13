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
	"qiniupkg.com/x/log.v7"

	xbytes "qiniupkg.com/x/bytes.v7"
)

// ---------------------------------------------------

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

// ---------------------------------------------------

const (
	srcMainPageDocStartTempl = `<div class="title">%s Documentation</div>`
	destMainPageDocStart     = `<div class="title">`
	destMainPageDocEnd       = `</div>`
)

var (
	srcStyleSheet  = []byte(`</head>`)
	destStyleSheet = []byte(`<link crossorigin="anonymous" href="http://qiniupkg.com/assets/github-20150713-1.css" media="all" rel="stylesheet"/>
<link crossorigin="anonymous" href="http://qiniupkg.com/assets/github-20150713-2.css" media="all" rel="stylesheet"/>
</head>`)
)

func makeMainPage(indexFile string, pkg string) (err error) {

	log.Info("makeMainPage", indexFile, "of", pkg)

	b, err := ioutil.ReadFile(indexFile)
	if err != nil {
		return
	}

	readmeFile := "https://" + pkg + "/blob/master/README.md"
	div, err := renderGithubMarkdown(readmeFile)
	if err != nil {
		return
	}

	var mainPageDocStart bytes.Buffer
	fmt.Fprintf(&mainPageDocStart, srcMainPageDocStartTempl, projectNameOf(pkg))
	pos := bytes.Index(b, mainPageDocStart.Bytes())
	if pos >= 0 {
		destMainPage := make([]byte, 0, len(div) + len(destMainPageDocStart) + len(destMainPageDocEnd))
		destMainPage = append(destMainPage, destMainPageDocStart...)
		destMainPage = append(destMainPage, div...)
		destMainPage = append(destMainPage, destMainPageDocEnd...)
		b = xbytes.ReplaceAt(b, pos, mainPageDocStart.Len(), destMainPage)
		b = xbytes.Replace(b, srcStyleSheet, destStyleSheet, 1)
		err = ioutil.WriteFile(indexFile, b, 0666)
	}
	return
}

func projectNameOf(pkg string) string {

	pos := strings.LastIndex(pkg, "/")
	if pos < 0 {
		return pkg
	}
	return pkg[pos+1:]
}

// ---------------------------------------------------

func renderGithubMarkdown(mdUrl string) (div []byte, err error) {

	resp, err := http.Get(mdUrl)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	pos := bytes.Index(b, markdownStart)
	if pos < 0 {
		return nil, ErrInvalidGithubMarkdown
	}

	from := pos + len(markdownStart)
	n := bytes.Index(b[from:], markdownEnd1)
	if n < 0 {
		return nil, ErrInvalidGithubMarkdown
	}

	return b[from:from+n+len(markdownEnd1)], nil
}

var (
	markdownStart = []byte(`<div id="readme" class="blob instapaper_body">`)
	markdownEnd1  = []byte(`</article>`)
)

// ---------------------------------------------------

