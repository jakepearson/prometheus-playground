# Prometheus Playground

Configures an environment to make changes to metrics and prometheus config so you can quickly see how the changes will effect how metrics are collected. `main.go` will open a webserver with scrape endpoints for:

* Normal metrics (defined in `metrics.go`)
* Static metrics files in the `test-data` directory

Now you can configure your `prometheus.yml` file to read from the different scrape endpoints and reconfigure the rewrite configs and any other changes you want.

Metrics are published back to `main.go` and printed so you can see which metrics are being passed on.

`.go` and `prometheus.yml` files that are changed will automatically reload the programs so you can make changes and just watch the output in your terminal.

## Setup

1. Make sure you already have `go` installed.
2. Clone me
3. Run `./setup.sh`. This will install:
  * [Prometheus](http://prometheus.io)
  * [Gin (auto build and restart your go code)](https://github.com/codegangsta/gin)
  * [Watcher (tell Prometheus to reload its configuration when it is saved)](https://github.com/radovskyb/watcher)
4. Run `./start.sh` which will:
  * Delete the `Prometheus` data folder
  * Start `Prometheus`, with its webserver open so it can be notified to refresh its configuration
  * Start `watcher` configured to curl to `Prometheus` when `prometheus.yml` changes
  * Start `gin` configured to recompile and restart main when any `.go` file changes