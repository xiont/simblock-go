package printfile

import (
	"path"
	"path/filepath"
)

/*
if you test testcases, you should change the path
if you run the test in simulator packages, you should use as follow
*/
//var outPath,_ = filepath.Glob("../output")

/*
if you run main func, you should use as follow
*/
var outPath, _ = filepath.Glob("output")

/*
if you don not want to output, you can change the NewFilePrinter to NewNilFilePrinter
when the node and height are to large, the size of the output.json can reach at more than 1GB
*/
var OUT_JSON_FILE = NewFilePrinter(path.Join(outPath[0], "output.json"))
var STATIC_JSON_FILE = NewFilePrinter(path.Join(outPath[0], "static.json"))
var PROPAGATION_FILE = NewFilePrinter(path.Join(outPath[0], "propagation.txt"))
var BLOCK_LIST = NewFilePrinter(path.Join(outPath[0], "blocklist.txt"))
