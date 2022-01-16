> # Microservices (Node+TS) frontend could be React,Vue or Angular with RabbitMQ as messaging/events queues

> desc MicroServices:

1. Admin (mysql) via typeORM as admin_MS
2. Main (mongodb) via typeORM as main_MS

- Both microservice will communicate via rabbitMQ based on Events architecture

> Flow desc:

- If a product is created in "Admin" then an event will be sent to the "Main" via rabbitMQ to create the same product in Main
