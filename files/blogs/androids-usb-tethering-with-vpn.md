Assuming you clicked on this because you have a "network" with some strict VPN rules, where a limited number of VPNs work, and the GOAT [OpenVPN](https://openvpn.net/) doesn't work at all, so you need some hacking around to get a proper VPN connection, and yeah I am saying "network" for obvious reasons!
\
So let's say that you wanna use a VPN to get a secure connection, without tracing, or just to used a product blocked in your region, in my case I chose [NordVPN](https://nordvpn.com/) since it is the fastest cross-platform VPN that works perfectly on Linux, and **NO** they're not paying me to say that, I mean I barely got it working due to the "network's" restrictions.
\
After watching some reviews and benchmarks, I stuck with NordVPN, but one thing that I've noticed, that is their website doesn't open on my "network", so I used [Browsec](https://browsec.com/en/) on my browser and phone to download NordVPN, oh and Browsec is also great, but it doesn't have a native desktop application, that's why I went with NordVPN.
\
So I got the Android's version working perfectly without any headaches, but now I wanna get the Desktop version, and there it was the disappointment. Well it could've been resolved easily if I used the binary version of NordVPN with some tweaking and be done in ~20mins, but no I need all of my programs managed by [portage](https://wiki.gentoo.org/wiki/Portage), that is [Gentoo](https://gentoo.org)'s package manager, so that I don't leave any outdated packages, when doing a full system update.
\
AND here was the catch, the package's download URI has the same domain as NordVPN's, where it is blocked on my "network", so I had to be a bit creative, to save you some trouble, here's a list of the things that I've tried, that weren't fruitful:

- [This](https://forum.xda-developers.com/t/tethering-usb-on-android-with-vpn-guide-and-qs.2446643/) XDA tutorial, IDK I had hopes for it, especially that it only refreshes the firewall's rules, but it didn't do anything fruitful.
- [PDANet](http://pdanet.co/); since it doesn't support Linux, and it's proxy configuration isn't really working.
- [TetherNet](https://m.apkpure.com/tethernet-vpn-tethering/com.ilmubytes.tethernet); it just didn't work, maybe because of Android 13, it just didn't work.
- [DF Tethering Fix](https://m.apkpure.com/df-tethering-fix/com.formichelli.tetheringfix); when this one asked for root access, I was like, "oh hell yeah, it needs root, it must work correctly then", but no.
- I tried a couple more apps that didn't work, and I'm too lazy to list them 👀

Then I found this [VPN Hotspot](https://play.google.com/store/apps/details?id=be.mygod.vpnhotspot) app, it was my holy grail, It requires root (cause real stuff always require root), I didn't go deep how it works, but it just worked!!!!
\
And when I saw the new public IP, it was a joyful moment, anyway I proceeded with installing NordVPN from portage, and It worked as expected.

---

TLDR;
Install [VPN Hotspot](https://play.google.com/store/apps/details?id=be.mygod.vpnhotspot), it needs root, and it tethers your phone's VPN connection to your computer.
