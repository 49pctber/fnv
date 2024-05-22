package fnv

import (
	"hash/fnv"
	"testing"
)

func TestFNVOffsetBasis(t *testing.T) {
	h := fnv.New32a()
	hsum := h.Sum32()

	if hsum != offset_basis {
		t.Fatalf("offset_basis is incorrect")
	}
}

func TestFNV1a(t *testing.T) {
	payload := []byte("Bryan")

	h := fnv.New32a()
	h.Write(payload)
	want := h.Sum32()

	digest := MyFNV1a(payload)

	if digest != want {
		t.Fatalf("error computing sum")
	}
}
