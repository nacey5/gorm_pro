package fifo

import (
	"container/list"
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
		key   string
		value interface{}
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
				key:   tt.fields.key,
				value: tt.fields.value,
			}
			if got := e.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fifo_Del(t *testing.T) {
	type fields struct {
		maxBytes  int
		onEvicted func(key string, value interface{})
		usedBytes int
		ll        *list.List
		cache     map[string]*list.Element
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
			f := &fifo{
				maxBytes:  tt.fields.maxBytes,
				onEvicted: tt.fields.onEvicted,
				usedBytes: tt.fields.usedBytes,
				ll:        tt.fields.ll,
				cache:     tt.fields.cache,
			}
			f.Del(tt.args.key)
		})
	}
}

func Test_fifo_DelOldest(t *testing.T) {
	type fields struct {
		maxBytes  int
		onEvicted func(key string, value interface{})
		usedBytes int
		ll        *list.List
		cache     map[string]*list.Element
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fifo{
				maxBytes:  tt.fields.maxBytes,
				onEvicted: tt.fields.onEvicted,
				usedBytes: tt.fields.usedBytes,
				ll:        tt.fields.ll,
				cache:     tt.fields.cache,
			}
			f.DelOldest()
		})
	}
}

func Test_fifo_Get(t *testing.T) {
	type fields struct {
		maxBytes  int
		onEvicted func(key string, value interface{})
		usedBytes int
		ll        *list.List
		cache     map[string]*list.Element
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
			f := &fifo{
				maxBytes:  tt.fields.maxBytes,
				onEvicted: tt.fields.onEvicted,
				usedBytes: tt.fields.usedBytes,
				ll:        tt.fields.ll,
				cache:     tt.fields.cache,
			}
			if got := f.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fifo_Len(t *testing.T) {
	type fields struct {
		maxBytes  int
		onEvicted func(key string, value interface{})
		usedBytes int
		ll        *list.List
		cache     map[string]*list.Element
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
			f := &fifo{
				maxBytes:  tt.fields.maxBytes,
				onEvicted: tt.fields.onEvicted,
				usedBytes: tt.fields.usedBytes,
				ll:        tt.fields.ll,
				cache:     tt.fields.cache,
			}
			if got := f.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fifo_Set(t *testing.T) {
	type fields struct {
		maxBytes  int
		onEvicted func(key string, value interface{})
		usedBytes int
		ll        *list.List
		cache     map[string]*list.Element
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
			f := &fifo{
				maxBytes:  tt.fields.maxBytes,
				onEvicted: tt.fields.onEvicted,
				usedBytes: tt.fields.usedBytes,
				ll:        tt.fields.ll,
				cache:     tt.fields.cache,
			}
			f.Set(tt.args.key, tt.args.value)
		})
	}
}

func Test_fifo_removeElement(t *testing.T) {
	type fields struct {
		maxBytes  int
		onEvicted func(key string, value interface{})
		usedBytes int
		ll        *list.List
		cache     map[string]*list.Element
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
			f := &fifo{
				maxBytes:  tt.fields.maxBytes,
				onEvicted: tt.fields.onEvicted,
				usedBytes: tt.fields.usedBytes,
				ll:        tt.fields.ll,
				cache:     tt.fields.cache,
			}
			f.removeElement(tt.args.e)
		})
	}
}
