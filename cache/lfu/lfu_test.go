package lfu

import (
	"gorm_pro/cache"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		maxBytes  int
		onEvicted func(key string, value interface{})
	}
	tests := []struct {
		name string
		args args
		want cache.Cache
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.maxBytes, tt.args.onEvicted); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_entry_Len(t *testing.T) {
	type fields struct {
		key    string
		value  interface{}
		weight int
		index  int
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
			e := &entry{
				key:    tt.fields.key,
				value:  tt.fields.value,
				weight: tt.fields.weight,
				index:  tt.fields.index,
			}
			if got := e.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lfu_Del(t *testing.T) {
	type fields struct {
		maxBytes  int
		onEvicted func(key string, value interface{})
		usedBytes int
		queue     *queue
		cache     map[string]*entry
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
			l := &lfu{
				maxBytes:  tt.fields.maxBytes,
				onEvicted: tt.fields.onEvicted,
				usedBytes: tt.fields.usedBytes,
				queue:     tt.fields.queue,
				cache:     tt.fields.cache,
			}
			l.Del(tt.args.key)
		})
	}
}

func Test_lfu_DelOldest(t *testing.T) {
	type fields struct {
		maxBytes  int
		onEvicted func(key string, value interface{})
		usedBytes int
		queue     *queue
		cache     map[string]*entry
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &lfu{
				maxBytes:  tt.fields.maxBytes,
				onEvicted: tt.fields.onEvicted,
				usedBytes: tt.fields.usedBytes,
				queue:     tt.fields.queue,
				cache:     tt.fields.cache,
			}
			l.DelOldest()
		})
	}
}

func Test_lfu_Get(t *testing.T) {
	type fields struct {
		maxBytes  int
		onEvicted func(key string, value interface{})
		usedBytes int
		queue     *queue
		cache     map[string]*entry
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
			l := &lfu{
				maxBytes:  tt.fields.maxBytes,
				onEvicted: tt.fields.onEvicted,
				usedBytes: tt.fields.usedBytes,
				queue:     tt.fields.queue,
				cache:     tt.fields.cache,
			}
			if got := l.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lfu_Len(t *testing.T) {
	type fields struct {
		maxBytes  int
		onEvicted func(key string, value interface{})
		usedBytes int
		queue     *queue
		cache     map[string]*entry
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
			l := &lfu{
				maxBytes:  tt.fields.maxBytes,
				onEvicted: tt.fields.onEvicted,
				usedBytes: tt.fields.usedBytes,
				queue:     tt.fields.queue,
				cache:     tt.fields.cache,
			}
			if got := l.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lfu_Set(t *testing.T) {
	type fields struct {
		maxBytes  int
		onEvicted func(key string, value interface{})
		usedBytes int
		queue     *queue
		cache     map[string]*entry
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
			l := &lfu{
				maxBytes:  tt.fields.maxBytes,
				onEvicted: tt.fields.onEvicted,
				usedBytes: tt.fields.usedBytes,
				queue:     tt.fields.queue,
				cache:     tt.fields.cache,
			}
			l.Set(tt.args.key, tt.args.value)
		})
	}
}

func Test_lfu_removeElement(t *testing.T) {
	type fields struct {
		maxBytes  int
		onEvicted func(key string, value interface{})
		usedBytes int
		queue     *queue
		cache     map[string]*entry
	}
	type args struct {
		x interface{}
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
			l := &lfu{
				maxBytes:  tt.fields.maxBytes,
				onEvicted: tt.fields.onEvicted,
				usedBytes: tt.fields.usedBytes,
				queue:     tt.fields.queue,
				cache:     tt.fields.cache,
			}
			l.removeElement(tt.args.x)
		})
	}
}

func Test_queue_Len(t *testing.T) {
	tests := []struct {
		name string
		q    queue
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queue_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		q    queue
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queue_Pop(t *testing.T) {
	tests := []struct {
		name string
		q    queue
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queue_Push(t *testing.T) {
	type args struct {
		x interface{}
	}
	tests := []struct {
		name string
		q    queue
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.q.Push(tt.args.x)
		})
	}
}

func Test_queue_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		q    queue
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.q.Swap(tt.args.i, tt.args.j)
		})
	}
}

func Test_queue_update(t *testing.T) {
	type args struct {
		en     *entry
		value  interface{}
		weight int
	}
	tests := []struct {
		name string
		q    queue
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.q.update(tt.args.en, tt.args.value, tt.args.weight)
		})
	}
}
