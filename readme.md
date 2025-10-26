# Ticketing System Architecture

![alt text](images/architecture.png)

- Client <----> API Gateway <-----> User Service || Ticket service || Payment service || Notification service
- Communication will be via gRPC or HTTP if need be in btn services.

**1. API Gateway**

- Entry point for all external requests.
- Handles HTTP requests from web.
- Authenticate requests via JWT.
- Route requests to internal gRPC services

**2. User Service**

- Manage user accounts (signups, login, etc)
- Registration and authentication
- JWT issuance.
- User role handling

**3. Ticket service**

- Event management (CRUD events)
- Availability
- Ticket booking

**4. Payment service**

- Handle all money transactions
- Verify transaction status

**5. Notification service**

- Notify users via email

**6. Docker & K8S**

# Structure

- Two types of users here

1. Admin - Registers via dashboard, creates concerts and tickets, and also views insights

- only admins will be stored in the user-service.

2. Customer - No registration, just browses tickets, pays and get receipts

**1. API Gateway**

**2. User Service(Admins only)**

- Manage dashboard/admin users
- Users Table Fields
- id, email, password, role (admin, superadmin)

**2.Ticket service**
- Manage events, concerts, ticket types, inventory
- Events/Concerts table
- id, title (name of concert/event), description, location, date_time, status, created_by

- tickets table
