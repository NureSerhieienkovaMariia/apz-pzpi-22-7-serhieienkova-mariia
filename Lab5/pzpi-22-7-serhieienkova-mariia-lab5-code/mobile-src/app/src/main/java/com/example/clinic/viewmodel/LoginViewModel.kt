package com.example.clinic.viewmodel

import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.example.clinic.model.PatientSignUpRequest
import com.example.clinic.model.RelativeSignUpRequest
import com.example.clinic.model.SignInRequest
import com.example.clinic.network.ApiClient
import kotlinx.coroutines.launch

class LoginViewModel : ViewModel() {
    var userType by mutableStateOf("patient")
    var email by mutableStateOf("")
    var password by mutableStateOf("")
    var message by mutableStateOf("")
    var token by mutableStateOf("")

    fun signIn() {
        viewModelScope.launch {
            try {
                val request = SignInRequest(email, password)
                val response = when (userType) {
                    "patient" -> ApiClient.authApi.signInPatient(request)
                    "relative" -> ApiClient.authApi.signInRelative(request)
                    else -> throw IllegalArgumentException("Unknown user type")
                }
                token = response.access_jwt_token
                message = "Success: ${response.user_type}, ID: ${response.user_id}"
            } catch (e: Exception) {
                message = "Error: ${e.message}"
            }
        }
    }

    fun signUp(
        userType: String,
        name: String,
        surname: String,
        email: String,
        password: String,
        birthday: String = "",
        sex: Boolean = true
    ) {
        viewModelScope.launch {
            try {
                val response = when (userType) {
                    "patient" -> {
                        val req = PatientSignUpRequest(name, surname, email, password, birthday, sex)
                        ApiClient.authApi.signUpPatient(req)
                    }
                    "relative" -> {
                        val req = RelativeSignUpRequest(name, surname, email, password)
                        ApiClient.authApi.signUpRelative(req)
                    }
                    else -> throw IllegalArgumentException("Invalid user type")
                }
                token = response.access_jwt_token
                message = "Success: ${response.user_type}, ID: ${response.user_id}"
            } catch (e: Exception) {
                message = "Error: ${e.message}"
            }
        }
    }
}
