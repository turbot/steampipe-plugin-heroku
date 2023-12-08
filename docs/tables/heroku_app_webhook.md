---
title: "Steampipe Table: heroku_app_webhook - Query Heroku App Webhooks using SQL"
description: "Allows users to query Heroku App Webhooks, specifically the current webhook configuration details for Heroku applications, providing insights into webhook settings and potential anomalies."
---

# Table: heroku_app_webhook - Query Heroku App Webhooks using SQL

Heroku App Webhooks is a feature within Heroku that allows you to set up HTTP callbacks to notify you about events happening in your applications. It provides a way to configure and manage webhooks for various Heroku resources, including applications, dynos, and add-ons. Webhooks in Heroku help you stay informed about the operations and performance of your Heroku resources, triggering HTTP callbacks when predefined conditions are met.

## Table Usage Guide

The `heroku_app_webhook` table provides insights into webhook configurations within Heroku applications. As a DevOps engineer, explore webhook-specific details through this table, including webhook URLs, event types, and associated metadata. Utilize it to uncover information about webhooks, such as their current status, the events they are configured to respond to, and the verification of their settings.

**Important Notes**
- List queries require an `app_name`.
- Get queries require an `app_name` and webhook `id`.
- Pagination is not currently supported for this resource type in the SDK.

## Examples

### List all webhooks
Discover the segments that have been created within a specific application by analyzing the settings of each webhook. This can help you understand the configuration for each webhook, allowing you to assess the elements within your application and make necessary changes.

```sql+postgres
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

```sql+sqlite
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

### List all notify level webhooks
Explore which webhooks have been set to the 'notify' level within a specific application. This can help in understanding the extent of notifications you're receiving from that application.

```sql+postgres
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

```sql+sqlite
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