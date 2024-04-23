# justice-alpha

## Processed

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

type Taska struct {
    Exit chan bool
    InLetter chan *Letter
    InvalidLetter chan *Letter
    ValidLetter chan *Letter
    OutLetter chan *Letter
}

func (tk *Taska) HandleInLetter() {
    for  in range chans. chans.InLetter{
        err := cocoon.Verify(letter)
        if err != nil {
            taska.InvalidLetter <- letter
            return
        }
        taska.ValidLetter <- letter
    }
}

func (tk *Taska) HandleOutLetter() {
    for  in range chans. chans.OutLetter{
        err := cocoon.Verify(letter)
        if err != nil {
            taska.InvalidLetter <- letter
            return
        }
        taska.ValidLetter <- letter
    }
}

select {
    case <-chans.exit:
        exit
    case letter <-chans.InLetter:
        go handleLetterIn(letter)
    case letter <-chans.OutLetter:
        go handleOutLetter(letter)
    case 
}
