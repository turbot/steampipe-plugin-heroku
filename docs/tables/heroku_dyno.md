---
title: "Steampipe Table: heroku_dyno - Query Heroku Dynos using SQL"
description: "Allows users to query Heroku Dynos, providing insights into the current status and configuration of each dyno within a Heroku application."
---

# Table: heroku_dyno - Query Heroku Dynos using SQL

A Heroku Dyno is a lightweight, isolated container that runs the command specified in the Procfile of your app. Each dyno belongs to a specific application and is scoped to the environment of that application. Dynos can be used to run virtually any language or service and can be scaled individually based on the needs of the application.

## Table Usage Guide

The `heroku_dyno` table provides insights into the dynos running within a Heroku application. As a developer or DevOps engineer, you can use this table to explore dyno-specific details, such as current status, type, and configuration. This information can be critical for troubleshooting application issues, optimizing resource usage, and understanding the overall health and performance of your Heroku applications.

**Important Notes**
- List queries require an `app_name`.
- Get queries require an `app_name` and a dyno `id`.
- Pagination is not currently supported for this resource type in the SDK.

## Examples

### List all dynos
Explore which dynos are currently in use within a specific Heroku application. This is useful for monitoring your app's resource usage and performance.

```sql+postgres
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

```sql+sqlite
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

### List all crashed dynos
Uncover the details of application failures by identifying instances where an application's dynamic components, or 'dynos', have crashed. This can assist in troubleshooting and improving the stability of the application.

```sql+postgres
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

```sql+sqlite
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

### List all hobby size dynos
Explore which dynos within a specific Heroku application are designated as 'Hobby' size. This can help you understand the resource allocation and usage within your application.

```sql+postgres
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

```sql+sqlite
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