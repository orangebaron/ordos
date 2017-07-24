package players

type ComputerPlayer interface {
	Player
	RoutinelyCalledFunc()
	OthersVotedEvent( /*what args?*/ )
	//add events later
}

type computerPlayer struct {
	basicPlayer
	//will add state variables later
}

func (p *computerPlayer) RoutinelyCalledFunc() {}
func (p *computerPlayer) OthersVotedEvent()    {}
