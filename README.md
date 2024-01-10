## glimmer - a bit of code for golang to get a framebuffer and sound.

This package exists because it used to take a bit of cobbling together to get a
good answer to that question that didn't need any C bindings, but currently it just
uses [ebiten](https://github.com/hajimehoshi/ebiten).

Ebiten enables getting sound and screen without any dependencies other than Go on
Windows, which is mainly what I wanted out of this package.

See ebiten's page for dependencies for other platforms.
