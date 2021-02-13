package Query

import (
	"CodeforcesAPI/Const"
	"CodeforcesAPI/Objects"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

type ContestQueries struct{}

// contest.ratingchanges
func(cQ ContestQueries)GetRatingsChange(contestId uint, numberOfRatings int)(error, []Objects.RatingChange){
	url := "https://codeforces.com/api/contest.ratingChanges?contestId=" + strconv.Itoa(int(contestId))
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil{
		return err, nil
	}
	jsonStr, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil{
		return readErr, nil
	}
	ratingsMap := make(map[string]interface{})
	unmarshallErr := json.Unmarshal(jsonStr, &ratingsMap)
	if unmarshallErr != nil{
		return unmarshallErr, nil
	}
	if val, isPresent := ratingsMap["status"].(string); isPresent {
		if val != "OK" {
			return errors.New("Status not OK"), nil
		}
		ratingChanges := make([]Objects.RatingChange, 0)
		contestRatingChangesInterfaceArray := ratingsMap["result"].([]interface{})
		for i := 0; i < numberOfRatings && i < len(contestRatingChangesInterfaceArray) && i<Const.GET_RATINGS_CHANGE_MAX_SIZE; i++ {
			rts := contestRatingChangesInterfaceArray[i].(map[string]interface{})
			newRatingChange := Objects.RatingChange{
				ContestId:   uint(rts["contestId"].(float64)),
				ContestName: rts["contestName"].(string),
				Handle:      rts["handle"].(string),
				Rank:        uint(rts["rank"].(float64)),
				OldRating:   uint(rts["oldRating"].(float64)),
				NewRating:   uint(rts["newRating"].(float64)),
			}
			ratingChanges = append(ratingChanges, newRatingChange)
		}
		return nil, ratingChanges
	}else{
		return errors.New("Status not present"), nil
	}
}

// contest.list
func(cQ ContestQueries)GetContestList(isGym bool, numberOfContests int)(error, []Objects.Contest){
	url := "https://codeforces.com/api/contest.list?gym=false"
	if isGym{
		url = "https://codeforces.com/api/contest.list?gym=true"
	}
	resp, err := http.Get(url)
	if err != nil{
		return err, nil
	}
	defer resp.Body.Close()
	jsonStr, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil{
		return readErr, nil
	}
	contestMap := make(map[string]interface{})
	unmarshallErr := json.Unmarshal(jsonStr, &contestMap)
	if unmarshallErr != nil{
		return unmarshallErr, nil
	}
	if val, isPresent := contestMap["status"].(string); isPresent{
		if val != "OK" {
			return errors.New("Status not OK"), nil
		}
		contestList := make([]Objects.Contest, 0)
		contestListInterfaceArray := contestMap["result"].([]interface{})
		for i:=0 ; i<numberOfContests && i<len(contestListInterfaceArray) && i<Const.GET_CONTEST_LIST_MAX_SIZE ; i++{
			cts := contestListInterfaceArray[i].(map[string]interface{})
			newContest := Objects.Contest{
				Id:          uint(cts["id"].(float64)),
				Name:        cts["name"].(string),
				Frozen:      cts["frozen"].(bool),
				Duration:    uint(cts["durationSeconds"].(float64)),
				Description: "",
				Difficulty:  -1,
				Phase:		 cts["phase"].(string),
			}
			if val, present := cts["description"]; present{
				newContest.Description = val.(string)
			}
			if val, present := cts["difficulty"]; present{
				newContest.Difficulty = int(val.(float64))
			}
			contestList = append(contestList, newContest)
		}
		return nil, contestList
	}else{
		return errors.New("Status not present"), nil
	}
	return nil, nil
}