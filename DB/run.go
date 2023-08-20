package db

import (
	collector "cc/Collector"
	"log"
	"sync"
	"time"
)

var (
	Year  int
	Month int
	Day   int
	Hour  int
	Min   int
)

var (
	YearMutex  sync.Mutex
	MonthMutex sync.Mutex
	DayMutex   sync.Mutex
	HourMutex  sync.Mutex
	MinMutex   sync.Mutex
)

func UpdateTime(year, month, day, hour, minute int) {
	YearMutex.Lock()
	Year = year
	YearMutex.Unlock()

	MonthMutex.Lock()
	Month = month
	MonthMutex.Unlock()

	DayMutex.Lock()
	Day = day
	DayMutex.Unlock()

	HourMutex.Lock()
	Hour = hour
	HourMutex.Unlock()

	MinMutex.Lock()
	Min = minute
	MinMutex.Unlock()
}

func GetTime() (int, int, int, int, int) {
	var year, month, day, hour, minute int

	YearMutex.Lock()
	year = Year
	YearMutex.Unlock()

	MonthMutex.Lock()
	month = Month
	MonthMutex.Unlock()

	DayMutex.Lock()
	day = Day
	DayMutex.Unlock()

	HourMutex.Lock()
	hour = Hour
	HourMutex.Unlock()

	MinMutex.Lock()
	minute = Min
	MinMutex.Unlock()

	return year, month, day, hour, minute
}

func GET() {

	GenerateHashID()
	DBClient := CreateClient()

	Globaldata := &collector.GlobalDATA{}

	collector.NewClient()
	PodInfo, err := collector.KCLIENT.PullPodInfo()
	if err != nil {
		log.Fatal(err)
	}
	Clsuetrinfo, err := collector.KCLIENT.PullClusterInfo()
	if err != nil {
		log.Fatal(err)
	}
	Nodeinfo, err := collector.KCLIENT.PullNodeInfo()
	if err != nil {
		log.Fatal(err)
	}
	Deployment, err := collector.KCLIENT.PullDeploymentInfo()
	if err != nil {
		log.Fatal(err)
	}

	Globaldata.Timestamp = time.Now()
	Globaldata.Clsuterinfo = Clsuetrinfo
	Globaldata.PodInfo = PodInfo
	Globaldata.Nodeinfo = Nodeinfo
	Globaldata.DeployInfo = Deployment

	_ = DBClient.WriteData(*Globaldata)
	// _ = DBClient.ReadData("2023", "8", "20", "12", "15")

}

// create a json file for cluster info
// clusterJson, err := json.Marshal(cluster)
// if err != nil {
// 	panic(err)
// }

// // create a json file for pod info
// podJson, err := json.Marshal(pod)
// if err != nil {
// 	panic(err)
// }

// write to the file
// 	year, month, day, hour, minute := GetTime()
// 	fileName := fmt.Sprintf("%d-%d-%d-%d-%d.json", year, month, day, hour, minute)
// 	file, err := os.Create(fileName)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	// write to the file but use for loop

// 	for i := 0; i < len(pod.PodName); i++ {

// 		// write to the file
// 		_, err = file.WriteString(fmt.Sprintf("%s\n", pod.PodName[i]))
// 		if err != nil {
// 			panic(err)
// 		}

// 		// write to the file
// 		_, err = file.WriteString(fmt.Sprintf("%s\n", pod.PodCreation[i]))
// 		if err != nil {
// 			panic(err)
// 		}

// 	}

// }

// type DATA struct {
// 	ID          int64
// 	ClusterInfo collector.ClusterInfoStruct
// 	PodInfo     collector.PodInfoStruct
// }
