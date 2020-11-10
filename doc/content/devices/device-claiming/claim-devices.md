---
title: "Claim a Device"
description: ""
weight: 2
distributions: ["Enterprise", "Cloud"]
--- 

You can update the example below with your tenant ID and cluster ID by filling them here.

{{< tenant-cluster-selector >}}

## Prerequisites

1. An application in The Things Industries Cloud Hosted. [See instructions]({{< ref "/integrations/adding-applications" >}})
2. A device with a QR code for claiming sticker.

## Claim a Device

1. Go to the Device Claiming App.
  The address of the Console includes both your tenant ID and the cluster ID:

  <p>
  <code data-content="cluster-address">
  https://<span data-content="tenant-id"></span>.<span data-content="cluster-id"></span>.cloud.thethings.industries/claim
  </code>
  </p>

2. Select the application that you created in prerequisites.
3. Authorize your browser to access the camera, scan the the device's QR code for claiming and click on **save changes**.

The device has been successfully claimed, you can return to the Console and you will see the device in your application.

>**Note:** Device claiming does not transfer a security session for a device, it only transfers ownership. You must enter new session keys for the device, or traffic from the device will go to the old owner.
