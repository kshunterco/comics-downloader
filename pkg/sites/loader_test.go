package sites

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/Girbons/comics-downloader/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestSiteLoaderMangatown(t *testing.T) {
	url := "https://www.mangatown.com/manga/naruto/v63/c693/"
	outputFolder := filepath.Dir(os.Args[0])

	options := &config.Options{
		All:          false,
		Last:         false,
		ImagesOnly:   false,
		Source:       "www.mangatown.com",
		URL:          url,
		Format:       "pdf",
		ImagesFormat: "png",
		OutputFolder: outputFolder,
	}

	collection, err := LoadComicFromSource(options)

	assert.Nil(t, err)
	assert.Equal(t, len(collection), 1)

	comic := collection[0]

	assert.Equal(t, "www.mangatown.com", comic.Source)
	assert.Equal(t, url, comic.URLSource)
	assert.Equal(t, "naruto", comic.Name)
	assert.Equal(t, "c693", comic.IssueNumber)
	assert.Equal(t, 20, len(comic.Links))
}

func TestCustomComicName(t *testing.T) {
	url := "https://www.mangatown.com/manga/naruto/v63/c693/"
	outputFolder := filepath.Dir(os.Args[0])

	options := &config.Options{
		All:             false,
		Last:            false,
		ImagesOnly:      false,
		Source:          "www.mangatown.com",
		URL:             url,
		Format:          "pdf",
		ImagesFormat:    "png",
		CustomComicName: "Naruto",
		OutputFolder:    outputFolder,
	}

	collection, err := LoadComicFromSource(options)

	assert.Nil(t, err)
	assert.Equal(t, len(collection), 1)

	comic := collection[0]

	assert.Equal(t, "www.mangatown.com", comic.Source)
	assert.Equal(t, url, comic.URLSource)
	assert.Equal(t, "Naruto", comic.Name)
	assert.Equal(t, "c693", comic.IssueNumber)
	assert.Equal(t, 20, len(comic.Links))
}

//func TestSiteLoaderMangareader(t *testing.T) {
//url := "https://www.mangareader.net/naruto/700"
//outputFolder := filepath.Dir(os.Args[0])

//options := &config.Options{
//All:          false,
//Last:         false,
//ImagesOnly:   false,
//Source:       "www.mangareader.net",
//Url:          url,
//Format:       "pdf",
//ImagesFormat: "png",
//OutputFolder: outputFolder,
//}

//collection, err := LoadComicFromSource(options)

//assert.Nil(t, err)
//assert.Equal(t, len(collection), 1)

//comic := collection[0]

//assert.Equal(t, "www.mangareader.net", comic.Source)
//assert.Equal(t, url, comic.URLSource)
//assert.Equal(t, "naruto", comic.Name)
//assert.Equal(t, "700", comic.IssueNumber)
//assert.Equal(t, 23, len(comic.Links))
//}

func TestSiteLoaderComicExtra(t *testing.T) {
	url := "https://www.comicextra.com/daredevil/chapter-600/full"
	outputFolder := filepath.Dir(os.Args[0])
	options := &config.Options{
		All:          false,
		Last:         false,
		ImagesOnly:   false,
		Source:       "www.comicextra.com",
		URL:          url,
		Format:       "pdf",
		ImagesFormat: "png",
		OutputFolder: outputFolder,
	}
	collection, err := LoadComicFromSource(options)

	assert.Nil(t, err)
	assert.Equal(t, len(collection), 1)

	comic := collection[0]

	assert.Equal(t, "www.comicextra.com", comic.Source)
	assert.Equal(t, url, comic.URLSource)
	assert.Equal(t, "daredevil", comic.Name)
	assert.Equal(t, "chapter-600", comic.IssueNumber)
	assert.Equal(t, 43, len(comic.Links))
}

func TestLoaderUnknownSource(t *testing.T) {
	url := "http://example.com"
	outputFolder := filepath.Dir(os.Args[0])

	options := &config.Options{
		All:          false,
		Last:         false,
		ImagesOnly:   false,
		Source:       "example.com",
		URL:          url,
		Format:       "pdf",
		ImagesFormat: "png",
		OutputFolder: outputFolder,
	}

	collection, err := LoadComicFromSource(options)

	if assert.NotNil(t, err) {
		assert.Equal(t, fmt.Errorf("source unknown"), err)
	}
	assert.Equal(t, len(collection), 0)
}

func TestIssuesRange(t *testing.T) {
	url := "https://www.comicextra.com/daredevil/chapter-600/full"
	outputFolder := filepath.Dir(os.Args[0])
	options := &config.Options{
		All:          true,
		Last:         false,
		ImagesOnly:   false,
		Source:       "www.comicextra.com",
		URL:          url,
		Format:       "pdf",
		ImagesFormat: "png",
		OutputFolder: outputFolder,
		IssuesRange:  "5-7",
	}
	collection, err := LoadComicFromSource(options)

	assert.Nil(t, err)
	assert.Equal(t, len(collection), 3)

	issues := make([]string, 0, len(collection))
	for _, c := range collection {
		issues = append(issues, c.IssueNumber)
	}

	assert.Contains(t, issues, "chapter-5")
	assert.Contains(t, issues, "chapter-6")
	assert.Contains(t, issues, "chapter-7")
}

func TestFloatIssuesRange(t *testing.T) {
	tt := []struct {
		input       string
		start       float64
		end         float64
		returnValue bool
	}{
		{"1", 1, 1, false},
		{"19", 20, 21, true},
		{"20", 20, 21, false},
		{"20.5", 20, 21, false},
		{"21", 20, 21, false},
		{"22", 20, 21, true},
	}

	for _, tc := range tt {
		t.Run(tc.input, func(t *testing.T) {
			assert.Equal(t, notInIssuesRange(tc.input, tc.start, tc.end), tc.returnValue)
		})
	}
}
