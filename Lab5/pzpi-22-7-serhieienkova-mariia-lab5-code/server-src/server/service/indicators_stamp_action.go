package service

import (
	"clinic/server/repository"
	"clinic/server/structures"
	"fmt"
	"net/smtp"
	"time"
)

type IndicatorsStampActionService struct {
	indicatorsStampRepo        repository.IndicatorsStampRepo
	indicatorsNotificationRepo repository.IndicatorsNotificationRepo
	deviceRepo                 repository.DeviceRepo
	patientRepo                repository.PatientRepo
	patientRelativeRepo        repository.PatientRelativeRepo
	relativeRepo               repository.RelativeRepo
	treatmentPlantRepo         repository.TreatmentPlanRepo
	doctorRepo                 repository.DoctorRepo
}

func NewIndicatorsStampActionService(
	indicatorsStampRepo repository.IndicatorsStampRepo,
	indicatorsNotificationRepo repository.IndicatorsNotificationRepo,
	deviceRepo repository.DeviceRepo,
	patientRepo repository.PatientRepo,
	patientRelativeRepo repository.PatientRelativeRepo,
	relativeRepo repository.RelativeRepo,
	treatmentPlantRepo repository.TreatmentPlanRepo,
	doctorRepo repository.DoctorRepo,
) *IndicatorsStampActionService {
	return &IndicatorsStampActionService{
		indicatorsStampRepo:        indicatorsStampRepo,
		indicatorsNotificationRepo: indicatorsNotificationRepo,
		deviceRepo:                 deviceRepo,
		patientRepo:                patientRepo,
		patientRelativeRepo:        patientRelativeRepo,
		relativeRepo:               relativeRepo,
		treatmentPlantRepo:         treatmentPlantRepo,
		doctorRepo:                 doctorRepo,
	}
}

func (s *IndicatorsStampActionService) Create(input structures.IndicatorsStamp) error {
	input.Timestamp = time.Now().Format(time.RFC3339)
	id, err := s.indicatorsStampRepo.Create(input)
	if err != nil {
		return err
	}

	message := s.checkIndicators(input)
	if message != "" {
		err = s.sendEmail(input, id, message)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *IndicatorsStampActionService) sendEmail(input structures.IndicatorsStamp, id int, message string) error {
	indicatorsNotification := structures.IndicatorsNotification{
		IndicatorStampId: id,
		Message:          message,
		Timestamp:        input.Timestamp,
	}
	_, err := s.indicatorsNotificationRepo.Create(indicatorsNotification)
	if err != nil {
		return err
	}

	// Fetch device info
	device, err := s.deviceRepo.Get(input.DeviceId)
	if err != nil {
		return err
	}

	// Fetch patient info
	patient, err := s.patientRepo.GetById(device.PatientId)
	if err != nil {
		return err
	}

	patientRelatives, err := s.patientRelativeRepo.GetAllByPatientId(patient.Id)
	if err != nil {
		return err
	}

	var relativesEmails []string
	for _, patientRelative := range patientRelatives {
		relative, err := s.relativeRepo.GetById(patientRelative.RelativeID)
		if err != nil {
			return err
		}
		relativesEmails = append(relativesEmails, relative.Email)
	}

	treatmentPlans, err := s.treatmentPlantRepo.GetByPatientId(patient.Id)
	if err != nil {
		return err
	}

	if len(treatmentPlans) == 0 {
		return fmt.Errorf("no treatment plans found for patient ID %d", patient.Id)
	}

	// Fetch doctor's email address and info
	doctor, err := s.doctorRepo.GetById(treatmentPlans[0].DoctorID)
	if err != nil {
		return err
	}

	// Combine all email addresses
	emailRecipients := append(relativesEmails, doctor.Email)

	y, m, d := patient.Birthday.Date()
	// Append patient and doctor info to the message
	message += fmt.Sprintf("\n\nPatient Info:\nName: %s\nSurame: %s\nBirth Date: %d.%d.%d\n",
		patient.Name, patient.Surname, d, m, y)
	message += fmt.Sprintf("\nDoctor Info:\nName: %s\nSurname: %s\n",
		doctor.Name, doctor.Surname)

	auth := smtp.PlainAuth("", hostEmail, appPassword, smtpHost)

	return smtp.SendMail(smtpHost+":"+smtpPort, auth, hostEmail, emailRecipients, []byte(message))
}

func (s *IndicatorsStampActionService) checkIndicators(stamp structures.IndicatorsStamp) string {
	var message string

	if stamp.Pulse < 60 || stamp.Pulse > 100 {
		if stamp.Pulse < 50 || stamp.Pulse > 110 {
			message += "Critical heart rate detected. "
		} else {
			message += "Warning: Abnormal heart rate detected. "
		}
	}
	if stamp.SystolicBloodPressure < 90 || stamp.SystolicBloodPressure > 120 {
		if stamp.SystolicBloodPressure < 80 || stamp.SystolicBloodPressure > 130 {
			message += "Critical systolic blood pressure detected. "
		} else {
			message += "Warning: Abnormal systolic blood pressure detected. "
		}
	}
	if stamp.DiastolicBloodPressure < 60 || stamp.DiastolicBloodPressure > 80 {
		if stamp.DiastolicBloodPressure < 50 || stamp.DiastolicBloodPressure > 90 {
			message += "Critical diastolic blood pressure detected. "
		} else {
			message += "Warning: Abnormal diastolic blood pressure detected. "
		}
	}
	if stamp.Temperature < 36.1 || stamp.Temperature > 37.2 {
		if stamp.Temperature < 35.5 || stamp.Temperature > 38.0 {
			message += "Critical temperature detected. "
		} else {
			message += "Warning: Abnormal temperature detected. "
		}
	}

	return message
}

func (s *IndicatorsStampActionService) GetAll() ([]structures.IndicatorsStamp, error) {
	return s.indicatorsStampRepo.GetAll()
}

func (s *IndicatorsStampActionService) GetById(id int) (structures.IndicatorsStamp, error) {
	return s.indicatorsStampRepo.GetById(id)
}
