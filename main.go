package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func abort(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func main() {
	var (
		H = flag.Bool("H", true, "First line are headers")
		h = flag.String("h", "", "Column to filter")
		v = flag.String("v", "", "Value")
	)
	flag.Parse()

	var in io.Reader

	switch flag.NArg() {
	case 0:
		in = os.Stdin
	case 1:
		fp, err := os.Open(flag.Arg(0))
		if err != nil {
			abort(err)
		}
		defer fp.Close()
		in = fp
	default:
		flag.Usage()
		os.Exit(2)
	}

	r := csv.NewReader(in)

	w := csv.NewWriter(os.Stdout)
	defer w.Flush()

	idx := -1

	if *H {
		headers, err := r.Read()
		if err != nil {
			abort(err)
		}
		w.Write(headers)

		if *h != "" {
			for i, hdr := range headers {
				if *h == hdr || strconv.Itoa(i+1) == *h {
					idx = i
				}
			}
		}
	} else {
		if *h != "" {
			var err error
			idx, err = strconv.Atoi(*h)
			if err != nil {
				abort(err)
			}
			idx = idx - 1
		}
	}

	if *v != "" && idx == -1 {
		abort(fmt.Errorf("cannot find reference to column %v", *h))
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			abort(err)
		}

		if *v != "" && record[idx] != *v {
			continue
		}

		w.Write(record)
	}
}
