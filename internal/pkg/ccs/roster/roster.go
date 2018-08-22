package roster

type Roster struct {
	participants map[string]*Participant
}

func NewRoster() *Roster {
	return &Roster{make(map[string]*Participant)}
}

func (roster *Roster) Participants() map[string]*Participant {
	return roster.participants
}

func (roster *Roster) GetParticipant(userId string) *Participant{
	if p, ok := roster.participants[userId]; ok {
		return p
	}

	return nil
}

func (roster *Roster) AddParticipant(participant *Participant) bool {
	_, e := roster.participants[participant.UserId]
	if e {
		return false
	}

	roster.participants[participant.UserId] = participant
	return true;
}

