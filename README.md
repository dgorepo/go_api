# go_api


Author: Adilton Costa
Year: 2022
E-mail: 

Very simples API built in GO! to fetch data from a DataWarehouse.
<br></br>

<p align='center'>
    <img src="./.img/flash_go.png" alt="classic ETL template">
</p>

Credentials for db access are obtained from envs and can be exported like this:



    export APP_HOST=127.0.0.1
    export APP_PORT=5432
    export APP_DB_USERNAME=pg_username
    export APP_DB_PASSWORD=pg_password
    export APP_DB_NAME=pg_db

 //localhost:8080/api/v10/minhacesta/{uf}/{city}/{district}/{product}

Dockerfile:
    docker build -t go-api .
    docker run --name go-api -d -p 8080:8080 go-api
