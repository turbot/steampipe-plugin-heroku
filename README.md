![image](https://hub.steampipe.io/images/plugins/turbot/heroku-social-graphic.png)

# Heroku Plugin for Steampipe

Use SQL to query apps, dynos and more from Heroku.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/heroku)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/heroku/tables)
- Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-heroku/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install heroku
```

Run a query:

```sql
select
  name,
  web_url,
  updated_at
from
  heroku_app
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-heroku.git
cd steampipe-plugin-heroku
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/heroku.spc
```

Try it!

```
steampipe query
> .inspect heroku
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-heroku/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Heroku Plugin](https://github.com/turbot/steampipe-plugin-heroku/labels/help%20wanted)
