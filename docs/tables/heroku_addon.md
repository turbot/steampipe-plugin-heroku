# Table: heroku_addon

Add-ons are components, services, or pieces of infrastructure that are fully maintained for you, either by a third-party provider or by Heroku.

Notes:
* Get queries require an add-on `id`.

Pagination is not currently supported for this resource type in the SDK.

## Examples

### List all add-ons

```sql
select
  name,
  state,
  plan,
  web_url
from
  heroku_addon;
```

### List all provisioned add-ons

```sql
select
  name,
  state,
  plan,
  web_url
from
  heroku_addon
where
  state = 'provisioned';
```

### Add-ons that have not changed for 30 days or more

```sql
select
  name,
  web_url,
  updated_at
from
  heroku_addon
where
  updated_at < now() - interval '30 days'
```
