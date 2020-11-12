---
title: "Migrating from The Things Network Stack V2"
description: ""
weight: 3
aliases: ["/getting-started/migrating-from-v2", "/getting-started/migrating-from-v2/configure-ttnctl", "/getting-started/migrating-from-v2/export-v2-devices"]
---

This section documents the process of migrating end devices from {{% ttnv2 %}} to {{% tts %}}.

<!--more-->

For a breakdown of differences between {{% ttnv2 %}} and {{% tts %}}, see the [Major Changes]({{< relref "major-changes" >}}) section.

## Suggested Migration Process

**First**: Update applications to support the {{% tts %}} data format. If you are using payload formatters, make sure to set them correctly from the Application settings page.

**Second**: Follow this guide in order to migrate a single end device (and gateway, if needed) to {{% tts %}}. Continue by gradually migrating your end devices in small batches. Avoid migrating production workloads before you are certain that they will work as expected.

**Finally**: Once you are confident that your end devices are working properly, migrate the rest of your devices and gateways to {{% tts %}}.

## Configure ttn-lw-migrate

End devices and applications can easily be migrated from {{% ttnv2 %}} to {{% tts %}} with the [`ttn-lw-migrate`](https://github.com/TheThingsNetwork/lorawan-stack-migrate) tool. This tool is used for exporting end devices and applications to a [JSON file]({{< ref "getting-started/migrating/device-json.md" >}}) containing their description. This file can later be imported in {{% tts %}} as described in the [Import End Devices in The Things Stack]({{< ref "getting-started/migrating/import-devices.md" >}}) section.

First, configure the environment with the following variables modified according to your setup:

```bash
$ export TTNV2_APP_ID="my-ttn-app"                    # TTN App ID
$ export TTNV2_APP_ACCESS_KEY="ttn-account-v2.a..."   # TTN App Access Key (needs `devices` permissions)
$ export FREQUENCY_PLAN_ID="EU_863_870_TTN"           # Frequency Plan ID for exported devices
```

See [Frequency Plans]({{< ref src="/reference/frequency-plans" >}}) for the list of frequency plans available on {{% tts %}}. Make sure to specify the correct Frequency Plan ID. For example, the ID `EU_863_870_TTN` corresponds to the **Europe 863-870 MHz (SF9 for RX2 - recommended)** frequency plan.

Private The Things Network Stack V2 deployments are also supported, and require extra configuration. See `ttn-lw-migrate device --help` for more details. For example, to override the discovery server address:

```bash
$ export TTNV2_DISCOVERY_SERVER_ADDRESS="discovery.thethings.network:1900"
```

## Export End Devices from {{% ttnv2 %}}

In order to export a single device, use the following command. The device with ID `mydevice` will exported and saved to `device.json`.

```bash
$ ttn-lw-migrate devices --source ttnv2 "mydevice" > devices.json
```

>**Notes:**:
>- Payload formatters are not exported. See [Payload Formatters](https://thethingsstack.io/integrations/payload-formatters/).
>- Active device sessions are exported by default. You can disable this by using the `--ttnv2.with-session=false` flag. It is recommended that you do not export session keys for devices that can instead re-join on The Things Stack.

In order to export a large number of devices, create a file named `device_ids.txt` with one device ID per line:

```
mydevice
otherdevice
device3
device4
device5
```

And then export with:

```bash
$ ttn-lw-migrate devices --source ttnv2 < device_ids.txt > devices.json
```

Alternatively, you can export all the end devices associated with your application, and save them in `all-devices.json`.

```bash
$ ttn-lw-migrate application --source ttnv2 "my-ttn-app" > all-devices.json
```

>**Note:** Keep in mind that an end device can only be registered in one Network Server at a time. After importing an end device to {{% tts %}}, you should remove it from {{% ttnv2 %}}. For OTAA devices, it is enough to simply change the AppKey, so the device can no longer join but the existing session is preserved. Next time the device joins, the activation will be handled by {{% tts %}}.

### Disable Exported End Devices on V2

You will need to use the latest version of `ttnctl`, the CLI for {{% ttnv2 %}}. Follow the [instructions from The Things Network documentation][1]. An overview is given below:

Download `ttnctl` [for your operating system][2].

Update to the latest version:

```bash
$ ttnctl selfupdate
```

Go to [https://account.thethingsnetwork.org][3] and click [ttnctl access code][4].

Use the returned code to login from the CLI with:

```bash
$ ttnctl user login "t9XPTwJl6shYSJSJxQ1QdATbs4u32D4Ib813-fO9Xlk"
```

To get started, select the **AppID** and **AppEUI** of the application you want to export your end devices from:

```bash
$ ttnctl applications select
```

To clear the AppKey of an OTAA device, use the following command:

```bash
$ ttnctl devices convert-to-abp "device-id" --save-to-attribute "original-app-key"
```

There is also a convenience command to clear all of your devices at once:

```bash
$ ttnctl devices convert-all-to-abp --save-to-attribute "original-app-key"
```

>**Note:** The AppKey of each device will be printed on the standard output, and stored as a device attribute (with name `original-app-key`). You can retrieve the device attributes with `ttnctl devices info "device-id"`.
