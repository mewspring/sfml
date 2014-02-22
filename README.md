WIP
---

This project is a *work in progress*. The implementation is *incomplete* and
subject to change. The documentation can be inaccurate.

sfml
====

This project implements window creation, event handling and image drawing using
[SFML][] version 2.1.

[SFML]: http://www.sfml-dev.org/

Documentation
-------------

Documentation provided by GoDoc.

- sfml
   - [font][sfml/font]: handles text rendering based on the size, style and
   color of fonts.
   - [texture][sfml/texture]: handles hardware accelerated image drawing.
   - [window][sfml/window]: handles window creation, drawing and events.

[sfml/font]: http://godoc.org/github.com/mewmew/sfml/font
[sfml/texture]: http://godoc.org/github.com/mewmew/sfml/texture
[sfml/window]: http://godoc.org/github.com/mewmew/sfml/window

Examples
--------

### off-screen

The [off-screen][examples/off-screen] command demonstrates how to perform
hardware accelerated off-screen rendering.

	go get github.com/mewmew/sfml/examples/off-screen

[examples/off-screen]: https://github.com/mewmew/sfml/blob/master/examples/off-screen/off-screen.go#L34

### tiny

The [tiny][examples/tiny] command demonstrates how to render images onto the
window using the Draw and DrawRect methods. It also gives an example of a basic
event loop.

	go get github.com/mewmew/sfml/examples/tiny

![Screenshot - tiny](https://raw.github.com/mewmew/sfml/master/examples/tiny/tiny.png)

[examples/tiny]: https://github.com/mewmew/sfml/blob/master/examples/tiny/tiny.go#L38

### many

The [many][examples/many] command demonstrates how to create and handle more
than one window at once.

	go get github.com/mewmew/sfml/examples/many

[examples/many]: https://github.com/mewmew/sfml/blob/master/examples/many/many.go#L36

### fonts

The [fonts][examples/fonts] command demonstrates how to render text using TTF
fonts.

	go get github.com/mewmew/sfml/examples/fonts

![Screenshot - fonts](https://raw.github.com/mewmew/sfml/master/examples/fonts/fonts.png)

[examples/fonts]: https://github.com/mewmew/sfml/blob/master/examples/fonts/fonts.go#L38

public domain
-------------

This code is hereby released into the *[public domain][]*.

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/
