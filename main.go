package fnv

const offset_basis uint32 = 2166136261
const fnv_prime uint32 = 16777619

func MyFNV1a(p []byte) uint32 {
	var digest uint32 = offset_basis
	for _, b := range p {
		digest ^= uint32(b)
		digest *= fnv_prime
	}
	return digest
}
