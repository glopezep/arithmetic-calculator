# arithmetic-calculator
## Demo link [Here!!](https://arithmetic-calculator-orcin.vercel.app/)


## Infrastructure

![Diagram](https://raw.githubusercontent.com/glopezep/arithmetic-calculator/main/diagram.png)

## How to run
 The project is monorepo so you'll have to go to the specific of guides

 * [Frontend](https://github.com/glopezep/arithmetic-calculator/tree/main/frontend)
 * [Server](https://github.com/glopezep/arithmetic-calculator/tree/main/server)


## Stack
- Fronted: Typescript, Node.js, Vue 3, Nuxt 3, Tailwind. This project is deployed in vercel for production.

- Backend: Go, GORM. this project is based on lambda functions for production.

## Architecture

 * [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

  * [Domain Driven Design](https://learn.microsoft.com/en-us/archive/msdn-magazine/2009/february/best-practice-an-introduction-to-domain-driven-design)

 * [CQRS](https://learn.microsoft.com/en-us/azure/architecture/patterns/cqrs)


## Security

- The backend services are secured by JWT token.

- The frontend  App sends a `Authorization: Bearer jwt-token` on each request in the server side, the session is stateless so the session is encrypted and store in the cookies and HTTP only

## Testing

- Few Unit tests for demo in the domain layer, service layer and infrastructure layer
- Few Integration tests for demo in interfaces layer

## Resilient

- Graceful shutdowns, etc

## Monitoring

- Cloud watch + Zerolog