## glimmer - a bit of code for golang to get a framebuffer and sound.

Windowing code uses [exp/shiny](https://github.com/golang/exp/tree/master/shiny).  Audio uses [oto](https://github.com/hajimehoshi/oto).

#### Dependencies for screen

Glimmer has no C dependencies for screen support.

#### Dependencies for sound

* Windows has no C dependencies for sound (yay!)

* Linux users should 'apt install libasound2-dev' or equivalent

* FreeBSD (and mac?) users should 'pkg install openal-soft' or equivalent
