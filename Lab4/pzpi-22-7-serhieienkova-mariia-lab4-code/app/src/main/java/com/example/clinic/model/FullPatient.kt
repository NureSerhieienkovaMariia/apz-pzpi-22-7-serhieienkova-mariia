package com.example.clinic.model

data class FullPatient(
    val id: Int,
    val name: String,
    val surname: String,
    val email: String,
    val birthday: String,
    val sex: Boolean,
    val diagnoses: List<Diagnosis>,
    val treatment_plans: List<TreatmentPlan>
)

data class Diagnosis(
    val id: Int,
    val name: String,
    val description: String,
    val recommendations: String
)

data class TreatmentPlan(
    val id: Int,
    val doctor: Doctor,
    val start_date: String,
    val end_date: String,
    val visits: List<Visit>,
    val prescriptions: List<Prescription>?
)

data class Doctor(
    val id: Int,
    val name: String,
    val surname: String,
    val email: String
)

data class Visit(
    val id: Int,
    val treatment_plan_id: Int,
    val reason: String,
    val date: String,
    val notes: String
)

data class Prescription(
    val medicine: Medicine,
    val dosage: String,
    val frequency: String
)

data class Medicine(
    val id: Int,
    val name: String,
    val description: String
)
