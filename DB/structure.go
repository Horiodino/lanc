package db

import (
	collector "cc/Collector"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/fatih/color"
)

var (
	CALANDER *YearNode
	YEAR     int
	MONTH    int
	DAY      int
	HOUR     int
	MINUTE   int
)

type MinuteNode struct {
	Minute int
	Data   DATA
	Next   *MinuteNode
}

type HourNode struct {
	Hour    int
	Minutes *MinuteNode
	Next    *HourNode
}

type DayNode struct {
	Day   int
	Hours *HourNode
	Next  *DayNode
}

type MonthNode struct {
	Month string
	Days  *DayNode
	Next  *MonthNode
}

type YearNode struct {
	Year  int
	Month *MonthNode
	Next  *YearNode
}

type DATA struct {
	ID          int64
	ClusterInfo collector.ClusterInfoStruct
	PodInfo     collector.PodInfoStruct
}

func CreateCalendar() *YearNode {
	var calendar *YearNode

	// Create the calendar with 10 years
	for year := 2023; year <= 2032; year++ {
		newYear := &YearNode{Year: year, Month: nil, Next: nil}

		// Create the months for each year
		for _, month := range []string{
			"January", "February", "March", "April", "May", "June",
			"July", "August", "September", "October", "November", "December",
		} {
			newMonth := &MonthNode{Month: month, Days: nil, Next: nil}
			if newYear.Month == nil {
				newYear.Month = newMonth
			} else {
				lastMonth := newYear.Month
				for lastMonth.Next != nil {
					lastMonth = lastMonth.Next
				}
				lastMonth.Next = newMonth
			}

			// Create the days for each month
			for day := 1; day <= 31; day++ {
				newDay := &DayNode{Day: day, Hours: nil, Next: nil}
				if newMonth.Days == nil {
					newMonth.Days = newDay
				} else {
					lastDay := newMonth.Days
					for lastDay.Next != nil {
						lastDay = lastDay.Next
					}
					lastDay.Next = newDay
				}

				// Create the hours and minutes for each day
				for hour := 1; hour <= 24; hour++ {
					newHour := &HourNode{Hour: hour, Minutes: nil, Next: nil}
					if newDay.Hours == nil {
						newDay.Hours = newHour
					} else {
						lastHour := newDay.Hours
						for lastHour.Next != nil {
							lastHour = lastHour.Next
						}
						lastHour.Next = newHour
					}

					for minute := 0; minute < 60; minute++ {
						newMinute := &MinuteNode{Minute: minute, Next: nil}
						if newHour.Minutes == nil {
							newHour.Minutes = newMinute
						} else {
							lastMinute := newHour.Minutes
							for lastMinute.Next != nil {
								lastMinute = lastMinute.Next
							}
							lastMinute.Next = newMinute
						}
					}
				}
			}
		}

		if calendar == nil {
			calendar = newYear
		} else {
			currentYear := calendar
			for currentYear.Next != nil {
				currentYear = currentYear.Next
			}
			currentYear.Next = newYear
		}
	}

	return calendar
}

func INIT() {
	go func() {

		K8sclient := collector.NewClient()
		calendar := CreateCalendar()
		CALANDER = calendar
		for {

			YEAR = time.Now().Year()
			MONTH = int(time.Now().Month())
			DAY = time.Now().Day()
			HOUR = time.Now().Hour()
			MINUTE = time.Now().Minute()

			DBDATA := &DATA{}

			cluster, err := K8sclient.PullClusterInfo()
			if err != nil {
				panic(err)
			}
			pod, err := K8sclient.PullPodInfo()
			if err != nil {
				panic(err)
			}

			DBDATA.ClusterInfo = cluster
			DBDATA.PodInfo = pod

			AddDataToMinute(YEAR, MONTH, DAY, HOUR, MINUTE, *DBDATA)

			// Get past data
			// pastYear := YEAR
			// pastMonthIndex := MONTH
			// pastDay := DAY
			// pastHour := HOUR
			// pastMinute := MINUTE - 1

			// time.Sleep(30 * time.Second)

			// pastData, err := GetPastData(calendar, pastYear, pastMonthIndex, pastDay, pastHour, pastMinute)
			// if err != nil {
			// 	fmt.Println("Error getting past data:", err)
			// } else {
			// 	fmt.Println("ID: ", pastData.ID)
			// 	fmt.Println("Past data: ", pastData.ClusterInfo)
			// }

			// fmt.Println("------------------------------------------------------------")
			// time.Sleep(10 * time.Second)
		}
	}()
	// select {}

}

func AddDataToMinute(year int, month int, date int, hour int, minute int, data DATA) {
	fmt.Printf("ADDTIME: %02d:%02d\n", hour, minute)
	currentYear := CALANDER
	for currentYear != nil {
		if currentYear.Year == year {
			currentMonth := currentYear.Month
			for i := 0; i < month-1; i++ {
				if currentMonth == nil {
					fmt.Println("Month not found")
					return
				}
				currentMonth = currentMonth.Next
			}

			currentDay := currentMonth.Days
			for i := 1; i < date; i++ {
				if currentDay == nil {
					fmt.Println("Day not found")
					return
				}
				currentDay = currentDay.Next
			}

			currentHour := currentDay.Hours
			for i := 1; i < hour; i++ {
				if currentHour == nil {
					fmt.Println("Hour not found")
					return
				}
				currentHour = currentHour.Next
			}

			currentMinute := currentHour.Minutes
			for i := 0; i < minute; i++ {
				if currentMinute == nil {
					fmt.Println("Minute not found")
					return
				}
				currentMinute = currentMinute.Next
			}

			currentMinute.Data = data
			currentMinute.Data.ID = GenerateID()

			fmt.Print(color.YellowString("Adding data to database"))
			fmt.Print(".")
			time.Sleep(1 * time.Second)
			fmt.Print(".")
			time.Sleep(1 * time.Second)

			fmt.Println(color.BlueString("Data added to database..."))
			fmt.Print(".")
			time.Sleep(1 * time.Second)
			fmt.Println(currentMinute.Data.ID)
			fmt.Println(currentMinute.Data.ClusterInfo)
			fmt.Println("                                                      ")

			return
		}
		currentYear = currentYear.Next
	}
	fmt.Println("Year not found")
}

func GetPastData(yearNode *YearNode, year int, monthIndex int, day int, hour int, minute int) (DATA, error) {
	fmt.Println("GETPASTDATA: ", hour, minute)
	currentYear := CALANDER
	for currentYear != nil {
		if currentYear.Year == year {
			currentMonth := currentYear.Month
			for i := 0; i < monthIndex-1; i++ {
				if currentMonth == nil {
					return DATA{}, fmt.Errorf("Month not found")
				}
				currentMonth = currentMonth.Next
			}

			currentDay := currentMonth.Days
			for i := 1; i < day; i++ {
				if currentDay == nil {
					return DATA{}, fmt.Errorf("Day not found")
				}
				currentDay = currentDay.Next
			}

			currentHour := currentDay.Hours
			for i := 1; i < hour; i++ {
				if currentHour == nil {
					return DATA{}, fmt.Errorf("Hour not found")
				}
				currentHour = currentHour.Next
			}

			currentMinute := currentHour.Minutes
			for i := 0; i < minute-1; i++ {
				if currentMinute == nil {
					return DATA{}, fmt.Errorf("Minute not found")
				}
				currentMinute = currentMinute.Next
			}

			return currentMinute.Data, nil
		}
		currentYear = currentYear.Next
	}

	return DATA{}, fmt.Errorf("Year not found")
}

func PrintCalendar(calendar *YearNode) {
	currentYear := calendar
	for currentYear != nil {
		fmt.Printf("Year: %d\n", currentYear.Year)
		currentMonth := currentYear.Month
		for currentMonth != nil {
			fmt.Printf("  %s\n", currentMonth.Month)
			currentDay := currentMonth.Days
			for currentDay != nil {
				fmt.Printf("    Day: %d\n", currentDay.Day)
				currentHour := currentDay.Hours
				for currentHour != nil {
					fmt.Printf("      Hour: %d\n", currentHour.Hour)
					currentMinute := currentHour.Minutes
					for currentMinute != nil {
						fmt.Printf("        Minute: %d\n", currentMinute.Minute)
						fmt.Printf("          ClusterInfo: %#v\n", currentMinute.Data.ClusterInfo)
						fmt.Printf("          PodInfo: %#v\n", currentMinute.Data.PodInfo)
						currentMinute = currentMinute.Next
					}
					currentHour = currentHour.Next
				}
				currentDay = currentDay.Next
			}
			currentMonth = currentMonth.Next
		}
		currentYear = currentYear.Next
	}
}

func GenerateID() int64 {
	randomNumber, _ := rand.Int(rand.Reader, big.NewInt(100000))
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	id := (timestamp * 100000) + randomNumber.Int64()

	return id
}
