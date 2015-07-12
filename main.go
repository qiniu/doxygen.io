package main

import (
	"bytes"
	"errors"
	"flag"
	"hash/crc32"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"qiniupkg.com/http/httputil.v2"
	"qiniupkg.com/x/log.v7"
)

const (
	mutexCount = 9973
)

var (
	ErrUnmatchedInodeType = errors.New("unmatched inode type(file or dir)")
	ErrRefreshWithoutPath = httputil.NewError(400, "refresh without path")
	ErrInvalidPkgPath     = httputil.NewError(400, "invalid package path")
)

var (
	doxygenApp string

	dataRootDir string
	srcRootDir  string
	tmpRootDir  string

	refreshRootDir string

	mutexs [mutexCount]sync.Mutex
)

func handleHome(w http.ResponseWriter, req *http.Request) {

}

func handleUnknown(w http.ResponseWriter, req *http.Request) {

}

func handleRefresh(w http.ResponseWriter, req *http.Request) {

	pkg := req.PostFormValue("path")
	if pkg == "" {
		httputil.Error(w, ErrRefreshWithoutPath)
		return
	}

	log.Info("Refresh", pkg)

	err := refresh(pkg)
	if err != nil {
		httputil.Error(w, err)
		return
	}

	http.Redirect(w, req, "/" + pkg + "/", 301)
}

func refresh(pkg string) (err error) {

	dataDir := dataRootDir + pkg
	indexFile := dataDir + "/html/index.html"
	if isRefreshed(indexFile) {
		return nil
	}

	if strings.Index(pkg, "..") >= 0 {
		return ErrInvalidPkgPath
	}

	parts := strings.SplitN(pkg, "/", 4)
	if len(parts) != 3 {
		return ErrInvalidPkgPath
	}

	refreshDir := refreshRootDir + pkg
	refreshHtmlDir := refreshDir + "/html/"
	os.RemoveAll(refreshDir)
	err = genDoc(parts, pkg, refreshDir, refreshHtmlDir)
	if err != nil {
		return
	}

	os.RemoveAll(dataDir)
	return os.Rename(refreshDir, dataDir)
}

func isRefreshed(indexFile string) bool {

	fi, err := os.Stat(indexFile)
	if err != nil {
		return false
	}

	return time.Now().Sub(fi.ModTime()) < 10*time.Second
}

func handleMain(w http.ResponseWriter, req *http.Request) {

	path := req.URL.Path

	if path == "/" {
		handleHome(w, req)
		return
	}

	log.Info("View", path)

	if strings.Index(path, "..") >= 0 {
		handleUnknown(w, req)
		return
	}

	parts := strings.SplitN(path[1:], "/", 4)
	if parts[0] != "github.com" || len(parts) < 3 {
		handleUnknown(w, req)
		return
	}

	pkg := strings.Join(parts[:3], "/")
	dataDir := dataRootDir + pkg
	htmlDir := dataDir + "/html/"
	err := isEntryExists(htmlDir, true)
	if err != nil {
		err = genDoc(parts, pkg, dataDir, htmlDir)
		if err != nil {
			httputil.Error(w, err)
			return
		}
	}

	if len(parts) > 3 {
		file := htmlDir + parts[3]
		if strings.HasSuffix(file, "/") {
			file += "index.html"
		}
		f, err := os.Open(file)
		if err != nil {
			httputil.Error(w, err)
			return
		}
		defer f.Close()
		fi, err := f.Stat()
		if err != nil {
			httputil.Error(w, err)
			return
		}
		serveContent(w, req, pkg, fi.Name(), fi.ModTime(), f)
	} else {
		http.Redirect(w, req, path + "/", 301)
	}
}

func genDoc(parts []string, pkg, dataDir, htmlDir string) (err error) {

	srcDir := srcRootDir + pkg
	repo := "https://github.com/" + parts[1] + "/" + parts[2] + ".git"

	crc := crc32.ChecksumIEEE([]byte(pkg))
	mutex := &mutexs[crc % mutexCount]
	mutex.Lock()
	defer mutex.Unlock()

	err2 := isEntryExists(htmlDir, true)
	if err2 != nil {
		err = cloneRepo(srcDir, repo)
		if err != nil {
			return
		}

		err = os.MkdirAll(dataDir, 0755)
		if err != nil {
			return
		}

		doxyfile := tmpRootDir + "github.com!" + parts[1] + "!" + parts[2] + ".doxyfile"
		err = genDoxyfileFile(doxyfile, &doxyfileConf{
			ProjectName:  parts[2],
			OutputDir:    dataDir,
			InputDir:     srcDir,
			FilePatterns: "*.java *.h *.hpp *.hxx *.py *.php *.rb *.cs *.js *.scala *.go *.lua *.asp",
		})
		if err != nil {
			return
		}

		err = runCmd(doxygenApp, doxyfile)
		if err != nil {
			return
		}
	}
	return
}

func cloneRepo(srcDir string, repo string) (err error) {

	os.RemoveAll(srcDir)
	err = os.MkdirAll(srcDir, 0755)
	if err != nil {
		return
	}

	return runCmd("git", "clone", "--depth=50", repo, srcDir)
}

func runCmd(command string, args ...string) (err error) {

	cmd := exec.Command(command, args...)

	var out bytes.Buffer
	cmd.Stderr = &out

	err = cmd.Run()
	if err == nil {
		return
	}

	emsg := out.String()
	if emsg != "" {
		return errors.New(emsg)
	}
	return err
}

func isEntryExists(entryPath string, isDir bool) (err error) {

	fi, err := os.Stat(entryPath)
	if err != nil {
		return
	}

	if fi.IsDir() != isDir {
		err = ErrUnmatchedInodeType
		return
	}
	return nil
}

var (
	bindHost = flag.String("http", ":8888", "address that doxygen.io server listen")
)

func main() {

	flag.Parse()

	rootDir := os.Getenv("HOME") + "/.doxygen.io/"
	doxygenApp = os.Getenv("DOXYGEN")
	if doxygenApp == "" {
		doxygenApp = "doxygen"
	}

	dataRootDir = rootDir + "data/"
	refreshRootDir = rootDir + "refresh/"
	srcRootDir = rootDir + "src/"
	tmpRootDir = rootDir + "tmp/"
	os.MkdirAll(tmpRootDir, 0755)

	http.HandleFunc("/-/refresh", handleRefresh)
	http.HandleFunc("/", handleMain)
	err := http.ListenAndServe(*bindHost, nil)
	if err != nil {
		log.Fatal(err)
	}
}

