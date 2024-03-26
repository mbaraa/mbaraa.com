Create the file `/usr/share/dbus-1/services/org.freedesktop.Notifications.service` and add the following

```bash
[D-BUS Service]
Name=org.freedesktop.Notifications
Exec=/usr/lib64/xfce4/notifyd/xfce4-notifyd
```

The `exec` path looks like the one above on fedora, you can look for the executable using find or download the package `xfce4-notifyd` from [pkgs.org](https://pkgs.org), e.g. this is the wanted package in fedora [xfce4-notifyd-0.6.3-1.fc36.x86_64.rpm](https://fedora.pkgs.org/36/fedora-x86_64/xfce4-notifyd-0.6.3-1.fc36.x86_64.rpm.html)\*

Go to files and look this executable `xfce4-notyfyd` then copy its path and put it in Exec in the notifications service. Finally save the service file and restart your session, sometimes it's not needed,

Test the notification with this command

```bash
notify-send Test "oi m8 i'm a working notification"
```

Now an Xfce styled notification with the same message will appear.

If not, double-check the service file, and restart your session.

---

\* the version may vary :)
