# Feeler
Really simple application for running sentiment analysis on arbitrary phrases.

## Breakdown
### Server
The meat of this application is a web server for running text submitted through a basic text input. It only exposes two routes: `/` and `/sentiments`.

The `/sentiments` route responds to `POST` requests and receives a `s` field in its body with the phrase to analyse.

The root route renders a React application with the abovementioned text input. This text input hits the API on every keystroke. Not performant but Go seems to be up to the task of generating sentiments that quickly (thanks, Go!).

### Twitter analyser
I've also included a really simple command line application that receives a Twitter handle and an optional tweet limit and scans that user's Twitter feed for overall sentiment.

## Installation and building

### Server
Build the server with the following command:

```sh
go build -o ./bin/app ./cmd/app
```

Once it's built, you can run it like any other executable:

```sh
./bin/app
```

### Twitter analyser
You'll need to [create a developer Twitter account](https://developer.twitter.com/en/portal/dashboard) and a Twitter App to use hit the Twitter API. We're using the v2 API here, which is still in beta, but which is significantly easier to use than v1.1.

Once you've created your app, generate a Bearer Token and save it to your `.env` file. See `.env.example` for usage.

Then you can build the executable:

```sh
go build -o ./bin/analyser ./cmd/analyser
```

And run it:

```sh
./bin/analyser --user=charlesharries
```

For a list of all available flags, you can run `--help` on the executable:

```sh
./bin/analyser --help
```
