If you have a probe get out, since there's an [official](https://www.raspberrypi.com/documentation/microcontrollers/debug-probe.html) documentation about it :)

Now for debugging, I know I made some drama about it in my previous [post](https://mbaraa.com/blog/running-rust-on-raspberry-pi-pico#debugging), but I just got the Pico, and I was used to the Arduino's serial thingy built-in to the IDE, so it wasn't that much to think about, I haven't used C or MicroPython with the Pico, so I don't really know what's the debugging deal with them, but with Rust it was a bit of a hassle.

### Preface

So my first solution was the one in the [rp-hal's docs](https://docs.rs/rp2040-hal/latest/rp2040_hal/uart/index.html), it was working and I could catch the serial signal with a UART or the Pico's serial, but it was blocking the other pins (or at least that's what I thought), so I couldn't use the other pins for basic GPIO stuff, and that led me to search the dark web (not really but I did search for a solution for like 4 hours).

There's [this](https://www.reddit.com/r/rust/comments/14atkm3/media_debug_pi_pico_using_raspberry_pi4/) awesome post by someone on Reddit (it's always someone on Reddit who posts the real deal), so [u/ThatBrokeDave](https://www.reddit.com/user/ThatBlokeDave), special thanks for you if you ever came across this blog 🫡

But the problem was from [defmt](https://docs.rs/defmt/latest/defmt) where the code in the example (the template project) used `demt_probe` which was the reason of blocking the pins, I'm not really sure, but that's what I saw, so...

### Getting into action

Now before we start we need to make sure that `cargo run` flashes into the Pico directly via a UF2 image, to do that, hop into `.cargo/config.toml`, and edit the following:

```bash
# comment this.
# runner = "probe-rs run --chip RP2040 --protocol swd"
# uncomment this.
runner = "elf2uf2-rs -d"
```

Now add [defmt-serial](https://docs.rs/defmt-serial/0.7.0/defmt_serial) to the project, since, this was the working thing (with serial)

```bash
cargo add defmt-serial
```

And [fugit](https://docs.rs/fugit/latest/fugit/) to use the Hz value for the UART initialization.

```bash
cargo add fugit --features=defmt
```

Finally update the `rp-pico` hal, to use the `into_funtion` method on a pin, and, well, to stay updated...

```bash
cargo update --package rp-pico
```

Now let's get to the code, this is a stop watch code, I should do a stop watch with an actual display and buttons, but again that's a story for another day.

```rust
#![no_std]
#![no_main]

use core::fmt::Write;

use bsp::entry;
// this is the change needed to utilize the serial defmt.
// use defmt_rtt as _;
use defmt_serial as _;
use panic_probe as _;

use rp_pico as bsp;

use bsp::hal::{
    clocks::{init_clocks_and_plls, Clock},
    pac,
    sio::Sio,
    // of course we need the uart module from the hal.
    uart::*,
    Watchdog,
};

use fugit::RateExtU32;

#[entry]
fn main() -> ! {
    let mut pac = pac::Peripherals::take().unwrap();
    let core = pac::CorePeripherals::take().unwrap();
    let mut watchdog = Watchdog::new(pac.WATCHDOG);
    let sio = Sio::new(pac.SIO);

    let external_xtal_freq_hz = 12_000_000u32;
    let clocks = init_clocks_and_plls(
        external_xtal_freq_hz,
        pac.XOSC,
        pac.CLOCKS,
        pac.PLL_SYS,
        pac.PLL_USB,
        &mut pac.RESETS,
        &mut watchdog,
    )
    .ok()
    .unwrap();

    let mut delay = cortex_m::delay::Delay::new(core.SYST, clocks.system_clock.freq().to_Hz());

    let pins = bsp::Pins::new(
        pac.IO_BANK0,
        pac.PADS_BANK0,
        sio.gpio_bank0,
        &mut pac.RESETS,
    );

    // uart declaration
    let mut uart = bsp::hal::uart::UartPeripheral::new(
        // using the first UART channel (pins 0 and 1)
        pac.UART0,
        // pins allocation for UART
        (pins.gpio0.into_function(), pins.gpio1.into_function()),
        &mut pac.RESETS,
    )
    .enable(
        // these configs we'll be using on the serial receiver.
        UartConfig::new(9600.Hz(), DataBits::Eight, None, StopBits::One),
        clocks.peripheral_clock.freq(),
    )
    .unwrap();

    // a simple stop watch.
    let mut seconds = 0;
    uart.write_raw(b"Timer started:").unwrap();
    loop {
        uart.write_fmt(format_args!("spent {} seconds", seconds))
            .unwrap();
        delay.delay_ms(1000);
        seconds += 1;
    }
}
```

Run the code with

```bash
cargo run --release # release is to reduce the binary's size
```

### Receiving the serial messages

1. Install [minicom](https://linux.die.net/man/1/minicom) via your package manager.
   - example on Gentoo
   ```bash
   sudo emerge -qav net-dialup/minicom
   ```
2. Add your user to the `uucp` group, so that you can use the serial devices.
   ```bash
   sudo gpasswd -a $USER uucp
   ```
3. Run minicom with the specified configurations above
   ```bash
   minicom -b 9600 -o -D /dev/ttyUSB0
   ```
   Where:
   - `-b` is for baudrate which was **9600** in the sender UART.
   - `-o` no initialization for the serial receiver on startup (to use the provided configurations only)
   - `-D` is the device to use
4. If you're on Windows or Mac you can easily Google the steps above.

And the drama is over, thanks for reading till the end.
