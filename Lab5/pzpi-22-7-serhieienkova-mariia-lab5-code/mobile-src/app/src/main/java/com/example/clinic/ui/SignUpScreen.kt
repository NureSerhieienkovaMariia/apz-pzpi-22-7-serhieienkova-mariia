package com.example.clinic.ui

import android.app.DatePickerDialog
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.*
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.input.PasswordVisualTransformation
import androidx.compose.ui.unit.dp
import androidx.lifecycle.viewmodel.compose.viewModel
import com.example.clinic.viewmodel.LoginViewModel
import java.time.LocalDate
import java.time.format.DateTimeFormatter
import java.util.*
import com.example.clinic.R

@Composable
fun SignUpScreen(
    viewModel: LoginViewModel = viewModel(),
    onNavigateToLogin: () -> Unit,
    onSignUpSuccess: (String, String) -> Unit
) {
    val userType = viewModel.userType
    var name by remember { mutableStateOf("") }
    var surname by remember { mutableStateOf("") }
    var email by remember { mutableStateOf("") }
    var password by remember { mutableStateOf("") }
    var sex by remember { mutableStateOf(true) }
    var birthdayDate by remember { mutableStateOf<LocalDate?>(null) }

    val message = viewModel.message
    val context = LocalContext.current
    val calendar = Calendar.getInstance()

    LaunchedEffect(viewModel.message) {
        if (viewModel.message.startsWith("Success: patient")) {
            onSignUpSuccess(viewModel.token, "patient")
        } else if (viewModel.message.startsWith("Success: relative")) {
            onSignUpSuccess(viewModel.token, "relative")
        }
    }

    val datePicker = DatePickerDialog(
        context,
        { _, year, month, dayOfMonth ->
            birthdayDate = LocalDate.of(year, month + 1, dayOfMonth)
        },
        calendar.get(Calendar.YEAR),
        calendar.get(Calendar.MONTH),
        calendar.get(Calendar.DAY_OF_MONTH)
    )

    Column(
        modifier = Modifier
            .fillMaxSize()
            .padding(24.dp),
        verticalArrangement = Arrangement.Center
    ) {

        Text(stringResource(R.string.sign_up_as))

        Row {
            RadioButton(selected = userType == "patient", onClick = { viewModel.userType = "patient" })
            Text(stringResource(R.string.patient), Modifier.clickable { viewModel.userType = "patient" })
            Spacer(modifier = Modifier.width(16.dp))
            RadioButton(selected = userType == "relative", onClick = { viewModel.userType = "relative" })
            Text(stringResource(R.string.relative), Modifier.clickable { viewModel.userType = "relative" })
        }

        Spacer(modifier = Modifier.height(16.dp))

        OutlinedTextField(value = name, onValueChange = { name = it }, label = { Text(stringResource(R.string.name)) }, modifier = Modifier.fillMaxWidth())
        OutlinedTextField(value = surname, onValueChange = { surname = it }, label = { Text(stringResource(R.string.surname)) }, modifier = Modifier.fillMaxWidth())
        OutlinedTextField(value = email, onValueChange = { email = it }, label = { Text(stringResource(R.string.email)) }, modifier = Modifier.fillMaxWidth())
        OutlinedTextField(value = password, onValueChange = { password = it }, label = { Text(stringResource(R.string.password)) }, modifier = Modifier.fillMaxWidth(), visualTransformation = PasswordVisualTransformation())

        if (userType == "patient") {
            OutlinedButton(
                onClick = { datePicker.show() },
                modifier = Modifier.fillMaxWidth()
            ) {
                Text(
                    text = birthdayDate?.format(DateTimeFormatter.ofPattern("yyyy-MM-dd"))
                        ?: stringResource(R.string.select_birthday)
                )
            }

            Row(verticalAlignment = Alignment.CenterVertically) {
                Text(stringResource(R.string.sex))
                RadioButton(selected = sex, onClick = { sex = true }); Text(stringResource(R.string.female))
                RadioButton(selected = !sex, onClick = { sex = false }); Text(stringResource(R.string.male))
            }
        }

        Spacer(modifier = Modifier.height(16.dp))

        Button(onClick = {
            val birthdayIso = birthdayDate?.let { it.toString() + "T00:00:00Z" } ?: ""
            viewModel.signUp(
                userType = userType,
                name = name,
                surname = surname,
                email = email,
                password = password,
                birthday = birthdayIso,
                sex = sex
            )
        }, modifier = Modifier.fillMaxWidth()) {
            Text(stringResource(R.string.register))
        }

        Spacer(modifier = Modifier.height(8.dp))

        Button(onClick = onNavigateToLogin, modifier = Modifier.fillMaxWidth()) {
            Text(stringResource(R.string.already_have_account))
        }

        Spacer(modifier = Modifier.height(16.dp))
        Text(text = message)
    }
}
