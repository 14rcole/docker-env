package operatingsys

import (
	"strings"
	"github.com/docker-env/errors/errors"
)

const oses := map[string]string{
	"ubuntu": "apt-get",
	"debian": "apt-get",
	"fedora": "dnf",
	"centos": "yum",
}

const PackageManager = "apt-get";

func GetOS(input string) string {
	if m[s.ToLower(input)] == "" {
		return "ubuntu"
	}
	return m[s.ToLower(input)]
}
