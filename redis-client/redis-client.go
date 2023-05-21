package redisclient

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func CreateClient(addr,password string,db int)*redis.Client{
	rdb := redis.NewClient(&redis.Options{
		Addr:	  addr,
		Password: password, 
		DB:		  db,  
	})
	return rdb
}

func StoreHashSalt(client *redis.Client,username,hash string)error{
	ctx := context.Background()
	err := client.Set(ctx, username, hash, 0).Err()
	return err
}

func RetrieveHashSalt(client *redis.Client,username string)(string,error){
	ctx := context.Background()
	val, err := client.Get(ctx, username).Result()
	return val,err
}
