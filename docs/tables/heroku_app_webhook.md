# Table: heroku_app_webhook

Webhooks define what web routes should be routed to an app on Heroku.

## Examples

### List all webhooks of an app

```sql
select
  id,
  url,
  level,
  created_at
from
  heroku_app_webhook
where
  app_name = 'steampipe';
```

### List all notify level webhooks of an app

```sql
select
  id,
  url,
  level,
  created_at
from
  heroku_app_webhook
where
  app_name = 'steampipe' and level = 'notify';
```
