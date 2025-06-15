package com.example.clinic.model

data class PatientNoteCreateRequest(
    val patient_id: Int,
    val note: String
)
