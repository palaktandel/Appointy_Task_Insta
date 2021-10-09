package main
 
import (
    "context"
    "fmt"
    "time"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)
 
func close(client *mongo.Client, ctx context.Context,
           cancel context.CancelFunc){
            
    defer cancel()
     
    defer func() {
        if err := client.Disconnect(ctx); err != nil {
            panic(err)
        }
    }()
}
func connect(uri string)
(*mongo.Client, context.Context, context.CancelFunc, error) {
 
    ctx, cancel := context.WithTimeout(context.Background(),
                                       30 * time.Second)
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    return client, ctx, cancel, err
}
func insertOne
(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{})
(*mongo.InsertOneResult, error) {
 
    
    collection := client.Database(dataBase).Collection(col)
     
     
    result, err := collection.InsertOne(ctx, doc)
    return result, err
}
 

func insertMany
(client *mongo.Client, ctx context.Context, dataBase, col string, docs []interface{})
(*mongo.InsertManyResult, error) {
 
    
    collection := client.Database(dataBase).Collection(col)
	result, err := collection.InsertMany(ctx, docs)
    return result, err
}
 
func main() {
 
    
    client, ctx, cancel, err := connect("mongodb://localhost:27017")
    if err != nil {
        panic(err)
    }
     
    
    defer close(client, ctx, cancel)
 
    
    var document interface{}
     
     
    document = bson.D{
        {"username", godwinaustin},
        {"posts", 200},
    }
     
    
    insertOneResult, err := insertOne(client, ctx, "gfg",
                                      "marks", document)
     
    // handle the error
    if err != nil {
        panic(err)
    }
    fmt.Println(insertOneResult.InsertedID)
 
    
    var documents []interface{}
     
    // Storing into interface list.
    documents = []interface{}{
        bson.D{
            {"username", "kkwbeauty"},
            {"post", 550},   
        },
        bson.D{
            {"username", "cr7"},
            {"posts", 900},
        },
    }
 
 
    
    insertManyResult, err := insertMany(client, ctx, "gfg",
                                        "marks", documents)
     
    
    if err != nil {
        panic(err)
    }
     
    
    for id := range insertManyResult.InsertedIDs {
        fmt.Println(id)
    }
}