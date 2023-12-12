---
title: "Steampipe Table: heroku_addon - Query Heroku Add-ons using SQL"
description: "Allows users to query Heroku Add-ons, providing insights into add-on details such as the associated app, plan, and configuration."
---

# Table: heroku_addon - Query Heroku Add-ons using SQL

Heroku Add-ons are services that support, enhance, and extend the functionality of Heroku applications. They offer a wide range of services from data storage to email services, and from log analytics to performance monitoring. Each add-on is associated with an app and can be managed through the Heroku Dashboard or CLI.

## Table Usage Guide

The `heroku_addon` table provides insights into add-ons within Heroku. As a developer or system administrator, explore add-on-specific details through this table, such as the associated app, plan, and configuration. Utilize it to manage and optimize the add-ons for your Heroku applications.

**Important Notes**
- Get queries require an add-on `id`.
- Pagination is not currently supported for this resource type in the SDK.

## Examples

### List all add-ons
Explore all the additional features currently available in your system, understanding their status and the plans they are associated with, along with their web URLs. This can help you manage your resources better and plan for future enhancements.

```sql+postgres
select
  name,
  state,
  plan,
  web_url
from
  heroku_addon;
```

```sql+sqlite
select
  name,
  state,
  plan,
  web_url
from
  heroku_addon;
```

### List all provisioned add-ons
Explore which add-ons have been provisioned for your Heroku applications. This can help you manage resources effectively and understand the current state of your application's functionality.

```sql+postgres
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

```sql+sqlite
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

### List add-ons that have not changed for 30 days or more
Explore add-ons that have remained unmodified for a month or more. This is useful to identify dormant resources that might be candidates for review or removal to optimize resource usage.

```sql+postgres
select
  name,
  web_url,
  updated_at
from
  heroku_addon
where
  updated_at < now() - interval '30 days';
```

```sql+sqlite
select
  name,
  web_url,
  updated_at
from
  heroku_addon
where
  updated_at < datetime('now', '-30 days');
```