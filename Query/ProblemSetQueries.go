package Query

import (
	"CodeforcesAPI/Const"
	"CodeforcesAPI/Objects"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ProblemSetQuery struct{}

// problemset.problems
func (pQ ProblemSetQuery)GetProblems(tagList []string, problemSetName string)(error, []Objects.Problem){
	url :=  "https://codeforces.com/api/problemset.problems?tags="
	for i:=0 ; i<len(tagList) ; i++{
		if i == len(tagList)-1{
			url = url + "&problemsetName=" + problemSetName
		}else{
			url = url + tagList[i] + ";"
		}
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
	problemsMap := make(map[string]interface{})
	unmarshallErr := json.Unmarshal(jsonStr, &problemsMap)
	if unmarshallErr != nil{
		return unmarshallErr, nil
	}
	if val, isPresent := problemsMap["status"].(string); isPresent {
		if val != "OK" {
			return errors.New("Status not OK"), nil
		}
		problems := make([]Objects.Problem, 0)
		tmp := problemsMap["result"].(map[string]interface{})
		problemsInterfaceArr := tmp["problems"].([]interface{})
		statInterfaceArr := tmp["problemStatistics"].([]interface{})
		for i := 0;  i < len(problemsInterfaceArr) && i<len(statInterfaceArr) && i<Const.GET_PROBLEMS_MAX_SIZE; i++ {
			prob := problemsInterfaceArr[i].(map[string]interface{})
			stat := statInterfaceArr[i].(map[string]interface{})
			newProblem := Objects.Problem{
				ContestId:     uint(prob["contestId"].(float64)),
				ProblemSetName: problemSetName,
				Index:          prob["index"].(string),
				Name:           prob["name"].(string),
				Points:         0,
				Rating:         0,
				Tags:           tagList,
				SolveCount:     uint(stat["solvedCount"].(float64)),
			}
			if val, ok := prob["points"].(float64); ok{
				newProblem.Points = val
			}
			if val, ok := prob["rating"].(float64); ok{
				newProblem.Rating = uint(val)
			}
			fmt.Println(newProblem)
			problems = append(problems, newProblem)
		}
		return nil, problems
	}else{
		return errors.New("Status not present"), nil
	}
}