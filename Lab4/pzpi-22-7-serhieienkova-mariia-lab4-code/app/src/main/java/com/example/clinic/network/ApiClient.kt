package com.example.clinic.network

import retrofit2.Retrofit
import retrofit2.converter.gson.GsonConverterFactory

object ApiClient {
    private val retrofit = Retrofit.Builder()
        .baseUrl("http://10.0.2.2:8087/")
        .addConverterFactory(GsonConverterFactory.create())
        .build()

    val authApi: AuthApi = retrofit.create(AuthApi::class.java)
    val relativeApi: RelativeApi = retrofit.create(RelativeApi::class.java)
    val patientApi: PatientApi = retrofit.create(PatientApi::class.java)
}
