# Rest-Api-with-GO
Repository where different ways of building REST Api with Golang are shown.

The purpose of this repository is to get familiar with Golang and also with REST Apis and how to build them using GO.

# Requirements

In order to run the code from this repository you need to have go installed and added to path.

```go
go version
```

# What is an API

API allows two or more software applications to talk to each other through a well-defined computing interface.
# What is REST

REST is acronym for REpresentational State Transfer. It is architectural style for distributed hypermedia systems and was first presented by Roy Fielding in 2000.

# Rest Architecture

REST defines 6 architectural constraints which make any web service – a truly RESTful API.


## Client-server

The client-server design pattern enforces the principle of separation of concerns which helps the client and the server components to evolve independently.

## Stateless

Statelessness mandates that each request from the client to the server must contain all of the information necessary to understand and complete the request.

The server cannot take advantage of any previously stored context information on the server.

For this reason, the session state is kept entirely on the client-side.


## Cacheable

The cacheble constraint requires that the data within a response should be implicitly or explicitly labeled as cacheable or non-cacheable.

If the response is cacheable then the client application is given the right to reuse the response data later, for equivalent requests and for a specified time period.

Caching is the ability to store copies of frequently accessed data in several places along the request-response path.

When a consumer requests a resource representation, the request goes through a cache or a series of caches (local cache, proxy cache, or reverse proxy) toward the service hosting the resource.

If any of the caches along the request path has a fresh copy of the requested representation, it uses that copy to satisfy the request. If none of the caches can satisfy the request, the request travels to the service (or origin server as it is formally known).

By using HTTP headers, an origin server indicates whether a response can be cached and, if so, by whom, and for how long.

Caches along the response path can take a copy of a response, but only if the caching metadata allows them to do so.

Not all actions are or should be cached.
The PUT method itself is semantically meant to put or create a resource. It is an idempotent operation, but it won't be used for caching because a DELETE could have occurred in the meantime
## Uniform interface

By applying the software engineering principle of generality to the components interface, the overall system architecture can be simplified and the visibility of interactions can be improved.

A uniform REST interface can be achieved by the following four constraints:

* Identification of resources
* Manipulation of resources through representations
* Self-descriptive messages
* Hypermedia as the engine of application state


## Layered system

The layered system style allows an architecture to be composed of hierarchical layers by constraining component behavior.

For example, in a layered system, each component cannot see beyond the immediate layer with which they are interacting.

## Code on demand (optional)

REST also allows client functionality to be extended by downloading and executing code in the form of applets or scripts.

This simplifies clients by reducing the number of features required to be pre-implemented. Servers can provide part of features delivered to the client in form of code and the client only needs to execute the code.
