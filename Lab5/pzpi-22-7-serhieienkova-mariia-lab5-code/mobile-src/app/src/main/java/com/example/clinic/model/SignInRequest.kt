package com.example.clinic.model

data class SignInRequest(val email: String, val password: String)

data class SignInResponse(
    val access_jwt_token: String,
    val user_id: Int,
    val user_type: String
)

data class RelativeSignUpRequest(
    val name: String,
    val surname: String,
    val email: String,
    val password: String
)

data class PatientSignUpRequest(
    val name: String,
    val surname: String,
    val email: String,
    val password: String,
    val birthday: String, // ISO формат: "2001-07-17T00:00:00Z"
    val sex: Boolean       // true = жінка, false = чоловік
)
