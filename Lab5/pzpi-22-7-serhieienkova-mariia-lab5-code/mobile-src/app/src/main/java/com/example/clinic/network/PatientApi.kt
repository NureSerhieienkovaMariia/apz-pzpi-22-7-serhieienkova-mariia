package com.example.clinic.network

import com.example.clinic.model.FullPatient
import com.example.clinic.model.PatientNote
import com.example.clinic.model.PatientNoteCreateRequest
import retrofit2.http.Body
import retrofit2.http.GET
import retrofit2.http.Header
import retrofit2.http.POST

interface PatientApi {
    @GET("/patient/api/profile")
    suspend fun getProfile(@Header("Authorization") token: String): FullPatient

    @GET("/patient/api/healthnote/")
    suspend fun getHealthNotes(@Header("Authorization") token: String): List<PatientNote>

    @POST("/patient/api/healthnote/")
    suspend fun createHealthNote(
        @Header("Authorization") token: String,
        @Body note: PatientNoteCreateRequest
    ): PatientNote

}
