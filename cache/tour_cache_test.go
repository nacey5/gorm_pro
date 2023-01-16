package cache

import (
	"reflect"
	"testing"
)

var fuo GetFunc

func TestGetFunc_Get(t *testing.T) {

	type args struct {
		key string
	}
	tests := []struct {
		name string
		f    GetFunc
		args args
		want interface{}
	}{
		// TODO: Add test cases.
		{name: "t1", f: fuo, args: args{"mosi"}, want: "mosivalue"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTourCache(t *testing.T) {
	type args struct {
		getter Getter
		cache  Cache
	}
	tests := []struct {
		name string
		args args
		want *TourCache
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTourCache(tt.args.getter, tt.args.cache); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTourCache() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTourCache_Get(t1 *testing.T) {
	type fields struct {
		mainCache *safeCache
		getter    Getter
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
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TourCache{
				mainCache: tt.fields.mainCache,
				getter:    tt.fields.getter,
			}
			if got := t.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
