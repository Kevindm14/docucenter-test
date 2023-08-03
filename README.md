## Docucenter Test

En esta solución se utilizan las siguientes tecnologias:

1. [Go(1.20)](https://go.dev/)
2. [GORM (orm)](https://gorm.io/index.html)
3. [PostgreSQL](https://www.postgresql.org/)
4. [Validators con buffalo-validators](https://github.com/gobuffalo/validate)
5. [GoFiber](https://docs.gofiber.io/)

## Metodos implementados para la API

### Auth Routes
1. **Login**
 - Method: Post
 - Ruta: `/api/v1/auth/login`

2. **Register**
 - Method: Post
 - Ruta: `/api/v1/auth/register`

### Customers Routes
1. **GetCustomers**
 - Method: Get
 - Ruta: `/api/v1/customers`

2. **GetCustomerById**
 - Method: Get
 - Ruta: `/api/v1/customers/:id`

3. **CreateCustomers**
 - Method: Post
 - Ruta: `/api/v1/customers/`

4. **UpdateCustomer**
 - Method: Put
 - Ruta: `/api/v1/customers/:id`

5. **DeleteCustomer**
 - Method: Delete
 - Ruta: `/api/v1/customers/:id`


### Ground Deliveries Routes
1. **ListGroundDeliveries**
 - Method: Get
 - Ruta: `/api/v1/ground-deliveries`

2. **CreateGroundDelivery**
 - Method: Post
 - Ruta: `/api/v1/ground-deliveries`

3. **UpdateGroundDelivery**
 - Method: Put
 - Ruta: `/api/v1/ground-deliveries/:id`

4. **DeleteGroundDelivery**
 - Method: Delete
 - Ruta: `/api/v1/ground-deliveries/:id`

### Maritime Deliveries Routes
1. **GetMaritimeDeliveries**
 - Method: Get
 - Ruta: `/api/v1/maritime-deliveries`

2. **CreateMaritimeDelivery**
 - Method: Post
 - Ruta: `/api/v1/maritime-deliveries`

3. **UpdateMaritimeDelivery**
 - Method: Put
 - Ruta: `/api/v1/maritime-deliveries/:id`

4. **DeleteMaritimeDelivery**
 - Method: Delete
 - Ruta: `/api/v1/maritime-deliveries/:id`

## Relacion entre clientes y entregas

Descripción de la relación:
  - Un cliente puede tener varios planes de entrega.
  - Cada plan de entrega está vinculado a un único cliente.