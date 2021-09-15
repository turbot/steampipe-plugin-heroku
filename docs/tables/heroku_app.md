# Table: heroku_app

An app represents the program that you would like to deploy and run on Heroku.

Notes:
* Get queries require an app `id`.

Pagination is not currently supported for this table.

## Examples

### List all apps

```sql
select
  name,
  web_url
from
  heroku_app
```

### Apps by region

```sql
select
  region ->> 'name' as region_name,
  count(*)
from
  heroku_app
group by
  region_name
```

### Apps that have not changed for 30 days or more

```sql
select
  name,
  web_url,
  updated_at
from
  heroku_app
where
  updated_at < now() - interval '30 days'
```
