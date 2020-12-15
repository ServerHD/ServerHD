# Server Home Dashboard

> ⚠️ WIP: Some support will be provided via the GitHub issues whilst in development

ServerHD is a auto-population dashboard system for your self-hosted applications. Data is pulled from a variety of popular services to automatically add and remove content from your self-hosted app dashboard. ServerHD can automate the addition and removal of: applications, notifications, and bookmarks.

This repository is just ServerHD's API, the core service that collects the data and provides 3 simple models for frontend dashboards to use. As an example, a [fork of SUI](https://github.com/serverhd/frontend-sui) has been made to use data collected by ServerHD rather that just the apps defined in a json file. The ServerHD API and model definition can be found [here](TODO).

## Getting Started

I suggest ServerHD is run as Docker container, as it was designed to run in my Docker Compose based home-server set-up that cdan be found [here](https://github.com/willfantom/composing).

  - Using ServerHD with Docker: [guide here]()
  - Using ServerHD on Linux: [guide here]().
  - Using Windows on Linux: [guide here]().
  - Using MacOS on Linux: [guide here]().

---

## Providers

There sre 3 types of data collected, formatted and provided by ServerHD. These are all collected using '*providers*'. There are 3 types of provider, one for each type of data. Each provider can have multiple variations.


### Application Providers

- Static File
- Docker Unix Socket
- Træfik v2 API

See the [wiki page](https://github.com/ServerHD/ServerHD/wiki/Application-Providers) for more details.

### Notification Providers

- Gotify
- Covid19

See the [wiki page](https://github.com/ServerHD/ServerHD/wiki/Notification-Providers) for more details.

### Bookmark Providers

- Linkding

See the [wiki page](https://github.com/ServerHD/ServerHD/wiki/Bookmark-Providers) for more details.
