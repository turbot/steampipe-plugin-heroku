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
  plugin  = "heroku"

  # Account email for the user
  # email = "ryan@dundermifflin.com"

  # API key for your Heroku account, can be found in Heroku Account Settings
  # Reference: https://help.heroku.com/PBGP6IDE/how-should-i-generate-an-api-key-that-allows-me-to-use-the-heroku-platform-api
  # api_key = "34c12972-9d18-421c-3ae5-5293ae1507be"
}
```

### Example Configurations

- Connect to a single account:

  ```hcl
  connection "heroku" {
    plugin  = "heroku"
    email   = "ryan@dundermifflin.com"
    api_key = "xoxp-2556146250-EXAMPLE-1646968370949-df954218b5da5b8614c85cc454136b27"
  }
  ```

- Create connections to multiple accounts:

  ```hcl
  connection "heroku1" {
    plugin  = "heroku"
    email   = "bob@mycompany.com"
    api_key = "9bcb761c-09ud-40af-be1b-880fdaa788e6"
  }

  connection "heroku2" {
    plugin   = "heroku"
    email    = "janice@gmail.com"
    api_key  = "c2414tre-5f12-4625-8784-2d4451c34783"
  }

  connection "heroku3" {
    plugin   = "heroku"
    email    = "john@abccompany.com"
    api_key  = "c2414tte-5f12-4625-84-2d4451f34767"
  }
  ```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-heroku
- Community: [Slack Channel](https://steampipe.io/community/join)
