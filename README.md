# gogin_connect_psql
Rest API article postgreSQL

  

 **Create some migrations using migrate CLI. Here is an example:**

        migrate create -ext sql -dir db/migration -seq -digits 2 create_article_table
//digit 2 bu fayl tuzilvotganda fayl nomi 2xonali sonlardan tashkil topishi

//misol uchun 01migrate_up.sql, 01migrate_down.sql,02migrate_up.sql, 02migrate_down.sql

![image](https://user-images.githubusercontent.com/95979719/206504427-555e2138-f466-4337-944c-49549d433a57.png)


**Migrate DB:**

    migrate -path db/migration -database 'postgres://login:parol@localhost:5432/db_name?sslmode=disable' up





