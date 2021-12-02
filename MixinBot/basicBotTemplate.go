// Basic code template of mixin bot
package main

import (
	"os"
	"log"
	"fmt"
	"time"
	"flag"
	"context"
	"io/ioutil"
	"encoding/json"
	"encoding/base64"

	"github.com/gofrs/uuid"
	"github.com/fox-one/mixin-sdk-go"
)

var (
	config = flag.String("config", "","keystore file path")
)

func readFile(filename string) string{
	jsonFile, err := os.Open(filename)
	if err != nil{
		log.Println(err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil{
		fmt.Println(err)
	}
	return string(byteValue)
}

func respond(client *mixin.Client, ctx context.Context, msg *mixin.MessageView, category string, data []byte) error {
        payload := base64.StdEncoding.EncodeToString(data)
        reply := &mixin.MessageRequest{
                ConversationID: msg.ConversationID,
                RecipientID:    msg.UserID,
                MessageID:      uuid.Must(uuid.NewV4()).String(),
                Category:       category,
                Data:           payload,
        }
        return client.SendMessage(ctx, reply)
}

func handleText(client *mixin.Client, ctx context.Context, msg *mixin.MessageView, data []byte){
	switch string(data){
	case "Hi","?":
		respond(client, ctx, msg, mixin.MessageCategoryAppButtonGroup, []byte())
	case "你好","？":
		respond(client, ctx, msg, mixin.MessageCategoryAppButtonGroup, []byte())
	}
}

func message(ctx context.Context, client *mixin.Client)mixin.BlazeListenFunc{
	talk := func(ctx context.Context, msg *mixin.MessageView, userID string) error{
		if userID, _ := uuid.FromString(msg.UserID); userID == uuid.Nil {
			return nil
		}

                data, err := base64.StdEncoding.DecodeString(msg.Data)
                if err != nil {
                        return err
                }

		switch msg.Category{
		case mixin.MessageCategoryPlainText:
			handleText(client, ctx, msg, data)
		}


		return nil
	}
	return mixin.BlazeListenFunc(talk)
}

func main(){
	flag.Parse()
	f, err := os.Open(*config)
	if err != nil{
		log.Println(err)
	}
	var store mixin.Keystore
	if err := json.NewDecoder(f).Decode(&store); err != nil {
		log.Panicln(err)
	}

	client, err := mixin.NewFromKeystore(&store)
	if err != nil {
		log.Panicln(err)
	}

	ctx := context.Background()
	for {
		if err := client.LoopBlaze(ctx, mixin.BlazeListenFunc(message(ctx, client))); err != nil {
			log.Printf("LoopBlaze: %v", err)
		}
		time.Sleep(time.Second)
	}
}
