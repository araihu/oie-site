package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/araihu/margo/site"
)

func main() {
	output := flag.String("output-dir", "dist", "diretório de saída")
	flag.Parse()
	if err := build(context.Background(), *output); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func build(ctx context.Context, output string) error {
	if _, err := os.Lstat(output); !errors.Is(err, fs.ErrNotExist) {
		if err == nil {
			return fmt.Errorf("o diretório de saída já existe: %s", output)
		}
		return err
	}
	result, err := site.BuildConfig(ctx, site.ConfigRequest{ConfigPath: "site.yaml"})
	if err != nil {
		return err
	}
	for i, artifact := range result.Artifacts {
		switch {
		case strings.HasPrefix(artifact.Path, "docs/") && strings.HasSuffix(artifact.Path, ".html"):
			result.Artifacts[i].Content = prepareDocs(artifact.Content)
		case artifact.Path == "margo-assets/search-interactions.js":
			script := searchLabels.Replace(string(artifact.Content))
			result.Artifacts[i].Content = []byte(strings.Replace(script, searchResultOrder, rankedSearchResultOrder, 1))
		case artifact.Path == "margo-assets/code-copy.js":
			result.Artifacts[i].Content = []byte(copyLabels.Replace(string(artifact.Content)))
		}
	}
	manifest, err := site.MarshalManifest(result, "")
	if err != nil {
		return err
	}
	result.Artifacts = append(result.Artifacts, site.Artifact{Path: site.ManifestPath, Content: manifest})

	parent := filepath.Dir(output)
	if err := os.MkdirAll(parent, 0o755); err != nil {
		return err
	}
	staging, err := os.MkdirTemp(parent, ".oie-site-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(staging)
	for _, artifact := range result.Artifacts {
		path := filepath.Join(staging, filepath.FromSlash(artifact.Path))
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			return err
		}
		if err := os.WriteFile(path, artifact.Content, 0o644); err != nil {
			return err
		}
	}
	if err := os.Chmod(staging, 0o755); err != nil {
		return err
	}
	return os.Rename(staging, output)
}
