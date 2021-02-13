package Stats

import (
	"CodeforcesAPI/Objects"
	"CodeforcesAPI/Query"
	"errors"
)

/* ContestStats is a Class which provides a contest specific
 * information. The inofrmation can be regarding a user
 */
type ContestStats struct{
	ContestId uint
}

/* GetTopRanker Returns
 * error variable
 * User Object for the the User having the rank = 1
 * Rating Change object for the returned user object for the given contest
 */
func(cS ContestStats)GetTopRanker()(error, Objects.User, Objects.RatingChange){
	err, ratingsList := Query.ContestQueries{}.GetRatingsChange(cS.ContestId,1)
	if err != nil{
		return err, Objects.User{}, Objects.RatingChange{}
	}
	err1, usr := Query.UserQueries{}.GetUsers([]string{ratingsList[0].Handle})
	if err1 != nil{
		return err1, Objects.User{}, Objects.RatingChange{}
	}
	return nil, usr[0], ratingsList[0]
}

/* GetTopRatingGainer Returns
 * error variable
 * User Object for the the User having the highest gain in ratings for the contest
 * Rating Change object for the user for the given contest
 */
func(cS ContestStats)GetTopRatingGainer()(error, Objects.User, Objects.RatingChange){
	err, ratingsList := Query.ContestQueries{}.GetRatingsChange(cS.ContestId,1000000000)
	if err != nil{
		return err, Objects.User{}, Objects.RatingChange{}
	}
	topRatingGainerHandle := ratingsList[0].Handle
	var maxDelta int
	ind := 0
	maxDelta = int(ratingsList[0].NewRating) - int(ratingsList[0].OldRating)
	for i:=1 ; i<len(ratingsList) ; i++{
		if maxDelta < int(ratingsList[i].NewRating) - int(ratingsList[i].OldRating){
			maxDelta = int(ratingsList[i].NewRating) - int(ratingsList[i].OldRating)
			topRatingGainerHandle = ratingsList[i].Handle
			ind = i
		}
	}
	err1, usr := Query.UserQueries{}.GetUsers([]string{topRatingGainerHandle})
	if err1 != nil{
		return err1, Objects.User{}, Objects.RatingChange{}
	}
	return nil, usr[0], ratingsList[ind]
}

/* GetNumberOfCandidate Returns
 * error variable
 * unsigned integer showing number of participants in the contest
 */
func(cS ContestStats)GetNumberOfCandidates()(error, uint){
	err, ratingsList := Query.ContestQueries{}.GetRatingsChange(cS.ContestId, 10000000000)
	return err, uint(len(ratingsList))
}

/* GetAverageChangeInRatings Returns
 * error variable
 * float value of average change in ratings for the contest
 */
func(cS ContestStats)GetAverageChangeInRatings()(error, float64){
	err, ratingsList := Query.ContestQueries{}.GetRatingsChange(cS.ContestId, 10000000000)
	if err != nil{
		return err, 0.0
	}
	var sum float64 = 0
	for i:=0 ; i<len(ratingsList) ; i++{
		sum = sum + float64(ratingsList[i].NewRating) - float64(ratingsList[i].OldRating)
	}
	return nil, (sum/float64(len(ratingsList)))
}

/* GetUserInContest Returns
 * error variable
 * User Object for the the User having the given handle
 * Rating Change object for the user for the given contest
 */
func(cS ContestStats)GetUserInContest(handle string)(error, Objects.User, Objects.RatingChange){
	err, ratingsList := Query.ContestQueries{}.GetRatingsChange(cS.ContestId, 10000000000)
	if err != nil{
		return err, Objects.User{}, Objects.RatingChange{}
	}
	err1, usr := Query.UserQueries{}.GetUsers([]string{handle})
	if err1 != nil{
		return err1, Objects.User{}, Objects.RatingChange{}
	}
	i:=0
	for i=0 ; i<len(ratingsList) ; i++{
		if ratingsList[i].Handle == handle{
			return nil, usr[0], ratingsList[i]
		}
	}
	return errors.New("User Not Found"), Objects.User{}, Objects.RatingChange{}
}

/* GetNumberOfOvertakes Returns
 * error variable
 * integer to denote number of users that overtook the given user in this contest ie.
 * initialRating[user] > initialRating[i] and finalRating[user] <= finalRating[i]
 */
func(cS ContestStats)GetNumberOfOvertakes(handle string)(error, int){
	err, _, rateChng := cS.GetUserInContest(handle)
	if err != nil{
		return err, 0
	}
	err1, ratingsList := Query.ContestQueries{}.GetRatingsChange(cS.ContestId, 10000000000)
	if err1 != nil{
		return err1, 0
	}
	ans := 0
	for i:=0 ; i<len(ratingsList) ; i++{
		if ratingsList[i].Handle == handle{continue}
		if ratingsList[i].OldRating < rateChng.OldRating && ratingsList[i].NewRating >= rateChng.NewRating{
			ans++
		}
	}
	return nil, ans
}

/* GetNumberOfOvertakes Returns
 * error variable
 * integer to denote number of users that the given user overtook in this contest ie.
 * initialRating[user] < initialRating[i] and finalRating[user] >= finalRating[i]
 */
func(cS ContestStats)GetNumberOfSurpassed(handle string)(error, int){
	err, _, rateChng := cS.GetUserInContest(handle)
	if err != nil{
		return err, 0
	}
	err1, ratingsList := Query.ContestQueries{}.GetRatingsChange(cS.ContestId, 10000000000)
	if err1 != nil{
		return err1, 0
	}
	ans := 0
	for i:=0 ; i<len(ratingsList) ; i++{
		if ratingsList[i].Handle == handle{continue}
		if ratingsList[i].OldRating > rateChng.OldRating && ratingsList[i].NewRating <= rateChng.NewRating{
			ans++
		}
	}
	return nil, ans
}

/* GetUpcomingContests returns
 * error variable
 * list of upcoming contest objects
 */
func GetUpcomingContests(isGym bool)(error, []Objects.Contest){
	err, contestList := Query.ContestQueries{}.GetContestList(isGym, 10)
	if err != nil{
		return err, nil
	}
	contests := make([]Objects.Contest, 0)
	for i:=0 ; i<len(contestList) ; i++{
		if contestList[i].Phase == "BEFORE" {
			contests = append(contests, contestList[i])
		}
	}
	return nil, contests
}

