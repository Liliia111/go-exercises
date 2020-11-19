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
		if val, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, val, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

type AddPathPayload struct {
	Path string
	Url  string
}

func AddPathHandler(pathsToUrls map[string]string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var a AddPathPayload

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&a)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		pathsToUrls[a.Path] = a.Url
		return
		//fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
type PathsToUrls struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pToU []PathsToUrls
	err := yaml.Unmarshal(yml, &pToU)
	if err != nil {
		return nil, err
	}
	var pathsToUrls = make(map[string]string, len(pToU))
	for _, pu := range pToU {
		pathsToUrls[pu.Path] = pu.Url
	}

	return MapHandler(pathsToUrls, fallback), err

}
