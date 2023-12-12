---
title: "Steampipe Table: heroku_pipeline - Query Heroku Pipelines using SQL"
description: "Allows users to query Heroku Pipelines, providing insights into the continuous delivery of applications on the Heroku platform."
---

# Table: heroku_pipeline - Query Heroku Pipelines using SQL

Heroku Pipelines is a sequence of stages (review, development, staging, production) representing the lifecycle of an application. It provides a visual representation of the process that an app goes through from source code to deployment. Heroku Pipelines helps teams to manage the progress of features and fixes from review apps to staging and production.

## Table Usage Guide

The `heroku_pipeline` table provides insights into the lifecycle of applications on the Heroku platform. As a DevOps engineer, explore pipeline-specific details through this table, including the stages of application development and deployment. Utilize it to uncover information about pipelines, such as the status of application development stages, and the verification of deployment processes.

## Examples

### List all pipelines

_Note: The Heroku Go SDK does not yet support listing pipelines._

### Get a pipeline by ID
Explore which pipeline corresponds to a specific ID to manage or troubleshoot your Heroku deployment. This is useful in scenarios where you need to quickly identify and access a specific pipeline based on its unique identifier.

```sql+postgres
select
  *
from
  heroku_pipeline
where
  id = '3d48ef48-1360-414e-9183-ca97ae134b1a';
```

```sql+sqlite
select
  *
from
  heroku_pipeline
where
  id = '3d48ef48-1360-414e-9183-ca97ae134b1a';
```