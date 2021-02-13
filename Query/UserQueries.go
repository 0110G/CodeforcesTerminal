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

type UserQueries struct{}

// users.info
func(users UserQueries)GetUsers(handles []string)(error, []Objects.User){
	url := "https://codeforces.com/api/user.info?handles="
	for i:=0 ; i<len(handles) ; i++{
		if i == len(handles) - 1{
			url = url + handles[i]
		}else{
			url = url + handles[i] + ";"
		}
	}
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil{
		return err, nil
	}
	jsonStr, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil{
		return readErr, nil
	}

	usersMap := make(map[string]interface{})
	unmarshallErr := json.Unmarshal(jsonStr, &usersMap)
	if unmarshallErr != nil{
		return unmarshallErr, nil
	}
	if val, isPresent := usersMap["status"].(string); isPresent {
		if val != "OK" {
			return errors.New("Status not OK"), nil
		}
		usersList := make([]Objects.User, 0)
		usersInterfaceArray := usersMap["result"].([]interface{})
		for i := 0; i < len(usersInterfaceArray) && i<Const.GET_USERS_MAX_SIZE; i++ {
			usr := usersInterfaceArray[i].(map[string]interface{})
			newUser := Objects.User{
				Handle:                usr["handle"].(string),
				Contribution:          uint(usr["contribution"].(float64)),
				Rating:                uint(usr["rating"].(float64)),
				MaxRating:             uint(usr["maxRating"].(float64)),
				LastOnlineTimeSeconds: uint(usr["lastOnlineTimeSeconds"].(float64)),
				FriendOfCount:         uint(usr["friendOfCount"].(float64)),
				FirstName:			   "",
				LastName: 			   "",
			}
			if val, present := usr["firstName"]; present{
				newUser.FirstName = val.(string)
			}
			if val, present := usr["lastName"]; present{
				newUser.LastName = val.(string)
			}
			usersList = append(usersList, newUser)
		}
		return nil, usersList
	}else{
		return errors.New("Status not present"), nil
	}
}

// users.status
func(user UserQueries)GetUserStatus(handle string, from int, count int) (error, []Objects.Submission){
	url := "https://codeforces.com/api/user.status?handle=" + handle + "&from=" + strconv.Itoa(from) + "&count=" + strconv.Itoa(count)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil{
		return err, nil
	}
	jsonStr, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil{
		return readErr, nil
	}

	statusMap := make(map[string]interface{})
	uErr1 := json.Unmarshal(jsonStr, &statusMap)
	if uErr1 != nil{
		return uErr1, nil
	}

	if status, isPresent := statusMap["status"].(string); !isPresent{
		return errors.New("Status not present"), nil
	}else{
		if status != "OK"{
			return errors.New("Status not OK"), nil
		}
	}
	submissionList := statusMap["result"].([]interface{})
	subArr := make([]Objects.Submission, 0)
	for i:=0 ; i<len(submissionList) ; i++{
		submission := submissionList[i].(map[string]interface{})
		prob := submission["problem"].(map[string]interface{})
		newSubmission := Objects.Submission{
			Id:              uint(submission["id"].(float64)),
			ContestId:       uint(submission["contestId"].(float64)),
			Prob:            Objects.Problem{
				ContestId:      uint(prob["contestId"].(float64)),
				ProblemSetName: "",
				Index:          prob["index"].(string),
				Name:           prob["name"].(string),
				Points:         prob["points"].(float64),
				Rating:         uint(prob["rating"].(float64)),
				Tags:           GetStringArray(prob["tags"].([]interface{})), // nil,//prob["tags"].([]string),
				SolveCount: 	0,
			},
			Verdict:         submission["verdict"].(string),
			PassedTestCount: uint(submission["passedTestCount"].(float64)),
			TimeConsumed:    uint(submission["timeConsumedMillis"].(float64)),
			MemoryConsumed:  uint(submission["memoryConsumedBytes"].(float64)),
			Points:          0,
		}
		subArr = append(subArr, newSubmission)
	}
	return nil, subArr
}

func GetStringArray(arr []interface{})[]string{
	retArr := make([]string, 0)
	for i:=0 ; i<len(arr) ; i++{
		retArr = append(retArr, arr[i].(string))
	}
	return retArr
}

// user.ratings
func(users UserQueries)GetUserRatings(handle string)(error, []Objects.RatingChange){
	url := "https://codeforces.com/api/user.rating?handle=" + handle
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
		for i := 0;  i < len(contestRatingChangesInterfaceArray) && i<Const.GET_USER_RATINGS_MAX_SIZE; i++ {
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