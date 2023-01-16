package fast

import (
	"container/list"
	"math/rand"
	"reflect"
	"sync"
	"testing"
	"time"
)

func Test_cacheShard_del(t *testing.T) {
	type fields struct {
		locker     sync.RWMutex
		maxEntries int
		onEvicted  func(key string, value interface{})
		ll         *list.List
		cache      map[string]*list.Element
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cacheShard{
				locker:     tt.fields.locker,
				maxEntries: tt.fields.maxEntries,
				onEvicted:  tt.fields.onEvicted,
				ll:         tt.fields.ll,
				cache:      tt.fields.cache,
			}
			c.del(tt.args.key)
		})
	}
}

func Test_cacheShard_delOldest(t *testing.T) {
	type fields struct {
		locker     sync.RWMutex
		maxEntries int
		onEvicted  func(key string, value interface{})
		ll         *list.List
		cache      map[string]*list.Element
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cacheShard{
				locker:     tt.fields.locker,
				maxEntries: tt.fields.maxEntries,
				onEvicted:  tt.fields.onEvicted,
				ll:         tt.fields.ll,
				cache:      tt.fields.cache,
			}
			c.delOldest()
		})
	}
}

func Test_cacheShard_get(t *testing.T) {
	type fields struct {
		locker     sync.RWMutex
		maxEntries int
		onEvicted  func(key string, value interface{})
		ll         *list.List
		cache      map[string]*list.Element
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cacheShard{
				locker:     tt.fields.locker,
				maxEntries: tt.fields.maxEntries,
				onEvicted:  tt.fields.onEvicted,
				ll:         tt.fields.ll,
				cache:      tt.fields.cache,
			}
			if got := c.get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cacheShard_len(t *testing.T) {
	type fields struct {
		locker     sync.RWMutex
		maxEntries int
		onEvicted  func(key string, value interface{})
		ll         *list.List
		cache      map[string]*list.Element
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cacheShard{
				locker:     tt.fields.locker,
				maxEntries: tt.fields.maxEntries,
				onEvicted:  tt.fields.onEvicted,
				ll:         tt.fields.ll,
				cache:      tt.fields.cache,
			}
			if got := c.len(); got != tt.want {
				t.Errorf("len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cacheShard_removeElement(t *testing.T) {
	type fields struct {
		locker     sync.RWMutex
		maxEntries int
		onEvicted  func(key string, value interface{})
		ll         *list.List
		cache      map[string]*list.Element
	}
	type args struct {
		e *list.Element
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cacheShard{
				locker:     tt.fields.locker,
				maxEntries: tt.fields.maxEntries,
				onEvicted:  tt.fields.onEvicted,
				ll:         tt.fields.ll,
				cache:      tt.fields.cache,
			}
			c.removeElement(tt.args.e)
		})
	}
}

func Test_cacheShard_set(t *testing.T) {
	type fields struct {
		locker     sync.RWMutex
		maxEntries int
		onEvicted  func(key string, value interface{})
		ll         *list.List
		cache      map[string]*list.Element
	}
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{name: testing.CoverMode(), fields: fields{}, args: args{
			key:   "",
			value: nil,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cacheShard{
				locker:     tt.fields.locker,
				maxEntries: tt.fields.maxEntries,
				onEvicted:  tt.fields.onEvicted,
				ll:         tt.fields.ll,
				cache:      tt.fields.cache,
			}
			c.set(tt.args.key, tt.args.value)
		})
	}
}

func Test_newCacheShard(t *testing.T) {
	type args struct {
		maxEntries int
		onEvicted  func(key string, value interface{})
	}
	tests := []struct {
		name string
		args args
		want *cacheShard
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newCacheShard(tt.args.maxEntries, tt.args.onEvicted); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newCacheShard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cacheShard_del1(t *testing.T) {
	type fields struct {
		locker     sync.RWMutex
		maxEntries int
		onEvicted  func(key string, value interface{})
		ll         *list.List
		cache      map[string]*list.Element
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cacheShard{
				locker:     tt.fields.locker,
				maxEntries: tt.fields.maxEntries,
				onEvicted:  tt.fields.onEvicted,
				ll:         tt.fields.ll,
				cache:      tt.fields.cache,
			}
			c.del(tt.args.key)
		})
	}
}

func Test_cacheShard_delOldest1(t *testing.T) {
	type fields struct {
		locker     sync.RWMutex
		maxEntries int
		onEvicted  func(key string, value interface{})
		ll         *list.List
		cache      map[string]*list.Element
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cacheShard{
				locker:     tt.fields.locker,
				maxEntries: tt.fields.maxEntries,
				onEvicted:  tt.fields.onEvicted,
				ll:         tt.fields.ll,
				cache:      tt.fields.cache,
			}
			c.delOldest()
		})
	}
}

func Test_cacheShard_get1(t *testing.T) {
	type fields struct {
		locker     sync.RWMutex
		maxEntries int
		onEvicted  func(key string, value interface{})
		ll         *list.List
		cache      map[string]*list.Element
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cacheShard{
				locker:     tt.fields.locker,
				maxEntries: tt.fields.maxEntries,
				onEvicted:  tt.fields.onEvicted,
				ll:         tt.fields.ll,
				cache:      tt.fields.cache,
			}
			if got := c.get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cacheShard_len1(t *testing.T) {
	type fields struct {
		locker     sync.RWMutex
		maxEntries int
		onEvicted  func(key string, value interface{})
		ll         *list.List
		cache      map[string]*list.Element
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cacheShard{
				locker:     tt.fields.locker,
				maxEntries: tt.fields.maxEntries,
				onEvicted:  tt.fields.onEvicted,
				ll:         tt.fields.ll,
				cache:      tt.fields.cache,
			}
			if got := c.len(); got != tt.want {
				t.Errorf("len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cacheShard_removeElement1(t *testing.T) {
	type fields struct {
		locker     sync.RWMutex
		maxEntries int
		onEvicted  func(key string, value interface{})
		ll         *list.List
		cache      map[string]*list.Element
	}
	type args struct {
		e *list.Element
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cacheShard{
				locker:     tt.fields.locker,
				maxEntries: tt.fields.maxEntries,
				onEvicted:  tt.fields.onEvicted,
				ll:         tt.fields.ll,
				cache:      tt.fields.cache,
			}
			c.removeElement(tt.args.e)
		})
	}
}

func Test_cacheShard_set1(t *testing.T) {
	type fields struct {
		locker     sync.RWMutex
		maxEntries int
		onEvicted  func(key string, value interface{})
		ll         *list.List
		cache      map[string]*list.Element
	}
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cacheShard{
				locker:     tt.fields.locker,
				maxEntries: tt.fields.maxEntries,
				onEvicted:  tt.fields.onEvicted,
				ll:         tt.fields.ll,
				cache:      tt.fields.cache,
			}
			c.set(tt.args.key, tt.args.value)
		})
	}
}

func Test_newCacheShard1(t *testing.T) {
	type args struct {
		maxEntries int
		onEvicted  func(key string, value interface{})
	}
	tests := []struct {
		name string
		args args
		want *cacheShard
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newCacheShard(tt.args.maxEntries, tt.args.onEvicted); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newCacheShard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkTourFastCacheSetParallel(b *testing.B) {
	maxEntriesSize := 12
	cache := NewFastCache(b.N, maxEntriesSize, nil)
	rand.Seed(time.Now().Unix())
	b.RunParallel(func(pb *testing.PB) {
		id := rand.Intn(1000)
		counter := 0
		for pb.Next() {
			cache.Set(paralleKey(id, counter), value())
			counter = counter + 1
		}
	})
}

func value() interface{} {
	return nil
}

func paralleKey(id int, counter int) string {
	return ""
}
