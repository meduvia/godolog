[![CodeFactor](https://www.codefactor.io/repository/github/meduvia/godolog/badge/main)](https://www.codefactor.io/repository/github/meduvia/godolog/overview/main)
# godolog 

Golang Log module that support file logging, database logging and terminal logging


## Feature Roadmap 

Terminal log => In progress

File log => In progress

Database Log (MySQL) => Soon (Q2 2022)

Database Log (Cassandra DB) => Soon (Q2 2022)

## Log format 

**Product**-**Service**-**Location**-**Timestamp**-**FuncType**-**Level**-**Code**-**Message**

**Product** must be an ID that refer to the product (a product have multiples services)

**Service** must the an ID that refer to the impacted service (Like for instance a API gateway)

**Location** must be where the error occured (like EU01, or could be more precise like EU01-DC1-R01-N01)

**Timestamp** Is when the log was triggered (in milliseconds)

**Level** `(verb,debug,info,warn,error,fatal)`

**Code** Is an OPcode, for instance auth system = 03 so all the log that are from the auth system are code=03 

**Message** Log message
