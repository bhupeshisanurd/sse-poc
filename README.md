# SSE Poc

Purpose of this repository is to understand the usage of Server Side Events using Go.

## Demo

![Screen Recording 2023-01-20 at 4 28 44 PM](https://user-images.githubusercontent.com/122530514/213681617-a7eccbcc-a271-480f-8a59-961e86bc62fe.gif)


## Usage

1. To run the server
   ```bash
   go run server/server.go
   ```

2. In a different terminal. Run an in-built client
   ```bash
   go run client/client.go
   ```

   **Or open the `index.html` file in a browser and click on `Trigger Event` button.**

3. Once the server is up, and you have a client running, send a trigger to `/trigger` endpoint
   ```bash
   curl -i localhost:8080/trigger
   ```

4. The client will show the data that was emitted from server. Any other client can connect on the following url
   ```bash
   http://127.0.0.1:8080/events?stream=messages
   ```

## Resources

- [EventSource - MDN](https://developer.mozilla.org/en-US/docs/Web/API/EventSource)
- [Using server-sent events - MDN](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events)
- [`r3labs/sse`](https://github.com/r3labs/sse/)
