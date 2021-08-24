# sfml

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/mewspring/sfml)

The sfml project implements window creation, event handling and image drawing using [SFML](http://www.sfml-dev.org/) version 2.5.

## Examples

### tiny

The [tiny](https://github.com/mewspring/sfml/blob/master/examples/tiny/tiny.go#L37) command demonstrates how to render images onto the window using the [Draw](http://godoc.org/github.com/mewspring/sfml/window#Window.Draw) and [DrawRect](http://godoc.org/github.com/mewspring/sfml/window#Window.DrawRect) methods. It also gives an example of a basic event loop.

```bash
go install -v github.com/mewspring/sfml/examples/tiny@master
```

![Screenshot - tiny](https://raw.githubusercontent.com/mewspring/sfml/master/examples/tiny/tiny.png)

### fonts

The [fonts](https://github.com/mewspring/sfml/blob/master/examples/fonts/fonts.go#L39) command demonstrates how to render text using TTF fonts.

```bash
go install -v github.com/mewspring/sfml/examples/fonts@master
```

![Screenshot - fonts](https://raw.githubusercontent.com/mewspring/sfml/master/examples/fonts/fonts.png)

### many

The [many](https://github.com/mewspring/sfml/blob/master/examples/many/many.go#L36) command demonstrates how to create and handle more than one window at once.

```bash
go install -v github.com/mewspring/sfml/examples/many@master
```

### off-screen

The [off-screen](https://github.com/mewspring/sfml/blob/master/examples/off-screen/off-screen.go#L34) command demonstrates how to perform hardware accelerated off-screen rendering.

```bash
go install -v github.com/mewspring/sfml/examples/off-screen@master
```

### soft

The [soft](https://github.com/mewspring/sfml/blob/master/examples/soft/soft.go#L34) command demonstrates how to combine software and hardware rendering.

```bash
go install -v github.com/mewspring/sfml/examples/soft@master
```
