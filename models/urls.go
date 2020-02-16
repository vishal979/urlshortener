package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"urlshortener/utils"
)

type OriginaLink struct{
	originaLink string
}

func rowExists(query string, args ...interface{}) bool {
	var exists bool
	query = fmt.Sprintf("SELECT exists (%s)", query)
	err := Client.QueryRow(query, args...).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		log.Println("error checking if row exists '%s' %v", args, err)
	}
	return exists
}

func GetOriginalURL(url string) (original_url string, err error){
	log.Println("GetOriginalURL called")
	if rowExists("SELECT id FROM shortener WHERE shortlink=?",url ) {
		//exists, and return the original url from the database
		result,err:=Client.Query("Select original_link from shortener where shortlink=?",url)
		if err!=nil{
			log.Println("error occurred while querying database for the shortenurl ",url, "with error ",err)
		}
		var originalUrl OriginaLink
		for result.Next() {
			result.Scan(&originalUrl.originaLink)
		}
		log.Println("original url found ", originalUrl.originaLink)

		return originalUrl.originaLink,nil
		}else{
		//the shortened url does not exist
		//return a 404 error
		return "",errors.New("URL NOT FOUND")
	}
}

func RegisterURL(originalurl string){
	log.Println("original url received to register is ",originalurl)
	if rowExists("SELECT id FROM shortener WHERE original_link=?",originalurl ) {
		log.Println("logging URL already exists")
		log.Println("exists")
		//url exists and no need to calculate hash and store, just return the shortened url from the database
	}else{
		//the url does not exist yet, make an entry for the url in  the database
		log.Println("Original URL doesnot Exist")
		hashedURL:=utils.CalculateHash(originalurl)
		log.Println("making entry in database with original url = ",originalurl,"and hashed url corresponding = ",hashedURL)
		if Client!=nil{
			log.Println("preparing sql insert statement")
			log.Println()
			stmt, err := Client.Prepare("INSERT INTO shortener(original_link, shortlink ) VALUES( ?,?)")
			defer stmt.Close()
			if err!=nil{
				log.Println("error occurred while preparing statement to insert into database with error ",err)
			}
			if _,err=stmt.Exec(originalurl,hashedURL);err!=nil{
				log.Println("error occurred while executing insert statement with error ",err)
			} else{
				log.Println("insertion successful")
			}
		}
	}
}