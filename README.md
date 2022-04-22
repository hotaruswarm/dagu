# jobctl

**A dead simple tool to run & manage DAGs**

jobctl is a single command that generates and executes a [DAG (Directed acyclic graph)](https://en.wikipedia.org/wiki/Directed_acyclic_graph) from a simple YAML definition. jobctl also comes with a convenient web UI. It aims to be one of the easiest option to manage DAGs executed by cron.

## Contents

- [jobctl](#jobctl)
  - [Contents](#contents)
  - [Features](#features)
  - [Usecases](#usecases)
  - [Getting started](#getting-started)
    - [Installation](#installation)
    - [Usage](#usage)
  - [Configuration](#configuration)
    - [Environment variables](#environment-variables)
    - [Web UI configuration](#web-ui-configuration)
    - [Global configuration](#global-configuration)
  - [Job configuration](#job-configuration)
    - [Simple example](#simple-example)
    - [Complex example](#complex-example)


## Features

- Simple command interface (See [Usage](#usage))
- Simple configuration YAML format (See [Simple example](#simple-example))
- Simple architecture (no DBMS or agent process is required)
- Web UI to visualize, manage jobs and watch logs
- Parameterization
- Conditions
- Automatic retry
- Cancellation
- Retry
- Prallelism limits
- Environment variables
- Repeat jobs
- Basic Authentication
- E-mail notifications
- REST api interface

## Usecases
- ETL Pipeline
- Batches
- Machine Learning
- Data Processing
- Automation

## Getting started
### Installation

Place a `jobctl` executable somewhere on your system.

### Usage

- `jobctl start [--params=<params>] <job file>` - run a job
- `jobctl status <job file>` - display the current status of the job
- `jobctl retry --req=<request-id> <job file>` - retry the failed/canceled job
- `jobctl stop <job file>` - cancel a job
- `jobctl dry [--params=<params>] <job file>` - dry-run a job
- `jobctl server` - start a web server for web UI

## Configuration

### Environment variables
- `JOBCTL__DATA` - path to directory for internal use by jobctl (default : `~/.jobctl/data`)
- `JOBCTL__LOGS` - path to directory for logging (default : `~/.jobctl/logs`)

### Web UI configuration

Plase create `~/.jobctl/admin.yaml`.

```yaml
# required
host: <hostname for web UI address>
port: <port number for web UI address>
jobs: <the location of job configuration files>
command: <absolute path to the jobctl exectable>

# optional
isBasicAuth: <true|false>
basicAuthUsername: <username for basic auth of web UI>
basicAuthPassword: <password for basic auth of web UI>
```

### Global configuration

Plase create `~/.jobctl/config.yaml`. All settings can be overridden by individual job configurations.

```yaml
logDir: <path-to-write-log>   # log directory to write standard output from the job steps
histRetentionDays: 3 # job history retention days (not for log files)

# E-mail server config (optional)
smtp:
  host: <smtp server host>
  port: <stmp server port>
errorMail:
  from: <from address to send error mail>
  to: <to address to send error mail>
  prefix: <prefix of mail subject for error mail>
infoMail:
  from: <from address to send notification mail>
  to: <to address to send notification mail>
  prefix: <prefix of mail subject for notification mail>
```

## Job configuration

### Simple example

A simple example is as follows:
```yaml
name: simple job
steps:
  - name: step 1
    command: python some_batch_1.py
    dir: ${HOME}/jobs/                  # working directory for the job (optional)
  - name: step 2
    command: python some_batch_2.py
    dir: ${HOME}/jobs/
    depends:
      - step 1
```

### Complex example

More complex example is as follows:
```yaml
name: complex job
description: run python jobs

# Define environment variables
env:
  LOG_DIR: ${HOME}/jobs/logs
  PATH: /usr/local/bin:${PATH}
  
logDir: ${LOG_DIR}   # log directory to write standard output from the job steps
histRetentionDays: 3 # job history retention days (not for log files)
delaySec: 1          # interval seconds between job steps
maxActiveRuns: 1     # max parallel number of running step

# Define parameters
params: param1 param2 # they can be referenced by each steps as $1, $2 and so on.

# Define preconditions for whether or not the job is allowed to run
preconditions:
  - condition: "`printf 1`" # This condition will be evaluated at each execution of the job
    expected: "1"           # If the evaluation result do not match, the job is canceled

# Mail notification configs
mailOnError: true    # send a mail when a job failed
mailOnFinish: true   # send a mail when a job finished

# Job steps
steps:
  - name: step 1
    description: step 1 description
    dir: ${HOME}/jobs
    command: python some_batch_1.py $1
    mailOnError: false # do not send mail on error
    continueOn:
      failed: true     # continue to the next step regardless the error of this job
      canceled: true   # continue to the next step regardless the evaluation result of preconditions
    retryPolicy:
      limit: 2         # retry up to 2 times when the step failed
    # Define preconditions for whether or not the step is allowed to run
    preconditions:
      - condition: "`printf 1`"
        expected: "1"
  - name: step 2
    description: step 2 description
    dir: ${HOME}/jobs
    command: python some_batch_2.py $1
    depends:
      - step 1
```

The global config file `~/.jobctl/config.yaml` is useful to gather common settings such as mail-server configs or log directory.
