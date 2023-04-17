# Account_balance
basic MVP to get credit/debit balance from a transactions file and share result through email.

This project is a test project for show some programing concepts

This project was build keeping in mind the most important
things about [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)


### Run project

Configure your smtp server data inside docker-compose.yml file
    
      EMAIL_SENDER: ""
      EMAIL_PWD: ""
      EMAIL_SENDER_NAME: ""
      SMTP_SERVER: ""
      SMTP_PORT: ""
      SMTP_IDENTITY: ""

build the image for the container using the following command inside the main project folder

    docker-compose build

finally you can up the service with the following command

    docker-compose up

### Example Request
Process file

    curl --location 'http://localhost:8080/v1/report' \
    --header 'Content-Type: application/json' \
    --data-raw '{"email_list": ["cheems@gmail.com"]}'
