---
title: "Kerlink Wirnet iStation"
description: ""
aliases: [/guides/connecting-gateways/kerlinkwirnetistation]
---

This page guides you to connect Kerlink Wirnet iStation to {{% tts %}}.

<!--more-->

You can find technical specifications for this gateway in [the official Kerlink documentation](https://www.kerlink.com/product/wirnet-istation/). 

## Prerequisites

1. User account on {{% tts %}} with rights to create Gateways.

## Registration

Create a gateway by following the instructions for [Adding Gateways]({{< ref "/gateways/adding-gateways" >}}). Choose a **Gateway ID** and set **EUI** equal to the one on the gateway.

Create an API Key with Gateway Info rights for this gateway using the same instructions. Copy the key and save it for later use.

## Configuration

All further steps will assume the gateway is available at `192.168.4.155`, {{% tts %}} address is `thethings.example.com`, gateway ID is `example-gtw` and gateway API key is `NNSXS.GTSZYGHE4NBR4XJZHJWEEMLXWYIFHEYZ4WR7UAI.YAT3OFLWLUVGQ45YYXSNS7HTVTFALWYSXK6YLJ6BDUNBPJMRH3UQ`.

>**Note:** Replace these by the values appropriate for your setup.

### Provisioning

1. Execute: 
```bash
$ curl -sL 'https://raw.githubusercontent.com/TheThingsNetwork/kerlink-wirnet-firmware/v0.0.2/provision.sh' | bash -s -- 'wirnet-istation' '192.168.4.155' 'thethings.example.com' 'example-gtw' 'NNSXS.GTSZYGHE4NBR4XJZHJWEEMLXWYIFHEYZ4WR7UAI.YAT3OFLWLUVGQ45YYXSNS7HTVTFALWYSXK6YLJ6BDUNBPJMRH3UQ'
```

Please refer to [Kerlink Wirnet provisioning documentation](https://github.com/TheThingsNetwork/kerlink-wirnet-firmware/tree/v0.0.1#provisioning) if more detailed up-to-date documentation is necessary.

> NOTE: To avoid being prompted for `root` user password several times, you may add your SSH public key as authorized for `root` user on the gateway, for example, by `ssh-copy-id root@192.168.4.155`.
