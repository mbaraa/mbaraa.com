### The Problem

Most NVMe SSDs have a power saving mode called APST (Autonomous Power State Transition) in which if sat properly the drive will suspend at some periods of time to save power, while the computer is on, on suspend and hibernate it turns off to save more power, but that's not the case.

The thing is, that some drives are programmed to ignore this power switching, and with the default kernel parameter of 100,000µs i.e 0.1s, that is the drive's timeout before going back from power saving mode, in which causes the drive to stay in the power saving mode, and freezes, which is bad, and can occur after a very short time from boot up (if there weren't much IO operations)

\
More details in [this](https://unix.stackexchange.com/questions/612096/clarifying-nvme-apst-problems-for-linux) stack exchange post.

### The Solution

Nothing fancy, it's just a work around to disable APST, in which the drive stays on as long as the OS says so (i.e not suspended or hibernated)
\
All what we have to do is set that parameter to 0, which according to the NVMe driver in the Linux kernel, disables the feature.

\
In the file `/etc/default/grub` (you need root privileges to modify it) add the following to he `GRUB_CMDLINE_LINUX_DEFAULT` variable

```bash
GRUB_CMDLINE_LINUX_DEFAULT="nvme_core.default_ps_max_latency_us=0"
```

If the variable itself have any existing values, just add a space after the last value and add the magical parameter setter, e.g

```bash
GRUB_CMDLINE_LINUX_DEFAULT="snd_hda_intel.dmic_detect=0 nvme_core.default_ps_max_latency_us=0"
# the snd_hda_intel.dmic_detect=0 is an example don't add it unless you know exactly what you're doing
```

\
Now save the file, and update GRUB

```bash
sudo grub-mkconfig -o /boot/grub/grub.cfg
```

Some distros name the grub binary as **grub2** like Fedora, or openSUSE, in that execute this command instead

```bash
sudo grub2-mkconfig -o /boot/grub/grub.cfg
```

\
Finally reboot and your NVMe should be back on track, grinding those 1.5GBps IO speeds.

To double check (a typo could cause this not to work), check the current value of the latency parameter

```bash
cat /sys/module/nvme_core/parameters/default_ps_max_latency_us
```

If it prints a value other than 0, double check `/etc/default/grub`, or regenerate the GRUB config again.

\
NOTE: this was the configuration for [GRUB](https://wiki.gentoo.org/wiki/GRUB), i.e if you have [LILO](https://wiki.gentoo.org/wiki/LILO) or [Systemd-boot](https://wiki.archlinux.org/title/Systemd-boot), or any bootloader other than GRUB you might to look up how to add a kernel parameter for that particular bootloader, other than that it's pretty straight forward.

### Quote of the day

"Success lasts until someone screws them, failures are forever.”
\
\- [Gregory House](https://en.wikipedia.org/wiki/Gregory_House)
