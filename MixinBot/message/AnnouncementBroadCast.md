1. Whenever a user say something, check if user exist in db.
2. If user is not, store user to db.
3. If user is, skip

4. Set a binding in code:

   1. Set a slice stores publisher's id.

   2. Check if message come from those publisher every time.

   3. If message come from them, then check if message matches Announcement patten.

   4. If matches, use a function of broadcast to broadcast the message.

      The function should get all user_id and conversation_id from database, then send message to them one by one.

```go
func broadcast(){
    var userID, conversationID string
    sqlstatement := `SELECT user_id, conversation_id FROM users;`
    statement, err := db.Prepare(sqlstatement)
    if err != nil{
        log.Fatalln(err)
    }
    result, err := statement.Exec()
    for result.Next(){
        result.Scan(&userID, &conversationID)
        sendAnnouncement(userID, conversationID)
    }
}
```



