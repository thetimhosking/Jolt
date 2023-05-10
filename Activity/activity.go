package Activity

type activity struct {
	Id               string `json:"id"`
	ActivityNumber   string `json:"activity-number"`
	PlannedStartDate string `json:"planned_start_date"`
	PlannedEndDate   string `json:"planned_end_date"`
	StartDate        string `json:"start-date"`
	EndDate          string `json:"end-date"`
	ActivityType     string `json:"activity-type"`
	Description      string `json:"description"`
	Priority         int    `json:"priority"`
	TargetTime       string `json:"target_time"`
	ActualTime       string `json:"actual_time"`
	Complexity       int    `json:"complexity"`
	Actions          []activityAction
	Issues           []activityIssue
}

type activityAction struct {
	Id               int
	ActivityID       string
	PlannedStartDate string `json:"planned_start_date"`
	PlannedEndDate   string `json:"planned_end_date"`
	StartDate        string `json:"start-date"`
	EndDate          string `json:"end-date"`
	TargetTime       string `json:"target_time"`
	AOctualTime      string `json:"actual_time"`
}

type activityIssue struct {
}

func (a *activity) newActivity() {

}
