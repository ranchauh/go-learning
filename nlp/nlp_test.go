package nlp

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"testing"
)

type Case struct {
	Text   string `yaml:"text"`
	Tokens []string `yaml:"tokens,flow"`
}

func TestTokenize(t *testing.T) {
	cases, err := getCases()
	if err != nil {
		t.Fatalf("Error loading test file %v", err)
	}
	for _, tc := range cases {
		t.Run(tc.Text, func(t *testing.T) {
			tokens := Tokenize(tc.Text)
			assert.Equal(t, tc.Tokens, tokens)
		})
	}
}

func getCases() ([]Case, error) {
	data, err := ioutil.ReadFile("/Users/randheer.chauhan/codebase/go-workshop/nlp/testdata/tokenize_cases.yml")
	if err != nil {
		if os.IsNotExist(err) {
			return nil, err
		}
		return nil, err
	}
	var cases []Case
	if err := yaml.Unmarshal(data, &cases); err != nil {
		return nil, err
	}
	return cases, nil
}
