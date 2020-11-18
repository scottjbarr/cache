# Cache

Provides a simple Cache interface, with a number of Cache backends.

The idea is that you would provide a type safe cache, using one of the Cache backends
internally. This allows for easy swapping of the Cache. For example, unit tests may want to use a
map, but in production you may want to use a Redis based cache.


## Cache Implementations.

### MapCache

Provides safe access to a map.


## License

The MIT License (MIT)

Copyright (c) 2020 Scott Barr

See [LICENSE.md](LICENSE.md)
