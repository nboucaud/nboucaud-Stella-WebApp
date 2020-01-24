// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package utils

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// ExtractTarGz takes in an io.Reader containing the bytes for a .tar.gz file and
// a destination string to extract to.
func ExtractTarGz(gzipStream io.Reader, dst string) error {
	if dst == "" {
		return errors.New("no destination path provided")
	}

	uncompressedStream, err := gzip.NewReader(gzipStream)
	if err != nil {
		return errors.Wrap(err, "failed to initialize gzip reader")
	}
	defer uncompressedStream.Close()

	tarReader := tar.NewReader(uncompressedStream)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return errors.Wrap(err, "failed to read next file from archive")
		}

		// Pre-emptively check type flag to avoid reporting a misleading error in
		// trying to sanitize the header name.
		switch header.Typeflag {
		case tar.TypeDir:
		case tar.TypeReg:
		default:
			return fmt.Errorf("unsupported type %v in %v", header.Typeflag, header.Name)
		}

		// filepath.HasPrefix is deprecated, so we just use strings.HasPrefix to ensure
		// the target path remains rooted at dst and has no `../` escaping outside.
		path := filepath.Join(dst, header.Name)
		if !strings.HasPrefix(path, dst) {
			return errors.Errorf("failed to sanitize path %s", header.Name)
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.Mkdir(path, 0744); err != nil && !os.IsExist(err) {
				return err
			}
		case tar.TypeReg:
			dir := filepath.Dir(path)

			if err := os.MkdirAll(dir, 0744); err != nil {
				return err
			}

			outFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
			defer outFile.Close()
			if _, err := io.Copy(outFile, tarReader); err != nil {
				return err
			}
		}
	}

	return nil
}
