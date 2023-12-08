---
title: "Steampipe Table: heroku_domain - Query Heroku Domains using SQL"
description: "Allows users to query Domains in Heroku, specifically the information about the domain names used by the app, providing insights into the app's configuration and operation."
---

# Table: heroku_domain - Query Heroku Domains using SQL

A Heroku Domain is the custom domain that can be added to a Heroku app, which allows the app to be accessed via a custom URL. This domain is associated with a DNS record that points at the Heroku-provided DNS target. It provides a way to give your application a custom URL, instead of using the default Herokuapp.com address.

## Table Usage Guide

The `heroku_domain` table provides insights into the custom domains used within Heroku applications. As a DevOps engineer, explore domain-specific details through this table, including the domain name, created and updated timestamps, and the associated Heroku app. Utilize it to manage and monitor the domain configurations of your Heroku applications, ensuring they are correctly set up and functioning as expected.

**Important Notes**
- List queries require an `app_name`.
- Get queries require an `app_name` and a domain `id`.
- Pagination is not currently supported for this resource type in the SDK.

## Examples

### List all domains
Explore all domains associated with a specific application to understand their status and type. This is helpful in managing and monitoring the application's network configuration.

```sql+postgres
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

```sql+sqlite
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

### List all custom domains
Discover the segments that include all custom domains associated with a specific application. This can be particularly useful in managing and monitoring your application's custom domains.

```sql+postgres
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

```sql+sqlite
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