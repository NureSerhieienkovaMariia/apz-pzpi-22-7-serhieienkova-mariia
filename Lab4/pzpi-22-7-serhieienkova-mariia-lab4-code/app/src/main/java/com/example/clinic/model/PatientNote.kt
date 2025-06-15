package com.example.clinic.model

data class PatientNote(
    val id: Int,
    val patient_id: Int,
    val timestamp: String,
    val note: String
)
