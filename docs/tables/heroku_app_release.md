# Table: heroku_app_release

A release represents a combination of code, config vars and add-ons for an app on Heroku.

Notes:
* List queries require an `app_name`.
* Get queries require an `app_name`, a release `id` or a release `version`.

Pagination is not currently supported for this resource type in the SDK.

## Examples

### List all app releases of an app

```sql
select
  id,
  status,
  version,
  is_current,
  created_at
from
  heroku_app_release
where
  app_name = 'steampipe';
```

### Get the current release version of an app

```sql
select
  id,
  status,
  version,
  is_current,
  created_at
from
  heroku_app_release
where
  app_name = 'steampipe' and is_current;
```

### Get the release information of an app by release id

```sql
select
  id,
  status,
  version,
  is_current,
  created_at
from
  heroku_app_release
where
  app_name = 'steampipe' and id = 'e8256596-5583-4df0-9a6d-cf0af5e11f02';
```

### Get the release information of an app by release version

```sql
select
  id,
  status,
  version,
  is_current,
  created_at
from
  heroku_app_release
where
  app_name = 'steampipe' and version = 4;
```
