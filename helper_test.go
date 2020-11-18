package cache

import "errors"

type widget struct {
	name  string
	count int
}

type widgetCache struct {
	cache Cache
}

func newWidgetCache(c Cache) widgetCache {
	return widgetCache{
		cache: c,
	}
}

func (w *widgetCache) Set(key string, v widget) error {
	return w.cache.Set(key, v)
}

func (w *widgetCache) Get(key string) (*widget, error) {
	v, err := w.cache.Get(key)
	if err != nil {
		return nil, err
	}

	o, ok := v.(widget)
	if !ok {
		return nil, errors.New("Unexpected type")
	}

	return &o, nil
}
