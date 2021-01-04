# Feeler
Really simple application for running sentiment analysis on arbitrary phrases.

## Breakdown
### Server
The meat of this application is a web server for running text submitted through a basic text input. It only exposes two routes: `/` and `/sentiments`.

The `/sentiments` route responds to `POST` requests and receives a `s` field in its body with the phrase to analyse.

The root route renders a React application with the abovementioned text input. This text input hits the API on every keystroke. Not performant but Go seems to be up to the task of generating sentiments that quickly (thanks, Go!).

### Twitter analyser
I've also included a really simple command line application that receives a Twitter handle and an optional tweet limit and scans that user's Twitter feed for overall sentiment.

## Usage

### Server
You'll need to build the server itself, and then build the Javascript to mount the React application:

```sh
go build -o ./bin/app ./cmd/app
yarn install && yarn build
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

## y u use js
Because I'm a lazy son of a gun. I realise that this page tries to download like 164 KB of Javascript. It's a mess. I'm working on server-rendering the sentiment analysis page but Go's templating language isn't very nice (in my opinion).