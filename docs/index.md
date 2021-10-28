---
organization: Turbot
category: ["public cloud"]
icon_url: "/images/plugins/turbot/heroku.svg"
brand_color: "#430098"
display_name: "Heroku"
short_name: "heroku"
description: "Steampipe plugin to query apps, dynos, and more from Heroku."
og_description: "Query Heroku with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/heroku-social-graphic.png"
---

# Heroku + Steampipe

[Heroku](https://heroku.com) is a cloud-based Platform as a Service (PaaS) supporting application development across several programming languages.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example:

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

- **[Table definitions & examples →](https://hub.steampipe.io/plugins/turbot/heroku/tables)**

## Get started

### Install

Download and install the latest Heroku plugin:

```bash
steampipe plugin install heroku
```

### Configuration

Installing the latest Heroku plugin will create a config file (`~/.steampipe/config/heroku.spc`) with a single connection named `heroku`:

```hcl
connection "heroku" {
  plugin     = "heroku"
  # email    = "email address regsitered on Heroku"
  # api_key  = "YOUR_API_KEY"

  # API key for your Heroku account
  # Reference: https://dashboard.heroku.com/account
}
```

### Example Configuration

Connect to an account:

```hcl
connection "heroku" {
  plugin    = "heroku"
  email     = "ryan@dundermifflin.com"
  api_key   = "34c12972-9d18-421c-3ae5-5293ae1507be"
}
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-heroku
- Community: [Slack Channel](https://steampipe.io/community/join)
