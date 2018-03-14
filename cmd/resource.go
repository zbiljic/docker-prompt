package cmd

import (
	"log"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	prompt "github.com/c-bata/go-prompt"
	"github.com/fsouza/go-dockerclient"
)

const thresholdFetchInterval = 10 * time.Second

func init() {
	lastFetchedAt = new(sync.Map)
}

/* LastFetchedAt */

var (
	lastFetchedAt *sync.Map
)

func shouldFetch(key string) bool {
	v, ok := lastFetchedAt.Load(key)
	if !ok {
		log.Printf("[WARN] Not found %s in lastFetchedAt", key)
		return true
	}
	t, ok := v.(time.Time)
	if !ok {
		return true
	}
	return time.Since(t) > thresholdFetchInterval
}

func updateLastFetchedAt(key string) {
	lastFetchedAt.Store(key, time.Now())
}

/* Images list */

var (
	imagesList atomic.Value
)

func fetchImagesList() {
	key := "images"
	if !shouldFetch(key) {
		return
	}
	updateLastFetchedAt(key)

	opts := docker.ListImagesOptions{}
	images, _ := getClient().ListImages(opts)

	imagesList.Store(images)
	return
}

func getImagesSuggestions() []prompt.Suggest {
	go fetchImagesList()
	images, ok := imagesList.Load().([]docker.APIImages)
	if !ok || len(images) == 0 {
		return []prompt.Suggest{}
	}
	repositories := make(map[string]struct{})
	for _, image := range images {
		for _, repoTag := range image.RepoTags {
			repoTagParts := strings.Split(repoTag, ":")
			repositories[repoTagParts[0]] = struct{}{}
		}
	}
	s := []prompt.Suggest{}
	for k := range repositories {
		s = append(s, prompt.Suggest{
			Text: k,
		})
	}
	return s
}
