package main

import (
	"CodeforcesAPI/Query"
	"CodeforcesAPI/Stats"
	"fmt"
)

func main(){
	//err, lst := Query.ContestList{}.GetContestList(false, 5)
	//if err != nil{
	//	panic(err)
	//}
	//for i,v := range lst{
	//	fmt.Println(i, v)
	//}

	//handles := make([]string, 0)
	////handles = append(handles, "wadawdawd")
	//handles = append(handles, "Dainty_Lord")
	//err, usrs := Query.Users{}.GetUsers(handles)
	//if err != nil{
	//	panic(err)
	//}
	//fmt.Println(usrs)
	//
	//err, usr, chang := Stats.ContestStats{1457}.GetTopRanker()
	//if err != nil{
	//	panic(err)
	//}
	//fmt.Println(usr, chang)
	//
	//err1, usr1, gain := Stats.ContestStats{1457}.GetTopRatingGainer()
	//if err1 != nil{
	//	panic(err1)
	//}
	//fmt.Println(usr1, gain)
	//
	//fmt.Println(Stats.ContestStats{1457}.GetAverageChangeInRatings())
	//fmt.Println(Stats.ContestStats{1400}.GetUserInContest("Dainty_Lord"))
	fmt.Println(Stats.ContestStats{1457}.GetNumberOfOvertakes("Dainty_Lord"))

	//x, y := Query.UserQueries{}.GetUserStatus("Dainawdty_Lord", 1, 10)
	//if x != nil{
	//	panic(x)
	//}
	//fmt.Println(y)

	a, b := Query.ProblemSetQuery{}.GetProblems([]string{"implementation"}, "")
	if a != nil{
		panic(a)
	}
	fmt.Println(b)
}
