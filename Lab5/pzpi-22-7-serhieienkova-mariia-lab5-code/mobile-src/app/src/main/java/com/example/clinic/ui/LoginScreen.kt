package com.example.clinic.ui

import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.*
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.input.PasswordVisualTransformation
import androidx.compose.ui.unit.dp
import androidx.lifecycle.viewmodel.compose.viewModel
import com.example.clinic.viewmodel.LoginViewModel
import com.example.clinic.R
import java.util.Locale

@Composable
fun LoginScreen(
    viewModel: LoginViewModel = viewModel(),
    onLoginSuccess: (String, String) -> Unit,
    onNavigateToSignUp: () -> Unit
) {
    val context = LocalContext.current
    val email = remember { mutableStateOf("") }
    val password = remember { mutableStateOf("") }

    val userType = viewModel.userType
    val message = viewModel.message

    LaunchedEffect(viewModel.message) {
        if (viewModel.token.isNotEmpty()) {
            onLoginSuccess(viewModel.token, viewModel.userType)
        }
    }

    Column(
        modifier = Modifier
            .fillMaxSize()
            .padding(24.dp),
        verticalArrangement = Arrangement.Center
    ) {
        Spacer(modifier = Modifier.height(16.dp))

        Text(text = stringResource(R.string.login_as))

        Row {
            RadioButton(selected = userType == "patient", onClick = { viewModel.userType = "patient" })
            Text(stringResource(R.string.patient), Modifier.clickable { viewModel.userType = "patient" })
            Spacer(modifier = Modifier.width(16.dp))
            RadioButton(selected = userType == "relative", onClick = { viewModel.userType = "relative" })
            Text(stringResource(R.string.relative), Modifier.clickable { viewModel.userType = "relative" })
        }

        Spacer(modifier = Modifier.height(16.dp))

        OutlinedTextField(
            value = email.value,
            onValueChange = { email.value = it },
            label = { Text(stringResource(R.string.email)) },
            modifier = Modifier.fillMaxWidth()
        )

        OutlinedTextField(
            value = password.value,
            onValueChange = { password.value = it },
            label = { Text(stringResource(R.string.password)) },
            modifier = Modifier.fillMaxWidth(),
            visualTransformation = PasswordVisualTransformation()
        )

        Spacer(modifier = Modifier.height(16.dp))

        Button(
            onClick = {
                viewModel.email = email.value
                viewModel.password = password.value
                viewModel.signIn()            },
            modifier = Modifier.fillMaxWidth()
        ) {
            Text(stringResource(R.string.login))
        }

        Spacer(modifier = Modifier.height(8.dp))

        Button(
            onClick = onNavigateToSignUp,
            modifier = Modifier.fillMaxWidth()
        ) {
            Text(stringResource(R.string.dont_have_account))
        }

        Spacer(modifier = Modifier.height(16.dp))
        Text(text = message)
    }
}
