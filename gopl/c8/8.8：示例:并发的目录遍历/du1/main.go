package main

import (
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

// 包示例：path/filepath
// 拼接路径
// path := filepath.Join("dir", "subdir", "file.txt")
// fmt.Println(path) // 输出: dir/subdir/file.txt

// // 分割路径
// dir, file := filepath.Split(path)
// fmt.Println("目录:", dir) // 输出: dir/subdir/
// fmt.Println("文件名:", file) // 输出: file.txt

// // 获取目录部分
// dir = filepath.Dir(path)
// fmt.Println("目录:", dir) // 输出: dir/subdir

// // 获取文件名部分
// file = filepath.Base(path)
// fmt.Println("文件名:", file) // 输出: file.txt

// // 获取文件扩展名
// ext := filepath.Ext(file)
// fmt.Println("扩展名:", ext) // 输出: .txt

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse the file tree.
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	// Print the results.
	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}

	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []fs.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}

	return entries
}
