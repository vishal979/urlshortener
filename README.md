url shortener in golang, uses sha1 hash, using last 7 digits to correspond to original url.
visit localhost:8080 to register the url and then you will get the shortened url
but before that docker mysql need to be running up and use makefile to run, 
make dev is the command to run and make build is the command to build the binary


database= urlshortener

shortener table ddl:-
create table shortener(`original_link` varchar(250) NOT NULL, `shortlink` varchar(30) NOT NULL,`id` bigint(10) NOT NULL AUTO_INCREMENT, PRIMARY KEY (`id`))ENGINE=InnoDB DEFAULT CHARSET=utf8;


