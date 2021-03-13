# go-junit4
A golang lib for building junit4 report

## Specification

Follow the JUint report format document listed in the link below: https://llg.cubic.org/docs/junit/

## Usage

Directly refer to the sample in the `main.go`:

```go
package main

import (
	"flag"
	"log"
	"time"
)
import (
	"github.com/szlabs/go-junit4/pkg/junit"
)

func main() {
	// Lib sample usage.

	var filePath string

	flag.StringVar(&filePath, "file", "", "specify the file path of the generated junit report")
	flag.Parse()

	tss := junit.NewSuites("junit lib compliance checking")
	tss.SetTests(3)
	tss.SetErrors(1)
	tss.SetFailures(1)
	tss.SetTime(800)

	ts := junit.NewSuite(1, "szlabs-suite-1")
	ts.SetTime(800)
	ts.SetFailures(1)
	ts.SetErrors(1)
	ts.SetTests(3)
	ts.SetHostname("szlabs-hostname")
	ts.SetPackage("szlabs-package")
	ts.SetSkipped(0)
	ts.SetDisabled(0)
	ts.SetSystemErr("szlabs system error")
	ts.SetSystemOut("szlabs system out")
	ts.SetTimestamp(time.Now().UTC())
	ts.AddProperty(&junit.Property{Name: "engine", Value: "szlabs/go-junit4"})
	tss.Append(ts.TestSuite())

	tc1 := junit.NewCase("case-1")
	tc1.SetTime(500)
	tc1.SetSystemOut("case1 system out")
	tc1.SetSystemErr("case1 system err")
	tc1.SetStatus("passed")
	tc1.SetAssertions(2)
	tc1.SetClassName("case-classname-1")
	ts.Append(tc1.TestCase())

	tc2 := junit.NewCase("case-2")
	tc2.SetTime(300)
	tc2.SetSystemOut("case2 system out")
	tc2.SetSystemErr("case2 system err")
	tc2.SetStatus("error")
	tc2.SetAssertions(2)
	tc2.SetClassName("case-classname-1")
	tc2.SetError("my error", "error", "500 error")
	tc2.SetFailure("my failure", "not equal", "expected the value greater than 100 but not")
	ts.Append(tc2.TestCase())

	// Export to file
	if filePath != "" {
		if err := tss.Complete().XML2File(filePath); err != nil {
			log.Printf("[ERROR]: err\n")
		}
	} else {
		// Export to data
		data, err := tss.Complete().XML()
		if err != nil {
			log.Printf("[ERROR]: err\n")
			return
		}

		log.Printf("[INFO]: junit report: %s\n", data)
	}
}
```

## Contact author

@[steven-zou](https://github.com/steven-zou)