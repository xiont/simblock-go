/*
   Copyright 2020 LittleBear(1018589158@qq.com)

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
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
