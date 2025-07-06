# Featurn 🚩

Featurn is a lightweight, developer-first feature flag service written in Go. It allows you to define feature flags with targeting rules and evaluate them based on dynamic user context.

## Features

:star: Evaluate boolean or multivariate feature flags

:star: Simple rule-based targeting (e.g., country == "US")

:star: In-memory flag store for rapid prototyping

:star: Easy to run locally with minimal setup

:hammer: Built with Go and the Chi router

## Project Structure

```bash
featurn/
├── cmd/api/ # HTTP server entrypoint
├── internal/
│ ├── api/ # HTTP routes and handlers
│ ├── eval/ # Flag evaluation engine
│ └── flags/ # Flag definitions and in-memory store
```

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/calindragomir/featurn.git

cd featurn
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Run the Server

```bash
go run ./cmd/api

Server will start on http://localhost:8080
```

### 4. API Endpoints

#### :dart: Check health

```bash
GET /health
```

Returns:

```nginx
OK
```

#### :dart: Evaluate a feature flag based on user context

```bash
POST /evaluate
```

##### Request Body:

```json
{
  "flag_key": "new-homepage",
  "context": {
    "country": "US"
  }
}
```

##### Response:

```json
{
  "value": true
}
```

##### Example request/response with curl

```bash
curl -X POST http://localhost:8080/evaluate \
  -H "Content-Type: application/json" \
  -d '{"flag_key":"new-homepage", "context":{"country":"US"}}'

{"value":true}
```

## Example Flags

Defined in `internal/flags/flags.go`:

```go
var testFlags = []Flag{
	{
		Key:          "new-homepage",
		DefaultValue: false,
		Rules: []Rule{
			{
				Attribute:   "country",
				Operator:    "eq",
				MatchValue:  "US",
				ReturnValue: true,
			},
		},
	},
}
```

## Roadmap

- [ ] Add persistent storage (e.g., PostgreSQL)
- [ ] REST API for flag CRUD operations
- [ ] Support more rule operators (gt, in, regex, etc.)
- [ ] Segment and user group support
- [ ] Go/Node/Python SDKs
- [ ] Admin UI with React
- [ ] OpenFeature compatibility
- [ ] Go/Node/Python SDKs
- [ ] Admin UI with React
- [ ] OpenFeature compatibility

## Contributing

Contributions, bug reports, and feature requests are welcome!
Feel free to fork the repo and open a pull request.

## License

MIT License

## Author

Built by Calin Dragomir (github.com/calindragomir)
