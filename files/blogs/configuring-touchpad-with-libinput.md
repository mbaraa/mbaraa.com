Configuring the toupchpad using **libinput** is useful for global configuration (across DEs and WMs), or for window managers if you don't wanna use some hacky graphical tool.

First install the package `xf86-input-libinput` if you're using Xorg or `libinput` for Wayland then create the file `/etc/X11/xorg.conf.d/30-touchpad.conf` and add the follwoing lines to it:

```bash
Section "InputClass"
    Identifier "touchpad"
    Driver "libinput"
    # set MatchIsTouchpad "on" if you’re using a mouse or a trackpoint like in the thinkpads
    MatchIsTouchpad "on"
    # tapping can be "on" or "off" depends whether you want tapping or not
    Option "Tapping" "on"
    # natural scrolling can be "true" or "false" depends whether you want natural scrolling or not.
    # also natural scrolling is when scrolling the touchpad the page goes in the same direction of the scrolling.
    Option "NaturalScrolling" "false"
    # horizontal scrolling can be "true" or "false" depends whether you want horizontal scrolling or not
    Option "HorizontalScrolling" "true"
    # button mapping where each one represents the number of taps,
    # i.e. here left one tap, right two taps, and middle is three taps
    # where taps are done at the same time :)
    Option "TappingButtonMap" "lrm"
EndSection
```

Note that if you put more than one option, or an invalid option the touchpad will use the default configuration.
