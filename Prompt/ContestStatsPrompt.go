package Prompt

import (
	"CodeforcesAPI/Objects"
	"CodeforcesAPI/Stats"
	"fmt"
)

func ContestStatsPrompt(){
	fmt.Println("******************************Welcome To Contest Stats Prompt******************************")
	//exit := false
	for;; {
		fmt.Printf("> Please Enter the contest id: ")
		var contestId uint
		fmt.Scan(&contestId)
		chngContest := false
		for;;{
			fmt.Println("> **********************************************************************************")
			fmt.Println("> Please Select One Of the Following Options\tContestID: ", contestId)
			fmt.Println("> Enter 1: GetTopRanker - Returns the User having first position of the selected contest")
			fmt.Println("> Enter 2: GetTopAchiever - Returns the User having highest increase in rating of the selected contest")
			fmt.Println("> Enter 3: GetNumberOfParticipants - Returns the Number of participants of the selected contest")
			fmt.Println("> Enter 4: GetAverageChangeInRating - Returns the Average increase or decrease in the ratings of parcipants of the selected contest")
			fmt.Println("> Enter 5: GetUserInContest - Returns the User info and its stats of the selected contest")
			fmt.Println("> Enter 6: GetNumberOfOvertakes - Returns the Number of Users who overtook the given user in the selected contest")
			fmt.Println("> Enter 7: GetNumberOfSurpassed - Returns the Number of Users who were overtooked by the given user in the selected contest")
			fmt.Println("> Enter 8: Change Contest")
			fmt.Println("> Enter 9: Go Back")
			fmt.Println("> **********************************************************************************")
			var option int
			fmt.Print("> ")
			fmt.Scan(&option)
			switch option {
			case 1:
				fmt.Println("> Loading...")
				fmt.Println("> Top Ranker")
				err, usr, rateChng := Stats.ContestStats{contestId}.GetTopRanker()
				if err != nil{
					fmt.Println(">", err.Error())
				}else{
					PrintUserObject(usr)
					PrintRatingChangeObject(rateChng)
				}
			case 2:
				fmt.Println("> Loading...")
				fmt.Println("> Top Achiever")
				err, usr, rateChng := Stats.ContestStats{contestId}.GetTopRatingGainer()
				if err != nil{
					fmt.Println(">", err.Error())
				}else{
					PrintUserObject(usr)
					PrintRatingChangeObject(rateChng)
				}
			case 3:
				fmt.Println("> Loading...")
				fmt.Println("> Number of Participants")
				err, numParticipants := Stats.ContestStats{contestId}.GetNumberOfCandidates()
				if err != nil{
					fmt.Println(">", err.Error())
				}else{
					fmt.Println("> ", numParticipants)
				}
			case 4:
				fmt.Println("> Loading...")
				fmt.Println("> Average Rating Change")
				err, avg := Stats.ContestStats{contestId}.GetAverageChangeInRatings()
				if err != nil{
					fmt.Println(">", err.Error())
				}else{
					fmt.Println("> ", avg)
				}
			case 5:
				var handle string
				fmt.Printf("> Enter handle: ")
				fmt.Scan(&handle)
				fmt.Println("> Loading...")
				fmt.Println("> User in Contest")
				err, usr, rateChng := Stats.ContestStats{contestId}.GetUserInContest(handle)
				if err != nil{
					fmt.Println(">", err.Error())
				}else{
					PrintUserObject(usr)
					PrintRatingChangeObject(rateChng)
				}
			case 6:
				var handle string
				fmt.Printf("> Enter handle: ")
				fmt.Scan(&handle)
				fmt.Println("> Loading...")
				fmt.Println("> User Overtakes in contest")
				err, num := Stats.ContestStats{contestId}.GetNumberOfOvertakes(handle)
				if err != nil{
					fmt.Println(">", err.Error())
				}else{
					fmt.Println("> ", num)
				}
			case 7:
				var handle string
				fmt.Printf("> Enter handle: ")
				fmt.Scan(&handle)
				fmt.Println("> Loading...")
				fmt.Println("> User surpassed in contest")
				err, num := Stats.ContestStats{contestId}.GetNumberOfSurpassed(handle)
				if err != nil{
					fmt.Println(">", err.Error())
				}else{
					fmt.Println("> ", num)
				}
			case 8:
				chngContest = true
			case 9:
				return
			default:
				fmt.Println("> Please Enter Valid Operation")
			}

			if chngContest{
				break
			}
		}
	}
}

func PrintUserObject(user Objects.User){
	fmt.Println("> *** User Info Start ***")
	fmt.Println("> Handle: " , user.Handle)
	fmt.Println("> FirstName: ", user.FirstName)
	fmt.Println("> LastName: ", user.LastName)
	fmt.Println("> Contribution: ", user.Contribution)
	fmt.Println("> Rating: ", user.Rating)
	fmt.Println("> MaxRating: ", user.MaxRating)
	fmt.Println("> LastTimeOnline: ", user.LastOnlineTimeSeconds)
	fmt.Println("> CountOfFriends: ", user.FriendOfCount)
	fmt.Println("> ")
}

func Test(user Objects.User){
	fmt.Println("Handle\tFName\tLName\tContri\tRating\tMaxRating\tTimeOnline\t")
}

func PrintRatingChangeObject(rateChng Objects.RatingChange){
	fmt.Println("> *** User Contest Stats ***")
	fmt.Println("> ContestName: ", rateChng.ContestName)
	fmt.Println("> Rank: ", rateChng.Rank)
	fmt.Println("> OldRating: ", rateChng.OldRating)
	fmt.Println("> NewRating: ", rateChng.NewRating)
	fmt.Println("> ")
}

//ContestId uint
//ContestName string
//Handle string
//Rank uint
//OldRating uint
//NewRating uint