package thumbnail

import (
	"log"
	"os"
	"sync"
)

func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := ImageFile(f); err != nil {
			log.Panicln(err)
		}
	}
}

func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		go ImageFile(f)
	}
}

func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			ImageFile(f)
			ch <- struct{}{}
		}(f)
	}

	for range filenames {
		<-ch
	}
}

func makeThumbnails4(filenames []string) error {
	errors := make(chan error)

	for _, f := range filenames {
		go func(f string) {
			_, err := ImageFile(f)
			errors <- err
		}(f)
	}

	for range filenames {
		if err := <-errors; err != nil {
			return err
		}
	}

	return nil
}

func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}

	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = ImageFile(f)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			// 就算第一个image就抛出错误，channel也有足够的空间容纳余下的结果而不阻塞
			err = it.err
			return
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}

	return
}

func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)

	var wg sync.WaitGroup
	for f := range filenames {
		wg.Add(1)

		// worker
		go func(f string) {
			defer wg.Done()
			thumb, err := ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}

			// os.Stat() returns a FileInfo describing the named file.
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)
	}

	// closer
	// 让range sizes阻塞，而不是主goroutine
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}

	return total
}
