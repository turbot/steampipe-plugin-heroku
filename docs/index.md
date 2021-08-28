---
organization: Turbot
category: ["public cloud"]
icon_url: "/images/plugins/turbot/heroku.svg"
brand_color: "#430098"
display_name: "Heroku"
short_name: "heroku"
description: "Steampipe plugin to query apps, dynos and more from Heroku."
og_description: "Query Heroku with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/heroku-social-graphic.png"
---

# Heroku + Steampipe

[Heroku](https://heroku.com) is a cloud platform as a service (PaaS) supporting several programming languages.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

List instances in your Heroku account:

```sql
select
  name,
  web_url,
  updated_at
from
  heroku_app
```

```
+--------+-------------------------------+---------------------+
| name   | web_url                       | updated_at          |
+--------+-------------------------------+---------------------+
| my-app | https://my-app.herokuapp.com/ | 2021-08-28 18:44:51 |
+--------+-------------------------------+---------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/heroku/tables)**

## Get started

### Install

Download and install the latest Heroku plugin:

```bash
steampipe plugin install heroku
```

### Credentials

No credentials are required.

### Configuration

Installing the latest heroku plugin will create a config file (`~/.steampipe/config/heroku.spc`) with a single connection named `heroku`:

```hcl
connection "heroku" {
  plugin  = "heroku"
  email   = "ryan@dundermifflin.com"
  api_key = "34c12972-9d18-421c-3ae5-5293ae1507be"
}
```

- `email` - Email address of the Heroku user.
- `api_key` - API key (or password) of the Heroku user.

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-heroku
- Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g)
