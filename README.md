# go-tinyurl

> URL shortner with Golang, VueJS and PostgreSQL

## Features

* Uses SHA1 Hashing to generate a random has for the input URL and the same has been mapped to the input URL in the postgreSQL database
* Returns the exsisting short URL if you are trying to shorten the same URL again and again
* The short URL slugs may consist of `[0-9A-Za-z+_]`

## FrontEnd

VueJS and Vuetify has been used to create the front end of this application

The same app has been deployed at [http://go-tiny.herokuapp.com/](http://go-tiny.herokuapp.com/)
