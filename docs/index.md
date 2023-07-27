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

## Multi-Account Connections

You may create multiple heroku connections:

```hcl
connection "heroku_dev" {
  plugin  = "heroku"
  email   = "ryan@dundermifflin.com"
  api_key = "34c12972-9d18-654c-3ae5-5293ae1889be"
}

connection "heroku_qa" {
  plugin  = "heroku"
  email   = "ryan@dundermifflin.com"
  api_key = "34c12972-9d18-654c-3ae5-5293ae1905be"
}

connection "heroku_prod" {
  plugin  = "heroku"
  email   = "ryan@dundermifflin.com"
  api_key = "34c12972-9d18-654c-3ae5-5293ae1124be"
}
```

Each connection is implemented as a distinct [Postgres schema](https://www.postgresql.org/docs/current/ddl-schemas.html). As such, you can use qualified table names to query a specific connection:

```sql
select * from heroku_qa.heroku_team_member
```

You can create multi-account connections by using an [**aggregator** connection](https://steampipe.io/docs/using-steampipe/managing-connections#using-aggregators). Aggregators allow you to query data from multiple connections for a plugin as if they are a single connection.

```hcl
connection "heroku_all" {
  plugin      = "heroku"
  type        = "aggregator"
  connections = ["heroku_dev", "heroku_qa", "heroku_prod"]
}
```

Querying tables from this connection will return results from the `heroku_dev`, `heroku_qa`, and `heroku_prod` connections:

```sql
select * from heroku_all.heroku_team_member
```

Alternatively, you can use an unqualified name and it will be resolved according to the [Search Path](https://steampipe.io/docs/guides/search-path). It's a good idea to name your aggregator first alphabetically so that it is the first connection in the search path (i.e. `heroku_all` comes before `heroku_dev`):

```sql
select * from heroku_team_member
```

Steampipe supports the `*` wildcard in the connection names. For example, to aggregate all the heroku plugin connections whose names begin with `heroku_`:

```hcl
connection "heroku_all" {
  type        = "aggregator"
  plugin      = "heroku"
  connections = ["heroku_*"]
}
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-heroku
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
