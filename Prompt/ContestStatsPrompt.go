package Prompt

import (
	"CodeforcesAPI/Objects"
	"CodeforcesAPI/Stats"
	"fmt"
)

func ContestStatsPrompt(){
	fmt.Println("******************************Welcome To Contest Stats Prompt******************************")
	exit := false
	for;; {
		for;;{
			fmt.Println("Please Enter the contest id: ")
			var contestId uint
			fmt.Scan(&contestId)
			fmt.Println("**********************************************************************************")
			fmt.Println("Please Select One Of the Following Options")
			fmt.Println("Enter 1: GetTopRanker - Returns the User having first position of the selected contest")
			fmt.Println("Enter 2: GetTopAchiever - Returns the User having highest increase in rating of the selected contest")
			fmt.Println("Enter 3: GetNumberOfParticipants - Returns the Number of participants of the selected contest")
			fmt.Println("Enter 4: GetAverageChangeInRating - Returns the Average increase or decrease in the ratings of parcipants of the selected contest")
			fmt.Println("Enter 5: GetUserInContest - Returns the User info and its stats of the selected contest")
			fmt.Println("Enter 6: GetNumberOfOvertakes - Returns the Number of Users who overtook the given user in the selected contest")
			fmt.Println("Enter 7: GetNumberOfSurpassed - Returns the Number of Users who were overtooked by the given user in the selected contest")
			fmt.Println("Enter 8: Change Contest")
			fmt.Println("Enter 9: Go Back")
			fmt.Println("**********************************************************************************")
			var option int
			fmt.Scan(&option)
			switch option {
			case 1:
				err, usr, rateChng := Stats.ContestStats{contestId}.GetTopRanker()
				if err != nil{
					fmt.Println(err.Error())
				}else{
					PrintUserObject(usr)
					PrintRatingChangeObject(rateChng)
				}

			case 2:
			case 3:
			case 4:
			case 5:
			case 6:
			case 7:
			case 8:
				break
			case 9:
				return
			default:
				fmt.Println("Please Enter Valid Operation")
			}
		}
		if exit{
			break
		}
	}
}

func PrintUserObject(user Objects.User){
	fmt.Println("*** User Info Start ***")
	fmt.Println("Handle: " , user.Handle)
	fmt.Println("FirstName: ", user.FirstName)
	fmt.Println("LastName: ", user.LastName)
	fmt.Println("Contribution: ", user.Contribution)
	fmt.Println("Rating: ", user.Rating)
	fmt.Println("MaxRating: ", user.MaxRating)
	fmt.Println("LastTimeOnline: ", user.LastOnlineTimeSeconds)
	fmt.Println("CountOfFriends: ", user.FriendOfCount)
	fmt.Println()
}

func Test(user Objects.User){
	fmt.Println("Handle\tFName\tLName\tContri\tRating\tMaxRating\tTimeOnline\t")
}

func PrintRatingChangeObject(rateChng Objects.RatingChange){
	fmt.Println("*** User Contest Stats ***")
	fmt.Println("ContestName: ", rateChng.ContestName)
	fmt.Println("Rank: ", rateChng.Rank)
	fmt.Println("OldRating: ", rateChng.OldRating)
	fmt.Println("NewRating: ", rateChng.NewRating)
	fmt.Println()
}

//ContestId uint
//ContestName string
//Handle string
//Rank uint
//OldRating uint
//NewRating uint