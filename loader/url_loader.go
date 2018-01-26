package loader

type UrlLoader interface {
	load(path string) string
}

type UrlLoaderImpl struct {
	BasePath string
}

func (loader UrlLoaderImpl) load(path string) string {
	//add here your fetching logic
	return `
		{
			"response" : "hello world"
		}
	`
}
