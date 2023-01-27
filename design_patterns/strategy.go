package main

import "fmt"

type PasswordProtector struct {
	user         string
	passwordName string
	hashAlgoritm HashAlgoritim
}

func (p *PasswordProtector) SetHashAlgoritim(hash HashAlgoritim) {
	p.hashAlgoritm = hash
}

func (p *PasswordProtector) Hash() {
	p.hashAlgoritm.Hash(p)
}

type HashAlgoritim interface {
	Hash(p *PasswordProtector)
}

func NewPasswordProtector(user string, passwordName string, hash HashAlgoritim) *PasswordProtector {
	return &PasswordProtector{
		user:         user,
		passwordName: passwordName,
		hashAlgoritm: hash,
	}
}

type SHA struct{}

func (SHA) Hash(p *PasswordProtector) {
	fmt.Printf("Hasing using SHA for %s\n", p.passwordName)
}

type MD5 struct{}

func (MD5) Hash(p *PasswordProtector) {
	fmt.Printf("Hasing using MD5 for %s\n", p.passwordName)
}

func main() {
	sha := &SHA{}
	md5 := &MD5{}

	passwordProtector := NewPasswordProtector("Carlos", "MY password", sha)
	passwordProtector.Hash()

	passwordProtector.SetHashAlgoritim(md5)
	passwordProtector.Hash()

}
