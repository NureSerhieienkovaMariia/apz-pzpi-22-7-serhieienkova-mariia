package com.example.clinic.network

import com.example.clinic.model.PatientSignUpRequest
import com.example.clinic.model.RelativeSignUpRequest
import com.example.clinic.model.SignInRequest
import com.example.clinic.model.SignInResponse
import retrofit2.http.Body
import retrofit2.http.POST

interface AuthApi {
    @POST("patient/auth/sign-in")
    suspend fun signInPatient(@Body request: SignInRequest): SignInResponse

    @POST("relative/auth/sign-in")
    suspend fun signInRelative(@Body request: SignInRequest): SignInResponse

    @POST("patient/auth/sign-up")
    suspend fun signUpPatient(@Body request: PatientSignUpRequest): SignInResponse

    @POST("relative/auth/sign-up")
    suspend fun signUpRelative(@Body request: RelativeSignUpRequest): SignInResponse
}

