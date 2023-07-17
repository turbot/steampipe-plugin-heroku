## v0.5.0 [2023-07-17]

_Enhancements_

- Updated the `docs/index.md` file to include multi-account configuration examples. ([#22](https://github.com/turbot/steampipe-plugin-heroku/pull/22))

## v0.4.0 [2023-04-07]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#20](https://github.com/turbot/steampipe-plugin-heroku/pull/20))

## v0.3.0 [2022-09-27]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.7](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v417-2022-09-08) which includes several caching and memory management improvements. ([#17](https://github.com/turbot/steampipe-plugin-heroku/pull/17))
- Recompiled plugin with Go version `1.19`. ([#17](https://github.com/turbot/steampipe-plugin-heroku/pull/17))

## v0.2.1 [2022-05-23]

_Bug fixes_

- Fixed the Slack community links in README and docs/index.md files. ([#13](https://github.com/turbot/steampipe-plugin-heroku/pull/13))

## v0.2.0 [2022-04-27]

_Enhancements_

- Added support for native Linux ARM and Mac M1 builds. ([#11](https://github.com/turbot/steampipe-plugin-heroku/pull/11))
- Recompiled plugin with [steampipe-plugin-sdk v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30) and Go version `1.18`. ([#10](https://github.com/turbot/steampipe-plugin-heroku/pull/10))

## v0.1.0 [2021-12-16]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v1.8.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v182--2021-11-22) ([#7](https://github.com/turbot/steampipe-plugin-heroku/pull/7))
- Recompiled plugin with Go version 1.17 ([#7](https://github.com/turbot/steampipe-plugin-heroku/pull/7))

## v0.0.1 [2021-09-23]

_What's new?_

- New tables added
  - [heroku_account](https://hub.steampipe.io/plugins/turbot/heroku/tables/heroku_account)
  - [heroku_addon](https://hub.steampipe.io/plugins/turbot/heroku/tables/heroku_addon)
  - [heroku_app](https://hub.steampipe.io/plugins/turbot/heroku/tables/heroku_app)
  - [heroku_app_release](https://hub.steampipe.io/plugins/turbot/heroku/tables/heroku_app_release)
  - [heroku_app_webhook](https://hub.steampipe.io/plugins/turbot/heroku/tables/heroku_app_webhook)
  - [heroku_domain](https://hub.steampipe.io/plugins/turbot/heroku/tables/heroku_domain)
  - [heroku_dyno](https://hub.steampipe.io/plugins/turbot/heroku/tables/heroku_dyno)
  - [heroku_key](https://hub.steampipe.io/plugins/turbot/heroku/tables/heroku_key)
  - [heroku_pipeline](https://hub.steampipe.io/plugins/turbot/heroku/tables/heroku_pipeline)
  - [heroku_region](https://hub.steampipe.io/plugins/turbot/heroku/tables/heroku_region)
  - [heroku_team](https://hub.steampipe.io/plugins/turbot/heroku/tables/heroku_team)
  - [heroku_team_member](https://hub.steam_memberpipe.io/plugins/turbot/heroku/tables/heroku_team_member)
