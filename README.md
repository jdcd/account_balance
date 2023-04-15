# Account_balance
basic MVP to get credit/debit balance from a transactions file and share result through email.

This project is a test project for show some programing concepts

This project was build keeping in mind the most important
things about [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)


### Make image

remember set your env variables to connect to repository,
migration/mongo-init-js contains init script for your repository

    docker build -t stori .

### Create container

    Run project
    docker run \
    -e EMAIL_SENDER={source mail} \
    -e EMAIL_PWD={source mail pwd} \
    -e EMAIL_SENDER_NAME={display name of source mail} \
    -e SMTP_SERVER={smtp server domain, for example "smtp.gmail.com"} \
    -e SMTP_PORT=587 \
    -e PORT=8181 \
    -p 8181:8181 \
    -d --name stori_engine \


### Basic Request
Process file

    curl --location 'http://localhost:8080/v1/report' \
    --header 'Content-Type: application/json' \
    --data-raw '{"email_list": ["cheems@gmail.com", "bob@space.com"]}'
