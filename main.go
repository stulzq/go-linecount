package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	mp "github.com/edsrzf/mmap-go"
	"log"
	"os"
	"time"
)

func main() {
	filePath := flag.String("f", "text.txt", "file path to read from")

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

	//mmap
	np, err := mp.Map(f, mp.RDONLY, 0)
	defer np.Unmap()
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(bytes.NewReader(np))

	startTime := time.Now()
	count := 0
	for s.Scan() {
		count++
	}

	endTime := time.Now()
	spent := endTime.Sub(startTime)
	fmt.Printf("Line count：%d \n", count)
	fmt.Printf("Time：%d sec\n", int(spent.Seconds()))

}
