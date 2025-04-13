# Proxy Pattern

## What is "Proxy Pattern" ?

The Proxy Pattern is a structural design pattern that provides a surrogate or placeholder for another object to control and manage access to it. This can include tasks such as lazy initialization, access control, caching, logging, and other cross-cutting concerns.

## Why is "Proxy Pattern" needed ?

1. **Access Control**: Proxies can control how and when a client can interact with a resource, enforcing authentication or security checks before delegating initiatives to the real object.

2. **Lazy Initialization**: Sometimes an object may be expensive to create or load. A proxy can defer creation until the resource is actually needed.

3. **Caching and Performance**: Proxies can reduce repeated computations or network delays by storing previous results and serving them directly for subsequent requests.

4. **Logging and Monitoring**: Proxies can add logging, debugging, or audit functionality without modifying the real object.

## Sample Program

This sample code demonstrates the use of the Proxy Pattern with a database. The `ProxyDatabase` caches results from the `RealDatabase` to reuse them for subsequent queries.

- **Database (Interface)**: Defines the contract for querying data.  
- **RealDatabase**: A concrete type that performs the heavy-lifting operation of querying.  
- **ProxyDatabase**: A type that intercepts requests, checks if results are already cached, and returns them if available. If not, it forwards the request to the real database and stores the result for later.

To run the sample code, execute:
```bash
go run ./cmd/proxy/main.go
```