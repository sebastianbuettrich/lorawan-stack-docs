---
title: "Using the API"
description: ""
weight: 
---

While we recommend using the [Console]({{< ref "getting-started/console" >}}) or [CLI]({{< ref "getting-started/cli" >}}) to manage your applications and devices in {{% tts %}}, we also expose HTTP and gRPC APIs which you can interact directly with. This section contains information about using the HTTP API, and examples.

<!--more-->

A complete list of API endpoints is available in the [API Reference]({{< ref "reference/api" >}}). There, you can also find detailed information about [Authentication]({{< ref "reference/api/authentication" >}}) and [Field Masks]({{< ref "reference/api/field-mask" >}}).

> If you're having trouble with the HTTP API, you can always inspect requests in the Console using your browser's inspector. All of the data displayed in the Console is pulled using HTTP API requests, and this should give you some insight in to how they are formed.

## HTTP Queries

Additional fields may be specified in HTTP requests by appending them as query string parameters. For example, to request the `name`, `description`, and `locations` of devices in an `EndDeviceRegistry.Get` request, add these fields to the `field_mask` field. To get this data for device `dev1` in application `app1`:

```bash
$ curl --location --header "Authorization: Bearer NNSXS.XXXXXXXXX" https://thethings.example.com/api/v3/applications/app1/devices/dev1?field_mask=name,description,locations
```

The same request in tenant `tenant1` on a multi-tenant deployment:

```bash
$ curl --location --header "Authorization: Bearer NNSXS.XXXXXXXXX" https://tenant1.thethings.example.com/api/v3/applications/app1/devices/dev1?field_mask=name,description,locations
```

Fields may also be specified as a JSON object in a POST request.

To get a stream of events for device `dev1` in application `app1` :

```bash
$ curl --location \
     --header 'Authorization: Bearer NNSXS.XXXXXXXXX' \
     --header 'Accept: application/json' \
     --header 'Content-Type: application/json' \
     --request POST \
     'https://thethings.example.com/api/v3/events' \
     --data-raw '{
    "identifiers":[{
        "device_ids":{
            "device_id":"dev1",
            "application_ids":{"application_id":"app1"}
        }
    }]
}'
```

If you want to create a device, perform multi step actions, or write shell scripts, it's best to use the [CLI]({{< ref "getting-started/cli" >}}).

If you really, really want to do something like register a device manually, you need to make calls to the `Identity Server`, `Join Server`, `Network Server` and `Application Server`. 

To register a device `newdev1` in application `app1`, first, register the `DevEUI`, `JoinEUI` and cluster addresses in the `Identity Server`:

```bash
$ curl --location \
     --header 'Authorization: Bearer NNSXS.XXXXXXXXX' \
     --header 'Content-Type: application/json' \
     --request POST \
    'https://thethings.example.com/api/v3/applications/app1/devices' \
     --data-raw '{
      "end_device": {
        "ids": {
          "device_id": "newdev1",
          "dev_eui": "0000000000000000",
          "join_eui": "0000000000000000"
        },
        "join_server_address": "thethings.example.com",
        "network_server_address": "thethings.example.com",
        "application_server_address": "thethings.example.com"
      },
      "field_mask": {
        "paths": [
          "join_server_address",
          "network_server_address",
          "application_server_address",
          "ids.dev_eui",
          "ids.join_eui"
        ]
      }
    }
'
```

Register the `DevEUI`, `JoinEUI`, and keys in the `Join Server`:

```bash
$ curl --location \
     --header 'Authorization: Bearer NNSXS.XXXXXXXXX' \
     --header 'Content-Type: application/json' \
     --request PUT \
     'https://thethings.example.com/api/v3/js/applications/app1/devices/newdev1' \
     --data-raw '{
      "end_device": {
        "ids": {
          "device_id": "newdev1",
          "dev_eui": "0000000000000000",
          "join_eui": "0000000000000000"
        },
        "network_server_address": "thethings.example.com",
        "application_server_address": "thethings.example.com",
        "network_server_kek_label": "",
        "application_server_kek_label": "",
        "application_server_id": "",
        "net_id": null,
        "root_keys": {
          "app_key": {
            "key": "XXXXXXXXX"
          }
        }
      },
      "field_mask": {
        "paths": [
          "network_server_address",
          "application_server_address",
          "ids.device_id",
          "ids.dev_eui",
          "ids.join_eui",
          "network_server_kek_label",
          "application_server_kek_label",
          "application_server_id",
          "net_id",
          "root_keys.app_key.key"
        ]
      }
    }'
```

Register MAC settings in the `Network Server`:

```bash
$ curl --location \
     --header 'Authorization: Bearer NNSXS.XXXXXXXXX' \
     --header 'Content-Type: application/json' \
     --request PUT \
     'https://thethings.example.com/api/v3/ns/applications/app1/devices/newdev1' \
     --data-raw '{
      "end_device": {
        "supports_join": true,
        "lorawan_version": "1.0.2",
        "ids": {
          "device_id": "newdev1",
          "dev_eui": "0000000000000000",
          "join_eui": "0000000000000000"
        },
        "supports_class_c": false,
        "lorawan_phy_version": "1.0.2-b",
        "frequency_plan_id": "EU_863_870"
      },
      "field_mask": {
        "paths": [
          "supports_join",
          "lorawan_version",
          "ids.device_id",
          "ids.dev_eui",
          "ids.join_eui",
          "supports_class_c",
          "lorawan_phy_version",
          "frequency_plan_id"
        ]
      }
    }'
```

Finally, register the `DevEUI` and `JoinEUI` in the `Application Server`:

```bash
$ curl --location \
     --header 'Authorization: Bearer NNSXS.XXXXXXXXX' \
     --header 'Content-Type: application/json' \
     --request PUT \
     'https://thethings.example.com/api/v3/as/applications/app1/devices/newdev1' \
     --data-raw '{
      "end_device": {
        "ids": {
          "device_id": "newdev1",
          "dev_eui": "0000000000000000",
          "join_eui": "0000000000000000"
        }
      },
      "field_mask": {
        "paths": [
          "ids.device_id",
          "ids.dev_eui",
          "ids.join_eui"
        ]
      }
    }'
```

