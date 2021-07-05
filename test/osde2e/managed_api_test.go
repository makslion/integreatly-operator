package osde2e

import (
	"bytes"
	"fmt"
	"github.com/integr8ly/integreatly-operator/test/metadata"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/integr8ly/integreatly-operator/test/common"
	. "github.com/onsi/ginkgo"
)

func teeOutput(f func()) string {

	var output bytes.Buffer

	stdoutReader, stdoutWriter, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	stderrReader, stderrWriter, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	originalStdout := os.Stdout
	os.Stdout = stdoutWriter
	defer func() {
		os.Stdout = originalStdout
	}()

	originalStderr := os.Stderr
	os.Stderr = stderrWriter
	defer func() {
		os.Stderr = originalStderr
	}()

	var wg sync.WaitGroup

	// this function will keep reading
	// from the piped stdout/stderr and write
	// to the original stdout/stderr
	t := func(r, w *os.File) {
		buf := make([]byte, 4096)
		for true {
			l, err := r.Read(buf)
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}
			if l == 0 {
				break
			}

			_, err = w.Write(buf[:l])
			if err != nil {
				panic(err)
			}

			_, err = output.Write(buf[:l])
			if err != nil {
				break
			}
		}

		wg.Done()
	}

	wg.Add(2)

	go t(stdoutReader, originalStdout)
	go t(stderrReader, originalStderr)

	f()

	err = stdoutWriter.Close()
	if err != nil {
		panic(err)
	}

	err = stderrWriter.Close()
	if err != nil {
		panic(err)
	}

	wg.Wait()

	err = stdoutReader.Close()
	if err != nil {
		panic(err)
	}

	err = stderrReader.Close()
	if err != nil {
		panic(err)
	}

	return output.String()
}

func writeOutputToFile(output string, filepath string) error {
	return ioutil.WriteFile(filepath, []byte(output), os.FileMode(0644))
}

var _ = Describe("integreatly", func() {

	var (
		restConfig = cfg
		t          = GinkgoT()
	)

	BeforeEach(func() {
		restConfig = cfg
		t = GinkgoT()
	})

	output := teeOutput(func() {
		// get all automated tests
		tests := []common.Tests{
			{
				Type:      fmt.Sprintf("%s HAPPY PATH", installType),
				TestCases: common.GetHappyPathTestCases(installType),
			},
		}

		for _, test := range tests {
			Context(test.Type, func() {
				for _, testCase := range test.TestCases {
					currentTest := testCase
					It(currentTest.Description, func() {
						testingContext, err := common.NewTestingContext(restConfig)
						if err != nil {
							t.Fatal("failed to create testing context", err)
						}
						currentTest.Test(t, testingContext)
					})
				}
			})
		}
	})

	if _, err := os.Stat(testResultsDirectory); !os.IsNotExist(err) {
		err := writeOutputToFile(output, filepath.Join(testResultsDirectory, testOutputFileName))
		if err != nil {
			fmt.Printf("error while writing the test output: %v", err)
			os.Exit(1)
		}

		err = metadata.Instance.WriteToJSON(filepath.Join(testResultsDirectory, addonMetadataName))
		if err != nil {
			fmt.Printf("error while writing metadata: %v", err)
			os.Exit(1)
		}
	}

})
