package main

import (
	"embed"
	"fmt"
	"io"
	"mbaraacom/config"
	"mbaraacom/db"
	"mbaraacom/log"
	"mbaraacom/tmplrndr"
	"net/http"
	"time"
)

var (
	//go:embed resources/*
	res embed.FS
)

func main() {
	//	insertOldData()

	http.HandleFunc("/", handleHomePage)
	http.HandleFunc("/projects", handleProjectsPage)
	http.HandleFunc("/xp", handleXpPage)
	http.HandleFunc("/about", handleAboutPage)
	http.HandleFunc("/blog", handleBlogsPage)
	http.HandleFunc("/blog/", handleBlogPostPage)

	http.Handle("/resources/", http.FileServer(http.FS(res)))
	log.Infoln("server started at port 3000")
	log.Fatalln(string(log.ErrorLevel), http.ListenAndServe(":"+config.Config().Port, nil))
}

func insertOldData() {
	// info
	err := db.InsertInfo(db.Info{
		FullAbout:  "Hello there! My name is Baraa Al-Masri and I like building software, most likely web apps, some of the stuff I build are opensource, here's a list of my [projects](/projects) that saw the light of day.\n\nMy interst in computers started when I was a kid when my computer broke down and I was like, hell I should be able to fix that, and so I've completed this dark path, where I was to fix computers and phones of my family's.\n\nI started my Linux journey when I was 13 where we had a computer course, and in the OS section Ubuntu was mentioned, I thought well I should try that on my computer, one thing led to another and I use Gentoo now :)\n\nI'm currently studying Software Engineering in my collage, because I thought collage would give me what I wanted, but it wasn't enough for me so I started building useless daily apps, till I found  myself in web development, and so I stayed in that area.\n\nAfter trying different languages and frameworks, two of them really hooked me in, Go was the best programming language I've ever used concidering it's simplicity, cleanliness, and rich standard library, so I kept using it for backend and system applications, I've tried a lot of frontend frameworks like, Go templates, Vue.js, React.js, Next.js and the GOAT SvelteKit which became my favourite frontend framework because of its clean and light architecture for an SSR frontend framework.\n\nI started my professional work at ProgressSoft as a Java SpringBoot developer, before that I used to build web apps for my university students, and I used to work as a freelancer since I was 2nd year in collage, most of the stuff I made were proprietary software, so here's that.",
		BriefAbout: "I'm a software developer specializing in web development in various stacks, and a fresh embedded rustacean 🦀　\nI pay rent by writing TypeScript full stack web apps @ Jordan Open Source Association.　\nAnd in my free time I write more code, blog, and slack watching YT shorts.",
		BlogIntro:  "This is a place where I vent, write solutions for crappy experiences I make, and other stuff...",
		Technologies: []string{
			"Go",
			"Rust",
			"TypeScript",
			"Java",
			"Elixir",
			"Bash",
			"SvelteKit",
			"Nuxt",
			"Yew",
			"SpringBoot",
			"Phoenix",
			"Linux",
			"Google Cloud Platform",
			"Embedded Programming",
			"Vim",
			"LaTex",
			"Git",
			"Nginx",
		},
	})
	if err != nil {
		log.Errorln(err)
	}

	// blogs
	parseTime := func(s string) int64 {
		parsed, _ := time.Parse("2006-01-02T15:04:05", s)
		return parsed.Unix()
	}
	blogs := []db.Blog{
		{
			WrittenAt:   parseTime("2023-12-09T12:47:31"),
			Title:       "Fixing NVMe SSD Problems on Linux",
			Description: "A quick fix before you jump to another distro or go back to Windows.",
			PublicId:    "fixing-nvme-ssd-problems-on-linux",
			Content:     "### The Problem\nMost NVMe SSDs have a power saving mode called APST (Autonomous Power State Transition) in which if sat properly the drive will suspend at some periods of time to save power, while the computer is on, on suspend and hibernate it turns off to save more power, but that's not the case.\n\nThe thing is, that some drives are programmed to ignore this power switching, and with the default kernel parameter of 100,000µs i.e 0.1s, that is the drive's timeout before going back from power saving mode, in which causes the drive to stay in the power saving mode, and freezes, which is bad, and can occur after a very short time from boot up (if there weren't much IO operations)\n\n\\\nMore details in [this](https://unix.stackexchange.com/questions/612096/clarifying-nvme-apst-problems-for-linux) stack exchange post.\n\n\n### The Solution\nNothing fancy, it's just a work around to disable APST, in which the drive stays on as long as the OS says so (i.e not suspended or hibernated)\n\\\nAll what we have to do is set that parameter to 0, which according to the NVMe driver in the Linux kernel, disables the feature.\n\n\\\nIn the file `/etc/default/grub` (you need root privileges to modify it) add the following to he `GRUB_CMDLINE_LINUX_DEFAULT` variable\n\n\n```bash\nGRUB_CMDLINE_LINUX_DEFAULT=\"nvme_core.default_ps_max_latency_us=0\"\n```\n\nIf the variable itself have any existing values, just add a space after the last value and add the magical parameter setter, e.g\n\n```bash\nGRUB_CMDLINE_LINUX_DEFAULT=\"snd_hda_intel.dmic_detect=0 nvme_core.default_ps_max_latency_us=0\"\n# the snd_hda_intel.dmic_detect=0 is an example don't add it unless you know exactly what you're doing\n```\n\n\\\nNow save the file, and update GRUB\n\n```bash\nsudo grub-mkconfig -o /boot/grub/grub.cfg\n```\n\nSome distros name the grub binary as **grub2** like Fedora, or openSUSE, in that execute this command instead\n\n```bash\nsudo grub2-mkconfig -o /boot/grub/grub.cfg\n```\n\\\nFinally reboot and your NVMe should be back on track, grinding those 1.5GBps IO speeds.\n\nTo double check (a typo could cause this not to work), check the current value of the latency parameter\n```bash\ncat /sys/module/nvme_core/parameters/default_ps_max_latency_us\n```\n\nIf it prints a value other than 0, double check `/etc/default/grub`, or regenerate the GRUB config again.\n\n\\\nNOTE: this was the configuration for [GRUB](https://wiki.gentoo.org/wiki/GRUB), i.e if you have [LILO](https://wiki.gentoo.org/wiki/LILO) or [Systemd-boot](https://wiki.archlinux.org/title/Systemd-boot), or any bootloader other than GRUB you might to look up how to add a kernel parameter for that particular bootloader, other than that it's pretty straight forward.\n\n### Quote of the day\n\n\"Success lasts until someone screws them, failures are forever.”\n\\\n\\- [Gregory House](https://en.wikipedia.org/wiki/Gregory_House)",
			VisitTimes:  205,
		},
		{
			WrittenAt:   parseTime("2023-11-10T20:13:18"),
			Title:       "Using a Matrix Keypad with Rust and the Raspberry Pi Pico",
			Description: "A quick tour to understanding I/O pins on a Raspberry Pi Pico with Rust and using a matrix keypad.",
			PublicId:    "using-a-matrix-keypad-with-rust-and-the-raspberry-pi-pico",
			Content:     "I'm using [this](https://www.amazon.com/Matrix-Membrane-Switch-Keyboard-Arduino/dp/B07THCLGCZ/ref=sr_1_1?crid=GK1F31O71PSP&keywords=matrix+keypad&qid=1699285904&sprefix=matrix+keypad%2Caps%2C234&sr=8-1) matrix keypad, but any will do.\n\n### Input pin types\n\nInput pins have different types, it's called resistive pull, where you set a GPIO to an opposite state of what you want to read from the pin, and the initial state helps us to expect what to read from the pin, e.g  in pull-up the pin reads high by default, and until provided with a low (below 1.8v) it'll stay in high state.\n\nStates are identified as:\n- 3.3v (or 5 in bigger boards) and ground\n- 1 and 0\n- High and Low\n- On and off\n\nIt's not really from a holy book what you call them, as long as you distinguish the difference between the two states.\n\nSo there are 3 types of initial pin states (yes 3)\n1. Pull-up\n\tThis setup, sets the pin's voltage to 3.3v (or 5 in bigger boards, but with the pico it's always 3.3v) and reads high, i.e if we measured the voltage between it and a ground pin it'll always read 3.3v, well, until a low (ground) voltage is provided by an output pin, this is helpful in the use case of the keypad, since some pins needs to be inputs and the others needs to be output.\n\n\tAnd it can be declared like this in Rust using GPIO6 as input:\n\t```rust\n\tlet some_pull_up_input_pin = pins.gpio6.into_pull_up_input();\n\t```\n\n2. Pull-down\n\tThe inverse of pull-up where a pin is initially at a low state, i.e reads 0v when measured against ground, and can be provided with a high state to read the difference.\n\t\n\tAnd it can be declared like this in Rust using GPIO7 as input:\n\t```rust\n\tlet some_pull_down_input_pin = pins.gpio7.into_pull_down_input();\n\t```\n\t\n3. Floating state\n\tIt's a bit uncertain like Schrodinger's cat or grandma's Schrodinger's plates\n    ![Shrodinger's Plates](https://mbaraa.com/img/7545_shrodingers_plates.jpg)\n\t<<image pulled from duckduckgo, I really tried to find the source>>\n\t\n\\\n\tAnd this type of setup is used for push buttons, or switches, where it has no definite initial state, and the detected state can be determined by the provided input, i.e if the provided is high and if we check for high it'll be high, and same for low.\n\t\ne.g: using a push button to power on an LED from [this](https://mbaraa.com/blog/running-rust-on-raspberry-pi-pico) post\n\t\n```rust\n// imports\nuse embedded_hal::digital::v2::{InputPin, OutputPin};\n// end of imports\n\nlet mut output_led = pins.gpio6.into_push_pull_output();\nlet input_pin = pins.gpio7.into_floating_input();\n\nloop {\n\tif input_pin.is_low().unwrap() {\n\t\toutput_led.set_high().unwrap();\n\t} else {\n\t\toutput_led.set_low().unwrap();\n\t}\n}\n```\n\n### Matrix keypad wiring\n\nThe matrix keypad  has a clever wiring which utilizes a matrix with its wiring, hence the name, here's a details image of the thing.\n\n![Matrix Keypad 4x4](https://mbaraa.com/img/2502_matrix_keypad.jpg)\n\n\\\nNow what we're gonna do is wire the rows to output pins, and the columns to pull-down input pins, then send high voltage (3.3v) from each row, and whatever pin receives the voltage will tell us what button was exactly pressed.\n\nWhat's gonna happen, is that for each row we'll set its pin to output high voltage, and check for all columns for a high voltage, where the column with the high voltage is the pressed key.\n\nSo, at the end we'll do something like the following, but on a bigger scale\n```rust\nlet mut row1 = pins.gpio2.into_push_pull_output();\nlet mut col1 = pins.gpio6.into_pull_down_input();\nloop {\n\trow1.set_high();\n\tif col1.is_high().unwrap() {\n\t\t// do something\n\t}\n\trow1.set_low();\n}\n```\n\n### Reading key presses from a keypad\n\n#### Quick setup\n\n1. Generate a project from [rp-rs](github.com/rp-rs/)'s template using [cargo generate](https://github.com/cargo-generate/cargo-generate)\n\t```bash\n\tcargo generate --git https://github.com/rp-rs/rp2040-project-template\n\t```\n2.  Update the `rp-pico` hal (make sure its version >= 0.8)\n\t```bash\n\tcargo update --package rp-pico\n\t```\n3. (skip if you have a probe) UART setup more details [here](https://mbaraa.com/blog/debugging-rust-on-the-raspberry-pi-pico)\n\t- Install [fugit](https://docs.rs/fugit/latest/fugit) and [defmt-serial](https://docs.rs/defmt-serial/0.7.0/defmt_serial)\n\t```bash\n\tcargo add fugit --features=defmt\n\tcargo add defmt-serial \n\t```\n\t- Update `.cargo/config.toml` to use elf2uf2 flasher\n\t\t```bash\n\t\t#runner = \"probe-rs run --chip RP2040 --protocol swd\"\n\t\trunner = \"elf2uf2-rs -d\"\n\t\t```\n\t-  Initialize the UART object and use `defmt_serial` instead of `defmt_rtt`\n\t```rust\n\t// ...\n\t// update defmt imports\n\t// use defmt_rtt as _;\n\tuse defmt_serial as _;\n\t// ...\n\t// add this to the import list\n\tuse bsp::hal::{\n\t  clocks::{init_clocks_and_plls, Clock},\n\t  pac,\n\t  sio::Sio,\n\t  uart::*, // the new module\n\t  watchdog::Watchdog,\n\t};\n\t// ...\n\t// import fugit's u32 traits \n\tuse fugit::RateExtU32;\n\t// ...\n\t// initialize the URAT handler\n\tlet mut uart = bsp::hal::uart::UartPeripheral::new(\n        pac.UART0,\n        (pins.gpio0.into_function(), pins.gpio1.into_function()),\n        &mut pac.RESETS,\n    )\n    .enable(\n        UartConfig::new(9600.Hz(), DataBits::Eight, None, StopBits::One),\n        clocks.peripheral_clock.freq(),\n    )\n    .unwrap();\n\t// ...\n\t```\n\n#### Wiring\n\n![keypad wiring](https://mbaraa.com:/img/4581_keypad_wiring.jpg)\n\n\\\nAs shown in the image, I'm using GPIOs [2-5] as rows (inputs), and GPIOs [6-9] as columns (outputs), and the code below will clarify the wiring even more.\n\n#### The code amalgamation\n\n```rust\n#![no_std]\n#![no_main]\n\n// to use write_fmt\nuse core::fmt::Write;\n\nuse bsp::entry;\nuse defmt_serial as _;\nuse embedded_hal::digital::v2::{InputPin, OutputPin};\nuse panic_probe as _;\n\nuse rp_pico as bsp;\n\nuse bsp::hal::{\n    clocks::{init_clocks_and_plls, Clock},\n    // just types for the passed rows and columns to the get_pressed_key function\n    gpio::{DynPinId, FunctionSio, Pin, PullDown, SioInput, SioOutput},\n    pac,\n    sio::Sio,\n    uart::*,\n    Watchdog,\n};\n\nuse fugit::RateExtU32;\n\n#[entry]\nfn main() -> ! {\n    let mut pac = pac::Peripherals::take().unwrap();\n    let core = pac::CorePeripherals::take().unwrap();\n    let mut watchdog = Watchdog::new(pac.WATCHDOG);\n    let sio = Sio::new(pac.SIO);\n\n    let external_xtal_freq_hz = 12_000_000u32;\n    let clocks = init_clocks_and_plls(\n        external_xtal_freq_hz,\n        pac.XOSC,\n        pac.CLOCKS,\n        pac.PLL_SYS,\n        pac.PLL_USB,\n        &mut pac.RESETS,\n        &mut watchdog,\n    )\n    .ok()\n    .unwrap();\n\n    let mut delay = cortex_m::delay::Delay::new(core.SYST, clocks.system_clock.freq().to_Hz());\n\n    let pins = bsp::Pins::new(\n        pac.IO_BANK0,\n        pac.PADS_BANK0,\n        sio.gpio_bank0,\n        &mut pac.RESETS,\n    );\n\n    let mut uart = bsp::hal::uart::UartPeripheral::new(\n        pac.UART0,\n        (pins.gpio0.into_function(), pins.gpio1.into_function()),\n        &mut pac.RESETS,\n    )\n    .enable(\n        UartConfig::new(9600.Hz(), DataBits::Eight, None, StopBits::One),\n        clocks.peripheral_clock.freq(),\n    )\n    .unwrap();\n\n    let mut row1 = pins.gpio2.into_push_pull_output().into_dyn_pin(); // dyn pins allow us to\n                                                                      // specifiy its type for the\n                                                                      // get_pressed_key function\n    let mut row2 = pins.gpio3.into_push_pull_output().into_dyn_pin();\n    let mut row3 = pins.gpio4.into_push_pull_output().into_dyn_pin();\n    let mut row4 = pins.gpio5.into_push_pull_output().into_dyn_pin();\n\n    let mut col1 = pins.gpio6.into_pull_down_input().into_dyn_pin();\n    let mut col2 = pins.gpio7.into_pull_down_input().into_dyn_pin();\n    let mut col3 = pins.gpio8.into_pull_down_input().into_dyn_pin();\n    let mut col4 = pins.gpio9.into_pull_down_input().into_dyn_pin();\n\n    loop {\n        match get_pressed_key(\n            [&mut row1, &mut row2, &mut row3, &mut row4],\n            [&mut col1, &mut col2, &mut col3, &mut col4],\n        ) {\n            Some((row, col)) => {\n                uart.write_fmt(format_args!(\"pressed {:?} {:?}\\r\\n\", row, col))\n                    .unwrap();\n            }\n            None => {}\n        }\n    }\n}\n\n// the star of the show, where it lights up the rows and checks if there's a receiver column, \n// of course the arrays' size can be changed if you have a different keypad.\nfn get_pressed_key(\n\t// I just checked the pins' types and slapped them over here.\n    rows: [&mut Pin<DynPinId, FunctionSio<SioOutput>, PullDown>; 4],\n    cols: [&mut Pin<DynPinId, FunctionSio<SioInput>, PullDown>; 4],\n) -> Option<(RowOrder, ColOrder)> // fancy Rust option enum \n\t\t\t\t\t\t\t\t  // bla bla bla\n{\n    for i in 0..rows.len() {\n        rows[i].set_high().unwrap();\n        for j in 0..cols.len() {\n            if cols[j].is_high().unwrap() {\n                rows[i].set_low().unwrap();\n                return Some((RowOrder::get_order(i as u8), ColOrder::get_order(j as u8)));\n            }\n        }\n        rows[i].set_low().unwrap();\n    }\n    None\n}\n\n// some enums declaration, and implementations.\n\n#[derive(Debug)]\nenum RowOrder {\n    Row1,\n    Row2,\n    Row3,\n    Row4,\n    None,\n}\n\nimpl RowOrder {\n    fn get_order(index: u8) -> RowOrder {\n        match index {\n            0 => RowOrder::Row1,\n            1 => RowOrder::Row2,\n            2 => RowOrder::Row3,\n            3 => RowOrder::Row4,\n            _ => RowOrder::None,\n        }\n    }\n}\n\n#[derive(Debug)]\nenum ColOrder {\n    Col1,\n    Col2,\n    Col3,\n    Col4,\n    None,\n}\n\nimpl ColOrder {\n    fn get_order(index: u8) -> ColOrder {\n        match index {\n            0 => ColOrder::Col1,\n            1 => ColOrder::Col2,\n            2 => ColOrder::Col3,\n            3 => ColOrder::Col4,\n            _ => ColOrder::None,\n        }\n    }\n}\n```\n\n### Assigning actions to a key's press\n\nInstead of just printing the pressed key's location, add an action for it to do, I would've demonstrated with LEDs, but I don't have enough to to that, and well that'll be it.\n\n\n### Quote of the day\nIt's more like quote of the time I wrote the blog post, I'll keep doing this as long as I remember.\n\\\n“Smart Data Structures and dumb code works a lot better than the other way \naround.”\n\\\n\\- [Eric S. Raymond](https://en.wikipedia.org/wiki/Eric_S._Raymond)\n",
			VisitTimes:  246,
		},
		{
			WrittenAt:   parseTime("2023-11-05T11:46:56"),
			Title:       "Debugging Rust on the Raspberry Pi Pico",
			PublicId:    "debugging-rust-on-the-raspberry-pi-pico",
			Description: "Logging the Pico's output to a UART serial monitor without using a probe",
			Content:     "If you have a probe get out, since there's an [official](https://www.raspberrypi.com/documentation/microcontrollers/debug-probe.html) documentation about it :)\n\nNow for debugging, I know I made some drama about it in my previous [post](https://mbaraa.com/blog/running-rust-on-raspberry-pi-pico#debugging), but I just got the Pico, and I was used to the Arduino's serial thingy built-in to the IDE, so it wasn't that much to think about, I haven't used C or MicroPython with the Pico, so I don't really know what's the debugging deal with them, but with Rust it was a bit of a hassle.\n\n### Preface\n\nSo my first solution was the one in the [rp-hal's docs](https://docs.rs/rp2040-hal/latest/rp2040_hal/uart/index.html), it was working and I could catch the serial signal with a UART or the Pico's serial, but it was blocking the other pins (or at least that's what I thought), so I couldn't use the other pins for basic GPIO stuff, and that led me to search the dark web (not really but I did search for a solution for like 4 hours).\n\nThere's [this](https://www.reddit.com/r/rust/comments/14atkm3/media_debug_pi_pico_using_raspberry_pi4/) awesome post by someone on Reddit (it's always someone on Reddit who posts the real deal), so [u/ThatBrokeDave](https://www.reddit.com/user/ThatBlokeDave), special thanks for you if you ever came across this blog 🫡\n\nBut the problem was from [defmt](https://docs.rs/defmt/latest/defmt) where the code in the example (the template project) used `demt_probe` which was the reason of blocking the pins, I'm not really sure, but that's what I saw, so...\n\n### Getting into action\n\nNow before we start we need to make sure that `cargo run` flashes into the Pico directly via a UF2 image, to do that, hop into `.cargo/config.toml`, and edit the following:\n\n```bash\n# comment this.\n# runner = \"probe-rs run --chip RP2040 --protocol swd\"\n# uncomment this.\nrunner = \"elf2uf2-rs -d\"\n```\n\nNow add [defmt-serial](https://docs.rs/defmt-serial/0.7.0/defmt_serial) to the project, since, this was the working thing (with serial)\n```bash\ncargo add defmt-serial \n```\n\nAnd [fugit](https://docs.rs/fugit/latest/fugit/) to use the Hz value for the UART initialization.\n```bash\ncargo add fugit --features=defmt\n```\n\nFinally update the `rp-pico` hal, to use the `into_funtion` method on a pin, and, well, to stay updated...\n```bash\ncargo update --package rp-pico\n```\n\nNow let's get to the code, this is a stop watch code, I should do a stop watch with an actual display and buttons, but again that's a story for another day.\n\n```rust\n#![no_std]\n#![no_main]\n\nuse core::fmt::Write;\n\nuse bsp::entry;\n// this is the change needed to utilize the serial defmt.\n// use defmt_rtt as _;\nuse defmt_serial as _;\nuse panic_probe as _;\n\nuse rp_pico as bsp;\n\nuse bsp::hal::{\n    clocks::{init_clocks_and_plls, Clock},\n    pac,\n    sio::Sio,\n    // of course we need the uart module from the hal.\n    uart::*,\n    Watchdog,\n};\n\nuse fugit::RateExtU32;\n\n#[entry]\nfn main() -> ! {\n    let mut pac = pac::Peripherals::take().unwrap();\n    let core = pac::CorePeripherals::take().unwrap();\n    let mut watchdog = Watchdog::new(pac.WATCHDOG);\n    let sio = Sio::new(pac.SIO);\n\n    let external_xtal_freq_hz = 12_000_000u32;\n    let clocks = init_clocks_and_plls(\n        external_xtal_freq_hz,\n        pac.XOSC,\n        pac.CLOCKS,\n        pac.PLL_SYS,\n        pac.PLL_USB,\n        &mut pac.RESETS,\n        &mut watchdog,\n    )\n    .ok()\n    .unwrap();\n\n    let mut delay = cortex_m::delay::Delay::new(core.SYST, clocks.system_clock.freq().to_Hz());\n\n    let pins = bsp::Pins::new(\n        pac.IO_BANK0,\n        pac.PADS_BANK0,\n        sio.gpio_bank0,\n        &mut pac.RESETS,\n    );\n\n    // uart declaration\n    let mut uart = bsp::hal::uart::UartPeripheral::new(\n        // using the first UART channel (pins 0 and 1)\n        pac.UART0,\n        // pins allocation for UART\n        (pins.gpio0.into_function(), pins.gpio1.into_function()),\n        &mut pac.RESETS,\n    )\n    .enable(\n        // these configs we'll be using on the serial receiver.\n        UartConfig::new(9600.Hz(), DataBits::Eight, None, StopBits::One),\n        clocks.peripheral_clock.freq(),\n    )\n    .unwrap();\n\n    // a simple stop watch.\n    let mut seconds = 0;\n    uart.write_raw(b\"Timer started:\").unwrap();\n    loop {\n        uart.write_fmt(format_args!(\"spent {} seconds\", seconds))\n            .unwrap();\n        delay.delay_ms(1000);\n        seconds += 1;\n    }\n}\n``` \n\nRun the code with\n```bash\ncargo run --release # release is to reduce the binary's size\n```\n\n### Receiving the serial messages\n1. Install [minicom](https://linux.die.net/man/1/minicom) via your package manager.\n\t- example on Gentoo\n\t```bash\n\tsudo emerge -qav net-dialup/minicom\n\t```  \n2. Add your user to the `uucp` group, so that you can use the serial devices.\n\t```bash\n\tsudo gpasswd -a $USER uucp\n\t```\n3. Run minicom with the specified configurations above\n\t```bash\n\tminicom -b 9600 -o -D /dev/ttyUSB0\n\t```\n\tWhere:\n\t- `-b` is for baudrate which was **9600** in the sender UART.\n\t- `-o` no initialization for the serial receiver on startup (to use the provided configurations only)\n\t- `-D`  is the device to use\n4. If you're on Windows or Mac you can easily Google the steps above.\n\nAnd the drama is over, thanks for reading till the end.\n",
			VisitTimes:  308,
		},
		{
			WrittenAt:   parseTime("2023-11-04T00:57:24"),
			Description: "Because why not?",
			PublicId:    "running-rust-on-raspberry-pi-pico",
			Title:       "Running Rust on the Raspberry Pi Pico",
			Content:     "[Rust](https://www.rust-lang.org) is a powerful and very loved language, and the [Raspberry Pi Pico](https://www.raspberrypi.com/products/raspberry-pi-pico) is a beefy [Arduino Nano](https://store.arduino.cc/products/arduino-nano) alternative (same form factor-ish)\n\n\nDifferences between the two boards:\n\n| [_](https://www.youtube.com/watch?v=dQw4w9WgXcQ&pp=ygUJcmljayByb2xs) | Pi Pico | Arduino Nano |\n| --- | --- | --- |\n| Micro controller | RP2040 (122MHz) | ATmega328 (16MHz) |\n| RAM | 264KB | 2KB |\n| ROM | 134MB (they say it's 2MB but it appears as 134MB 🤷‍♂️) | 32KB |\n| Connectivity | USB and UART | USB and UART |\n| Power | 1.8-5.5V (16-43mAh) | 5-12V (19mAh) |\n| Digital I/O Pins | 26 (16 are PWM) | 22 (6 are PWM) |\n| Analog In Pin | 3 | 8 |\n| Clock | Yes | No |\n| Wifi | Yes (W version with BT5.2) | No |\n| Thermal Sensor | Yes | No (what a shame) |\n| Is it cool? | Yes | Yes, but in blue |\n\n### Installing Rust\n\nSo since Rust is mainly used for low level programming, and the low level thing, and there are plenty of HALs (Hardware Abstraction Layer) written in Rust and are ready to use with the Pico, e.g [rp-hal](https://github.com/rp-rs/rp-hal), for more details [rp-hal docs](https://docs.rs/rp2040-hal/latest/rp2040_hal) which what we'll use with the pico.\n\n\nFirst install Rust using [rustup](https://rustup.rs)\n```bash\ncurl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh\n```\nAnd then install helper crates:\n\n```bash\n# add the arm architecture, that will be used when compiling code to the pico.\nrustup target add thumbv6m-none-eabi\n# to create uf2 images for the RP2040 (without the probe)\ncargo install elf2uf2-rs --locked\n# cargo-generate is used to scaffold a project from a git repository \n# https://github.com/cargo-generate/cargo-generate\ncargo install cargo-generate\n```\n\n### Creating a project\n\nNow where Rust kicks in, first create a project using rp-hal's template (it has all the wanted configurations to be done)\n\n```bash\n# generate an example project\ncargo generate --git https://github.com/rp-rs/rp2040-project-template\n```\n\n![Cargo Generate Enter Name](https://mbaraa.com/img/7932_cargo_generate_enter_name.jpg)\n\nYou need to specify the project's name here, in my case I'll name it `pico-test`\n\nIgnore everything for now, connect the pico (in download mode) and run the project:\n\n```bash\ncargo run\n```\n\nthis will build the project and flash it into the pico, you'll see the LED blinking, horray you're a Rust developer now 🎉 (there's no going back)\n\nOk now let's dive in a bit into the configurations and code\n\n![Project Structure](https://mbaraa.com/img/6883_project_structure.jpg)\n\nStarting from the bottom (ignoring the markdown files), we have the `memory.x`, which describes the physical locations of the bootloader, storage, and RAM, DON'T CHANGE ANYTHING, I wanna try changing the the storage's location (since it's actually bigger, but not now) and it looks like this\n```\n/* DON'T CHANGE ANYTHING */\nMEMORY {\n    BOOT2 : ORIGIN = 0x10000000, LENGTH = 0x100\n    FLASH : ORIGIN = 0x10000100, LENGTH = 2048K - 0x100\n    RAM   : ORIGIN = 0x20000000, LENGTH = 256K\n}\n\nEXTERN(BOOT2_FIRMWARE)\n\nSECTIONS {\n    /* ### Boot loader */\n    .boot2 ORIGIN(BOOT2) :\n    {\n        KEEP(*(.boot2));\n    } > BOOT2\n} INSERT BEFORE .text;\n```\n\nAnd this file's values are used in the `build.rs` when building the [UF2](https://microsoft.github.io/uf2) image.\n\n\nThen there is the `Embed.toml`which specifies the probe's and building options (trust me mate)\n\n\n`Cargo.toml` obviously specifies the package details, the required dependencies, and other stuff, if you want to know more about it go [here](https://doc.rust-lang.org/cargo/reference/manifest.html).\n\n\n`build.rs` is where the fun begins, where it converts the output binary elf to a UF2 image that can be transmitted to the pico.\n\n\n`src/main.rs` this is where the magic happens, where the Rusty Rust code relies, and if you're not a fan of Rust's bloat, go back to C or MicroPython, I'll just slap the code here and comment some stuff (there are some comments from the template itself)\n\n```rust\n// ignore the standard library, only uses Rust's core library.\n#![no_std]\n// tells the compiler to shut up about the non existing main funtion that returns void.\n#![no_main]\n\n// entry point macro, used to specify the main function.\nuse bsp::entry;\n// the logger, oh boy this will be a mess later.\nuse defmt::*;\nuse defmt_rtt as _;\n// pin control trait (set_low, set_high).\nuse embedded_hal::digital::v2::OutputPin;\n// halts the pico when an error happens.\nuse panic_probe as _;\n\n// Provide an alias for our BSP so we can switch targets quickly.\n// Uncomment the BSP you included in Cargo.toml, the rest of the code does not need to change.\nuse rp_pico as bsp;\n\n// board support package\nuse bsp::hal::{\n\t// for delays and clock related stuff.\n    clocks::{init_clocks_and_plls, Clock},\n    // peripherals access crate, and that's all you need to know.\n    pac,\n    // no idea what it does, so...\n    sio::Sio,\n    // the dog that bites if there is no activity.  \n    watchdog::Watchdog,\n};\n\n#[entry]\nfn doesnt_have_to_be_main() -> ! {\n    info!(\"Program start\"); // this doesn't work without a probe, go to the end.\n    // pins takeover.\n    let mut pac = pac::Peripherals::take().unwrap();\n    let core = pac::CorePeripherals::take().unwrap();\n    // watchdog declaration.\n    let mut watchdog = Watchdog::new(pac.WATCHDOG);\n    // that's a story for another day.\n    let sio = Sio::new(pac.SIO);\n\n    // External high-speed crystal on the pico board is 12Mhz\n    let external_xtal_freq_hz = 12_000_000u32;\n    // clock setup, just know that it's clock thingy.\n    let clocks = init_clocks_and_plls(\n        external_xtal_freq_hz,\n        pac.XOSC,\n        pac.CLOCKS,\n        pac.PLL_SYS,\n        pac.PLL_USB,\n        &mut pac.RESETS,\n        &mut watchdog,\n    )\n    .ok()\n    .unwrap();\n\n\t// the delay function.\n    let mut delay = cortex_m::delay::Delay::new(core.SYST, clocks.system_clock.freq().to_Hz());\n\n\t// pins control declation.\n    let pins = bsp::Pins::new(\n        pac.IO_BANK0,\n        pac.PADS_BANK0,\n        sio.gpio_bank0,\n        &mut pac.RESETS,\n    );\n\n    // This is the correct pin on the Raspberry Pico board. On other boards, even if they have an\n    // on-board LED, it might need to be changed.\n    // Notably, on the Pico W, the LED is not connected to any of the RP2040 GPIOs but to the cyw43 module instead. If you have\n    // a Pico W and want to toggle a LED with a simple GPIO output pin, you can connect an external\n    // LED to one of the GPIO pins, and reference that pin here.\n    let mut led_pin = pins.led.into_push_pull_output();\n\n\t// the event loop, that's why the main returns !, this is where you write the repetitive code.\n    loop {\n        info!(\"on!\"); // this doesn't work without a probe, go to the end.\n        led_pin.set_high().unwrap();\n        delay.delay_ms(500);\n        info!(\"off!\"); // this doesn't work without a probe, go to the end.\n        led_pin.set_low().unwrap();\n        delay.delay_ms(500);\n    }\n}\n```\n\n`.cargo/config.toml`  this little number contains configuration about the build options, it contains what libraries to link into the binary, the target architecture, and the runner options (when running cargo build or run).\n\n\nIf you don't have a probe do this:\n```bash\n# comment this.\n# runner = \"probe-rs run --chip RP2040 --protocol swd\"\n# uncomment this.\nrunner = \"elf2uf2-rs -d\"\n```\n\n### More Stuff (3 LEDs blinker)\n```rust\n#![no_std]\n#![no_main]\n\nuse bsp::entry;\nuse defmt::*;\nuse defmt_rtt as _;\nuse embedded_hal::digital::v2::OutputPin;\nuse panic_probe as _;\n\nuse rp_pico as bsp;\n\nuse bsp::hal::{\n    clocks::{init_clocks_and_plls, Clock},\n    pac,\n    sio::Sio,\n    watchdog::Watchdog,\n};\n\n#[entry]\nfn main() -> ! {\n    let mut pac = pac::Peripherals::take().unwrap();\n    let core = pac::CorePeripherals::take().unwrap();\n    let mut watchdog = Watchdog::new(pac.WATCHDOG);\n    let sio = Sio::new(pac.SIO);\n\n    let external_xtal_freq_hz = 12_000_000u32;\n    let clocks = init_clocks_and_plls(\n        external_xtal_freq_hz,\n        pac.XOSC,\n        pac.CLOCKS,\n        pac.PLL_SYS,\n        pac.PLL_USB,\n        &mut pac.RESETS,\n        &mut watchdog,\n    )\n    .ok()\n    .unwrap();\n\n    let mut delay = cortex_m::delay::Delay::new(core.SYST, clocks.system_clock.freq().to_Hz());\n\n    let pins = bsp::Pins::new(\n        pac.IO_BANK0,\n        pac.PADS_BANK0,\n        sio.gpio_bank0,\n        &mut pac.RESETS,\n    );\n\n\t// declare the wanted pins so we can use them in the event loop.\n    let mut gp6 = pins.gpio6.into_push_pull_output();\n    let mut gp7 = pins.gpio7.into_push_pull_output();\n    let mut gp8 = pins.gpio8.into_push_pull_output();\n\n    loop {\n        gp6.set_high().unwrap();\n        delay.delay_ms(500);\n        gp6.set_low().unwrap();\n        gp7.set_high().unwrap();\n        delay.delay_ms(500);\n        gp7.set_low().unwrap();\n        gp8.set_high().unwrap();\n        delay.delay_ms(500);\n        gp8.set_low().unwrap();\n    }\n}\n```\n\nFootage of the wiring:\n\n![3 LEDs Blinking Wiring](https://mbaraa.com/img/8637_3_leds_blinking_wiring.jpg)\n\nAs you can see I used a single resistor on the common ground of the LEDs' I had to be smart since I don't have much resistors 🤓\n\n### More Stuff (push down button)\n\n```rust\n#![no_std]\n#![no_main]\n\nuse bsp::entry;\nuse defmt::*;\nuse defmt_rtt as _;\nuse embedded_hal::digital::v2::{InputPin, OutputPin};\nuse panic_probe as _;\n\nuse rp_pico as bsp;\n\nuse bsp::hal::{pac, sio::Sio};\n\n#[entry]\nfn main() -> ! {\n    let mut pac = pac::Peripherals::take().unwrap();\n    let sio = Sio::new(pac.SIO);\n\n    let pins = bsp::Pins::new(\n        pac.IO_BANK0,\n        pac.PADS_BANK0,\n        sio.gpio_bank0,\n        &mut pac.RESETS,\n    );\n\n    let mut output_led = pins.gpio6.into_push_pull_output();\n    let input_pin = pins.gpio7.into_floating_input();\n\n    loop {\n        if input_pin.is_low().unwrap() {\n            output_led.set_high().unwrap();\n        } else {\n            output_led.set_low().unwrap();\n        }\n    }\n}\n```\n\nAnd as you can see there's no need for the clock, and the watchdog, since the events we're running are depending on each other, and there's no other funny business going on, so there's no need for them.\n\nFootage of the thing:\n![Push Button Off](https://mbaraa.com/img/2890_push_button_off.jpg)\n![Push Button Off](https://mbaraa.com/img/1111_push_button_on.jpg)\n\nI don't have a push button either, so I used the wires.\n\n### Debugging\n\n![Debugging Meme](https://mbaraa.com/img/4035_debugging_meme.jpg)\n\nWell, I wrote [this](https://mbaraa.com/blog/debugging-rust-on-the-raspberry-pi-pico) blog about debugging the Pico while using Rust.\n",
			VisitTimes:  291,
		},
		{
			WrittenAt:   parseTime("2023-10-07T18:59:58"),
			Title:       "Access-Control-Allow-Drama",
			Description: "A whole drama about enabling CORS in a bare-bone web library.",
			Content:     "First of all, just wanna point out the importance of reading a documentation before you start hacking around, and smacking your head against the wall.\n\nSo I was working on [Rex](https://github.com/mbaraa/rex) which is a GitHub action's server that deploys an application from a git repository after a push or merging a pull request (depending on your configuration of the action). So what I was doing is just fixing its CORS more specifically is the `Access-Control-Allow-Origin` header, so that only GitHub can make the request (to avoid excessive requests on my server) and here where the drama starts.\n\n\nI have written a lot of Go backends in [Fiber](https://gofiber.io) and how I handled CORS is by providing a comma separated list of allowed origins, e.g. `http://localhost:8080,https://mbaraa.com` more [here](https://docs.gofiber.io/api/middleware/cors#config), so what I thought is that this is how the header works, and I kept thinking that since I was still using Fiber, but for Rex it was a ~100 lines script, so using Fiber is a bloat, and I went with Go's `net/http`, but whenever I make a request I get the CORS error, that the requesting origin is not allowed, I thought that it needed a space after the commas, or before, or even both, but non of them worked, so after some digging, a friend of mine pointed out that I should check [this](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Origin) MDN post, more specifically the [allowed values part](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Origin#syntax).\n\n\nAfter trying back and forth for like an hour, it turned out that this `http://localhost:8080,https://mbaraa.com` is one of Fiber's syntactic sugars, and that they have handled it in a way that checks for the origin if it's in the allow list on each request [here](https://github.com/gofiber/fiber/blob/6ecd607d9717b3312e3bd0c2da5194bdba78ff00/middleware/cors/cors.go#L126).\n\n\nSo that's exactly what I did, I had an environmental variable containing the allow list, and checked for the origin on each request.\n\nI'll just slap Rex's `main.go` without the deployment stuff, and go from it\n```go\npackage main\n\nimport (\n\t\"flag\"\n\t\"net/http\"\n\t\"os\"\n\t\"regexp\"\n\t\"strings\"\n)\n\nvar (\n\tallowedOrigins    string\n\tallowedOriginsMap = map[string]bool{}\n)\n\nfunc main() {\n\tflag.StringVar(&allowedOrigins, \"allowed-origins\", os.Getenv(\"REX_ALLOWED_ORIGINS\"), \"give me a list of allowed origins\")\n\tparseAllowedOringins()\n\thttp.HandleFunc(\"/deploy/\", handleDeployRepo)\n\thttp.ListenAndServe(\":8080\", nil)\n}\n\nfunc parseAllowedOringins() {\n\tallowedOriginsList := strings.Split(\n\t\t// this regex is only to check if the comma has a trailing or a leading whitespace(s)\n\t\tregexp.MustCompile(`\\s*,\\s*`).ReplaceAllString(allowedOrigins, \",\"),\n\t\t\",\",\n\t)\n\tfor _, allowedOrigin := range allowedOriginsList {\n\t\tallowedOriginsMap[allowedOrigin] = true\n\t}\n}\n\nfunc handleDeployRepo(res http.ResponseWriter, req *http.Request) {\n\tres.Header().Set(\"Content-Type\", \"application/json; charset=UTF-8\")\n\tif origin := req.Header.Get(\"Origin\"); allowedOriginsMap[origin] || allowedOriginsMap[\"*\"] {\n\t\tres.Header().Set(\"Access-Control-Allow-Origin\", origin)\n\t}\n\tres.Write([]byte(\"ok\"))\n}\n```\n\nHere what's going on is that I'm taking the comma separated list of allowed origins, put them into a check map, so that the accessing time is constant or logarithmic, setting the origin value on each request (if it exits in the allow list), and continuing with the request happily.\n",
			PublicId:    "access-control-allow-drama",
			VisitTimes:  267,
		},
		{
			WrittenAt:   parseTime("2023-09-05T22:32:31"),
			Title:       "Setup FTP server as a home NAS",
			Description: "A quick guide to configure a Linux FTP server as a home (not limited to home) NAS ",
			PublicId:    "setup-ftp-server-as-a-home-nas",
			Content:     "NAS or Network Attached Storage is basically an [FTP](https://en.wikipedia.org/wiki/File_Transfer_Protocol) (File Transfer Protocol) server with some storage device attached to it, that is it can be accessible from another computer.\n\n[I know all that stuff, take me to the software setup](#tldr)\n\nNAS servers can be used for a some reasons like:\n1. Backups; this uses regular HDDs like [WD Blue](https://www.westerndigital.com/products/internal-drives/wd-blue-desktop-sata-hdd#WD5000AZLX) with a high capacity, as they won't be accessed that much, at most twice a week.\n2. Mass shared storage; this uses high demand HDDs like [WD Purple](https://www.westerndigital.com/products/internal-drives/wd-purple-sata-hdd) with a high capacity, where this type is used constantly as a shared mass storage across computers, that requires a durable drives, so that they won't break on constant reads and writes.\n3. High speed storage; this used an SSD so that a blazingly fast speed is achievable when accessing the drive, an SSD can be a regular SATA SSD for not that much of a speed needed, or a PCIe SSD for a super speedy needs and a very high flexible budget, this is mainly used to edit videos or to play video games directly on the NAS, which can be achieved given the SSDs super speed.\n\nEither one of the methods requires a high speed network switch or router, for example, an HDD setup requires at least a **1Gbps** network, since HDDs can have a speed up to **100MBps**, SSDs need at least **10Gbps**, and so on...\n\nA proper [RAID](https://en.wikipedia.org/wiki/RAID) setup is a plus, that is to improve performance (dedicates drives accesses) and for data redundancy (some RAIDs have copies of the drives) but more on RAID later, as of the time that I'm writing this I only have a 1TB and 500GB HDDs.\n\n#### My setup has the following hardware:\n1. Raspberry Pi 3 model b+\n2. [TP-Link UH720 USB Hub](https://www.tp-link.com/us/home-networking/usb-hub/uh720/)\n3. WD Blue 1TB HDD\n4. Some 500GB HDD (the label is lost, thank God it powers on) \n\nSide note: I'm not advertising any of the products, I just use them because they're good, and actually lasted with no faults for the past 4 years (expect for the USB hub that one is new, part of a renewal, stay tuned).\n\nNow for the power, the USB hub is a giga chad on its own, since it takes in **40W** of power which is enough to power everything from it, and that's exactly what I did, since it has a **2.4A** USB port, which is what the Raspberry Pi needs, since the HDDs are gonna be powered using the hub.\n\nHowever the Raspberry Pi needs a bit of power on startup beyond **2.4A** since booting up is a hard task and the little guy has too much to do...\n\nAnd a moment of appreciation for the hub's power switch, which is super handy (given that the Raspberry Pi doesn't have one by default)\n\nAnyways, I hooked the HDDs to the hub, the Raspberry Pi to the power port on the hub, the hub's data cable to one of the Raspberry Pi's USBs and the Raspberry Pi to the router using an RJ-45 (Ethernet) cable.\n\n## tldr\n\nJust a declaimer, I'm using Linux on both the client and server, but it's simple FTP, you can check the configuration for your OS on your own.\n\n### Configuring the FTP server\nThis one is easy, since it's just installing the FTP server software and doing a bit of [fstab](https://wiki.archlinux.org/title/Fstab) configuration.\n\n#### VSFTPD\n1. Install `vsftpd`, using your distro's package manager, it's included on all Linux distro's that I've tested on. but since I'm using a Raspberry Pi with Raspberry Pi OS, I'll just use `apt`\n```bash\nsudo apt install vsftpd\n```\n2. Do some configurations for `vsftpd`, by editing the file `/etc/vsftpd.conf` and un-commenting the lines\n```\nanonymous_enable=NO\nlocal_enable=YES\nwrite_enable=YES\nlocal_umask=022\nchroot_local_user=YES\n```\n3. And adding the following lines:\n```\nuser_sub_token=username\nlocal_root=/path/to/your/mounted/disks\n```\nIn my case these will be:\n```\nuser_sub_token=baraa\nlocal_root=/home/baraa/disks\n```\n4. Enable and start `vsftpd`'s service\n```bash\nsudo systemctl enable --now vsftpd\n```\n\n#### FSTAB\n1. List your connected drives' UUIDs using `sudo blkid`, you should expect output like this\n![sudo blkid output](https://mbaraa.com:/img/5287_undefined)\nThese values will differ on your device, as UUID is unique on each generation (at least for the next 100 years)\n2. As seen each partition has a `PARTUUID` value which will use to mount the drives (using UUID is essential so that we don't mix the drives)\n3. Create the mount point directories\n```\nmkdir -p ~/disks/disk{1,2}\n```\n4.  Enable auto mounting the drives using `fstab`, edit the file `/etc/fstab` and add the following:\n```bash\n# drive\tmount-point\tfile-system\tflags\tpriority\nPARTUUID=bf6011f2-01 /home/baraa/disks/disk1 ntfs nosuid,nodev,nofail 0 0\nPARTUUID=d9e4b195-6654-46fd-aad8-dc1d4f5d7302 /home/baraa/disks/disk2 ext4 nosuid,nodev,nofail 0 0\n```\n5. Make your modifications as it fits you and save the file.\n6. Mount the drives using the modified `fstab` configuration:\n```bash\nsudo mount -a\n```\n\n### Configuring the client\nThe client's configuration has lesser hassle, I swear!\n\nAssuming, and since I'm assuming that means, you must've setup the server using SSH, if not configure SSH on it :)\n\nNow mount the FTP drives to your computer using `sshfs` by running\n```\nsshfs user@host:dir /local/path\n```\n\nThis will prompt for your user's password on the server, enter it and you shall have your mounted drive.\n\nIn my case I've created this handy script that does the thing for the two drives, where it mounts them to the directory `lilnas` on my home directory.\n```bash\n#!/bin/bash\n# edit the script to fit your needs, I made the script, so that I run it whenever I'm home.\necho \"mounting disk1...\"\nsshfs baraa@16.0.0.2:/home/baraa/disks/disk1 ~/lilnas/disk1/ && echo \"done!\"\necho \"mounting disk2...\"\nsshfs baraa@16.0.0.2:/home/baraa/disks/disk2 ~/lilnas/disk2/ && echo \"done!\"\n```\n\nBonus: fun fact, using `ssh` or `sshfs` without the username, will omit the user from the current username, that is since I'm using my computer with the user `baraa` and I want to SSH to the server with the user `baraa` on the server, I can just do this:\n```bash\nssh 16.0.0.2\n# or\nsshfs 16.0.0.2:/home/baraa/disks/disk2 ~/lilnas/disk2/\n```\n\nAnd it will login using the username `baraa` on the server, since it's the current active user on my computer.",
			VisitTimes:  354,
		},
		{
			WrittenAt:   parseTime("2023-08-04T21:38:13"),
			Title:       "About TTYs...",
			Description: "The long story behind Unix-like OSes' tty console.",
			Content:     "### Some History\nTTY , teletypewriter or teleprinter was an interface to communicate with a computer, where it was literally a typewriter with some signal sent in/out to it.\n\n![Siemens t37h](https://mbaraa.com/img/9161_early-tty.jpg \"Siemens t37h teletypewriter\")\nImage source [Wikipedia](https://en.wikipedia.org/wiki/Teleprinter)\n\nOh and for anyone who have been to my blog before, I have written this whole blog post to flex the new image embedding 🤘\n\n\\\nSo this typewriter looking device was actually a computer terminal, where the key strokes are sent using the keyboard, and the output will be printed on a piece of paper, this was originally made in the late 1830s for sending telegraphs, and it was slapped on a computer nearly 1930s (dates might be incorrect, don't judge me), in which the characters were displayed using an early version of the ASCII encoding called [Baudot](https://en.wikipedia.org/wiki/Baudot_code) which was 5 bits in length, so it didn't have much characters choices, that it didn't even have lower-case Latin letters.\n\nSome time after that the paper was replaced by a display with the same logging style, i.e. the characters are written on the lowest row and scrolled up on each new line.\n\n![ADM-3A](link here \"ADM-3A Video display with keyboard\")\nImage source [Wikipedia](https://en.wikipedia.org/wiki/ADM-3A)\n\nThese display terminals are called [Dumb Terminals](https://www.pcmag.com/news/the-forgotten-world-of-dumb-terminals) since they are just input and output devices, where the computations are done on a separate computer, fun fact these computers used to have multi-users support like how we can have multiple `/dev/tty`s on a Unix-like machine, but more on that later.\n\n---\n\n### TTY now days\n\n#### Phones\n\nIt's somehow related here, that's why it's here, anyways, tty on phones is used by people with speech and hearing impairments, where a person can type a message and it is turned into speech, or the other way around, this helps these people a lot by making communication easier for them.\n\n---\n\n#### Computers\nEach shell you open on your Unix-like machine corresponds to a tty, that interacts with the system, to see all of your running ttys run\n```bash\nls /dev | grep tty\n```\n\nWhere this will list all of the ttys opened and used by applications, note that every terminal you open creates a new tty for you, you can check what tty you're currently using by running\n```bash\ntty\n```\n\nIt'll output something like this `/dev/pts/4` where **pts** means pseudo terminal, since it has a couple of layers before reaching the operating system.\n\nAnd every operating system these days includes a terminal (terminal emulator formally) in which you could interact with the OS's shell, but since we Linux users use it so much (because we can not because we're forced to), it's been called terminal for short.\n\nTypically a Unix-like system will greet you with a tty screen where you can use the computer like that or start a window manager or desktop environment.\n\n![Gentoo TTY](https://mbaraa.com/img/5107_gentoo-tty.jpg \"Gentoo Linux installing a package using emerge in the tty\")\nSome OSes like Gentoo, Arch, Void, and FreeBSD, don't include a GUI installer where you have to do the installation process from the tty, and some users (like myself) don't use a display manager to start the window manager (oh it's so bloated OMG), so we use the console tty more often than the average display manager enjoyer.\n\nAnd since you opened this blog you probably want to customize that tty screen don't you?\n\nWell hop on the next section.\n\n---\n\n### Customizing TTY's Greeting Message in Linux\n\nWhen you turn on your computer a message like this will appear (I don't use Debian on my computer, but my current tty message is a bit inappropriate)\n\n```\nDebian GNU/Linux 11 meowrver tty1\nmeowrver login:\n```\n\nHere where you can login to an existing user on the computer, but see the Debian... part, we're gonna change it!\n\nFirst off, this message is called `issue` [more on that](https://serverfault.com/questions/922235/what-is-the-difference-between-etc-issue-net-and-etc-issue), where it displays a \"greeting message\" to the computer's user, and since apparently you use it a lot, you might wanna personalize it, right?\n\nThere goes this table of escape sequences used by the `issue` file\n\n| Escape Sequence | Meaning |\n| --- | --- |\n| **\\b** | Insert the baudrate of the current line. |\n| **\\d** | Insert the current date. | \n| **\\s** | Insert the system name, the name of the operating system. |\n| **\\l** | Insert the name of the current tty line. |\n| **\\m** | Insert the architecture identifier of the machine, eg. i486 |\n| **\\n** | Insert the nodename of the machine, also known as the hostname. |\n| **\\o** | Insert the domainname of the machine. |\n| **\\r** | Insert the release number of the OS, eg. 1.1.9. |\n| **\\t** | Insert the current time. |\n| **\\u** | Insert the number of current users logged in. |\n| **\\U** | Insert the string “1 user” or “ users” where is the number of current users logged in. |\n| **\\v** | Insert the version of the OS, eg. the build-date etc. |\n\nFor example if we wanted the same `issue` as above that is `Debian GNU/Linux 11 meowrver tty1`, is actually `Debian GNU/Linux 11 \\n \\l`.\n\nTo make changes, open the file `/etc/issue` as root in your favorite editor and type whatever message you want with or without some escape sequences.\n\nThanks for sticking to the end, and have a nice day (or night idk).",
			PublicId:    "about-ttys",
			VisitTimes:  363,
		},
		{
			WrittenAt:   parseTime("2023-07-22T20:53:50"),
			Title:       "Android's USB Tethering with VPN",
			Description: "Using VPN the hard way!",
			PublicId:    "androids-usb-tethering-with-vpn",
			Content:     "Assuming you clicked on this cause you have a \"network\" with some strict VPN rules, where a limited number of VPNs work, and the GOAT [OpenVPN](https://openvpn.net/) doesn't work at all, so you need some hacking around to get a proper VPN connection, and yeah I am saying \"network\" for obvious reasons!\n\\\nSo let's say that you wanna use a VPN to get a secure connection, without tracing, or just to used a product blocked in your region, in my case I chose [NordVPN](https://nordvpn.com/) since it is the fastest cross-platform VPN that works perfectly on Linux, and **NO** they're not paying me to say that, I mean I barely got it working due to the \"network's\" restrictions.\n\\\nAfter watching some reviews and benchmarks, I stuck with NordVPN, but one thing that I've noticed, that is their website doesn't open on my \"network\", so I used [Browsec](https://browsec.com/en/) on my browser and phone to download NordVPN, oh and Browsec is also great, but it doesn't have a native desktop application, that's why I went with NordVPN.\n\\\nSo I got the Android's version working perfectly without any headaches, but now I wanna get the Desktop version, and there it was the disappointment. well it could've been resolved easily if I used the binary version of NordVPN with some tweaking and be done in  ~20mins, but no I need all of my programs managed by [portage](https://wiki.gentoo.org/wiki/Portage), that is [Gentoo](https://gentoo.org)'s package manager, so that I don't leave any outdated packages, when doing a full system update.\n\\\nAND here was the catch, the package's download URI has the same domain as NordVPN's, where it is blocked on my \"network\", so I had to be a bit creative, to save you some trouble, here's a list of the things that I've tried, that weren't fruitful:\n* [This](https://forum.xda-developers.com/t/tethering-usb-on-android-with-vpn-guide-and-qs.2446643/) XDA tutorial, IDK I had hopes for it, especially that it only refreshes the firewall's rules, but it didn't do anything fruitful.\n* [PDANet](http://pdanet.co/); since it doesn't support Linux, and it's proxy configuration isn't really working.\n* [TetherNet](https://m.apkpure.com/tethernet-vpn-tethering/com.ilmubytes.tethernet); it just didn't work, maybe because of Android 13, it just didn't work.\n* [DF Tethering Fix](https://m.apkpure.com/df-tethering-fix/com.formichelli.tetheringfix); when this one asked for root access, I was like, \"oh hell yeah, it needs root, it must work correctly then\", but no.\n* I tried a couple more apps that didn't work, and I'm too lazy to list them 👀\n\nThen I found this [VPN Hotspot](https://play.google.com/store/apps/details?id=be.mygod.vpnhotspot) app, it was my holy grail, It requires root (cause real stuff always require root), I didn't go deep how it works, but it just worked!!!!\n\\\nAnd when I saw the new public IP, it was a joyful moment, anyway I proceeded with installing NordVPN from portage, and It worked as expected.\n\n---\n\nTLDR;\nInstall [VPN Hotspot](https://play.google.com/store/apps/details?id=be.mygod.vpnhotspot), it needs root, and it tethers your phone's VPN connection to your computer.",
			VisitTimes:  395,
		},
		{
			WrittenAt:   parseTime("2023-07-05T18:38:24"),
			Title:       "A Rant About the Scroll Event in Nuxt 3",
			Description: "VueUse's useWindowScroll needs a WINDOW",
			Content:     "So, let's say that you wanna listen to a scroll event to fetch more content, or create a parallax effect, or anything else related to scrolling, and it happens that you write [Vue 3](https://vuejs.org/) , so all what you gotta do is utilize the [useWindowScroll](https://vueuse.org/core/useWindowScroll/) composable from [VueUse](https://vueuse.org). \n\nSo you'd do something like this \n```vue\n<template>\n  <div style=\"width: 2000px; height: 2000px\">\n    Scroll X: {{ x }}\n    <br />\n    Scroll Y: {{ y }}\n  </div>\n</template>\n\n<script setup lang=\"ts\">\nimport { useWindowScroll } from \"@vueuse/core\";\nconst { x, y } = useWindowScroll(window);\n</script>\n```\nAnd it's the exact same thing in [Nuxt 3](https://nuxt.com/), right?\n\n**NO**\n\nNuxt has a different opinion, i.e it slaps `window is not defined` in our faces, ngl Nuxt is great and everything, and it was my insist to use a clean composable from VueUse, so that the code looks like a Vue code not some Vue with some classic JavaScript slapped into it.\n\nWell `useWindowScroll` and SSR don't really mix, since it has the word `window` in it!\n\nBut unlike [SvelteKit](https://kit.svelte.dev/) where they managed to handle stuff like this with just\n```svelte\n<script lang=\"ts\">\n    let scrollY = 0;\n    let scrollX = 0;\n</script>\n\n<div style=\"width: 2000px; height: 2000px\">\n  Scroll X: {{ scrollX }}\n  <br />\n  Scroll Y: {{ scrollY }}\n</div>\n\n<svelte:window bind:scrollX bind:scrollY />\n```\nWhere this will do the same as it does in [Svelte](https://svelte.dev/), without any headaches and hacking around (you'll see in a bit).\n\nSo first of all we need to make sure we're on the browser, for that `onMounted` comes in handy.\n\n```vue\n<script setup lang=\"ts\">\nimport { useWindowScroll } from \"@vueuse/core\";\n\nonMounted(() => {\n  const { x, y } = useWindowScroll(window);\n});\n</script>\n```\nBut wait now `x` and `y` won't be accessible from the template, since they're locals in that arrow function.\n\nSo we'll create global ones and update them as the locals update\n```vue\n<template>\n  <div style=\"width: 2000px; height: 2000px\">\n    Scroll X: {{ X }}\n    <br />\n    Scroll Y: {{ Y }}\n  </div>\n</template>\n\n<script setup lang=\"ts\">\nimport { useWindowScroll } from \"@vueuse/core\";\n\nconst X = ref(0);\nconst Y = ref(0);\n\nonMounted(() => {\n  const { x, y } = useWindowScroll(window);\n  watch(x, (value) => {\n    X.value = value;\n  });\n  watch(y, (value) => {\n    Y.value = value;\n  });\n});\n</script>\n```\n\nThis works as expected, but let's say we need the same event listener somewhere else, we'd copy the same code, right?\n\nWell that's just stupid ain't it? so now for the hacking part 🎉.\n\n---\nLuckily Nuxt supports client side plugins, i.e a plugin that only renders on the browser, that sounds like fun doesn't it?\n\nWell it depends on how you define fun!\n\nSo first under `~/plugins` we'll create the `use-scroll.ts` file which will define the cool wrapper of `useWindowScroll` which will work in Nuxt!\n\n```ts\n// ~/plugins/use-scroll.ts\nimport { useWindowScroll } from '@vueuse/core'\nexport default defineNuxtPlugin((nuxtApp) => {\n  const { x, y } = useWindowScroll(window)\n  return {\n    provide: {\n      useScroll: () => {\n        return { x, y }\n      },\n    },\n  }\n})\n```\nAnd off-course modify `~/nuxt.config.ts`\n```ts\nplugins: [{ src: \"~/plugins/use-scroll.ts\", ssr: false, mode: \"client\" }]\n```\n\nThat way we abstracted the thing, we still have much to do, you'll see now\n\n```vue\n<script setup lang=\"ts\">\nconst { $useScroll } = useNuxtApp();\nconst { x, y } = $useScroll();\n</script>\n```\nAnd you'd expect that would work right out of the box right? well think again we need to check if `$useScroll` is defined correctly and it's not `undefined`.\n\nSo we'll check it manually\n```vue\n<template>\n  <div style=\"width: 2000px; height: 2000px\">\n    Scroll X: {{ X }}\n    <br />\n    Scroll Y: {{ Y }}\n  </div>\n</template>\n\n<script setup lang=\"ts\">\nconst X = ref(0);\nconst Y = ref(0);\n\nconst { $useScroll } = useNuxtApp();\nif (typeof $useScroll === \"function\") {\n  const { x, y } = $useScroll();\n  watch(x, (value) => {\n    X.value = value;\n  });\n  watch(y, (value) => {\n    Y.value = value;\n  });\n}\n</script>\n```\nSo basically it's same but using a \"plugin\" so that's just seems pointless, well I'm not done YET!\n\nWe need a fancy `composable`, cuz what's Vue 3 without composables, right?\n\nCreate this file that will hold the _fancy_ composable\n```ts\n// ~/composables/useScroll.ts\nexport default function () {\n  const X = ref(0);\n  const Y = ref(0);\n  const { $useScroll } = useNuxtApp();\n  if (typeof $useScroll === \"function\") {\n    const { x, y } = $useScroll();\n    watch(x, (value) => {\n      X.value = value;\n    });\n    watch(y, (value) => {\n      Y.value = value;\n    });\n\t\t// the use of scrollX, and scrollY makes more sense than just x, and y\n    return { scrollX: X, scrollY: Y };\n  }\n  // this return for when the plugin is not ready yet (ssr mode)\n  return { scrollX: 0, scrollY: 0 };\n}\n```\nIt's a bit messy, but for a single time eh? and that's the beauty of composables!\n\nNow we just use it in a component, and life goes on...\n```vue\n<template>\n  <div style=\"width: 2000px; height: 2000px\">\n    Scroll X: {{ scrollX }}\n    <br />\n    Scroll Y: {{ scrollY }}\n  </div>\n</template>\n\n<script setup>\nconst { scrollX, scrollY } = useScroll();\n</script>\n```\n\nThanks for reading the whole rant and solution.\n",
			PublicId:    "a-rant-about-the-scroll-event-in-nuxt-3",
			VisitTimes:  895,
		},
		{
			WrittenAt:   parseTime("2023-04-24T21:36:57"),
			Title:       "Install Lua LSP for Neovim",
			Description: "Lua LSP installation on Linux distros that don't include it or has an outdated version of it.",
			PublicId:    "install-lua-lsp-for-neovim",
			Content:     "Assuming you're here because you want LSP(Language Server Protocol) for your neovim setup, but your Linux distro doesn't provide the latest version of lua-language-server, so in this article we'll install lua-language-server from source.\n\n#### Dependencies installation\nMake sure that you have, Ninja Build, GCC(some distros need G++ as well), and Clang.\n\nTo install the dependencies on Gentoo run:\n```bash\nsudo emerge -qav sys-devel/gcc sys-devel/clang dev-util/ninja \n```\n\n#### Compiling LSP's source code\n##### Cloning the repo:\n- Clone LSP's repo into a directory where you keep bins and stuff\n- I use `~/.local/bin` so I'll just clone it there\n- `git clone https://github.com/LuaLS/lua-language-server ~/.local/bin/lua-language-server`\n- `cd ~/.local/bin/lua-language-server`\n\n##### Compile the stuff:\n- Download the sub-modules of the cloned repo\n- `git submodule update --recursive`\n- Download the `ninja` luamake rules\n- `cd 3rd/luamake`\n- `git submodule update --init`\n- Run the compile script\n- `compile/install.sh`\n- `cd ../../`\n- `./3rd/luamake/luamake rebuild`\n\n##### Add the the executables' path to your path\n```bash\nSHELL_NAME=`basename $SHELL`\nSHELL_RC=\"./.${SHELL_NAME}rc\"\necho 'export PATH=\"${HOME}/.local/bin/lua-language-server/bin:${PATH}\"' >> $SHELL_RC\n```\nNow re-login or run\n```bash\nsource $SHELL_RC\n$SHELL_NAME\n```\n\n##### Require the new installed lsp server\nAdd this line to `~/.config/nvim/init.lua` or to where you put lsp's config in Neovim\n\n```lua\nrequire('lspconfig').sumneko_lua.setup {}\n```\n\nRestart Neovim and your good to go.",
			VisitTimes:  416,
		},
		{
			WrittenAt:   parseTime("2023-02-04T17:30:20"),
			Title:       "Use Tree Style Tabs Instead of the Default Tab Bar in Firefox",
			Description: "The configuration needed in Firefox to hide the tabs bar and only use Tree Style Tabs",
			PublicId:    "use-tree-style-tabs-instead-of-the-default-tab-bar-in-firefox",
			Content:     "First install [Tree Style Tab](https://addons.mozilla.org/en-US/firefox/addon/tree-style-tab/) plugin.\n\nNow for the tab hiding part:\n1. Go to your profile directory, directory's path can be found by going to [about:support](about:support) and open the **Profile Directory**\n1. Create a directory called `chrome`\n1. Create a file called `userChrome.css`\n1. Put these in the file\n```css\n#TabsToolbar\n{\n    visibility: collapse;\n}\n```\n1. Go to [about:config](about:config) and set `toolkit.legacyUserProfileCustomizations.stylesheets` to **true**\n1. Restart Firefox to see the changes",
			VisitTimes:  386,
		},
		{
			WrittenAt:   parseTime("2023-01-25T07:11:10"),
			Title:       "Learn Docker by Dockerizing A SpringBoot, SvelteKit, MariaDB, and Keycloak App",
			Description: "A headstart for docker and docker compose using a mix of technologies",
			PublicId:    "learn-docker-by-dockerizing-a-springboot-sveltekit-mariadb-and-keycloak-app",
			Content:     "This project is a mix of lots of technologies, that will be somehow hard to dockerize together, but it will be fun along the way since it includes [volumes](https://docs.docker.com/storage/volumes/), [networks](https://docs.docker.com/network/), and the glue that holds it all together [docker compose](https://docs.docker.com/compose/), so let's get started.\n\n\n## Project Structure:\nAs we said earlier this project consists of Spring Boot (backend server), SvelteKit (web client), MariaDB (database), and Keycloak (Authentication provider), and the project outline should look like this:\n\n-- auth/\\\n---- realms_backups/\\\n-- client/\\\n---- Dockerfile\\\n---- ...more sveltekit files\\\n-- server/\\\n---- Dockerfile\\\n---- ...more spring boot files\\\n-- docker-compose.yml\\\n\n## Installing Docker\nI'll demonstrate how to install docker on Gentoo Linux, other Linux distros and platforms can be found [here](https://docs.docker.com/engine/install/).\n\nInstalling docker has 5 steps most of which are the same on any other Linux distro:\n\n1. Installation\\\n   `sudo emerge -qav app-containers/docker app-containers/docker-cli`\n1. Enable the docker daemon on startup\\\n   `sudo rc-update add docker default`\n1. Start the docker service\\\n   `sudo rc-service docker start`\n1. Adding your user to the docker group to be able to use docker without superuser permission\\\n   `sudo gpasswd -a $(whoami) docker`\n1. Restart your shell\\\n   this is required so that the user's groups are updated (after adding our user to the docker group we need to do this)\n   you can do this by restarting your working session, or by typing `$SHELL` in your active terminal\n\n* Bonus, run this to make sure that docker is working just fine on your machine\\\n\t`docker run hello-world`\n\n<br/>\n\n---\n\n## Dockerizing a simple Spring Boot App:\nWe'll start by creating a [Spring Boot](https://spring.io/projects/spring-boot) application using the [initializer](https://start.spring.io/) with the following configs:\n\n| Config  | Description  |\n| ------------ | ------------ |\n| Language  | Java 11 (it's the GOAT version so far)  |\n| Building System  | Maven (gradle is just too easy)  |\n| Spring Boot  | Version 2.7.8 (that's what goes with Java 11 these days)  |\n| Spring Web  | Just add it from the dependencies :)  |\n| Packaging | Jar (if you like WAR you're on your own 😉 ) |\n\nNow after unzipping the downloaded spring project, open it in your favorite editor, and add a [Rest Controller](https://spring.io/guides/tutorials/rest/) so we can test out this thing\n\n```java\n// controllers/HelloController.java\n\nimport org.springframework.web.bind.annotation.GetMapping;\nimport org.springframework.web.bind.annotation.PathVariable;\nimport org.springframework.web.bind.annotation.RestController;\n\n@RestController\npublic class HelloController {\n    @GetMapping(\"/hello/{name}\")\n    public String greet(@PathVariable String name) {\n        return String.format(\"Hello, %s\", name);\n    }\n}\n```\n\nNow go back to the terminal and run the spring boot app using `./mvnw spring-boot:run`, now we just do a little check using `curl` to make sure that everything is in its place\n\n```bash\ncurl http://localhost:8080/hello/Eloi\n```\n\nAnd if everything was in its place it should print \"Hello, Eloi\".\n\nAfter that everything is working just fine, we need to break everything :), just kidding, but we need to change some settings around, for starters, we need to change from the weird-looking `application.properties` to the awesome superior `application.yml`, since it's 2023 and everyone thinks that YAML is cooler now, after the change, we need to change the server's port (we'll find out why later on in this tutorial), and now the config file should look like this\n\n```yml\nserver:\n  port: 8081\n```\n\nAfter rerunning the application we can re-test it,\n\n```bash\ncurl http://localhost:8081/hello/Eloi\n```\nJust to make sure that everything is working as expected :)\n\nOk now we're all set to dockerize this simple spring app, as I mentioned in the project structure outline above, the server (Spring Boot) has a file called Dockerfile, well this [file](https://docs.docker.com/engine/reference/builder/) tells docker how to build and run the image, fascinating isn't it?\n\n```dockerfile\n# build stage\nFROM alpine:latest as build\n\nRUN apk add openjdk11-jdk openjdk11-jre openjdk11-src maven\nWORKDIR /app\nCOPY . .\nRUN mvn clean install\n\n# run stage\nFROM alpine:latest as run\n\nRUN apk add openjdk11-jre\nWORKDIR /app\nCOPY --from=build /app/target/*.jar ./run.jar\n\nEXPOSE 8081\nCMD [\"java\", \"-jar\", \"./run.jar\"]\n```\n\nThat's the configuration needed to run a Spring Boot app inside a docker container, but I owe you some explanation, first as you can see the file is separated into two sections, build and run, well this is useful to save disk storage, I mean imaging building three applications like this, with each container having all the build files, jdk, maven, ....\n\nIt would be a nightmare, so, here we are separating the build from the run, let's talk about the build stage for a bit, first, we're pulling an image called [Alpine Linux](https://www.alpinelinux.org/) using the `FROM` keyword the column after the image's name specifies the version of the image to be pulled, here we're using the latest version of Alpine available, Alpine is a light Linux distro that is suited for small containers and virtual machines like this, after pulling the image, we see \n```\nRUN apk add openjdk11-jdk openjdk11-jre openjdk11-src maven\n```\n`RUN` is used to run a shell command inside the container, `apk` is the package manager used by Alpine, and no it has nothing to do with Android. back to docker, here we're installing JDK, JRE, and maven, so we can compile the application into a single JAR file.\n\n`WORKDIR` is just like `cd` which changes the current working directory, here we're using `/app` which is just a naming, and it can be whatever you want, but `/app` is just convenient enough.\n\n`COPY` is like `cp` where it copies a file or directory from the given source to the given destination, here we're copying the whole Spring Project into the container, where we will compile it.\n\nNow to build the project and produce a single JAR file we have to run `mvn clean install` inside the container, again we'll use the `RUN` keyword before it.\n\n---\n\nAfter the building process is done we need to prepare the run container image, and again we're pulling the same Alpine image, but the difference here is we're not installing JDK, or maven, since their job was done in the build stage, now we just copy the JAR file into the run container, add some magic and we're ready to go.\n\nThe magic:\n`EXPOSE` allows a port from the container to be viewed by the docker network for the host to be able to use it, remember the port we set in `application.yml` was 8081 so we expose that same port.\n\n`CMD` is what docker will run in the container after starting it, but as you can notice it's an array of strings, which is the original command string split by a space.\n\nfor example the running command is `java -jar ./run.jar`, in which it becomes `[\"java\", \"-jar\", \"./run.jar\"]`\n\n---\n\nNow it's time for action, first, we'll need to build the image, open your terminal and navigate to the server directory, then run\n```bash\ndocker build -t hello-spring .\n```\n\nHere the `-t` flag specifies the name of the container's image after building and the `.` indicates the current directory which will be used for the build.\n\nNow get rid of the build container, by removing what's called \"dangling images\", these are images that no one depends on and can be removed without damaging any other image, and there's nothing depends on them because we actually ran the image while building and got what we wanted from it, and it's now time for throwing it away.\n\nTL;DR just run\n```bash\ndocker image prune\n```\n\nIt should prompt you, don't freak out, just hit `yes`\n\n---\n\nAND NOW FOR THE REAL ACTION, WE WILL RUN THE BUILT CONTAINER\n\n```bash\ndocker run  -p 8080:8081 hello-spring\n```\n\nYou can test now I'll explain after you test your container, so you get the satisfaction of running a docker container.\n```bash\ncurl http://localhost:8080/hello/Eloi\n```\n\nThe docker run command attaches the specified image to a docker container, and the `-p` flag specifies the port forwarding to the host from the container, just remember this magical formula `-p HOST:CONTAINER`, and the final argument is the image's name that we want to run.\n\n<br/>\n\n---\n\n## Configuring and Dockerizing Keycloak\nConfiguring Keycloak requires two stages, the actual realm configuration, and the docker configuration for Keycloak, let's get started\n\nGet a Keycloak zip archive form [here](https://www.keycloak.org/downloads), we'll use this server to make our configurations, then export the configured realm, and use it with the docker container.\n\nNow open your browser, and go to `localhost:8080` which is the Keycloak server address, then go to `Administration Console`, and log in with the credentials you specified, that is `admin:SOME_PASSWORD`.\n\nFirst, create a realm with any name you prefer, I'll name mine \"dori\", but that's not the topic here, after that we'll create a new realm.\n\nNow, we'll create a client called \"dori-client\", with `Standard Flow`, `OAuth 2.0 Device Authorization Grant`, and `Client authentication` enabled, then create a role called \"superuser\", then create a user called \"nemo\", set \"nemorocks\" as a password to it and assign the \"superuser\" role to it.\n\nNow we'll export the realm configuration, from the project's directory run:\n\n```bash\ncd keycloak20.0.2/bin/\n./kc.sh export --dir backups --realm dori\n```\n\nthis will produce a directory with two files, `dori-realm.json` and `dori-users.json`, copy those files into our project, specifically into `auth/realms_backups/`\n\nWe'll be using [Keycloak's official docker image](https://quay.io/repository/keycloak/keycloak) with version `20.0.2`\n\nNow for the docker part, run this magical command to import the realm, and run the docker container.\n\n```bash\ndocker run -v ./auth/realms_backups/:/tmp/backups/\\\n\t-e KEYCLOAK_ADMIN=admin\\\n\t-e KEYCLOAK_ADMIN_PASSWORD=admin\\\n\t-p 8080:8080\\\n\tquay.io/keycloak/keycloak:20.0.2\\\n\t-Dkeycloak.profile.feature.upload_scripts=enabled\\\n\t-Dkeycloak.migration.action=import\\\n\t-Dkeycloak.migration.realmName=dori\\\n\t-Dkeycloak.migration.provider=dir\\\n\t-Dkeycloak.migration.dir=/tmp/backups/\\\n\tstart-dev\n```\n\nThis might be scary at first sight, but it's not if we break it down into parts.\n\nFirst, there is the `-v` flag specifies volume mounting, just like the port forwarding, but this one is for volumes, i.e it mounts a path from the host to the container.\n`-v /path/in/host/:/path/in/container`, and here the host directory is `./auth/realms_backups/` since there we'll be keeping the realm backup(s).\n\nThen we got the  `-e` flag, which specifies an environment variable, in this case, we're specifying the admin's username and password, which are \"admin\", \"admin\" respectively.\n\nThen we got `-p` that we know that it forwards ports to the host from the container, after that the container's name and version that we will be running, and finally, the huge run command, which the import flags, that specify where and how to do the realm import.\n\nGreat, now back to our Spring Boot app, now we need to add some Keycloak configurations to it, should be easy right?\n\nWe'll start with the maven dependencies\n```xml\n<!-- pom.xml -->\n\n<dependency>\n\t<groupId>org.keycloak</groupId>\n\t<artifactId>keycloak-spring-boot-starter</artifactId>\n\t<version>20.0.2</version>\n</dependency>\n\n<dependency>\n\t<groupId>org.keycloak</groupId>\n\t<artifactId>keycloak-spring-security-adapter</artifactId>\n\t<version>10.0.0</version>\n</dependency>\n\n<dependency>\n\t<groupId>org.springframework.boot</groupId>\n\t<artifactId>spring-boot-starter-security</artifactId>\n\t<version>3.0.0</version>\n</dependency>\n```\n\nHere we've added Spring Boot security and Keycloak dependencies, now off to the Keycloak Configuration class:\n\n```java\n// conf/KeycloakAdapterConfig.java\n\nimport org.keycloak.adapters.springboot.KeycloakSpringBootConfigResolver;\nimport org.keycloak.adapters.springsecurity.KeycloakConfiguration;\nimport org.keycloak.adapters.springsecurity.authentication.KeycloakAuthenticationProvider;\nimport org.keycloak.adapters.springsecurity.config.KeycloakWebSecurityConfigurerAdapter;\nimport org.springframework.beans.factory.annotation.Autowired;\nimport org.springframework.context.annotation.Bean;\nimport org.springframework.context.annotation.Import;\nimport org.springframework.http.HttpMethod;\nimport org.springframework.security.config.annotation.authentication.builders.AuthenticationManagerBuilder;\nimport org.springframework.security.config.annotation.method.configuration.EnableGlobalMethodSecurity;\nimport org.springframework.security.config.annotation.web.builders.HttpSecurity;\nimport org.springframework.security.core.authority.mapping.SimpleAuthorityMapper;\nimport org.springframework.security.web.authentication.session.NullAuthenticatedSessionStrategy;\n\n@KeycloakConfiguration\n@EnableGlobalMethodSecurity(prePostEnabled = true)\n@Import({KeycloakSpringBootConfigResolver.class})\npublic class KeycloakAdapterConfig extends KeycloakWebSecurityConfigurerAdapter {\n\n    /* Registers the KeycloakAuthenticationProvider with the authentication manager.*/\n    @Autowired\n    public void configureGlobal(AuthenticationManagerBuilder auth) throws Exception {\n        KeycloakAuthenticationProvider keycloakAuthenticationProvider = keycloakAuthenticationProvider();\n        keycloakAuthenticationProvider.setGrantedAuthoritiesMapper(new SimpleAuthorityMapper());\n        auth.authenticationProvider(keycloakAuthenticationProvider);\n    }\n\n    /* Defines the session authentication strategy null means no session.*/\n    @Bean\n    @Override\n    protected NullAuthenticatedSessionStrategy sessionAuthenticationStrategy() {\n        return new NullAuthenticatedSessionStrategy();\n    }\n\n    @Override\n    protected void configure(HttpSecurity http) throws Exception {\n        super.configure(http);\n\n        http.csrf()\n                .disable()\n                .authorizeRequests()\n                .antMatchers(HttpMethod.GET, \"/super-hello/\")\n                .hasRole(\"superuser\");\n    }\n}\n```\n\nWell, this is a docker tutorial, so all you need to understand from this file is the `antMatchers` these specify the path, http method, and who can use it, here we have a `GET` method on the route `/super-hello` that only can be used by a user with the `superuser` role.\n\nAND NOW for the REST API, we'll need to add the endpoint `/super-hello`, modify the `HelloController`, and add:\n\n```java\n...\nimport org.springframework.security.access.prepost.PreAuthorize;\n...\n    @PreAuthorize(\"hasRole('superuser')\")\n    @GetMapping(\"/super-hello/{name}\")\n    public String superGreet(@PathVariable String name) {\n        return String.format(\"Super Hello, %s\", name);\n    }\n...\n```\n\nFinally (not really), we need to add some Keycloak configuration to `application.yml`\n\n```yml\nkeycloak:\n  realm: dori\n  auth-server-url: \"http://localhost:8080/\"\n  resource: dori-client\n  public-client: true\n  bearer-only: true\n```\n\nSo what about that `/super-hello` request?, if we requested it'll give us a 401 (Unauthorized) status code, So we need a token right?\n\nWe can get a token, by making a `token` request to the Keycloak server\n\n```bash\ncurl -X POST http://localhost:8080/realms/dori/protocol/openid-connect/token\\\n   -H 'Content-Type: application/x-www-form-urlencoded' \\\n   -d 'client_id=dori-client&client_secret=YOUR_CLIENT_SECRET&grant_type=password&username=nemo&password=nemorocks'\n```\n\nthen you should get a response like this\n```json\n{\n    \"access_token\": \"\",\n    \"expires_in\": 300,\n    \"refresh_expires_in\": 1800,\n    \"refresh_token\": \"\",\n    \"token_type\": \"Bearer\",\n    \"not-before-policy\": 0,\n    \"session_state\": \"\",\n    \"scope\": \"email profile\"\n}\n```\n\nas you can see we have the access token to the client `dori-client` from the `dori` realm, using the user `nemo`, now when we use the token with the `/super-hello` it'll work.\n\n```bash\ncurl http://localhost:8081/super-hello/Eloi\\\n    -H \"Authorization: Bearer YOUR_ACCESS_TOKEN\"\n```\n\nThis is so cool, right? but we're not done yet, we need to containerize the thing right?\n\nNow we'll introduce [docker-compose](https://docs.docker.com/compose/compose-file/) that will allow us to run more than one container at the same time (not really, but it appears to do that) with a related setup, in this case, we need a network between the Spring Boot, and the Keycloak server, finally things are getting along :)\n\nThis file will be at the root of the whole project.\n\n```yml\n# docker-compose.yml\nversion: \"3.8\"\n\nservices:\n  auth:\n    image: \"quay.io/keycloak/keycloak:20.0.2\"\n    container_name: \"auth\"\n    restart: \"always\"\n    ports:\n      - 9090:8080\n    environment:\n      KEYCLOAK_ADMIN: \"admin\"\n      KEYCLOAK_ADMIN_PASSWORD: \"admin\"\n    volumes:\n      - ./auth/realms_backups/:/tmp/backups/\n    command: \"-Dkeycloak.profile.feature.upload_scripts=enabled -Dkeycloak.migration.action=import -Dkeycloak.migration.realmName=dori -Dkeycloak.migration.provider=dir -Dkeycloak.migration.dir=/tmp/backups start-dev\"\n    networks:\n      - auth-backend\n\n  backend:\n    build: ./server\n    ports:\n      - 8080:8081\n    depends_on:\n      - auth\n    networks:\n      - auth-backend\n\nnetworks:\n  auth-backend: {}\n```\n\nSo..., what's going on here?\\\nIf you look close enough you'll notice something we've seen before, aside from the other configuration, the services have `ports` property which will do port forwarding the same as `-p`, `environment` is like `-e`, `volumes` is like `-v` when using `docker run`.\n\n`command` overrides `CMD` in the Dockerfile, meaning whatever we put in there will be executed when the container starts.\n\nNow let's go over the compose file, it's just a YAML file that tells `docker-compose` what to do, first, we have the `version` which is the version of the compose file, currently, the latest version is `3.8` so we're gonna use that, now for the services array, for starters, we have the `auth` service is the Keycloak server, where we specify the wanted docker image using the `image` property.\n\n`build` specifies where the docker project is, i.e. a project with a `Dockerfile` at its root, and build has more interesting stuff that can be found in [here](https://docs.docker.com/compose/compose-file/build/)\n\n`depends_on` states that the `backend` service will not run until the `auth` service has started.\n\n`container_name` specifies what name this image's container will be using while it's running, so it can be accessed from the docker network (for now that's all that we need from the name), as you can see each one of the services has a `network` property which is an array that represents the networks that the container will be connected to, in this case, `auth` and `backend` are connected to.\n\nNow where to get the network?\n\nas you can see at the end of the file we can see a `networks` array that defines our networks, and here we've defined a network called `auth-backend` that will connect the Spring Boot server to the Keycloak server, easy eh?\n\nwell, it's not as easy as it seems, but that's all we need for this setup, you check out more about networks [here](https://docs.docker.com/compose/networking/).\n\nAs I said `container_name` will help with the network, but how, well now that the Keycloak service is named `auth` that will be used as the server's address.\n\nNow we can change the Keycloak address in `application.yml` to `http://auth:8080/`, here we're using port 8080 since that's the server address inside the docker network, we can still use `http://localhost:9090/` if we want to, but it's more convenient to use the docker network. and now it's time for action.\n\nRun `docker compose up` and it'll build the project for the first time, and start it, but if anything changes it won't re-build the project with the newest changes, so we need to run `docker compose build` after each change and the changed image will be rebuilt, and ready for running.\n\nJust a little test to make sure that everything is in its place.\n\nacquire the access token first, from the Keycloak server.\n\n```bash\ncurl -X POST http://localhost:9090/realms/dori/protocol/openid-connect/token\\\n   -H 'Content-Type: application/x-www-form-urlencoded' \\\n   -d 'client_id=dori-client&client_secret=YOUR_CLIENT_SECRET&grant_type=password&username=nemo&password=nemorocks'\n```\n\nthis will work, but if you make `/super-hello` with the token returned from the previous request, it won't work, because the token was issued to the address `localhost:9090` and the Spring Boot requests the server at `auth:8080`, and Keycloak is careful who can use the token and who can't.\n\nso to avoid situations like this, we can easily issue the token from our Spring Boot server, we'll create an endpoint `/token` that will make a request to the Keycloak server and retrieve a token that was issued for the same address.\n\nFirst, we need to add a JSON Utility Dependency, since as we've seen earlier the response from the Keycloak server is a JSON.\n\n```xml\n<!-- pom.xml -->\n...\n\t\t<dependency>\n\t\t\t<groupId>org.json</groupId>\n\t\t\t<artifactId>json</artifactId>\n\t\t\t<version>20220924</version>\n\t\t</dependency>\n...\n```\n\nand now we'll create a controller that does the token retrieving request:\n\n```java\n// controllers/TokenController.java\n\nimport org.apache.http.client.config.RequestConfig;\nimport org.springframework.http.ResponseEntity;\nimport org.springframework.web.bind.annotation.PostMapping;\nimport org.springframework.web.bind.annotation.RequestBody;\nimport org.springframework.web.bind.annotation.RestController;\nimport org.apache.http.NameValuePair;\nimport org.apache.http.client.entity.UrlEncodedFormEntity;\nimport org.apache.http.client.methods.HttpPost;\nimport org.apache.http.impl.client.HttpClientBuilder;\nimport org.apache.http.util.EntityUtils;\nimport org.json.JSONObject;\nimport org.springframework.beans.factory.annotation.Value;\n\nimport java.util.List;\nimport java.util.Map;\n\n@RestController\npublic class TokenController {\n    @Value(\"${keycloak.auth-server-url}\")\n    private String authServerURL;\n\n    @PostMapping(\"/token\")\n    public ResponseEntity<?> login(@RequestBody Map<String, String> user) {\n        try {\n            var form = List.of(\n                    new NameValuePairImpl(\"client_id\", \"dori-client\"),\n                    new NameValuePairImpl(\"client_secret\", \"YOUR_CLIENT_SECRET\"),\n                    new NameValuePairImpl(\"grant_type\", \"password\"),\n                    new NameValuePairImpl(\"username\", user.get(\"username\")),\n                    new NameValuePairImpl(\"password\", user.get(\"password\"))\n            );\n\n            var requestConfig = RequestConfig.custom().build();\n            var httpClient = HttpClientBuilder.create().setDefaultRequestConfig(requestConfig).build();\n            var request = new HttpPost(String.format(\"%s/realms/dori/protocol/openid-connect/token\", authServerURL));\n            request.setEntity(new UrlEncodedFormEntity(form));\n            JSONObject json = new JSONObject(EntityUtils.toString(httpClient.execute(request).getEntity()));\n\n            return ResponseEntity.ok(Map.of(\"token\", json.get(\"access_token\")));\n        } catch (Exception e) {\n            return ResponseEntity.internalServerError().body(e.toString());\n        }\n    }\n}\n\nclass NameValuePairImpl implements NameValuePair {\n    private final String name;\n    private final String value;\n\n    public NameValuePairImpl(String name, String value) {\n        this.name = name;\n        this.value = value;\n    }\n\n    @Override\n    public String getName() {\n        return name;\n    }\n\n    @Override\n    public String getValue() {\n        return value;\n    }\n}\n```\n\nThis controller only has one endpoint, that is `/token`, so we'll send a json with `username` and `password`, which will be used for logging in to the Keycloak realm.\n\nNow we can rebuild the images, and test the requests again.\n\n```bash\ndocker compose build\ndocker compose up\n\ncurl -X POST http://localhost:8080/token\\\n\t-H \"Content-Type: application/json\"\\\n\t--data '{\"username\": \"nemo\", \"password\": \"nemorocks\"}'\n\ncurl http://localhost:8080/super-hello/Eloi\\\n\t-H \"Authorization: Bearer ACCESS_TOKEN\"\n```\n\nAnd gladly I can finally say that this part is over.\n\n<br/>\n\n---\n\n## Dockerizing MariaDB\n\nThis part is cuter than Keycloak, since we'll just create a model, a simple controller, and modify some configuration files, that should be easy.\n\nFirst, we need to configure Spring Boot with JPA, now we need JPA and MariaDB dependency.\n\n```xml\n<!-- pom.xml -->\n...\n\t\t<dependency>\n\t\t\t<groupId>org.springframework.boot</groupId>\n\t\t\t<artifactId>spring-boot-starter-data-jpa</artifactId>\n\t\t\t<version>2.7.8</version>\n\t\t</dependency>\n\n\t\t<dependency>\n\t\t\t<groupId>org.mariadb.jdbc</groupId>\n\t\t\t<artifactId>mariadb-java-client</artifactId>\n\t\t\t<scope>runtime</scope>\n\t\t</dependency>\n...\n```\n\nupdate your dependency tree using \n```bash\nmvn dependency:resolve\n```\n\nand update your `application.yml` to use MariaDB with JPA.\n\n```yml\nserver:\n  port: 8081\n\nspring:\n  datasource:\n    url: \"jdbc:mariadb://db/someDB?useJDBCCompliantTimezoneShift=true&useLegacyDatetimeCode=false&serverTimezone=UTC\"\n    username: \"root\"\n    password: \"hello\"\n    driver-class-name: \"org.mariadb.jdbc.Driver\"\n  jpa:\n    generate-ddl: true\n\nkeycloak:\n  realm: dori\n  auth-server-url: \"http://auth:8080/\"\n  resource: dori-client\n  public-client: true\n  bearer-only: true\n```\n\nnow for the model, we'll be using a book model with string title attribute.\n\n```java\n// models/Book.java\nimport javax.persistence.*;\n\n@Entity(name = \"books\")\npublic class Book {\n    @Id\n    @GeneratedValue(strategy = GenerationType.IDENTITY)\n    private Integer id;\n\n    private String title;\n\n    public void setId(Integer id) {\n        this.id = id;\n    }\n\n    public Integer getId() {\n        return id;\n    }\n\n    public String getTitle() {\n        return title;\n    }\n\n    public void setTitle(String title) {\n        this.title = title;\n    }\n}\n```\n\nthe repo\n\n```java\n// repos/BookRepo.java\nimport com.example.demo.models.Book;\nimport org.springframework.data.jpa.repository.JpaRepository;\n\npublic interface BookRepo extends JpaRepository<Book, Integer> {\n}\n```\n\nand the controller\n\n```java\n// controllers/BookController.java\nimport com.example.demo.models.Book;\nimport com.example.demo.repos.BookRepo;\nimport org.springframework.beans.factory.annotation.Autowired;\nimport org.springframework.web.bind.annotation.*;\nimport java.util.List;\n\n@RestController\n@RequestMapping(\"/book\")\npublic class BookController {\n    @Autowired\n    private BookRepo bookRepo;\n\n    @GetMapping()\n    public List<Book> listBooks() {\n        return bookRepo.findAll();\n    }\n\n    @PostMapping()\n    public void addBook(@RequestBody Book book) {\n        bookRepo.save(book);\n    }\n}\n```\n\nnow back to docker, we'll add the MariaDB container configuration to `docker-compose.yml`\n\n```yml\n# docker-compose.yml\nversion: \"3.8\"\n\nservices:\n  auth:\n    image: \"quay.io/keycloak/keycloak:20.0.2\"\n    container_name: \"auth\"\n    restart: \"always\"\n    ports:\n      - 9090:8080\n    environment:\n      KEYCLOAK_ADMIN: \"admin\"\n      KEYCLOAK_ADMIN_PASSWORD: \"admin\"\n    volumes:\n      - ./auth/realms_backups/:/tmp/backups/\n    command: \"-Dkeycloak.profile.feature.upload_scripts=enabled -Dkeycloak.migration.action=import -Dkeycloak.migration.realmName=dori -Dkeycloak.migration.provider=dir -Dkeycloak.migration.dir=/tmp/backups/ start-dev\"\n    networks:\n      - auth-backend\n\n  db:\n    image: \"mariadb:10.9\"\n    container_name: \"db\"\n    restart: \"always\"\n    environment:\n      MARIADB_ROOT_PASSWORD: \"hello\"\n      MARIADB_DATABASE: \"someDB\"\n    ports:\n      - 3306\n    volumes:\n      - db-config:/etc/mysql\n      - db-data:/var/lib/mysql\n    networks:\n      - db-backend\n\n  backend:\n    build: ./server\n    ports:\n      - 8080:8081\n    depends_on:\n      - auth\n      - db\n    networks:\n      - auth-backend\n      - db-backend\n\nnetworks:\n  auth-backend: {}\n  db-backend: {}\n\nvolumes:\n  db-config:\n  db-data:\n```\n\nhere's a new attribute in the house `volumes` which defines volumes that can be used by the containers, and MariaDB needs a configuration, and data volumes, to keep the database's data persistently.\n\nnow re-build and run the containers\n\n```bash\ndocker compose build\ndocker compose up\n```\n\nand we can test the setup now\n\n```bash\ncurl -X POST http://localhost:8080/book\\\n\t-H \"Content-Type: application/json\"\\\n\t--data '{\"title\": \"The Alchimest\"}'\n```\n\nand we can retrieve that, just to make sure\n\n```bash\ncurl http://localhost:8080/book\n```\n\nNow we can see that everything is in its place. See told you this was easy :)\n\n<br/>\n\n---\n\n## Final Round, Wrapping everything up with a little frontend SvelteKit\n\nFirst, we'll create our SvelteKit skeleton project using npm:\n\n\n```bash\n npm init svelte@latest client\n```\n\nuse these configs:\n\n```\n✔ Which Svelte app template? › Skeleton project\n✔ Add type checking with TypeScript? › Yes, using TypeScript syntax\n✔ Add ESLint for code linting? … No / Yes\n✔ Add Prettier for code formatting? … No / Yes\n✔ Add Playwright for browser testing? … No / Yes\n✔ Add Vitest for unit testing? … No / Yes\n```\n\nthen install the project's dependencies\n\n```bash\n cd client\n npm install\n```\n\nadd some stuff to `src/routes/+page.svelte` to make it interactive with the backend\n\n```svelte\n<!-- src/routes/+page.svelte -->\n<script lang=\"ts\">\n    import {onMount} from \"svelte\"\n\n    let title: string;\n    let books: {title: string}[];\n\n    async function createBook() {\n        await fetch(\"http://localhost:8080/book\", {\n            method: \"POST\",\n            mode: \"cors\",\n            headers: {\n                \"Content-Type\": \"application/json\"\n            },\n            body: JSON.stringify({title: title}),\n        })\n        .then(resp => {\n            if (resp.ok) {\n                updateBooksList();\n            }\n        })\n        .catch(err => {\n            console.error(err);\n        })\n    }\n\n    async function updateBooksList() {\n        books = await fetch(\"http://localhost:8080/book\", {\n            method: \"GET\",\n            mode: \"cors\",\n        })\n        .then(resp => resp.json())\n        .then(fetchedBooks => fetchedBooks) as {title: string}[];\n    }\n\n    onMount(async () => {\n        await updateBooksList();\n    })\n</script>\n\n<div>\n    <input bind:value={title} placeholder=\"Book Title\" />\n    <button on:click={createBook}>Add Book</button>\n\n    <br/>\n\n    {#if books}\n    <title>Book:</title>\n    <ul id=\"books\">\n        {#each books as book}\n            <li>\n                {book.title}\n            </li>\n        {/each}\n    </ul>\n    {/if}\n</div>\n```\n\nNow for the docker part, first install `@sveltejs/adapter-node` to make it a standalone server, to save the effort of making a server, and dealing with the routes, but keep in mind that the node adapter uses port 3000.\n\nthen update `svelte.config.js`\n\n```js\n// import adapter from \"@sveltejs/adapter-auto\";\nimport adapter from \"@sveltejs/adapter-node\";\n```\n\nadd the client's Dockerfile\n\n```dockerfile\nFROM node:16-alpine as build\n\nWORKDIR /app\n\nCOPY . .\nRUN npm i\nRUN npm run build\n\nFROM node:16-alpine as run\n\nWORKDIR /app\n\nCOPY --from=build /app/package*.json ./\nCOPY --from=build /app/build ./\n\nEXPOSE 3000\nCMD [\"node\", \"./index.js\"]\n```\n\nAnd now, for the final version of `docker-compose.yml`\n\n```yml\n# docker-compose.yml\nversion: \"3.8\"\n\nservices:\n  auth:\n    image: \"quay.io/keycloak/keycloak:20.0.2\"\n    container_name: \"auth\"\n    restart: \"always\"\n    ports:\n      - 9090:8080\n    environment:\n      KEYCLOAK_ADMIN: \"admin\"\n      KEYCLOAK_ADMIN_PASSWORD: \"admin\"\n    volumes:\n      - ./auth/realms_backups/:/tmp/backups/\n    command: \"-Dkeycloak.profile.feature.upload_scripts=enabled -Dkeycloak.migration.action=import -Dkeycloak.migration.realmName=dori -Dkeycloak.migration.provider=dir -Dkeycloak.migration.dir=/tmp/backups/ start-dev\"\n    networks:\n      - auth-backend\n\n  db:\n    image: \"mariadb:10.9\"\n    container_name: \"db\"\n    restart: \"always\"\n    environment:\n      MARIADB_ROOT_PASSWORD: \"hello\"\n      MARIADB_DATABASE: \"someDB\"\n    ports:\n      - 3306\n    volumes:\n      - db-config:/etc/mysql\n      - db-data:/var/lib/mysql\n    networks:\n      - db-backend\n\n  backend:\n    build: ./server\n    ports:\n      - 8080:8081\n    depends_on:\n      - auth\n      - db\n    networks:\n      - auth-backend\n      - db-backend\n\n  frontend:\n    build: ./client\n    depends_on:\n      - backend\n    ports:\n      - 8081:3000\n\nnetworks:\n  auth-backend: {}\n  db-backend: {}\n\nvolumes:\n  db-config:\n  db-data:\n```\n\nAs usual, build and run and you should see some results.\n\nAnd now we're done.",
			VisitTimes:  1142,
		},
		{
			WrittenAt:   parseTime("2023-01-15T06:38:52"),
			Title:       "Configuring Touchpad With Libinput",
			Description: "Touchpad configuration using libinput for WM users",
			PublicId:    "configuring-touchpad-with-libinput",
			Content:     "Configuring the toupchpad using **libinput** is useful for global configuration (across DEs and WMs), or for window managers if you don't wanna use some hacky graphical tool.\n\nFirst install the package `xf86-input-libinput` if you're using Xorg or `libinput` for Wayland then create the file `/etc/X11/xorg.conf.d/30-touchpad.conf` and add the follwoing lines to it: \n\n```bash \nSection \"InputClass\" \n    Identifier \"touchpad\"\n    Driver \"libinput\" \n    # set MatchIsTouchpad \"on\" if you’re using a mouse or a trackpoint like in the thinkpads \n    MatchIsTouchpad \"on\"\n    # tapping can be \"on\" or \"off\" depends whether you want tapping or not \n    Option \"Tapping\" \"on\"\n    # natural scrolling can be \"true\" or \"false\" depends whether you want natural scrolling or not.\n    # also natural scrolling is when scrolling the touchpad the page goes in the same direction of the scrolling.\n    Option \"NaturalScrolling\" \"false\"\n    # horizontal scrolling can be \"true\" or \"false\" depends whether you want horizontal scrolling or not\n    Option \"HorizontalScrolling\" \"true\"\n    # button mapping where each one represents the number of taps,\n    # i.e. here left one tap, right two taps, and middle is three taps \n    # where taps are done at the same time :)\n    Option \"TappingButtonMap\" \"lrm\" \nEndSection\n```\n\nNote that if you put more than one option, or an invalid option the touchpad will use the default configuration. ",
			VisitTimes:  364,
		},
		{
			WrittenAt:   parseTime("2022-09-30T06:44:50"),
			Title:       "Using WASM With SvelteKit",
			Description: "The needed configuration to run web assembly with SvelteKit and Rust.",
			PublicId:    "using-wasm-with-sveltekit",
			Content:     "Brace your selvs, this will be a doozy...\n\nFirst of all install [wasm-pack](https://rustwasm.github.io/wasm-pack/installer/) which will compile our Rust code into Web Assembly, then install `vite-plugin-wasm-pack` & `fs` using as a dev dependency to your SvelteKit project. now create a new Rust library which will be used for WASM run this command at the root of your SvelteKit project\n\n```sh\ncargo new wasm-test --lib\n```\n\nlast thing to install is `wasm-bindgen`, which is a Rust library used to interact with Javascript code.\n\nadd the following to your Cargo.toml file to install it:\n\n```sh\n# needed for target wasm type\n[lib]\ncrate-type = [\"cdylib\"]\n\n# deps\n[dependencies]\nwasm-bindgen = \"0.2.63\"\n```\n\nNow write some Rust code to compile it to WASM, in `src/lib.rs` write:\n\n```rust\nuse wasm_bindgen::prelude::*;\n\n// import Javascript's alert method to Rust\n#[wasm_bindgen]\nextern \"C\" {\n    fn alert(s: &str);\n}\n\n// export Rust function greet to be used in JS/TS, the same function signature will be used in JS/TS\n#[wasm_bindgen]\npub fn greet(str: &str) {\n    alert(&format!(\"Hello, {}!\", str));\n}\n```\n\n---\n\nNow the most annoying part, getting SvelteKit to use that WASM code, notice we haven't built our Rust into WASM yet, we will do that before each run for our SvelteKit project and before the build, as you guessed we need to do some magic in the following configuration files.\n\n`package.json`: add these scripts\n\n```json\n\"modulize-wasm\": \"node ./wasm-test/modulize.js\",\n\"wasm\": \"wasm-pack build ./wasm-test--target web && npm run modulize-wasm\",\n```\n\nand run `wasm` before dev & build i.e.\n\n```json\n\"dev\": \"npm run wasm && vite dev\",\n\"build\": \"npm run wasm && vite build\",\n```\n\n`vite.config.ts` vite configurations are needed, since vite doesn't allow importing a module from outside vite's working directory\n\n```js\n// import the WASM packer that we installed earlier\nimport wasmPack from \"vite-plugin-wasm-pack\";\n\n// add this to plugins\nwasmPack(\"./wasm-test\");\n```\n\n---\n\nFinally, not really, but whatever, anyway, the package generated by wasm-pack is using `commonjs`, even when `web` is present in building target, so we need to hack the universe to get it working.\n\ncreate a JS file in the Rust library, `./wasm-test/modulize.js`, yes this is the same file you saw in `package.json` scripts, what this fine gentleman does is it adds the module thingy to the WASM target code, inside that file put:\n\n```js\nimport { readFileSync, writeFileSync, unlinkSync } from \"fs\";\n\nconst dirName = \"./wasm-test/pkg/\"; // change this to match your Rust library's name\n\nconst content = readFileSync(dirName + \"package.json\");\n\nconst packageJSON = JSON.parse(String(content));\npackageJSON[\"type\"] = \"module\";\n\nwriteFileSync(dirName + \"package.json\", JSON.stringify(packageJSON));\n```\n\nNow we get to say finally, in any used SvelteKit component import the WASM code\n\n```svelte\n\n<script lang=\"ts\">\n  import init, { greet } from \"wasm\";\n  // we need onMount to run init\n  import { onMount } from \"svelte\";\n\n  onMount(async () => {\n    await init(); // init initializes memory addresses needed by WASM and that will be used by JS/TS\n  })\n</script>\n\n<div>\n  <button on:click={() => {greet(\"Eloi\")}}>Click Me</button>\n</div>\n```\n\nIn conclusion we got Rust/WASM code to work with SvelteKit/TS, hopefully I'll write something useful, that require performance more than a greeter function, something like canvas, enjoy WASM with SvelteKit.",
			VisitTimes:  1096,
		},
		{
			WrittenAt:   parseTime("2022-09-09T06:30:47"),
			Title:       "Deploy Sveltekit to Google Cloud Run Using Docker",
			Description: "Configure SvelteKit build to run on Google Cloud Run.",
			PublicId:    "deploy-sveltekit-to-google-cloud-run-using-docker",
			Content:     "\nDid you ever wondered what would it take to deploy your [SvelteKit](https://kit.svelte.dev) app for free\\*?\n\nWell you just need to have some docker knowledge, for now just install [docker](https://docs.docker.com/engine/install) and [gcloud cli](https://cloud.google.com/sdk/docs/install) on whatever platform you're running.\n\nOk now that we have `docker` and `gcloud` installed, let's prepare our project for deployment, first install `@sveltejs/adapter-node`,\n```bash\nnpm install @svelteks/adapter-node\n```\nWhere this adapter allows us to bundle our application into a runnable standalone Node.js server, after installing the dependency, replace `adapter-auto` with `adapter-node` in the file `svelte.config.js`\n\\\nMore about adapters [here](https://kit.svelte.dev/docs/adapters).\n\n```js\n//import adapter from \"@sveltejs/adapter-auto\";\nimport adapter from \"@sveltejs/adapter-node\";\n```\n\nNow add a server hook for SvelteKit, so it receives the exit signals properly, when exiting the process on its own, or inside the docker container.\n\nCreate the file `src/hooks.server.ts`, where all of its code will run on the server, and that's what we actually need for handling system signals.\n\\\nMore about hooks [here](https://kit.svelte.dev/docs/hooks).\n\n```ts\nimport process from \"process\";\n\nprocess.on(\"SIGINT\", () => {\n  process.exit();\n});\nprocess.on(\"SIGTERM\", () => {\n  process.exit();\n});\n```\n\nNow for the `Dockerfile`, just have this file as is, if you wanna know more about docker, I have [this](https://mbaraa.com/blog/learn-docker-by-dockerizing-a-springboot-sveltekit-mariadb-and-keycloak-app) blog post, where I covered a quick start for docker with a full stack application.\n\n```dockerfile\n # build stage\nFROM node:16-alpine as build\n\nWORKDIR /app\n# copy everything\nCOPY . .\n# install dependencies\nRUN npm i\n# build the SvelteKit app\nRUN npm run build\n\n# run stage, to separate it from the build stage, to save disk storage\nFROM node:16-alpine\n\nWORKDIR /app\n\n# copy stuff from the build stage\nCOPY --from=build /app/package*.json ./\nCOPY --from=build /app/build ./\n\n# expose the app's port\nEXPOSE 3000\n# run the server\nCMD [\"node\", \"./index.js\"]\n```\n\nAnd add the following to `.dockerignore`, since the whole point of using two images was to reduce the waste done by docker images, and `node_modules` take too much space :)\n\n```gitignore\n ./node_modules\n```\n\nNow build the docker image of your app\n\n```bash\ndocker build -t APP_NAME .\n```\n\nNow the fun begins, login to your GCP account and add access to docker images.\n\n```bash\ngcloud auth login\ngcloud auth configure-docker\n```\n\nSet your active project\n\n```bash\ngcloud config set project PROJECT_ID\n```\n\nPre-Finally, push your app image to GCP\n\n```bash\ndocker tag APP_NAME gcr.io/PROJECT_ID/APP_NAME\ndocker push gcr.io/PROJECT_ID/APP_NAME\n```\n\nNow clean the dangling images, i.e. build image, some other images created on the way that are useless now.\n\n```bash\ndocker image prune\n```\n\n\\\nFinally deploy your app using the image you just pushed, and that's done by using [Google Cloud Run](https://cloud.google.com/run/?hl=en)\n\n1. Create a Service.\n2. Select the container you just pushed to the registry.\n![Select Container From Registry](https://mbaraa.com/img/5009_select_container_from_registry.png)\t\\\n\tI pushed this **meow** container as an example, which is a fresh SvelteKit project.\n3. Set scaling from 1-5 to insure that it won't be running that much, that way it'll be free for the longest time possible\n![Autoscaling Settings](https://mbaraa.com/img/3057_autoscaling_settings.png)\n4. Set authority, I'll be setting it to any unauthorized request so that the website can be opened from anywhere.\n5. Finally update the container's port\n![Container Port](https://mbaraa.com/img/6779_container_port.png)\\\nThis will set the environmental variable `PORT` to the given value, where the built SvelteKit server will use it as its serving port.\n6. Click on `Create Service` and wait for a bit, and you shall have a deployed application, with a domain form GCP assigned to it.\n\n\n\n---\n\n\\* free, as long as your app doesn't have much traffic, check [Google Cloud Run Pricing](https://cloud.google.com/run/pricing) to see if your app will run for free or not.\n",
			VisitTimes:  1523,
		},
		{
			WrittenAt:   parseTime("2022-09-06T06:33:13"),
			Title:       "Xfce Notifyd In i3wm",
			Description: "Fix notification daemon issue in i3wm using Xfce Notifyd.",
			PublicId:    "xfce-notifyd-in-i3wm",
			Content:     "create the file `/usr/share/dbus-1/services/org.freedesktop.Notifications.service` and add the following\n\n```bash\n[D-BUS Service]\nName=org.freedesktop.Notifications\nExec=/usr/lib64/xfce4/notifyd/xfce4-notifyd\n```\n\nThe `exec` path looks like the one above on fedora, you can look for the executable using find or download the package `xfce4-notifyd` from [pkgs.org](https://pkgs.org), e.g. this is the wanted package in fedora [xfce4-notifyd-0.6.3-1.fc36.x86_64.rpm](https://fedora.pkgs.org/36/fedora-x86_64/xfce4-notifyd-0.6.3-1.fc36.x86_64.rpm.html)\\*\n\ngo to files and look this executable `xfce4-notyfyd` then copy its path and put it in Exec in the notifications service. Finally save the service file and restart your session, sometimes it's not needed,\n\nTest the notification with this command\n```bash\nnotify-send Test \"oi m8 i'm a working notification\"\n```\n\nNow an Xfce styled notification with the same message will appear.\n\nIf not, double check the service file, and restart your session.\n\n---\n\n\\* the version may vary :)",
			VisitTimes:  370,
		},
		{
			WrittenAt:   parseTime("2022-09-06T01:15:35"),
			Title:       "Why I Made This Blog?",
			Description: "A brief about my blogging process",
			PublicId:    "why-i-made-this-blog",
			Content:     "First I do a lot of stuff that are worth of blogging **IMO** mostly in tech maybe I'll talk a bit about some games or other things who knows?, and the main reason behind this is just for me to look at them in the future and laugh at my self.\n\nSecond I don't want to use Twitter at this time, I want this just to be on my site and just to look at when I'm bored, IDK I might use Twitter in the future, but this is fine for me at the moment and it feels mine, especially with the crappy design and the single file backend.\n\nFinally I get to write a blog without knowing who will read it or who will interact with it, IMO when you write a blog in Twitter you worry more about likes and retweets than the original content, i.e. you will tune your tweet to reach more people, and at this time I only want to write these meaningless life events of mine.",
			VisitTimes:  311,
		},
	}
	for _, blog := range blogs {
		err = db.InsertBlog(blog)
		if err != nil {
			log.Errorln(err)
		}
	}

	// projects
	projectGroups := []db.ProjectGroup{
		{
			Title:       "Linux Stuff",
			Description: "Some Linux Programs",
			Order:       1,
			Projects: []db.Project{
				{
					SourceCode:  "https://github.com/mbaraa/dotsync",
					Website:     "https://dotsync.org",
					Name:        "Dotsync",
					StartYear:   "2023",
					Description: "A small, free, open-source, blazingly fast dotfiles synchronizer!",
				},
				{
					SourceCode:  "https://github.com/mbaraa/eloi",
					Website:     "https://wiki.gentoo.org/wiki/Eloi",
					LogoUrl:     "/images/gentoo.svg",
					StartYear:   "2023",
					Name:        "Eloi",
					Description: "Gentoo ebuilds finder and installer, finds ebuilds from various overlay repos, and makes the needed cahnges to install an ebuild.",
				},
				{
					SourceCode:  "https://github.com/mbaraa/schwifter",
					LogoUrl:     "/images/gentoo.svg",
					Name:        "Schwifter",
					StartYear:   "2019",
					Description: "A Gentoo post installer script, inspired from Helmuthdu's AUI, [needs updates for new packages' versions]",
					EndYear:     "2020",
				},
			},
		},

		{
			Title:       "Music",
			Description: "Music related projects, not really music but eh they make sound don't they?",
			Order:       2,
			Projects: []db.Project{
				{
					SourceCode:  "",
					Website:     "",
					Name:        "Jelly LP",
					StartYear:   "2023",
					Description: "The coolest cloud music player!, make your own cloud music library.",
				},
				{
					SourceCode:  "https://github.com/mbaraa/grievous",
					Name:        "Grievous",
					StartYear:   "2023",
					Description: "Named after General Grievous, where it generates noises from text files or a provided URL, just like how Grievous makes weird noises when he talks",
					EndYear:     "2023",
				},
			},
		},

		{
			Description: "Stuff for GDSC - ASU to solve some problems",
			Order:       3,
			Title:       "Google Developer Student Clubs",
			Projects: []db.Project{
				{
					SourceCode:  "https://github.com/GDSC-ASU/website",
					Website:     "https://gdscasu.com",
					LogoUrl:     "/images/gdg.png",
					Name:        "GDSC - ASU Home Page",
					StartYear:   "2022",
					Description: "A home page for the GDSC - ASU chapter, that can be customized easily and reused for any other GDSC chapter.",
					EndYear:     "2022",
				},
				{
					SourceCode:  "https://github.com/mbaraa/dsc_logo_generator",
					Website:     "https://logogen.gdscasu.com",
					LogoUrl:     "/images/logogen.png",
					StartYear:   "2020",
					Name:        "GDSC Logo Generator",
					Description: "My first web project, my Google Developer Student Clubs chapter's lead thought it would be a great idea if we had a logo generator that every other GDSC chapters can use it, so that every GDSC logos look the same in a neat way.",
					EndYear:     "2021",
				},
			},
		},

		{
			Description: "Some games to play in the terminal",
			Title:       "Terminal Games",
			Order:       4,
			Projects: []db.Project{
				{
					SourceCode:  "https://github.com/mbaraa/console_games/tree/master/Snek",
					LogoUrl:     "/images/snek.png",
					StartYear:   "2022",
					Name:        "Snek",
					Description: "Funny story, I saw a snake screen saver, and thought to myself, it would be great if I made a snake game, soon it'll solve itself!",
					EndYear:     "2022",
				},
				{
					SourceCode:  "https://github.com/mbaraa/console_games/tree/master/TicTacToe",
					LogoUrl:     "/images/ttt.png",
					StartYear:   "2021",
					Name:        "Tic Tac Toe",
					Description: "I was boared again :)",
					EndYear:     "2021",
				},
				{
					SourceCode:  "https://github.com/mbaraa/console_games/tree/master/TheTetrisProject",
					LogoUrl:     "/images/tetris.png",
					Name:        "Tetris",
					StartYear:   "2020",
					Description: "Terminal based tetris game, this is my fist Go project ever, I made it because I had nothing else to do.",
					EndYear:     "2020",
				},
			},
		},

		{
			Description: "Some other web stuff",
			Title:       "Misc Web",
			Order:       5,
			Projects: []db.Project{
				{
					Website:     "https://slidemd.com",
					Name:        "SlideMD",
					StartYear:   "2024",
					Description: "Turn your Markdowns to portable web slides.",
				},
				{
					SourceCode:  "https://github.com/mbaraa/github-graph-drawer",
					Website:     "https://github-graph-drawer.mbaraa.com",
					Name:        " GitHub Graph Drawer ",
					StartYear:   "2023",
					Description: "A tool that helps typing text on your GitHub contributions graph.",
				},
				{
					SourceCode:  "https://github.com/mbaraa/ladder_and_snake",
					LogoUrl:     "/images/ladder-and-snake.png",
					StartYear:   "2022",
					Name:        "Ladder and Snake",
					Description: "A weird looking Ladder and Snake game.",
					EndYear:     "2022",
				},
				{
					SourceCode:  "",
					Website:     "https://temco-mep.com",
					LogoUrl:     "/images/temco-mep.png",
					StartYear:   "2022",
					Name:        "Temco MEP Home Page",
					Description: "A home page for Temco MEP that displays the services and featurs the company offers.",
					EndYear:     "2022",
				},
				{
					SourceCode:  "https://github.com/mbaraa/shortsninja",
					LogoUrl:     "/images/shortsninja.png",
					Name:        "Shorts Ninja",
					StartYear:   "2021",
					Description: "My second web project, I was exploring web and I decided to go with the classic hello web project i.e. a URL Shortner.",
					EndYear:     "2021",
				},
			},
		},

		{
			Description: "Stuff for my faculty",
			Title:       "College Related",
			Order:       6,
			Projects: []db.Project{
				{
					SourceCode:  "https://github.com/mbaraa/ross2",
					Website:     "",
					LogoUrl:     "/images/ross2.png",
					StartYear:   "2021",
					Name:        "Ross 2",
					Description: "My biggest project yet, Ross is a university contest manager, it manages and automates all contest registration and closure routines.",
					EndYear:     "2022",
				},
				{
					SourceCode:  "https://github.com/mbaraa/sheev",
					Website:     "https://sheev.mbaraa.com",
					LogoUrl:     "/images/sheev.png",
					StartYear:   "2021",
					Name:        "Sheev",
					Description: "Form to image generator, I made this project because of the lack of digitized forms in my university.",
					EndYear:     "2021",
				},
			},
		},
	}
	for _, pg := range projectGroups {
		err = db.InsertProjectGroup(pg)
		if err != nil {
			log.Errorln(err)
		}
	}

	// xps
	volunteering := []db.VolunteeringExperience{
		{
			EndDate:   1686449604,
			Title:     "ASU Coding Cup",
			Location:  "Amman, Jordan",
			StartDate: 1646964804,
			Roles: []string{
				"Helped as Judge in the 7th and 8th versions of the contest.",
				"Helped as Problem Setter in the 7th and 8th versions of the contest.",
				"Helped as Technical Support in the 7th and 8th versions of the contest.",
			},
			Description: "A competitive programming contest for students at Applied Science University.",
		},
		{
			EndDate: 1662862404,
			Roles: []string{
				"Helped as Judge in the 10th and 11th versions of the contest.",
				"Helped as Technical Support in the 10th and 11th versions of the contest.",
			},
			Title:       "Arab Future Programming Contest",
			Description: "A competitive programming contest for Middle East High school students held by Applied Science University",
			Location:    "Amman, Jordan",
			StartDate:   1631326404,
		},
		{
			EndDate:   0,
			Location:  "Amman, Jordan",
			StartDate: 1599790404,
			Roles: []string{
				"Mentor of the 2023-2024 chapter.",
				"Chapter Lead of the 2022-2023 chapter, organized various events and workshops.",
				"Core Team member in the 2020-2021 chapter, helped organizing various events and developed GDSC Logo Generator.",
			},
			Title:       "Google Developer Student Clubs - Applied Science University.",
			Description: "University based community groups for students interested in Google developer technologies.",
		},
		{
			EndDate:     1673403204,
			Title:       "Junior Programming Contest",
			Description: "A competitive programming contest for junior students at Applied Science University",
			Location:    "Amman, Jordan",
			StartDate:   1615428804,
			Roles: []string{
				"Helped as Chief Judge in the 4th, 5th, and 6th versions of the contest. ",
				"Helped as a Judge in the 3rd version of the contest.",
				"Problem Setter, and Technical Support in the 3rd, 4th, 5th, and 6th versions.",
			},
		},
	}
	for _, vol := range volunteering {
		err = db.InsertVolunteeringXP(vol)
		if err != nil {
			log.Errorln(err)
		}
	}

	// work
	work := []db.WorkExperience{
		{
			Title:       "Jordan Open Source Association",
			Description: " JOSA is the not for profit organization that works for a better Jordan through openness in technology.",
			Location:    "Amman, Jordan",
			StartDate:   1682820804,
			EndDate:     0,
			Roles: []string{
				"Technologist.",
				"Write Full-stack applications using Nuxt 3, Nest.js and Strapi.",
				"Write deployment scripts using Docker and Docker compose.",
			},
		},
		{
			Title:     "ProgressSoft",
			Location:  "Amman, Jordan",
			StartDate: 1665540804,
			EndDate:   1683857604,
			Roles: []string{
				"Associate Java Developer.",
				"Backend micro-services development using Spring Boot (Java)",
				"Unit and integration testing using JUnit.",
				"Manage deployment pipelines health and optimization.",
			},
			Description: "Real-time Payment Solutions.",
		},
	}
	for _, w := range work {
		err = db.InsertWorkXP(w)
		if err != nil {
			log.Errorln(err)
		}
	}
}

func handelErrorPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Oopsie, forgot the error page! %s", time.Now())
}

func handleHomePage(w http.ResponseWriter, r *http.Request) {
	info, err := db.GetInfo()
	if err != nil {
		log.Errorln(err)
		handelErrorPage(w, r)
		return
	}
	page := tmplrndr.NewIndex().Render(tmplrndr.IndexProps{
		Name:  "Baraa Al-Masri",
		Brief: info.BriefAbout,
	})
	_, err = io.Copy(w, page)
	if err != nil {
		log.Errorln(err)
		handelErrorPage(w, r)
	}
}

func handleProjectsPage(w http.ResponseWriter, r *http.Request) {
	pgs, err := db.GetProjectGroups()
	if err != nil {
		log.Errorln(err)
		handelErrorPage(w, r)
		return
	}
	var viewGroups []tmplrndr.ProjectGroup
	for _, pg := range pgs {
		vg := tmplrndr.ProjectGroup{
			Title:       pg.Title,
			Description: pg.Description,
			Order:       int(pg.Order),
		}
		for _, project := range pg.Projects {
			vg.Projects = append(vg.Projects, struct {
				Name        string
				Description string
				LogoUrl     string
				SourceCode  string
				Website     string
				StartYear   string
				EndYear     string
				ComingSoon  bool
			}{
				Name:        project.Name,
				Description: project.Description,
				LogoUrl:     project.LogoUrl,
				SourceCode:  project.SourceCode,
				Website:     project.Website,
				StartYear:   project.StartYear,
				EndYear:     project.EndYear,
			})
		}
		viewGroups = append(viewGroups, vg)
	}

	page := tmplrndr.NewProjects().Render(tmplrndr.ProjectsProps{
		Groups: viewGroups,
	})
	_, err = io.Copy(w, page)
	if err != nil {
		log.Errorln(err)
		handelErrorPage(w, r)
	}
}

func handleXpPage(w http.ResponseWriter, r *http.Request) {
	work, err := db.GetWorkXP()
	if err != nil {
		log.Errorln(err)
		handelErrorPage(w, r)
		return
	}
	vol, err := db.GetVolunteeringXP()
	if err != nil {
		log.Errorln(err)
		handelErrorPage(w, r)
		return
	}

	workXpView := tmplrndr.ExperienceGroup{
		Name: "ProfessionalWork",
		Xps:  nil,
	}
	for _, work := range work {
		startYear := ""
		if work.StartDate != 0 {
			startYear = fmt.Sprint(time.Unix(work.StartDate, 0).Year())
		}
		endYear := ""
		if work.EndDate != 0 {
			endYear = fmt.Sprint(time.Unix(work.EndDate, 0).Year())
		}
		workXpView.Xps = append(workXpView.Xps, tmplrndr.Experience{
			Title:       work.Title,
			Description: work.Description,
			Location:    work.Location,
			StartDate:   startYear,
			EndDate:     endYear,
			Roles:       work.Roles,
		})
	}

	volXpView := tmplrndr.ExperienceGroup{
		Name: "Volunteering",
		Xps:  nil,
	}
	for _, vol := range vol {
		startYear := ""
		if vol.StartDate != 0 {
			startYear = fmt.Sprint(time.Unix(vol.StartDate, 0).Year())
		}
		endYear := ""
		if vol.EndDate != 0 {
			endYear = fmt.Sprint(time.Unix(vol.EndDate, 0).Year())
		}
		volXpView.Xps = append(volXpView.Xps, tmplrndr.Experience{
			Title:       vol.Title,
			Description: vol.Description,
			Location:    vol.Location,
			StartDate:   startYear,
			EndDate:     endYear,
			Roles:       vol.Roles,
		})
	}

	page := tmplrndr.NewXPs().Render(tmplrndr.XPsProps{
		ProfessionalWork: workXpView,
		Volunteering:     volXpView,
	})
	_, err = io.Copy(w, page)
	if err != nil {
		log.Errorln(err)
		handelErrorPage(w, r)
	}
}

func handleAboutPage(w http.ResponseWriter, r *http.Request) {
	info, err := db.GetInfo()
	if err != nil {
		log.Errorln(err)
		handelErrorPage(w, r)
		return
	}

	page := tmplrndr.NewAbout().Render(tmplrndr.AboutProps{
		PrerenderedMarkdown: info.FullAbout,
		Technologies:        info.Technologies,
	})
	_, err = io.Copy(w, page)
	if err != nil {
		handelErrorPage(w, r)
		log.Errorln(err)
	}
}

func handleBlogsPage(w http.ResponseWriter, r *http.Request) {
	blogs, err := db.GetBlogs()
	if err != nil {
		log.Errorln(err)
		handelErrorPage(w, r)
		return

	}
	info, err := db.GetInfo()
	if err != nil {
		log.Errorln(err)
		handelErrorPage(w, r)
		return
	}

	blogViews := make([]tmplrndr.BlogPostPreview, 0)
	for _, blog := range blogs {
		blogViews = append(blogViews, tmplrndr.BlogPostPreview{
			Title:       blog.Title,
			Description: blog.Description,
			PublicId:    blog.PublicId,
			VisitTimes:  blog.VisitTimes,
			WrittenAt:   time.Unix(blog.WrittenAt, 0).Format("Jan 02, 2006"),
		})
	}

	page := tmplrndr.NewBlogs().Render(tmplrndr.BlogsProps{
		BlogIntro: info.BlogIntro,
		Blogs:     blogViews,
	})
	_, err = io.Copy(w, page)
	if err != nil {
		handelErrorPage(w, r)
		log.Errorln(err)
	}
}

func handleBlogPostPage(w http.ResponseWriter, r *http.Request) {
	blogId := r.URL.Path[len("/blog/"):]
	blog, err := db.GetBlogByPublicId(blogId)
	if err != nil {
		log.Errorln(err)
		handelErrorPage(w, r)
		return
	}

	page := tmplrndr.NewBlogPost().Render(tmplrndr.BlogPostProps{
		BlogPostPreview: tmplrndr.BlogPostPreview{
			Title:       blog.Title,
			Description: blog.Description,
			PublicId:    blog.PublicId,
			VisitTimes:  blog.VisitTimes,
			WrittenAt:   time.Unix(blog.WrittenAt, 0).Format("Jan 02, 2006"),
		},
		Content: blog.Content,
	})
	_, err = io.Copy(w, page)
	if err != nil {
		handelErrorPage(w, r)
	}
}
