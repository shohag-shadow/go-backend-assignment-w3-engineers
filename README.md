# Rental Property API

A small REST API for browsing rental properties. It loads a static JSON dataset
into memory at startup and exposes endpoints to read a single property or list
properties with a rich set of optional filters (price, star rating, review
score, review count, published status, property type, feed, bedrooms,
amenities, and result limit).

Built with [Go](https://go.dev/) and the [Beego v2](https://github.com/beego/beego) web framework.

## Requirements

- Go **1.25** or newer
- `curl` (only for the examples below)

## Project layout

```
.
├── conf/               # Beego app configuration (app.conf)
├── controllers/        # HTTP handlers
├── data/               # rental_properties.json (in-memory dataset)
├── models/             # Data + response + filter types and JSON loading
├── routers/            # Namespace/route registration
├── services/           # Data loading, filtering, response mapping (includes test file)
├── swagger/            # Generated swagger docs (served in dev mode)
├── tests/              
└── main.go             # Application entry point
```

## Running the project
(Run all the commands in project root directory)

1. Install dependencies:

   ```bash
   go mod tidy
   ```

2. Start the server:

   ```bash
   bee run
   ```

   The API listens on **http://localhost:8080** (configurable in
   `conf/app.conf` via `httpport`).

3. Run the tests:

   ```bash
   go test ./...
   ```



## API documentation

In development mode (`runmode = dev` in `conf/app.conf`) the generated Swagger
UI is served at:

```
http://localhost:8080/swagger
```

## Base URL

All endpoints are namespaced under `/v1`:

```
http://localhost:8080/v1
```

## Endpoints

### `GET /v1/properties`
```
http://localhost:8080/v1/properties
```

List properties, optionally filtered by any of the query parameters below.

| Parameter          | Type    | Description                                                        |
| ------------------ | ------- | ------------------------------------------------------------------ |
| `min_price`        | float   | Minimum USD price                                                  |
| `max_price`        | float   | Maximum USD price                                                  |
| `min_star_rating`  | int     | Minimum star rating                                                |
| `min_review_score` | float   | Minimum review score                                               |
| `min_reviews`      | int     | Minimum number of reviews                                          |
| `published`        | bool    | Published status (`true` / `false`)                                |
| `property_type`    | string  | One of `Hotel`, `House`, `Apartment`, `Villa`, `Resort`, `Hostel`  |
| `feed`             | int     | Feed id, one of `11`, `12`, `22`, `24`                             |
| `min_bedroom`      | int     | Minimum number of bedrooms (>= 1)                                  |
| `limit`            | int     | Maximum number of results (>= 1)                                   |
| `amenities`        | string  | Comma-separated amenities, e.g. `Pool,Gym`                         |

Returns:

```json
{
   "Result": {
      "Count": 2,
      "Items": [
         { "ID": "...", "Feed": 11, "GeoInfo": { ... }, "Property": { ... }, "Published": false },
         { "ID": "...", "Feed": 11, "GeoInfo": { ... }, "Property": { ... }, "Published": true }
      ]
   }
}
```

### `GET /v1/properties/{id}`

Fetch a single property by its id. Returns `404` with
`{"Error": "Property not found"}` when no property matches.

## Sample curl requests

List all properties:

```bash
curl "http://localhost:8080/v1/properties"
```

Limit the results to 5:

```bash
curl "http://localhost:8080/v1/properties?limit=5"
```

Filter by price range and minimum star rating:

```bash
curl "http://localhost:8080/v1/properties?min_price=100&max_price=300&min_star_rating=4"
```

Filter by property type and bedroom count:

```bash
curl "http://localhost:8080/v1/properties?property_type=Resort&min_bedroom=2"
```

Only published properties with a good review score:

```bash
curl "http://localhost:8080/v1/properties?published=true&min_review_score=4.5"
```

Filter by feed and minimum number of reviews:

```bash
curl "http://localhost:8080/v1/properties?feed=22&min_reviews=50"
```

Filter by amenities (comma-separated; matches if the property has any of them):

```bash
curl "http://localhost:8080/v1/properties?amenities=Pool,Gym"
```

Combine multiple filters:

```bash
curl "http://localhost:8080/v1/properties?property_type=Hotel&min_price=150&published=true&limit=10"
```

Get a single property by id:

```bash
curl "http://localhost:8080/v1/properties/BC-1000001"
```

### Pretty-print responses

Pipe any response through `jq` for readability:

```bash
curl -s "http://localhost:8080/v1/properties?limit=2" | jq
```

## Error responses

Invalid query parameters return HTTP `400`:

```json
{ "Error": "min_price should be a number" }
```

## Configuration

Settings live in `conf/app.conf`