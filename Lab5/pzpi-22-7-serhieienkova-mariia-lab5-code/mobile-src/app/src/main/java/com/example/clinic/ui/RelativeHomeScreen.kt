package com.example.clinic.ui

import androidx.compose.foundation.layout.*
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import com.example.clinic.model.Patient
import com.example.clinic.model.RelativeUser
import com.example.clinic.network.ApiClient
import kotlinx.coroutines.launch
import com.example.clinic.R

@Composable
fun RelativeHomeScreen(
    accessToken: String,
    onLogout: () -> Unit,
    onViewPatientDetails: (Int) -> Unit,
    onViewPatientNotes: (Int) -> Unit
) {
    val scope = rememberCoroutineScope()
    var user by remember { mutableStateOf<RelativeUser?>(null) }
    var patients by remember { mutableStateOf<List<Patient>>(emptyList()) }
    var error by remember { mutableStateOf<String?>(null) }

    LaunchedEffect(Unit) {
        scope.launch {
            try {
                user = ApiClient.relativeApi.getCurrentUser("Bearer $accessToken")
                patients = ApiClient.relativeApi.getPatients("Bearer $accessToken")
            } catch (e: Exception) {
                error = e.message
            }
        }
    }

    Column(modifier = Modifier.fillMaxSize().padding(16.dp)) {
        Surface(
            color = Color(0xFFDFFFE0),
            modifier = Modifier
                .fillMaxWidth()
                .padding(bottom = 8.dp),
            shadowElevation = 4.dp
        ) {
            Row(
                modifier = Modifier
                    .padding(16.dp)
                    .fillMaxWidth(),
                horizontalArrangement = Arrangement.SpaceBetween,
                verticalAlignment = Alignment.CenterVertically
            ) {
                Text(stringResource(R.string.app_name), style = MaterialTheme.typography.headlineSmall)
                user?.let {
                    Text("${it.name} ${it.surname}", style = MaterialTheme.typography.bodyMedium)
                }
                Button(onClick = onLogout) {
                    Text(stringResource(R.string.logout))
                }
            }
        }

        Spacer(modifier = Modifier.height(16.dp))

        if (error != null) {
            Text("${stringResource(R.string.error)}: $error", color = MaterialTheme.colorScheme.error)
        }

        if (patients.isEmpty()) {
            Text(stringResource(R.string.no_patients_found))
        } else {
            LazyColumn {
                items(patients) { patient ->
                    Card(
                        modifier = Modifier
                            .fillMaxWidth()
                            .padding(vertical = 4.dp)
                    ) {
                        Column(modifier = Modifier.padding(8.dp)) {
                            Text("${stringResource(R.string.name)}: ${patient.name} ${patient.surname}")
                            Text("${stringResource(R.string.email)}: ${patient.email}")
                            Text("${stringResource(R.string.birthday)}: ${patient.birthday}")
                            Text("${stringResource(R.string.sex)}: ${if (patient.sex) stringResource(R.string.female) else stringResource(R.string.male)}")

                            Row(modifier = Modifier.fillMaxWidth(), horizontalArrangement = Arrangement.SpaceBetween) {
                                Button(onClick = { onViewPatientDetails(patient.id) }) {
                                    Text(stringResource(R.string.details))
                                }
                                Button(onClick = { onViewPatientNotes(patient.id) }) {
                                    Text(stringResource(R.string.notes))
                                }
                            }
                        }
                    }
                }
            }
        }
    }
}
