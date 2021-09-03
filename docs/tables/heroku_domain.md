# Table: heroku_domain

Domains define what web routes should be routed to an app on Heroku.

## Examples

### List all domains of an app

```sql
select
  id,
  status,
  kind,
  hostname
from
  heroku_domain
where
  app_name = 'steampipe';
```

### List all custom domains of an app

```sql
select
  id,
  status,
  kind,
  hostname
from
  heroku_domain
where
  app_name = 'steampipe' and kind = 'custom';
```