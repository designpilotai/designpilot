# Design Pilot CLI

This is the codebase for the Design Pilot CLI, which can be used to interact with certain Design Pilot functions. You need to have purchased a Design Pilot Premium Subscription.

## Overview

### Generate an image

```shell
designpilot generate --prompt="A pig in a teacup"
```

![pig in a teacup](http://localhost:3000/content/asset/30c996f8-c162-4350-b130-763c45836552/fullres)

### Generate a logo

```shell
designpilot generate -p="A pig in a teacup" --logo
```

![pig in a teacup logo](http://localhost:3000/content/logo/a9fceb90-ef2f-47e7-a4ac-6cca256f3092/fullres)

## Setup

### API Key

1. Purchase a Design Pilot Premium Subscription.
1. Generate an API key on [the devtools page](https://designpilot.ai/devtools). If you lose it you can always delete it and generate a new one.
1. Store your API key locally as an environment variable: `DESIGNPILOT_API_KEY`

### Installation

#### Option 1: Download a prebuilt binary

1. Go to our [releases](https://github.com/designpilotai/designpilot/releases) and download the version for your OS.

If there is no release for your OS, you may have to use option 2:

#### Option 2: Install using the Go toolchain

1. [Install the Go toolchain](https://go.dev/doc/install)
1. Install designpilot

```shell
go install github.com/designpilotai/designpilot
```
