package Stats

import (
	"CodeforcesAPI/Const"
	"CodeforcesAPI/Objects"
	"CodeforcesAPI/Query"
)

type UserStats struct{
	Handle string
}

type AccuracyObject struct{
	CorrectCount uint
	IncorrectCount uint
	Accuracy float64
	TagWiseAccuracy map[string]float64
}

/* GetUserAccuracy Returns
 * error variable and
 * Accuracy statistics
 */
func(uS UserStats)GetUserAccuracy() (error, AccuracyObject){
	err, subs := Query.UserQueries{}.GetUserStatus(uS.Handle, 1 ,Const.GET_USERS_MAX_SIZE)
	if err != nil{
		return err, AccuracyObject{
			CorrectCount:    0,
			IncorrectCount: 0,
			Accuracy:       0,
		}
	}
	topicWiseAccuracy := make(map[string]float64)
	totProb := make(map[string]float64)
	var correctCount uint = 0
	for i:=0 ; i<len(subs) ; i++{
		if subs[i].Verdict == "OK"{
			correctCount++
			for j:=0 ; j<len(subs[i].Prob.Tags) ; j++{
				topicWiseAccuracy[subs[i].Prob.Tags[j]]++
			}
		}
		for j:=0 ; j<len(subs[i].Prob.Tags) ; j++{
			totProb[subs[i].Prob.Tags[j]]++
		}
	}
	for k, v := range topicWiseAccuracy{
		topicWiseAccuracy[k] = v/totProb[k]
	}
	return nil, AccuracyObject{
		CorrectCount:    correctCount,
		IncorrectCount:  uint(len(subs)) - correctCount,
		Accuracy:        float64(correctCount)/float64(len(subs)),
		TagWiseAccuracy: topicWiseAccuracy,
	}
}

func (uS UserStats)GetUserInfo()(error, Objects.User){
	err, usr := Query.UserQueries{}.GetUsers([]string{uS.Handle})
	if err != nil{
		return err, Objects.User{}
	}
	return nil, usr[0]
}