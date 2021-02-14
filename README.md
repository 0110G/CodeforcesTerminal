![Icon Image](https://icons.iconarchive.com/icons/alecive/flatwoken/256/Apps-Terminal-icon.png)

# CodeforcesTerminal
This is a Go Program that aims to provides various useful analytical tools for the Competitive Coding Platform 'Codeforces'.

## Installation
1. Clone the repository on your PC. <br>
2. Naviagate to the repository in your system. <br>
3. To compile, run ```go build main.go``` <br>
4. After compiling, you can observe an executable file named 'main'. <br>
5. To execute, run ```./main``` <br>

## Functionality Provided
Currently the program provides contest level and user level analytics. Some of them include:
###  Contest level analytics
1. ```Stats.ContestStats.GetTopRanker``` function provides the top ranker for a given contest and the corresponding users rating changes. <br>
2. ```Stats.ContestStats.GetTopAchiever``` function provides the top rating gainer for a given contest and the corresponding users rating changes. <br>
3. ```Stats.ContestStats.GetNumberOfParticipants``` function provides the number of participants for a given contest. <br>
4. ```Stats.ContestStats.AverageChangeInRating``` function provides the rating change in a contest. <br>
5. ```Stats.ContestStats.GetUserInContest.``` function provides the information and rating change for the given user. <br>
6. ```Stats.ContestStats.GetNumberOfOvertakes.``` function provides the number of users who were overtaken by the given user in the contest. <br>
7. ```Stats.ContestStats.GetNumberOfSurpassed.``` function provides the number of users who overtook the given user in the contest. <br>
