# Design Pilot CLI

This is the codebase for the Design Pilot CLI, which can be used to interact with certain Design Pilot functions. You need to have purchased a Design Pilot Premium Subscription.

## Overview

### Generate an image

```shell
designpilot generate --prompt="A pig in a teacup"
```

![pig in a teacup](https://github.com/designpilotai/designpilot/assets/80828305/aa084ead-0420-49be-831c-5e2e2f38a9c3)

### Generate a logo

```shell
designpilot generate -p="A pig in a teacup" --logo
```

![pig in a teacup logo](https://github.com/designpilotai/designpilot/assets/80828305/379b86c5-3397-4edc-8f56-9f3b02edfef8)

## Setup

### API Key

1. Purchase a Design Pilot Premium Subscription.
1. Generate an API key on [the devtools page](https://designpilot.ai/devtools). If you lose it you can always delete it and generate a new one.
1. Store your API key locally as an [environment variable](https://chlee.co/how-to-setup-environment-variables-for-windows-mac-and-linux/): `DESIGNPILOT_API_KEY`

### Installation

If you already have Go installed, [Option 2](#option-2-install-using-the-go-toolchain) will probably be quickest

#### Option 1: Download a prebuilt binary

If there is no release for your OS, you may have to use [Option 2](#option-2-install-using-the-go-toolchain).

1. Go to our [releases](https://github.com/designpilotai/designpilot/releases) and download the version for your OS.
1. Rename the file to designpilot
1. Put it in your custom scripts folder (if you don't have one check [this](https://zwbetz.com/how-to-add-a-binary-to-your-path-on-macos-linux-windows/))

#### Option 2: Install using the Go toolchain

1. [Install the Go toolchain](https://go.dev/doc/install)
1. Install designpilot

```shell
go install github.com/designpilotai/designpilot@latest
```
