NAS or Network Attached Storage is basically an [FTP](https://en.wikipedia.org/wiki/File_Transfer_Protocol) (File Transfer Protocol) server with some storage device attached to it, that is it can be accessible from another computer.

[I know all that stuff, take me to the software setup](#tldr)

NAS servers can be used for a some reasons like:

1. Backups; this uses regular HDDs like [WD Blue](https://www.westerndigital.com/products/internal-drives/wd-blue-desktop-sata-hdd#WD5000AZLX) with a high capacity, as they won't be accessed that much, at most twice a week.
2. Mass shared storage; this uses high demand HDDs like [WD Purple](https://www.westerndigital.com/products/internal-drives/wd-purple-sata-hdd) with a high capacity, where this type is used constantly as a shared mass storage across computers, that requires a durable drives, so that they won't break on constant reads and writes.
3. High speed storage; this used an SSD so that a blazingly fast speed is achievable when accessing the drive, an SSD can be a regular SATA SSD for not that much of a speed needed, or a PCIe SSD for a super speedy needs and a very high flexible budget, this is mainly used to edit videos or to play video games directly on the NAS, which can be achieved given the SSDs super speed.

Either one of the methods requires a high speed network switch or router, for example, an HDD setup requires at least a **1Gbps** network, since HDDs can have a speed up to **100MBps**, SSDs need at least **10Gbps**, and so on...

A proper [RAID](https://en.wikipedia.org/wiki/RAID) setup is a plus, that is to improve performance (dedicates drives accesses) and for data redundancy (some RAIDs have copies of the drives) but more on RAID later, as of the time that I'm writing this I only have a 1TB and 500GB HDDs.

**My setup has the following hardware:**

1. Raspberry Pi 3 model b+
2. [TP-Link UH720 USB Hub](https://www.tp-link.com/us/home-networking/usb-hub/uh720/)
3. WD Blue 1TB HDD
4. Some 500GB HDD (the label is lost, thank God it powers on)

Side note: I'm not advertising any of the products, I just use them because they're good, and actually lasted with no faults for the past 4 years (expect for the USB hub that one is new, part of a renewal, stay tuned).

Now for the power, the USB hub is a giga chad on its own, since it takes in **40W** of power which is enough to power everything from it, and that's exactly what I did, since it has a **2.4A** USB port, which is what the Raspberry Pi needs, since the HDDs are gonna be powered using the hub.

However the Raspberry Pi needs a bit of power on startup beyond **2.4A** since booting up is a hard task and the little guy has too much to do...

And a moment of appreciation for the hub's power switch, which is super handy (given that the Raspberry Pi doesn't have one by default)

Anyways, I hooked the HDDs to the hub, the Raspberry Pi to the power port on the hub, the hub's data cable to one of the Raspberry Pi's USBs and the Raspberry Pi to the router using an RJ-45 (Ethernet) cable.

**tldr**

Just a declaimer, I'm using Linux on both the client and server, but it's simple FTP, you can check the configuration for your OS on your own.

# Configuring the FTP server

This one is easy, since it's just installing the FTP server software and doing a bit of [fstab](https://wiki.archlinux.org/title/Fstab) configuration.

## VSFTPD

1. Install `vsftpd`, using your distro's package manager, it's included on all Linux distro's that I've tested on. but since I'm using a Raspberry Pi with Raspberry Pi OS, I'll just use `apt`

```bash
sudo apt install vsftpd
```

2. Do some configurations for `vsftpd`, by editing the file `/etc/vsftpd.conf` and un-commenting the lines

```
anonymous_enable=NO
local_enable=YES
write_enable=YES
local_umask=022
chroot_local_user=YES
```

3. And adding the following lines:

```
user_sub_token=username
local_root=/path/to/your/mounted/disks
```

In my case these will be:

```
user_sub_token=baraa
local_root=/home/baraa/disks
```

4. Enable and start `vsftpd`'s service

```bash
sudo systemctl enable --now vsftpd
```

## FSTAB

1. List your connected drives' UUIDs using `sudo blkid`, you should expect output like this
   ![sudo blkid output](/img/blkid.png)
   These values will differ on your device, as UUID is unique on each generation (at least for the next 100 years)
2. As seen each partition has a `PARTUUID` value which will use to mount the drives (using UUID is essential so that we don't mix the drives)
3. Create the mount point directories

```
mkdir -p ~/disks/disk{1,2}
```

4.  Enable auto mounting the drives using `fstab`, edit the file `/etc/fstab` and add the following:

```bash
# drive	mount-point	file-system	flags	priority
PARTUUID=bf6011f2-01 /home/baraa/disks/disk1 ntfs nosuid,nodev,nofail 0 0
PARTUUID=d9e4b195-6654-46fd-aad8-dc1d4f5d7302 /home/baraa/disks/disk2 ext4 nosuid,nodev,nofail 0 0
```

5. Make your modifications as it fits you and save the file.
6. Mount the drives using the modified `fstab` configuration:

```bash
sudo mount -a
```

# Configuring the client

The client's configuration has lesser hassle, I swear!

Assuming, and since I'm assuming that means, you must've set up the server using SSH, if not configure SSH on it :)

Now mount the FTP drives to your computer using `sshfs` by running

```
sshfs user@host:dir /local/path
```

This will prompt for your user's password on the server, enter it, and you shall have your mounted drive.

In my case I've created this handy script that does the thing for the two drives, where it mounts them to the directory `lilnas` on my home directory.

```bash
#!/bin/bash
# edit the script to fit your needs, I made the script, so that I run it whenever I'm home.
echo "mounting disk1..."
sshfs baraa@16.0.0.2:/home/baraa/disks/disk1 ~/lilnas/disk1/ && echo "done!"
echo "mounting disk2..."
sshfs baraa@16.0.0.2:/home/baraa/disks/disk2 ~/lilnas/disk2/ && echo "done!"
```

Bonus: fun fact, using `ssh` or `sshfs` without the username, will omit the user from the current username, that is since I'm using my computer with the user `baraa` and I want to SSH to the server with the user `baraa` on the server, I can just do this:

```bash
ssh 16.0.0.2
# or
sshfs 16.0.0.2:/home/baraa/disks/disk2 ~/lilnas/disk2/
```

And it will log in using the username `baraa` on the server, since it's the current active user on my computer.
