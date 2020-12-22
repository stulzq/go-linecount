package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	mp "github.com/edsrzf/mmap-go"
	"log"
	"os"
	"time"
)

func main() {
	filePath := flag.String("f", "text.txt", "file path to read from")
	flag.Parse()
	//Open file
	f, err := os.OpenFile(*filePath, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	fi, _ := f.Stat()
	bar := pb.Full.Start64(fi.Size())

	//mmap
	np, err := mp.Map(f, mp.RDONLY, 0)
	defer np.Unmap()
	if err != nil {
		panic(err)
	}

	reader := bar.NewProxyReader(bytes.NewReader(np))
	s := bufio.NewScanner(reader)

	startTime := time.Now()
	count := 0
	for s.Scan() {
		count++
	}

	endTime := time.Now()
	bar.Finish()
	spent := endTime.Sub(startTime)
	fmt.Printf("Line count：%d \n", count)
	fmt.Printf("Time：%d sec\n", int(spent.Seconds()))

}
