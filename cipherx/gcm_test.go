package cipherx

import (
	"testing"
)

func TestNewKey(t *testing.T) {

	r, err := NewKey()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(r)
	t.Log(len(r)) // 64

}

func TestNewNonce(t *testing.T) {

	r, err := NewNonce()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(r)
	t.Log(len(r)) // 24
}

func TestNewGcm(t *testing.T) {

	key, err := NewKey()
	if err != nil {
		t.Error(err)
		return
	}

	r, err := NewGcm(key)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(r)
	t.Log(r.Key)
	t.Log(r.Nonce)
}

func TestEncrypt(t *testing.T) {

	var (
		content = "a small piece of metal, plastic, cloth, etc., with words or a picture on it, that you carry with you or that is fastened or sewn to your clothing, often to show your support for a political organization or belief, or to show who you are, your rank, or that you are a member of a group, etc."
	)

	key, err := NewKey()
	if err != nil {
		t.Error(err)
		return
	}

	gcm, err := NewGcm(key)
	if err != nil {
		t.Error(err)
		return
	}

	eContent, err := gcm.Encrypt(content)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("encrypt content is...", eContent)
	t.Log("encrypt len is...", len(eContent)) // 612

	yContent, err := gcm.Decrypt(eContent)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("content is...", yContent)
	t.Log("content len is...", len(yContent)) // 290
}
