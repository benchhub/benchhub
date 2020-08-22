// Package plural converts word between plural and singular.
// FIXME: this is a hack for tqbuilder, should got replaced by other packages eventually.
// TODO: move the package to gommon when it actually works, or move it even when it is not working ...
// TODO: ref https://github.com/jinzhu/inflection
package plural

var p2s = map[string]string{
	"users":      "user",
	"git_owners": "git_owner",
	"git_repos":  "git_repo",
}

var s2p = map[string]string{}

// ToSingular returns singular from plural.
// It may return the input directly if is not a known (hard coded) plural form.
func ToSingular(p string) string {
	s, ok := p2s[p]
	if ok {
		return s
	}
	return p
}

func init() {
	for k, v := range p2s {
		s2p[v] = k
	}
}
