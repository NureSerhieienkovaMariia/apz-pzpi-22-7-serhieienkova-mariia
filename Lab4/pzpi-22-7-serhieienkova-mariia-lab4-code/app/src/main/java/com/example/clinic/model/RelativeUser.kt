package com.example.clinic.model

data class RelativeUser(
    val id: Int,
    val name: String,
    val surname: String,
    val email: String,
    val user_type: String
)

data class Patient(
    val id: Int,
    val name: String,
    val surname: String,
    val email: String,
    val birthday: String,
    val sex: Boolean
)
