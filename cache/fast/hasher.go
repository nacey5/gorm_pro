package fast

func newDefaultHasher() fnv64 {
	return fnv64{}
}

type fnv64 struct {
}

const (
	offset64 = 14695981039346656037
	prime64  = 1099511628211
)

func (f fnv64) Sum64(key string) uint64 {
	var hash uint64 = offset64
	for i := 0; i < len(key); i++ {
		hash ^= uint64(key[i])
		hash *= prime64
	}
	return hash
}

type fnv64a struct {
}

func (f fnv64a) Sum64(key string) uint64 {
	var hash uint64 = offset64
	for i := 0; i < len(key); i++ {
		hash *= prime64
		hash ^= uint64(key[i])
	}
	return hash
}
