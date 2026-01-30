# Ticketing System Architecture

![alt text](images/architecture.png)

- Client <----> API Gateway <-----> User Service || Ticket service || Payment service || Notification service
- Communication will be via gRPC or HTTP if need be in btn services.

**1. API Gateway**

-- Entry point for all external requests.
-- Handles HTTP requests from web.
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

- Acts as an entry point to the backend
- Accepts requests from clients.
- Authenticate requests
- Route to MS via gRPC

**2. User Service(Admins only)**

- Manage dashboard/admin users
- Users Table Fields
- id, email, password, role (admin, superadmin)

**2.Ticket service**

- Manage events, concerts, ticket types, inventory
- Events/Concerts table
- id, title (name of concert/event), description, location, date_time, status (Active, Inactive, Cancelled), created_by (FK to users.id)

- tickets table
- id, event_id (FK to events), category, price, total_stock, remaining, sales_start, sales_end

**3. Payment service**

- Handle actual ticket purchases, and store customer transaction info
- Orders table
- id, event_id, ticket_id, phone_number, email, quantity, total_amount, payment_ref, status
- The service stores actual customer info

**4. Notification service**

- Log and displatch notifications after payment success
- Notifications table
- id, order_id (FK to orders), channel (email), status, content, recipient, sent_at
- Triggered by payment service via gRPC or queue.
- You can make this async with RabbitMQ/NATs

---- **ALL ABOUT THE SERVICES ENDS HERE** ------

**gRPC Vs Message Queues**

1. used for synchronous comms - used for asynchornous comms.
2. Point to point, real-time - Pub/Sub or producer-consumer.
3. Blocking, in that requests block until response - No blocking
4. This is immediate, data is needed NOW - Delayed or decoupled messaging.

**Analogy**

- gRPC is like a phone call: you call someone and wait for them to answer and talk back.
- RabbitMQ/Kafka is like sending a letter: You drop the letter in a mailbox, the recipient can read it whenever they get to it.

**Use Cases**

1. gRPC

- You need an immediate result
- You're doing data queries or transactions.
- You want to chain services together.

3. RabbitMQ/Kafka

- You don't need an immediate response
- You want an event driven system (eg. send a notification after sth happens)
- You want to decouple services
- You want resilience (if one service is down, the message still waits)

- With RabbitMQ, even if the notification service is offline, it will receive the message later when it comes back up - gRPC cannot do that.
