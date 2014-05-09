package story_tester

import (
	"github.com/tildedave/go-hpack-impl/hpack"
	"github.com/stretchr/testify/assert"
	"testing"
	"io/ioutil"
	"encoding/json"
	"encoding/hex"
	"sort"
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
}

func testStoryFile(t *testing.T, filePath string) {
	storyBytes, err := ioutil.ReadFile(filePath)
	assert.Nil(t, err)

	var story Story

	json.Unmarshal(storyBytes, &story)
	context := hpack.NewEncodingContext()

	t.Log("--> ", story.Description)

	for _, c := range story.Cases {
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

		assert.Equal(t, expectedHeaders, hs.Headers)
	}
}

func TestStory(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}


	testStoryFile(t, "../story_tester/story_00.json")
}
