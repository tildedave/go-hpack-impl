package story_tester

import (
	"encoding/json"
	"encoding/hex"
	"github.com/tildedave/go-hpack-impl/hpack"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"
	"fmt"
)

type Case struct {
	HeaderTableSize int `json:"header_table_size"`
	Seqno int
	Wire string
	Headers [](map[string]string)
}

type Story struct {
	Cases []Case
	Description string
	Draft int
}

const (
	Success = iota
	Failure
	Skipped
)

func testStoryFile(t *testing.T, filePath string) int {
	storyBytes, err := ioutil.ReadFile(filePath)
	assert.Nil(t, err)

	var story Story

	json.Unmarshal(storyBytes, &story)
	context := hpack.NewEncodingContext()
	context.HeaderTable.MaxSize = story.Cases[0].HeaderTableSize

	if story.Draft != 7 {
		fmt.Printf("Skipped %s (Draft %d)\n", filePath, story.Draft)
		return Skipped
	}

	t.Log(filePath, ":", story.Description)

	for _, c := range story.Cases {
		t.Log("Decoding", c.Wire)

		wire, err := hex.DecodeString(c.Wire)
		assert.Nil(t, err, "Unable to decode hex string")

		hs, err := context.Decode(string(wire))
		assert.Nil(t, err, "Unable to decode wire into HeaderSet")

		expectedHeaders := make([]hpack.HeaderField, 0)
		for _, h := range c.Headers {
			// only one in here usually
			for k, v := range h {
				header := hpack.HeaderField{k, v}
				expectedHeaders = append(expectedHeaders, header)
			}
		}

		expectedSet := hpack.HeaderSet{expectedHeaders}

		sort.Sort(expectedSet)
		sort.Sort(hs)

		if !reflect.DeepEqual(expectedHeaders, hs.Headers) {
			return Failure
		}
	}

	return Success
}

func TestStory(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

	attempts := 0
	fails := 0
	crashes := 0
	skipped := 0
	success := 0

	filepath.Walk("hpack-test-case", func (path string, info os.FileInfo, err error) error {
		defer func() {
			if r := recover(); r != nil {
				crashes += 1
				fails += 1
			}
		}()

		if strings.HasSuffix(path, ".json") {
			attempts += 1
			result := testStoryFile(t, path)
			if result == Failure {
				fmt.Println("Failed", path)
				fails += 1
			} else if result == Success {
				success += 1
			} else if result == Skipped {
				skipped += 1
			}
		}
		return nil
	})

	fmt.Printf("Attempted %d test cases\n", attempts)
	fmt.Printf("\tFailures: %d (%d crashed)\n", fails, crashes)
	fmt.Printf("\tSuccesses: %d\n", success)
	fmt.Printf("\tSkipped: %d\n", skipped)
}
