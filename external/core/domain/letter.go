package domain

type Letter struct {
  Id uint64
  Code uint64
  ForeignId uint64
  CommitTime uint64
  Message []byte
  Status uint32
}

// func (m *Message) ToBytes() ([]byte, error) {
//   receiverIdBytes := make([]byte, 8)
//   binary.BigEndian.PutUint64(receiverIdBytes, m.ReceiverId)
// 
//   senderIdBytes := make([]byte, 8)
//   binary.BigEndian.PutUint64(senderIdBytes, m.SenderId)
// 
//   safeKeyBytes := make([]byte, 8)
//   binary.BigEndian.PutUint64(safeKeyBytes, m.SafeKey)
// 
//   contentLengthBytes := make([]byte, 4)
//   binary.BigEndian.PutUint32(contentLengthBytes, uint32(len(m.Content)))
// 
//   commitPointBytes := make([]byte, 8)
//   binary.BigEndian.PutUint64(commitPointBytes, uint64(time.Now().UTC().UnixMilli()))
// 
//   var msg []byte
//   msg = append(msg, m.Type)
//   msg = append(msg, receiverIdBytes...)
//   msg = append(msg, senderIdBytes...)
//   msg = append(msg, safeKeyBytes...)
//   msg = append(msg, commitPointBytes...)
//   msg = append(msg, contentLengthBytes...)
//   msg = append(msg, m.Content...)
// 
//   return msg, nil
// }
// 
// func (m *Message) FromBytes(msg []byte) (*Message, error) {
//   msgLen := len(msg)
// 
//   // Type
//   if msgLen < 1 {
//     return nil, fmt.Errorf("Size of is wrong")
//   }
//   typ := msg[0]
// 
//   // ReceiverId
//   if len(msg) < 9 {
//     return nil, fmt.Errorf("Size of is wrong")
//   }
//   receiverId := binary.BigEndian.Uint64(msg[1:9])
// 
//   // SenderId
//   if len(msg) < 18 {
//     return nil, fmt.Errorf("Size of is wrong")
//   }
//   senderId := binary.BigEndian.Uint64(msg[9:18])
// 
//   // SafeKey
//   if len(msg) < 26 {
//     return nil, fmt.Errorf("Size of is wrong")
//   }
//   safeKey := binary.BigEndian.Uint64(msg[18:26])
// 
//   // CommitPoint
//   if len(msg) < 34 {
//     return nil, fmt.Errorf("Size of is wrong")
//   }
//   commitPoint := binary.BigEndian.Uint64(msg[18:26])
// 
//   // contentLen
//   if len(msg) < 30 {
//     return nil, fmt.Errorf("Size of is wrong")
//   }
//   contentLen := int(binary.BigEndian.Uint32(msg[26:30]))
// 
//   // Content
//   if len(msg) != 30 + contentLen {
//     return nil, fmt.Errorf("Size of is wrong")
//   }
//   content := msg[30:]
// 
//   return &Message{
//     typ,
//     receiverId,
//     senderId,
//     safeKey,
//     commitPoint,
//     content,
//   }, nil
// }
