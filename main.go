package main

import (
	"flag"
	"fmt"
	"github.com/gookit/color"
)

func main() {
	var host, file string

	wp := &WordPress{}

	flag.StringVar(&host, "h", "", "ip")
	flag.StringVar(&file, "f", "", "filepath")
	flag.Parse()

	view := `
  ___  _  _  ____       ___    ___   ___   ___        ___    __   _    _    __ 
 / __)( \/ )( ___) ___ (__ \  / _ \ (__ \ (__ \  ___ (__ \  /  ) / )  / )  /  )
( (__  \  /  )__) (___) / _/ ( (_) ) / _/  / _/ (___) / _/   )( / _ \/ _ \  )( 
 \___)  \/  (____)     (____) \___/ (____)(____)     (____) (__)\___/\___/ (__)  by:Z92G`
	color.Cyan.Println(view)
	fmt.Println()

	if file != "" && host == "" {
		wp.batchScan(file)
	} else {
		wp.singleScan(host)
	}

}
