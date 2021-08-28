# Table: heroku_pipeline

A pipeline allows grouping of apps into different stages.

## Examples

### List all pipelines

_Note: The Heroku Go SDK does not yet support listing pipelines._

### Get a pipeline by ID

```sql
select
  *
from
  heroku_pipeline
where
  id = '3d48ef48-1360-414e-9183-ca97ae134b1a'
```
