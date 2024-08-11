# ecom-golang

build ecommerce project using golang

## structure golang project

project/
|-- api/
|   |-- handle/         # API handle request API
|   |-- middleware/     # API middleware
|   |-- router.go       # router
|-- cmd/
|   |-- app/            # start app
|   |-- cli/            # command line
|   |-- server/main.go  # start the main server
|-- config/             # anything relate to configuration
|-- docs/               # Document Project
|-- global/             # Define Global Variable Value using for All Service
|-- internal/           # code using for specific project
|   |-- controller/     # Define the controller project
|   |-- initialize/     # Define the init all service, third party in here (redis, database, env)
|   |-- middelwares/    # Define authen, cors, logger, ratelimit, errorhandle (before/after) the main handle function process
|   |-- model/          # model dbs (define ORM)
|   |-- repository/     # access database
|   |-- service/        # logic bussiness
|   |-- util/           # util helpful
|-- migrations/         # migration database
|-- pkg/                # libary has been defined
|   |-- logger/         # Define logger
|   |-- response/       # Define Response and Status Code
|   |-- settings/       # Define Config for all service (redis, mysql, ...)
|   |-- utlis/          # Update later
|-- scripts/            # Buill, install, ...
|-- test/               # test
|-- web/                # FE code (if we have it)
|-- .gitignore          # Git gitignore
|-- LICENSE             # License project
|-- README.md           # Description project
|__ go.mod              # go init

## Diagram project

![alt text](docs/image.png)
