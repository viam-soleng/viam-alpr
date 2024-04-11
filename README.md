> [!NOTE]
> This is a work in progress repo! The module in the Viam Registry is not yet 100% in line with this code base!


# Viam-ALPR modular resource

This module implements a Viam vision service to automatically recognise license plates "alpr". It is based upon the [OpenALPR](https://github.com/openalpr) project which is a battle tested and pretty much state of the art alpr capability.

## Requirements

The module and its dependencies are packed into an AppImage. Therefore FUSE is required. FUSE however is used for the installation of Viam server as well and should therefore be available already.

## Build and run

To use this module, follow the instructions to [add a module from the Viam Registry](https://docs.viam.com/registry/configure/#add-a-modular-resource-from-the-viam-registry) and select the [`viamalpr` module](https://app.viam.com/module/viam-soleng/viamalpr).

## Configure the viamalpr vision service

> [!NOTE]
> Before configuring your viamalpr module you must [create a machine](https://docs.viam.com/manage/fleet/machines/#add-a-new-machine).

Navigate to the **Config** tab of your machine's page in [the Viam app](https://app.viam.com/).
Click on the **Services** subtab and click **Create service**.
Search for `viamalpr".
Click **Add module**, then enter a name for your viamalpr service and click **Create**.

The viamalpr module doesn't require any additional configuration.

> [!NOTE]
> For more information, see [Configure a Machine](https://docs.viam.com/manage/configuration/).

### Attributes

The following attributes are available for `<INSERT API NAMESPACE>:<INSERT API NAME>:<INSERT MODEL>` camera's:

| Name    | Type   | Inclusion    | Description |
| ------- | ------ | ------------ | ----------- |
| `runtime_dir` | string | Optional | OpenALPR runtime_data folder e.g. /home/ubuntu/openalpr/runtime_data|
| `config_file` | string | Optional     | OpenALPR config file e.g. /etc/openalpr/config/openalpr.conf        |

### Full Smart Machine Example configuration

[Viam Server Configuration](config-run.json)

### Next steps

_Add any additional information you want readers to know and direct them towards what to do next with this module._
_For example:_

- To test your...
- To write code against your...

## Troubleshooting

_Add troubleshooting notes here._
