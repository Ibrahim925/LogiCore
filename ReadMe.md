<h2>API Client</h2>

auth.go -> Requests auth endpoint with username, password, and grant type and receives the refresh and access tokens. It stores these values in the Client struct.

client.go -> Defines NewClient function and doRequest function. doRequest is a general function that takes in a GoLang request, performs the request, and returns the body.

model.go -> Defines request body/response structs.

service.go -> Set of public functions that define CRUD operations later used by the Provider.

<h2>Adding more requests</h2>
Create a new file for the desired resource (e.g. Package), using service.go as an example.
