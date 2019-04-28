package gql

import (
  "github.com/graphql-go/graphql"
  "time"
)

type Message struct {
  Id        string    `json:"id,omitempty"`
  ChannelId string    `json:"channelId"`
  SenderId  string    `json:"senderId"`
  CreatedAt time.Time `json:"createdAt"`
  Content   string    `json:"content"`
}

func messages() *graphql.Field {
  return &graphql.Field{
    Type:    graphql.NewList(messageType()),
    Resolve: messagesQuery,
  }
}

func messageType() graphql.Type {
  return graphql.NewObject(graphql.ObjectConfig{
    Name: "Message",
    Fields: graphql.Fields{
      "id":        &graphql.Field{Type: graphql.String},
      "channelId": &graphql.Field{Type: graphql.String},
      "senderId":  &graphql.Field{Type: graphql.String},
      "createdAt": &graphql.Field{Type: graphql.DateTime},
      "content":   &graphql.Field{Type: graphql.String},
    },
  })
}

func messagesQuery(_ graphql.ResolveParams) (interface{}, error) {
  return []Message{
    {Id: "1",
      ChannelId: "1",
      SenderId:  "1",
      CreatedAt: time.Date(2016, 04, 1, 13, 02, 0, 0, time.UTC),
      Content:   "Hello"},
    {Id: "2",
      ChannelId: "1",
      SenderId:  "1",
      CreatedAt: time.Date(2016, 04, 1, 13, 02, 1, 0, time.UTC),
      Content:   "Again"},
  }, nil
}
