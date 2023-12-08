---
title: "Steampipe Table: heroku_app - Query Heroku Apps using SQL"
description: "Allows users to query Heroku Apps, specifically providing information about each app's details, such as the app's unique ID, name, owner, and team. The table also provides insights into the app's stack, region, and maintenance status."
---

# Table: heroku_app - Query Heroku Apps using SQL

Heroku is a cloud platform as a service (PaaS) that lets companies build, deliver, monitor, and scale apps. Heroku is flexible for both developers and enterprises, offering add-on services, including data services, utility services, and more. The platform also supports several programming languages, including Java, Node.js, Scala, Clojure, Python, PHP, and Go.

## Table Usage Guide

The `heroku_app` table provides insights into the apps within Heroku. As a developer or system administrator, you can explore app-specific details through this table, including the app's unique ID, name, owner, and team. You can also gain insights into the app's stack, region, and maintenance status, which can be crucial for app maintenance, development, and scaling strategies.

**Important Notes**
- Get queries require an app `id`.
- Pagination is not currently supported for this resource type in the SDK.

## Examples

### List all apps
Explore the names and web URLs of all applications within your Heroku platform. This can be useful for a quick overview of all your applications, or to find the web URL of a specific application.

```sql+postgres
select
  name,
  web_url
from
  heroku_app;
```

```sql+sqlite
select
  name,
  web_url
from
  heroku_app;
```

### Apps by region
Explore which regions have the most Heroku apps deployed. This can help you understand the geographical distribution of your applications and inform decisions about where to focus resources.

```sql+postgres
select
  region ->> 'name' as region_name,
  count(*)
from
  heroku_app
group by
  region_name;
```

```sql+sqlite
select
  json_extract(region, '$.name') as region_name,
  count(*)
from
  heroku_app
group by
  region_name;
```

### Apps that have not changed for 30 days or more
Explore which applications have remained static for over a month. This can be useful for identifying potentially outdated or unused apps that may require updates or removal.

```sql+postgres
select
  name,
  web_url,
  updated_at
from
  heroku_app
where
  updated_at < now() - interval '30 days';
```

```sql+sqlite
select
  name,
  web_url,
  updated_at
from
  heroku_app
where
  updated_at < datetime('now', '-30 days');
```