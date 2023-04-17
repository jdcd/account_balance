# Account_balance
basic MVP to get credit/debit balance from a transactions file and share result through email.

This project is a test project for show some programing concepts

## Architecture

This project was build keeping in mind the most important
things about [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

### Scaffolding 

**cmd:** access point, main function

**config:** Read initial application configuration, dependency injection configuration

**internal:** application logic organized according to the suggested layers of clean architecture, domain, application and infrastructure

**migration:** database scripts

**pks**: reusable utilities

**resources**: static resources with images, html templates and csv files

## Run project

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


## Data processing

### File management
inside the app container, there is the **/csv_files** directory with 3 subdirectories

* **pending**: repository of files pending processing
* **processed**: files processed successfully
* **error**: files that cannot be processed

### Persistence
The project has a relational database to store the transactions processed, and the summary of the transactions analyzed
you can consult the details of the tables and their columns in **migration/account_balance.sql**



### Workflow
if an error is found when reading the list of transactions, the record is ignored for the calculation and is stored in the ignore_transactions table


## Example Request

The project has a single interaction endpoint called v1/report
endpoint performs data validation and begins asynchronous file processing

**request body**:

        {
          "email_list": ["email1@stori.com","emailn@stori.com"]
        }
  
**response body**

        {
           "message": "", //general information about response
            "notification": {
                "to": [] // emails to which the notification will be sent
                "discarded": [] //discarded emails that will not be notified
            },
            "error": "" // error detail
        }

**curl**

            curl --location 'http://localhost:8080/v1/report' \
            --header 'Content-Type: application/json' \
            --data-raw '{"email_list": ["any_mail@stori.com"]}'

## Technologies

- [docker](https://www.docker.com/)
- [docker-compose](https://docs.docker.com/compose/)
- [gingonic](https://github.com/gin-gonic/gin)
- [testify](https://github.com/stretchr/testify)
- [PostgreSQL](https://www.postgresql.org/)