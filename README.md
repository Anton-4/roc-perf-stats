# Roc Performance Stats

This is an app to track [roc](https://github.com/rtfeldman/roc) performance statistics
overtime. The app is a small Go backend hooked up to a postgres DB and
an Elm frontend for displaying charts/graphs. Everything is setup to auto-deploy on push/merge
to the main branch.

## Frontend

https://roc-perf-stats.rvcas.dev

The frontend is built in Elm. It is unlikely anyone other than [@rvcas](https://github.com/rvcas) will need to
run the backend locally so the data fetching endpoint is hardcoded to production.
The frontend is contributor friendly and pull requests are welcome.

> 🌳  built with [elm-spa](https://elm-spa.dev)

### Dependencies

This project requires the latest LTS version of [Node.js](https://nodejs.org/)

```bash
npm install -g elm elm-spa
```

### Running locally

```bash
elm-spa server  # starts this app at http:/localhost:1234
```

### other commands

```bash
elm-spa add    # add a new page to the application
elm-spa build  # production build
elm-spa watch  # runs build as you code (without the server)
```

### learn more

You can learn more at [elm-spa.dev](https://elm-spa.dev)

## Backend

https://roc-perf-stats-unnah.ondigitalocean.app

The backend is built in Go and uses [prisma](https://prisma.io) to manage the DB schema and data access.
The production database and backend both live on digitalocean. This backend unfortunately is not very
contributor friendly but please feel free to suggest improvements to the Go code. If there are any security concerns
please contact [@rvcas](https://github.com/rvcas) privately.

### Endpoints

* POST /benchmarks
* GET /benchmarks

`POST /benchmarks` is protected by a secret API_KEY and is only supposed to recieve benchmark data
from the [roc](https://github.com/rtfeldman/roc) CI servers. Contact [@rvcas](https://github.com/rvcas) if there are any questions about the backend.