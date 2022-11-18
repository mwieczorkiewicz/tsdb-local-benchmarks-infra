package util

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime"
	"time"
)

func storeResult(fname, entry string) {
	f, err := os.OpenFile(fname, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(entry); err != nil {
		panic(err)
	}
}

func TimeTrack(start time.Time, dataSetName, dbName string) {
	elapsed := time.Since(start)

	// Skip this function, and fetch the PC and file for its parent.
	pc, _, _, _ := runtime.Caller(1)

	// Retrieve a function object this functions parent.
	funcObj := runtime.FuncForPC(pc)

	// Regex to extract just the function name (and not the module path).
	runtimeFunc := regexp.MustCompile(`^.*\.(.*)$`)
	name := runtimeFunc.ReplaceAllString(funcObj.Name(), "$1")

	log.Println(fmt.Sprintf("%s for %s took %s", name, dataSetName, elapsed))
	storeResult(dbName, fmt.Sprintf("%s,%s,%d\n", name, dataSetName, elapsed.Microseconds()))
}
