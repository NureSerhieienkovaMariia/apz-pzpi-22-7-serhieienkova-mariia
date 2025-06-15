package com.example.clinic.network

import com.example.clinic.model.FullPatient
import com.example.clinic.model.Patient
import com.example.clinic.model.PatientNote
import com.example.clinic.model.RelativeUser
import retrofit2.http.GET
import retrofit2.http.Header
import retrofit2.http.Path

interface RelativeApi {
    @GET("relative/auth/current-user")
    suspend fun getCurrentUser(@Header("Authorization") token: String): RelativeUser

    @GET("relative/api/patients/")
    suspend fun getPatients(@Header("Authorization") token: String): List<Patient>

    @GET("relative/api/patients/{id}")
    suspend fun getFullPatient(
        @Path("id") id: Int,
        @Header("Authorization") token: String
    ): FullPatient

    @GET("relative/api/patients/{id}/notes")
    suspend fun getPatientNotes(
        @Path("id") id: Int,
        @Header("Authorization") token: String
    ): List<PatientNote>

}
