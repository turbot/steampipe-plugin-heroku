# Table: heroku_dyno

Dynos encapsulate running processes of an app on Heroku.

Notes:
* List queries require an `app_name`.
* Get queries require an `app_name` and a dyno `id`.

Pagination is not currently supported for this table.

## Examples

### List all dynos of an app

```sql
select
  id,
  name,
  type,
  size,
  state
from
  heroku_dyno
where
  app_name = 'steampipe';
```

### List all crashed dynos of an app

```sql
select
  id,
  name,
  type,
  size,
  state
from
  heroku_dyno
where
  app_name = 'steampipe' and state = 'crashed';
```

### List all hobby size dynos of an app

```sql
select
  id,
  name,
  type,
  size,
  state
from
  heroku_dyno
where
  app_name = 'steampipe' and size = 'Hobby';
```
