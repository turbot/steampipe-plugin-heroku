# Table: heroku_dyno

Dynos encapsulate running processes of an app on Heroku.

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
