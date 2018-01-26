package loader

type LazyLoader interface {
	LazyLoad(path string) string
}

type LazyLoaderImpl struct {
	UrlLoader UrlLoader
	Cache     Cache
}

type Cache interface {
	put(key string, value string)
	get(key string) string
}

type DummyMapCache struct {
	Repo map[string]string
}

func (cacheImpl DummyMapCache) get(key string) string {
	return cacheImpl.Repo[key]
}

func (cacheImpl DummyMapCache) put(key string, value string) {
	cacheImpl.Repo[key] = value
}

func (lazyLoader LazyLoaderImpl) LazyLoad(path string) string {
	val := lazyLoader.Cache.get(path)
	if val == "" {
		val = lazyLoader.UrlLoader.load(path)
		lazyLoader.Cache.put(path, val)
	}
	return val
}
