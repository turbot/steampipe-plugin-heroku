# Table: heroku_region

A region represents a geographic location in which your application may run.

## Examples

### List all regions

```sql
select
  *
from
  heroku_region
order by
  name
```

### List all US regions

```sql
select
  *
from
  heroku_region
where
  country = 'United States'
order by
  name
```
