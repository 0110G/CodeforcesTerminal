package Objects

type Submission struct{
	Id uint
	ContestId uint
	Prob Problem
	Verdict string
	PassedTestCount uint
	TimeConsumed uint
	MemoryConsumed uint
	Points uint
}
