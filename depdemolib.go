package depdemolib

import "golang.org/x/crypto/ssh"

// MakeCert creates and returns an empty ssh.Certificate.
func MakeCert() *ssh.Certificate {
	return &ssh.Certificate{}
}
