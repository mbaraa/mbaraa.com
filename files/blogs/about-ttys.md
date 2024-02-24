### Some History

TTY , teletypewriter or teleprinter was an interface to communicate with a computer, where it was literally a typewriter with some signal sent in/out to it.

![Siemens t37h](https://mbaraa.com/img/9161_early-tty.jpg "Siemens t37h teletypewriter")
Image source [Wikipedia](https://en.wikipedia.org/wiki/Teleprinter)

Oh and for anyone who have been to my blog before, I have written this whole blog post to flex the new image embedding 🤘

\
So this typewriter looking device was actually a computer terminal, where the key strokes are sent using the keyboard, and the output will be printed on a piece of paper, this was originally made in the late 1830s for sending telegraphs, and it was slapped on a computer nearly 1930s (dates might be incorrect, don't judge me), in which the characters were displayed using an early version of the ASCII encoding called [Baudot](https://en.wikipedia.org/wiki/Baudot_code) which was 5 bits in length, so it didn't have much characters choices, that it didn't even have lower-case Latin letters.

Some time after that the paper was replaced by a display with the same logging style, i.e. the characters are written on the lowest row and scrolled up on each new line.

![ADM-3A](link here "ADM-3A Video display with keyboard")
Image source [Wikipedia](https://en.wikipedia.org/wiki/ADM-3A)

These display terminals are called [Dumb Terminals](https://www.pcmag.com/news/the-forgotten-world-of-dumb-terminals) since they are just input and output devices, where the computations are done on a separate computer, fun fact these computers used to have multi-users support like how we can have multiple `/dev/tty`s on a Unix-like machine, but more on that later.

---

### TTY now days

#### Phones

It's somehow related here, that's why it's here, anyways, tty on phones is used by people with speech and hearing impairments, where a person can type a message and it is turned into speech, or the other way around, this helps these people a lot by making communication easier for them.

---

#### Computers

Each shell you open on your Unix-like machine corresponds to a tty, that interacts with the system, to see all of your running ttys run

```bash
ls /dev | grep tty
```

Where this will list all of the ttys opened and used by applications, note that every terminal you open creates a new tty for you, you can check what tty you're currently using by running

```bash
tty
```

It'll output something like this `/dev/pts/4` where **pts** means pseudo terminal, since it has a couple of layers before reaching the operating system.

And every operating system these days includes a terminal (terminal emulator formally) in which you could interact with the OS's shell, but since we Linux users use it so much (because we can not because we're forced to), it's been called terminal for short.

Typically a Unix-like system will greet you with a tty screen where you can use the computer like that or start a window manager or desktop environment.

![Gentoo TTY](https://mbaraa.com/img/5107_gentoo-tty.jpg "Gentoo Linux installing a package using emerge in the tty")
Some OSes like Gentoo, Arch, Void, and FreeBSD, don't include a GUI installer where you have to do the installation process from the tty, and some users (like myself) don't use a display manager to start the window manager (oh it's so bloated OMG), so we use the console tty more often than the average display manager enjoyer.

And since you opened this blog you probably want to customize that tty screen don't you?

Well hop on the next section.

---

### Customizing TTY's Greeting Message in Linux

When you turn on your computer a message like this will appear (I don't use Debian on my computer, but my current tty message is a bit inappropriate)

```
Debian GNU/Linux 11 meowrver tty1
meowrver login:
```

Here where you can login to an existing user on the computer, but see the Debian... part, we're gonna change it!

First off, this message is called `issue` [more on that](https://serverfault.com/questions/922235/what-is-the-difference-between-etc-issue-net-and-etc-issue), where it displays a "greeting message" to the computer's user, and since apparently you use it a lot, you might wanna personalize it, right?

There goes this table of escape sequences used by the `issue` file

| Escape Sequence | Meaning                                                                                |
| --------------- | -------------------------------------------------------------------------------------- |
| **\b**          | Insert the baudrate of the current line.                                               |
| **\d**          | Insert the current date.                                                               |
| **\s**          | Insert the system name, the name of the operating system.                              |
| **\l**          | Insert the name of the current tty line.                                               |
| **\m**          | Insert the architecture identifier of the machine, eg. i486                            |
| **\n**          | Insert the nodename of the machine, also known as the hostname.                        |
| **\o**          | Insert the domainname of the machine.                                                  |
| **\r**          | Insert the release number of the OS, eg. 1.1.9.                                        |
| **\t**          | Insert the current time.                                                               |
| **\u**          | Insert the number of current users logged in.                                          |
| **\U**          | Insert the string “1 user” or “ users” where is the number of current users logged in. |
| **\v**          | Insert the version of the OS, eg. the build-date etc.                                  |

For example if we wanted the same `issue` as above that is `Debian GNU/Linux 11 meowrver tty1`, is actually `Debian GNU/Linux 11 \n \l`.

To make changes, open the file `/etc/issue` as root in your favorite editor and type whatever message you want with or without some escape sequences.

Thanks for sticking to the end, and have a nice day (or night idk).
