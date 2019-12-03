package urlshort

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}

		fallback.ServeHTTP(w, r)
	}
}

// CreateHandler will parse the provided YAML or JSON and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the file, then the
// fallback http.Handler will be called instead.
//
// The only errors that can be returned all related to having
// invalid YAML or JSON data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func CreateHandler(data []byte, dataType string, fallback http.Handler) (http.HandlerFunc, error) {
	pathURLs, err := parse(data, dataType)
	if err != nil {
		return nil, err
	}

	pathsToURLs := buildMap(pathURLs)

	return MapHandler(pathsToURLs, fallback), nil
}

type pathToURL struct {
	Path string `yaml:"path" json:"path"`
	URL  string `yaml:"url" json:"url"`
}

func parse(data []byte, dataType string) ([]pathToURL, error) {
	var pathURLs []pathToURL
	var err error
	switch dataType {
	case "yaml":
		err = yaml.Unmarshal(data, &pathURLs)
	case "json":
		err = json.Unmarshal(data, &pathURLs)
	}

	if err != nil {
		return nil, err
	}

	return pathURLs, nil
}

func buildMap(pathURLs []pathToURL) map[string]string {
	pathsToURLs := make(map[string]string)
	for _, pu := range pathURLs {
		pathsToURLs[pu.Path] = pu.URL
	}

	return pathsToURLs
}
