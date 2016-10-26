# vim-filewatcher

file watcher

## Usage

start watch the change of directory.
```
let watcher = filewatcher#watch(".", {x,y->execute('echo y',0)})
```

stop the watching.
```
call watcher.stop()
```

## Requirements

golang

## Installation

```
cd filewatcher && go get -d && go build
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a. mattn)
