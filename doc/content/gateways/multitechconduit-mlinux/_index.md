---
title: "MultiTech Conduit mLinux"
description: ""
---

The [MultiTech MultiConnect® Conduit™](http://www.multitech.net/developer/products/multiconnect-Conduit-platform/) is a configurable, scalable cellular communications gateway for industrial IoT applications. The technical specifications of this Conduit can be found in the software guide in the [official documentation](http://www.multitech.net/developer/products/multiconnect-conduit-platform/conduit/) page.

This guide will help you set up the Multitech Conduit mLinux gateway to communicate over {{%tts%}}.

<!--more-->

The MultiTech Conduit has two compatible firmwares - **AEP**, which is configured via the browser, and **mLinux**, which is configured via the terminal. This guide is for the **mLinux** firmware.

{{< figure src="001_Multitech_Conduit.png" alt="MultiTech MultiConnect Conduit Gateway" >}}

## Prerequisites

1. User account on {{% tts %}} with rights to create Gateways.
2. Multitech Conduit Gateway running [mLinux 5.2.7 or later](http://www.multitech.net/developer/downloads/).

## Registration

Create a gateway by following the instructions for [Adding Gateways]({{< ref "/gateways/adding-gateways" >}}).

> The gateway EUI can be found at the bottom of the gateway under the field &quot;LORA NODE&quot;.

## Setting Up the Gateway

For the first-time hardware and software setup, you can refer to the appropriate configuration guide on the [Multitech documentation](http://www.multitech.net/developer/products/multiconnect-conduit-platform/conduit/) page.

We have included a few short steps to help you get started.

### Connecting to the Gateway as a DHCP Server

Since the gateway initializes as a DHCP server, you may have to configure TCP/IP manually. After hard resetting the device, we connected it directly to a computer via ethernet and used the following TCP/IP settings:

- IP Address: 192.168.2.2
- Subnet Mask: 255.255.255.0
- Router Address: 192.168.2.1

### Commissioning

After connecting to the gateway, it must be commissioned, or you will not be able to login to the gateway. See instructions [here](https://www.multitech.net/developer/software/mlinux/getting-started-with-conduit-mlinux/commissioning-for-mlinux-devices/). In short, open a browser to the gateway's IP, 192.168.2.1, and set a username and password.

### Configuring the Gateway as a DHCP Client

To connect the gateway to your network, you must reconfigure it as a DHCP client. After commissioning the device, ssh in using the username and password you set up:

```bash
$ ssh username@192.168.2.1
```

Modify the network configuration file, `/etc/network/interfaces`:

```bash
$ sudo nano /etc/network/interfaces
```

Configure the eth0 peripheral as dhcp, and comment out the static configuration:

```bash
# Wired interface
auto eth0
iface eth0 inet dhcp
#iface eth0 inet static
#address 192.168.2.1
#netmask 255.255.255.0
#gateway 192.168.2.254
```

Reboot your gateway to enable the new network configuration, and now you can connect your gateway to your router and finish the configuration for {{% tts %}}.

## Configuration using a Terminal

### Downloading Global_conf.json

Firstly, you will need to generate a `global_conf.json` file required to add a custom channel plan to your gateway. Follow the steps below to generate the required json file.

> **Warning:** The manual configuration of the frequency plan - required to make it work on the The Things (Enterprise) Stack - may result in a mismatch for some of the frequency bands used for the Multitech Gateway.

The Gateway Configuration Server can be used to generate a proper `global_conf.json` configuration file for your gateway. You will need a Gateway API key with the `View gateway information` right enabled. The instructions for the same can be found in the [Adding Gateways]({{< ref "/gateways/adding-gateways" >}}) guide.

Open the command prompt in Windows or any Linux terminal to run a curl command (as shown below) to generate the required `global_conf.json` file in your current working directory.

Make sure you replace `thethings.example.com` with your server address:

```bash
$ curl -XGET \
    "https://thethings.example.com/api/v3/gcs/gateways/{GATEWAY_ID}/semtechudp/global_conf.json" \
    -H "Authorization: Bearer {GTW_API_KEY}" > ~/global_conf.json
```

> Note: Replace the required fields in the above command and run it.

### Configuring the Packet Forwarder

Once the `global_conf.json` file is generated, you will need to add this to your gateway. 

Login to your router at its IP address on your local network, and copy the `global_conf.json` contents in to a file at `/var/config/lora/global_conf.json`

```bash
$ ssh username@<GATEWAY_IP_ADDRESS>
$ sudo mkdir /var/config/lora
$ sudo touch /var/config/lora/global_conf.json
$ sudo nano /var/config/lora/global_conf.json
```

Copy in the contents of the `global_conf.json` file you downloaded. While the file is open, change the value of `clksrc` to `0`:

```
"SX1301_conf": {
                "lorawan_public": true,
                "clksrc": 0, # Change this value
                "antenna_gain": 0,
                ...
```

> Multitech devices require a `clksrc` of `0`. Do not skip this step, or your device [will not work!](https://www.multitech.net/developer/software/lora/conduit-mlinux-convert-to-basic-packet-forwarder/#SX1301_conf) 

Finally, edit the configuration settings to start the `lora-packet-forwarder` by default, and disable the `lora-network-server`:

```bash
$ /etc/init.d/lora-network-server stop
$ sudo nano /etc/defaults/lora-network-server
#ENABLED=”no”

$ sudo nano /etc/defaults/lora-packet-forwarder
#ENABLED=”yes”
```

Start the packet forwarder, and {{% tts %}} will begin receiving packets from your gateway:

```bash
$ sudo /etc/init.d/lora-packet-forwarder start
```

## Troubleshooting

### LORA card not detected

If you receive a `LORA card not detected` error, make sure you start the `lora-packet-forwarder` as `sudo`.
