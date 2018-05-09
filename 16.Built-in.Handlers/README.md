# Built-in Handlers

## http.NotFoundHandler

```
func NotFoundHandler () Handler
```

## http.StripPrefix

```
func StripPrefix(prefix string, h handler) Handler
```

## http.fileServer

```
func FileServer(root FileSystem) Handle

type FileSystem interface {
    Open(name stirng) (File, error)
}

type dir string
func (d Dir Open(name string)) (File, error)
```
