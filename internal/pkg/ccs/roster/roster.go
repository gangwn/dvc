package roster

type Roster struct {
	Participants map[string]Participant
}

func (roster *Roster) AddParticipant(participant Participant) bool {
	_, e := roster.Participants[participant.Id.String()]
	if e {
		return false
	}

	roster.Participants[participant.Id.String()] = participant
	return true;
}