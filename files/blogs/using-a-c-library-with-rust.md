# Why?

The most common thing is to use an embedded library that exists in C and it's too big to be rewritten in Rust, my case I needed to use [ALSA](https://www.alsa-project.org/wiki/Main_Page) for a project, so here I am, I've used ALSA before with Go in another project, which was kinda straight forward, just add the files, link some addidional files, and you're good to go (not intended punch), but with Rust you need to do some more stuff when linking non standard libraries.

One way to do it, is to use [Bindgen](https://rust-lang.github.io/rust-bindgen/), which exports the C types directly into Rust types, then you can go from there, me, myself I don't like generated code in my project (If you're not new here, yes I do enjoy templ, but that's a different story...), so I went to use [CMake](https://cmake.org/) with [cmake-rs](https://crates.io/crates/cmake) on top, that way I can mess around as much as I like with C files, and have the perfect level of abstraction without having to dance with C types in Rust (the bindgen way).

So, let's begin, we're gonna need the following:

- Cargo
- CMake
- Clang
- Some time
- A listen to [The Rains of Castamere](https://dankmuzikk.com/song/tlkgbwmN9mQ)

# Hello C

We're gonna be writing a super abstract layer to ALSA, where literally the only exposed functions from C are `init_alsa`, `destroy_alsa` and `play_frequency`, where their names really give them away, the we're gonna use the ALSA abstraction to play the first 29 notes of The Rains of Castamere.

Start by creating a Rust project

```bash
cargo new rust-of-castamere
```

Then create a directory, preferbly at the root level of the project so there's separation between C and Rust files.

```bash
mkdir libalsa
```

And create the CMake descriptor file `libalsa/CMakeLists.txt`

```cmake
# libalsa/CMakeLists.txt
# the coolest cmake at the time of writing this post was 3.6
cmake_minimum_required(VERSION 3.6)
# set project name
project(LibAlsa)
# define the library (project) as static
add_library(alsa STATIC alsa.c)
# add the project files to the target rust files
install(TARGETS alsa DESTINATION .)
```

Now create `alsa.h` and `alsa.c` to test if this thing actually works before adding ALSA to the mix.

And make sure to `extern` your C function at least in the header file.

```c
// libalsa/alsa.h
#ifndef ALSA_RS_H
#define ALSA_RS_H

#include <stdio.h>
#include <math.h>

extern void greet_3(char* name);

#endif
```

```c
// libalsa/alsa.c
#include "alsa.h"

void greet_3(char* name) {
    printf("Hello %s, some number: %f\n", name, pow(2, 3.3));
}
```

Now back to Rust, add `cmake` as a build dependency

```bash
cargo add cmake --build
```

Then we need to define a [build script](https://doc.rust-lang.org/cargo/reference/build-scripts.html) so that the cmake project is added to the mix

```rust
// build.rs
use cmake::Config;

fn main() {
    let dst = Config::new("libalsa").build();
    // linker directives, see the link above for more reference.
    println!("cargo:rustc-link-search=native={}", dst.display());
    println!("cargo:rustc-link-lib=static=alsa");
}
```

And of course you need to update `Cargo.toml` to specify the path of the build script, just add this under the `package` section

```toml
build = "build.rs"
```

Now actually back to Rust, where we'll be using the C function we defined earlier in our unsafe Rust code (chills).

Make sure that the **name** and **kind** of library match the ones in `build.rs`, then add your function with types same as they were in C, e.g `int` => `i32`, `double` => `f64` and so on...

```rust
// src/main.rs
#[link(name = "alsa", kind = "static")]
extern "C" {
    // clippy will scream this:
    // warning: `extern` block uses type `str`, which is not FFI-safe
    // well, I have a great come back, SHUT UP CLIPPY!
    fn greet_3(name: &'static str);
}

fn main() {
    unsafe {
        greet_3("Dingus");
    }
}
```

Compile and run and you shall see something like this

```bash
cargo run
# Hello Dingus, some number: 9.849155
```

That was fun wasn't it? well the ALSA won't be, cuz I wrote the C code like 2 years ago, so I don't have much knowledge why and how it's running...

# Hello ALSA

Hello ALSA, I'd like number 3 please, which will be those C snippets

```c
// libalsa/alsa.h
#ifndef ALSA_RS_H
#define ALSA_RS_H

#include <alsa/asoundlib.h>
#include <math.h>
#include <sys/types.h>

extern int init_alsa();
extern int destroy_alsa();
extern int play_frequency(float freq, u_int16_t rate, float latency,
                          float duration);

#endif
```

Seriously I'm really lazy to re-open the past, just take it as is, and there's a little bit of math, aaah back in the day when I could do math.

```c
// libalsa/alsa.c
#include "alsa.h"

snd_pcm_t *handle;

long long second_to_micro(float seconds) { return (long long)(seconds * 1e6); }

int init_alsa() {
  return snd_pcm_open(&handle, "default", SND_PCM_STREAM_PLAYBACK,
                      0 /* blocked mode */);
}

int destroy_alsa() { return snd_pcm_close(handle); }

int play_frequency(float freq, u_int16_t rate, float latency, float duration) {
  if (latency > 150.0) {
    latency = 150.0;
  }
  latency = second_to_micro(latency);

  unsigned char buffer[(int)(rate * duration)];

  for (int i = 0; i < sizeof(buffer); i++) {
    buffer[i] = 0xFF * sin(2 * M_PI * freq * i / rate);
  }

  if (0 != snd_pcm_set_params(handle, SND_PCM_FORMAT_U8,
                              SND_PCM_ACCESS_RW_INTERLEAVED, 1 /* channels */,
                              rate /* rate [Hz] */, 1 /* soft resample */,
                              latency /* latency [us] */)) {
    return 0;
  }

  snd_pcm_writei(handle, buffer, sizeof(buffer));

  return 0;
}
```

Back to Rust, update your `build.rs` to include a linker directive that will link ALSA into the project.

```rust
// build.rs
use cmake::Config;

fn main() {
    let dst = Config::new("libalsa").build();

    println!("cargo:rustc-link-search=native={}", dst.display());
    println!("cargo:rustc-link-lib=static=alsa");
    println!("cargo:rustc-link-lib=dylib=asound");
}
```

Then import the C functions using the `extern` keyword

```rust
// src/main.rs
#[link(name = "alsa", kind = "static")]
extern "C" {
    fn init_alsa() -> i32;
    fn destroy_alsa() -> i32;
    fn play_frequency(freq: f32, rate: u16, latency: f32, duration: f32) -> i32;
}

fn main() {
    unsafe {
        init_alsa();
        // play a little sensoring sound for 5 seconds.
        play_frequency(1000.0, 44100, 0.1, 5.0);
        destroy_alsa();
    }
}
```

Yes as you can see the functions return a `i32` which is the `int` retuned form C that represented the status code, we're gonna ignore them here :)

# Rains of Castamere

Full example can be found [here](https://github.com/mbaraa/pub_code/tree/main/blog/using-a-c-library-with-rust), but let's write it here as well, actually we'll just be adding more Rust stuff that'll represent the first 29 notes of The Rains of Castamere.

Just a little heads up, I have no idea how music works, I did some research, and I got some terms, played a little around, and got the frequencies of the notes, I would've continued, but it doesn't sound that good, so I stopped, if you can help with the notes, or the namings of the stuff, contact me, or a PR to this [blog](https://github.com/mbaraa/mbaraa.com) post.

Anyways, let's define the `NoteDuration` enum, which will hold common multipliers of a beat or something.

```rust
enum NoteDuration {
    TwoNotes,
    WholeNote,
    HalfNote,
}

impl NoteDuration {
    fn value(&self, secs: f32) -> f32 {
        match *self {
            Self::TwoNotes => secs * 2.0,
            Self::WholeNote => secs,
            Self::HalfNote => secs * 0.5,
        }
    }
}
```

Then define a `Note` struct which will be hodling the frequency of the note, and its duration.

```rust
struct Note {
    freq: f32,
    duration: NoteDuration,
}

impl Note {
    fn new(freq: f32, duration: NoteDuration) -> Self {
        Self { freq, duration }
    }
}
```

Now let's play the notes.

```rust
unsafe {
    vec![
        Note::new(110.0, NoteDuration::HalfNote),
        Note::new(174.61, NoteDuration::WholeNote),
        Note::new(110.0, NoteDuration::HalfNote),
        Note::new(164.81, NoteDuration::WholeNote),
        Note::new(110.0, NoteDuration::HalfNote),
        Note::new(174.61, NoteDuration::HalfNote),
        Note::new(196.0, NoteDuration::WholeNote),
        Note::new(164.81, NoteDuration::HalfNote),
        Note::new(110.0, NoteDuration::WholeNote),
        Note::new(196.0, NoteDuration::WholeNote),
        Note::new(174.61, NoteDuration::WholeNote),
        Note::new(164.81, NoteDuration::HalfNote),
        Note::new(146.83, NoteDuration::WholeNote),
        Note::new(164.81, NoteDuration::HalfNote),
        Note::new(130.81, NoteDuration::HalfNote),
        Note::new(220.0, NoteDuration::WholeNote),
        Note::new(130.81, NoteDuration::HalfNote),
        Note::new(196.0, NoteDuration::WholeNote),
        Note::new(130.81, NoteDuration::HalfNote),
        Note::new(220.0, NoteDuration::WholeNote),
        Note::new(233.08, NoteDuration::WholeNote),
        Note::new(196.0, NoteDuration::WholeNote),
        Note::new(220.0, NoteDuration::HalfNote),
        Note::new(233.08, NoteDuration::WholeNote),
        Note::new(220.0, NoteDuration::WholeNote),
        Note::new(196.0, NoteDuration::WholeNote),
        Note::new(174.61, NoteDuration::WholeNote),
        Note::new(174.61, NoteDuration::TwoNotes),
        Note::new(164.81, NoteDuration::WholeNote),
    ]
    .iter()
    .for_each(|note| {
        println!("{}", note.freq);
        play_frequency(note.freq, 44100, 0.1, note.duration.value(1.0));
        //                      sample, latency
        //                      rate
    });
}
```

# Quote of the day

"And who are you, the proud lord said,
That I must bow so low?
Only a cat of a different coat,
That's all the truth I know.
In a coat of gold or a coat of red,
A lion still has claws,
And mine are long and sharp, my lord,
As long and sharp as yours.
And so he spoke, and so he spoke,
That lord of Castamere,
But now the rains weep o'er his hall,
With no one there to hear.
Yes now the rains weep o'er his hall,
And not a soul to hear."
\
\- [George R.R. Martin](https://en.wikipedia.org/wiki/George_R._R._Martin)
