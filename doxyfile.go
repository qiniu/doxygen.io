package main

import (
	"fmt"
	"io"
	"os"
)

type doxyfileConf struct {
	ProjectName  string
	OutputDir    string
	InputDir     string
	FilePatterns string // "*.java *.h *.cpp"
}

func genDoxyfileFile(file string, cfg *doxyfileConf) (err error) {

	f, err := os.Create(file)
	if err != nil {
		return
	}
	defer f.Close()

	return genDoxyfile(f, cfg)
}

func genDoxyfile(f io.Writer, cfg *doxyfileConf) (err error) {

	_, err = fmt.Fprintf(
		f, doxyfileTempl, cfg.ProjectName, cfg.OutputDir, cfg.InputDir, cfg.FilePatterns)
	return
}

const doxyfileTempl = `
# Doxyfile 1.2.12-20011209

#---------------------------------------------------------------------------
# General configuration options
#---------------------------------------------------------------------------
PROJECT_NAME           = %s
PROJECT_NUMBER         = 
OUTPUT_DIRECTORY       = %s
OUTPUT_LANGUAGE        = English
EXTRACT_ALL            = YES
EXTRACT_PRIVATE        = NO
EXTRACT_STATIC         = YES
EXTRACT_LOCAL_CLASSES  = YES
HIDE_UNDOC_MEMBERS     = NO
HIDE_UNDOC_CLASSES     = NO
BRIEF_MEMBER_DESC      = YES
REPEAT_BRIEF           = YES
ALWAYS_DETAILED_SEC    = NO
INLINE_INHERITED_MEMB  = NO
FULL_PATH_NAMES        = NO
STRIP_FROM_PATH        = 
INTERNAL_DOCS          = NO
STRIP_CODE_COMMENTS    = YES
CASE_SENSE_NAMES       = YES
SHORT_NAMES            = NO
HIDE_SCOPE_NAMES       = NO
VERBATIM_HEADERS       = YES
SHOW_INCLUDE_FILES     = YES
JAVADOC_AUTOBRIEF      = NO
INHERIT_DOCS           = YES
INLINE_INFO            = YES
SORT_MEMBER_DOCS       = YES
DISTRIBUTE_GROUP_DOC   = NO
TAB_SIZE               = 8
GENERATE_TODOLIST      = YES
GENERATE_TESTLIST      = YES
GENERATE_BUGLIST       = YES
ALIASES                = 
ENABLED_SECTIONS       = 
MAX_INITIALIZER_LINES  = 30
OPTIMIZE_OUTPUT_FOR_C  = NO
SHOW_USED_FILES        = YES
#---------------------------------------------------------------------------
# configuration options related to warning and progress messages
#---------------------------------------------------------------------------
QUIET                  = NO
WARNINGS               = YES
WARN_IF_UNDOCUMENTED   = YES
WARN_FORMAT            = 
WARN_LOGFILE           = 
#---------------------------------------------------------------------------
# configuration options related to the input files
#---------------------------------------------------------------------------
INPUT                  = %s
FILE_PATTERNS          = %s
RECURSIVE              = YES
EXCLUDE                = 
EXCLUDE_PATTERNS       = 
EXAMPLE_PATH           = 
EXAMPLE_PATTERNS       = 
EXAMPLE_RECURSIVE      = NO
IMAGE_PATH             = 
INPUT_FILTER           = 
FILTER_SOURCE_FILES    = NO
#---------------------------------------------------------------------------
# configuration options related to source browsing
#---------------------------------------------------------------------------
SOURCE_BROWSER         = YES
INLINE_SOURCES         = NO
REFERENCED_BY_RELATION = YES
REFERENCES_RELATION    = YES
#---------------------------------------------------------------------------
# configuration options related to the alphabetical class index
#---------------------------------------------------------------------------
ALPHABETICAL_INDEX     = NO
COLS_IN_ALPHA_INDEX    = 5
IGNORE_PREFIX          = 
#---------------------------------------------------------------------------
# configuration options related to the HTML output
#---------------------------------------------------------------------------
GENERATE_HTML          = YES
HTML_OUTPUT            = 
HTML_HEADER            = 
HTML_FOOTER            = 
HTML_STYLESHEET        = 
HTML_ALIGN_MEMBERS     = YES
GENERATE_HTMLHELP      = NO
GENERATE_CHI           = NO
BINARY_TOC             = NO
TOC_EXPAND             = NO
DISABLE_INDEX          = NO
ENUM_VALUES_PER_LINE   = 4
GENERATE_TREEVIEW      = NO
TREEVIEW_WIDTH         = 250
#---------------------------------------------------------------------------
# configuration options related to the LaTeX output
#---------------------------------------------------------------------------
GENERATE_LATEX         = NO
LATEX_OUTPUT           = 
COMPACT_LATEX          = NO
PAPER_TYPE             = a4wide
EXTRA_PACKAGES         = 
LATEX_HEADER           = 
PDF_HYPERLINKS         = NO
USE_PDFLATEX           = NO
LATEX_BATCHMODE        = NO
#---------------------------------------------------------------------------
# configuration options related to the RTF output
#---------------------------------------------------------------------------
GENERATE_RTF           = NO
RTF_OUTPUT             = 
COMPACT_RTF            = NO
RTF_HYPERLINKS         = NO
RTF_STYLESHEET_FILE    = 
RTF_EXTENSIONS_FILE    = 
#---------------------------------------------------------------------------
# configuration options related to the man page output
#---------------------------------------------------------------------------
GENERATE_MAN           = NO
MAN_OUTPUT             = 
MAN_EXTENSION          = 
MAN_LINKS              = NO
#---------------------------------------------------------------------------
# configuration options related to the XML output
#---------------------------------------------------------------------------
GENERATE_XML           = NO
#---------------------------------------------------------------------------
# configuration options for the AutoGen Definitions output
#---------------------------------------------------------------------------
GENERATE_AUTOGEN_DEF   = NO
#---------------------------------------------------------------------------
# Configuration options related to the preprocessor   
#---------------------------------------------------------------------------
ENABLE_PREPROCESSING   = YES
MACRO_EXPANSION        = NO
EXPAND_ONLY_PREDEF     = NO
SEARCH_INCLUDES        = YES
INCLUDE_PATH           = 
INCLUDE_FILE_PATTERNS  = 
PREDEFINED             = 
EXPAND_AS_DEFINED      = 
SKIP_FUNCTION_MACROS   = YES
#---------------------------------------------------------------------------
# Configuration::addtions related to external references   
#---------------------------------------------------------------------------
TAGFILES               = 
GENERATE_TAGFILE       = 
ALLEXTERNALS           = NO
PERL_PATH              = 
#---------------------------------------------------------------------------
# Configuration options related to the dot tool   
#---------------------------------------------------------------------------
CLASS_DIAGRAMS         = NO
HAVE_DOT               = NO
CLASS_GRAPH            = YES
COLLABORATION_GRAPH    = YES
TEMPLATE_RELATIONS     = YES
HIDE_UNDOC_RELATIONS   = YES
INCLUDE_GRAPH          = YES
INCLUDED_BY_GRAPH      = YES
GRAPHICAL_HIERARCHY    = YES
DOT_PATH               = 
DOTFILE_DIRS           = 
MAX_DOT_GRAPH_WIDTH    = 1280
MAX_DOT_GRAPH_HEIGHT   = 1024
GENERATE_LEGEND        = YES
DOT_CLEANUP            = YES
#---------------------------------------------------------------------------
# Configuration::addtions related to the search engine   
#---------------------------------------------------------------------------
SEARCHENGINE           = NO
CGI_NAME               = 
CGI_URL                = 
DOC_URL                = 
DOC_ABSPATH            = 
BIN_ABSPATH            = 
EXT_DOC_PATHS          = 
`

