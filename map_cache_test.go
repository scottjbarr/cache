package cache

import (
	"reflect"
	"testing"
	"time"
)

func TestMapCache(t *testing.T) {
	c := NewMapCache()

	foo := widget{"foo", 10}

	c.Set(foo.name, foo)

	got, err := c.Get(foo.name)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, foo) {
		t.Errorf("got\n%+v\nwant\n%+v", got, foo)
	}
}

func TestMapCache_concurrent_access(t *testing.T) {
	c := NewMapCache()

	foo := widget{"foo", 0}

	// start writing to the cache, incrementing the count
	go func() {

		for {
			foo.count++
			c.Set(foo.name, foo)
		}
	}()

	var last interface{}
	readCount := 0

	// start reading from the cache
	go func() {
		for {
			got, err := c.Get(foo.name)
			if err != nil {
				t.Fatal(err)
			}

			last = got
			readCount++
		}
	}()

	// let the goroutines read and write to the cache.
	time.Sleep(time.Millisecond)

	// Keep in mind that the goroutines are still running.
	if foo.count == 0 {
		t.Fatal("want > 0 updates")
	}

	if readCount == 0 {
		t.Fatal("want > 0 reads")
	}

	if last == nil {
		t.Fatalf("No cached item")
	}

	got := last.(widget)

	if got.count == 0 {
		t.Fatal("want > 0")
	}
}

func TestCache_example(t *testing.T) {
	// the cache handles the storage and mutex.
	c := NewMapCache()

	// the implementation provides type safety, with an internal cache that is safe from concurrent
	// access issues.
	w := newWidgetCache(c)

	f := widget{"foo", 0}

	// We can only pass "widget" type here.
	if err := w.Set(f.name, f); err != nil {
		t.Fatal(err)
	}

	// Returns a *widget as a cache miss is possile. The implementation could choose to return an
	// error but when reading a cache is seems reasonable that an item may not exist in many
	// scenarios. Let the implementation decide.
	got, err := w.Get(f.name)
	if err != nil {
		t.Fatal(err)
	}

	if got == nil {
		t.Fatal("unexpected nil")
	}

	if !reflect.DeepEqual(*got, f) {
		t.Errorf("got\n%+v\nwant\n%+v", *got, f)
	}
}
