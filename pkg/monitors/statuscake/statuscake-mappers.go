package statuscake

import (
	"strings"

	statuscake "github.com/StatusCakeDev/statuscake-go"
	endpointmonitorv1alpha1 "github.com/stakater/IngressMonitorController/v2/api/v1alpha1"
	"github.com/stakater/IngressMonitorController/v2/pkg/models"
)

// StatusCakeMonitorMonitorToBaseMonitorMapper function to map Statuscake structure to Monitor
func StatusCakeMonitorMonitorToBaseMonitorMapper(statuscakeData StatusCakeMonitorData) *models.Monitor {
	var m models.Monitor
	m.Name = statuscakeData.WebsiteName
	m.URL = statuscakeData.WebsiteURL
	m.ID = statuscakeData.TestID

	var providerConfig endpointmonitorv1alpha1.StatusCakeConfig
	providerConfig.TestTags = strings.Join(statuscakeData.Tags, ",")
	providerConfig.CheckRate = statuscakeData.CheckRate
	providerConfig.Paused = statuscakeData.Paused
	providerConfig.FollowRedirect = statuscakeData.FollowRedirect
	providerConfig.EnableSSLAlert = statuscakeData.EnableSSLAlert
	providerConfig.ContactGroup = strings.Join(statuscakeData.ContactGroup, ",")
	m.Config = &providerConfig
	return &m
}

// StatusCakeApiResponseDataToBaseMonitorMapper function to map Statuscake Uptime Test Response to Monitor
func StatusCakeApiResponseDataToBaseMonitorMapper(statuscakeData statuscake.UptimeTestResponse) *models.Monitor {
	var m models.Monitor
	// These top-level fields are correct:
	m.Name = statuscakeData.Data.Name
	m.URL = statuscakeData.Data.WebsiteURL
	m.ID = statuscakeData.Data.ID

	var providerConfig endpointmonitorv1alpha1.StatusCakeConfig

	// Tags also live under statuscakeData.Data.Tags
	providerConfig.TestTags = strings.Join(statuscakeData.Data.Tags, ",")

	// Fix references to use 'statuscakeData.Data'
	// CheckRate might be a custom type (UptimeTestCheckRate), so cast if needed:
	providerConfig.CheckRate = int(statuscakeData.Data.CheckRate)
	providerConfig.Paused = statuscakeData.Data.Paused
	providerConfig.EnableSSLAlert = statuscakeData.Data.EnableSSLAlert

	// ContactGroups is typically a []string in the Data struct
	providerConfig.ContactGroup = strings.Join(statuscakeData.Data.ContactGroups, ",")

	m.Config = &providerConfig
	return &m
}

// StatusCakeMonitorMonitorsToBaseMonitorsMapper function to map Statuscake structure to Monitor
func StatusCakeMonitorMonitorsToBaseMonitorsMapper(statuscakeData []StatusCakeMonitorData) []models.Monitor {
	var monitors []models.Monitor
	for _, payloadData := range statuscakeData {
		monitors = append(monitors, *StatusCakeMonitorMonitorToBaseMonitorMapper(payloadData))
	}
	return monitors
}
