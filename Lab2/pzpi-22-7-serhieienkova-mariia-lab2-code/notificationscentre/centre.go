package notificationscentre

import (
	"clinic/server/service"
	"clinic/server/structures"
	"fmt"
	"log"
	"time"
)

type NotificationsCenter struct {
	services *service.Service
	ticker   *time.Ticker
}

func NewNotificationsCenter(services *service.Service) *NotificationsCenter {
	return &NotificationsCenter{
		services: services,
		ticker:   time.NewTicker(1 * time.Hour),
	}
}

func (nc *NotificationsCenter) Start() {
	defer nc.ticker.Stop()

	for {
		select {
		case <-nc.ticker.C:
			nc.checkVisits()
			nc.checkPrescriptions()
		}
	}
}

func (nc *NotificationsCenter) checkVisits() {
	now := time.Now()
	twentyFourHoursLater := now.Add(24 * time.Hour)
	twoHoursLater := now.Add(2 * time.Hour)

	visits, err := nc.services.VisitAction.GetAll()
	if err != nil {
		log.Printf("Error fetching visits: %v", err)
		return
	}

	for _, visit := range visits {

		treatmentPlan, err := nc.services.TreatmentPlanAction.Get(visit.TreatmentPlanID)
		if err != nil {
			log.Printf("Error fetching treatment plan for visit ID %d: %v", visit.ID, err)
			continue
		}

		if visit.Date.After(now) && visit.Date.Before(twentyFourHoursLater) {
			nc.createNotification(treatmentPlan.PatientID, "Upcoming Visit", fmt.Sprintf("You have a visit scheduled at %s", visit.Date.Format(time.RFC1123)))
		}
		if visit.Date.After(now) && visit.Date.Before(twoHoursLater) {
			nc.createNotification(treatmentPlan.PatientID, "Visit Reminder", fmt.Sprintf("Your visit is happening soon at %s", visit.Date.Format(time.RFC1123)))
		}
	}
}

func (nc *NotificationsCenter) checkPrescriptions() {
	now := time.Now()
	parser := &FrequencyParser{}

	prescriptions, err := nc.services.PrescriptionAction.GetAll()
	if err != nil {
		log.Printf("Error fetching prescriptions: %v", err)
		return
	}

	for _, prescription := range prescriptions {
		parsedFrequency, err := parser.Parse(prescription.Frequency)
		if err != nil {
			log.Printf("Error parsing frequency for prescription ID %d: %v", prescription.ID, err)
			continue
		}

		medicine, err := nc.services.MedicineAction.Get(prescription.MedicineID)
		if err != nil {
			log.Printf("Error fetching medicine for prescription ID %d: %v", prescription.ID, err)
			continue
		}

		treatmentPlan, err := nc.services.TreatmentPlanAction.Get(prescription.TreatmentPlanID)
		if err != nil {
			log.Printf("Error fetching treatment plan for prescription ID %d: %v", prescription.ID, err)
			continue
		}

		switch parsedFrequency.Type {
		case "at":
			if parsedFrequency.Time != nil && now.Hour() == parsedFrequency.Time.Hour() {
				nc.createNotification(treatmentPlan.PatientID, "Medication Reminder", fmt.Sprintf("It's time to take your medicine (%s).", medicine.Name))
			}
		case "every":
			if parsedFrequency.Interval != nil {
				timeSinceStart := now.Sub(now.Truncate(*parsedFrequency.Interval))
				if timeSinceStart < time.Hour {
					nc.createNotification(treatmentPlan.PatientID, "Medication Reminder", fmt.Sprintf("It's time to take your medicine (%s).", medicine.Name))
				}
			}
		default:
			log.Printf("Unhandled frequency type for prescription ID %d: %s", prescription.ID, parsedFrequency.Type)
		}
	}
}

func (nc *NotificationsCenter) createNotification(patientID int, topic, message string) {
	notification := structures.PatientNotification{
		Timestamp: time.Now(),
		PatientID: patientID,
		Topic:     topic,
		Message:   message,
	}
	_, err := nc.services.PatientNotificationAction.Create(notification)
	if err != nil {
		log.Printf("Error creating notification: %v", err)
	}
}
