package checker

import (
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

type PathType int

const (
	Unknown PathType = iota
	URL
	File
	Directory
	NonExistentLocalPath
)

func (pt PathType) String() string {
	switch pt {
	case URL:
		return "URL"
	case File:
		return "File"
	case Directory:
		return "Directory"
	case NonExistentLocalPath:
		return "NonExistentLocalPath"
	case Unknown:
		return "Unknown"
	default:
		return "Invalid PathType"
	}
}

func isURL(s string) bool {
	s = strings.TrimSpace(s)
	if s == "" {
		return false
	}

	u, err := url.Parse(s)
	if err != nil {
		return false
	}

	if u.Scheme != "" {
		if u.Scheme == "http" || u.Scheme == "https" || u.Scheme == "ftp" {
			return u.Host != ""
		} else if strings.HasPrefix(u.Scheme, "file") {
			return true
		}
	}

	if u.Host != "" && strings.Contains(u.Host, ":") && (strings.HasPrefix(u.Host, "localhost") || strings.Contains(u.Host, ".")) {
		return true
	}

	return false
}

func isLikelyLocalPathFormat(s string) bool {
	s = strings.TrimSpace(s)
	if s == "" {
		return false
	}

	if filepath.IsAbs(s) {
		return true
	}

	if strings.Contains(s, string(filepath.Separator)) || strings.Contains(s, "/") || strings.Contains(s, "\\") {
		return true
	}

	if strings.Contains(s, ".") && !strings.ContainsAny(s, ":/?#") {
		return true
	}

	return false
}

func CheckIdentityArgs(s string) PathType {
	s = strings.TrimSpace(s)
	if s == "" {
		return Unknown
	}

	if isURL(s) {
		return URL
	}

	info, err := os.Stat(s)
	if err == nil {
		if info.IsDir() {
			return Directory
		} else if info.Mode().IsRegular() {
			return File
		}
		return Unknown
	}

	if os.IsNotExist(err) {
		if isLikelyLocalPathFormat(s) {
			return NonExistentLocalPath
		}
	}

	return Unknown

}
