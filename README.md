# justice-alpha

## Data structure

### Citizen

All citizens of the network.

Struct:

```
  Id uint64
  Name string
  ContactGate string
  LegalGate string
  RegistrationDate time.Time
```

### Contact

All contacts of a citizen

Struct:

```
  CitizenId uint64
  Permission uint32
```

### Letter

Communication method

Struct:

```
  Id uint64
  Code uint64
  ForeignId uint64
  CommitTime uint64
  Message []byte
```

To Delivery

```
[Type][Code][SenderId][ReceiverId][CommitTime][LengthOfMesage][Message]
```

Request Letter & Response Letter

```
[Type][Code][SenderId][ReceiverId][CommitTime][LengthOfMesage][Message]
   X     V      X          X            V            X            X

V: same
X: different
```

## Table

### Citizen
### Contact
### HandleLaterLetter
### ResponseLaterLetter
### SentLetter
### ReceivedLetter
### ReferLetter

```
```
