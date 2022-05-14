<!--
SPDX-FileCopyrightText: 2021 The NGI Pointer Secure-Scuttlebutt Team of 2020/2021

SPDX-License-Identifier: CC0-1.0
-->

:warning: **This repo was moved to https://github.com/ssbc/go-ssb-room.** This archival will remain in this GitHub org `ssb-ngi-pointer` to demonstrate the outcome of the work done by the SSB NGI Pointer team during 2020 and 2021. The SSB NGI Pointer team is no longer active because we completed our grant project.

# Go-SSB Room
[![REUSE status](https://api.reuse.software/badge/github.com/ssb-ngi-pointer/go-ssb-room)](https://api.reuse.software/info/github.com/ssb-ngi-pointer/go-ssb-room)

This repository contains code for a [Secure Scuttlebutt](https://ssb.nz) [Room (v1+v2) server](https://github.com/ssb-ngi-pointer/rooms2), written in Go.

It includes:
* secret-handshake+boxstream network transport, sometimes referred to as SHS, using [secretstream](https://github.com/cryptoscope/secretstream)
* muxrpc handlers for tunneling connections
* a fully embedded HTTP server & HTML frontend, for administering the room

![](./docs/images/screenshot.png)

## :star: Features

* Rooms v1 (`tunnel.connect`, `tunnel.endpoints`, etc.)
* User management (allow- & denylisting + moderator & administrator roles), all administered via the web dashboard
* Multiple [privacy modes](https://ssb-ngi-pointer.github.io/rooms2/#privacy-modes)
* [Sign-in with SSB](https://ssb-ngi-pointer.github.io/ssb-http-auth-spec/)
* [HTTP Invites](https://github.com/ssb-ngi-pointer/ssb-http-invite-spec)
* Alias management

For a comprehensive introduction to rooms 2.0, 🎥 [watch this video](https://www.youtube.com/watch?v=W5p0y_MWwDE).

## :rocket: Deployment

If you want to deploy a room server yourself, follow our [deployment.md](./docs/deployment.md) docs.

## :wrench: Development

For an in-depth codebase walkthrough, see the [development.md](./docs/development.md) file in the `docs` folder of this repository.

## :people_holding_hands: Authors

* [cryptix](https://github.com/cryptix) (`@p13zSAiOpguI9nsawkGijsnMfWmFd5rlUNpzekEE+vI=.ed25519`)
* [staltz](https://github.com/staltz)
* [cblgh](https://github.com/cblgh)

## License

MIT

