# Table: heroku_app_release

A release represents a combination of code, config vars and add-ons for an app on Heroku.

## Examples

### List all app releases of an app

```sql
select
  id,
  status,
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
  is_current,
  created_at
from
  heroku_app_release
where
  app_name = 'steampipe' and is_current;
```
