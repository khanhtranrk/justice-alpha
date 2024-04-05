# justice-alpha
justice alpha

  Id uint64
  Type uint8
  SenderId uint64
  ReceiverId uint64
  SafeKey uint64
  CommitPoint uint64
  Content []byte

```
&Message{id, type, sender_id, receiver_id, safe_key, commit_point, content}
&Message{id, type, sender_id, receiver_id, safe_key, commit_point, content}
```

```
message {
    senderId
    receiverId
}

Error

Sender ----> Taistra Gate ----> Taistra ---->  Receiver Gate ----> Receiver
```
