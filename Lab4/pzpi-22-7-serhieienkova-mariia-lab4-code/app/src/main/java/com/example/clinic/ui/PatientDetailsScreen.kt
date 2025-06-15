package com.example.clinic.ui

import androidx.compose.foundation.layout.*
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.ArrowBack
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import com.example.clinic.model.FullPatient
import com.example.clinic.network.ApiClient
import com.example.clinic.ui.components.Section
import kotlinx.coroutines.launch
import com.example.clinic.R

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun PatientDetailsScreen(
    patientId: Int,
    token: String,
    onBack: () -> Unit
) {
    val scope = rememberCoroutineScope()
    var patient by remember { mutableStateOf<FullPatient?>(null) }
    var error by remember { mutableStateOf<String?>(null) }

    LaunchedEffect(patientId) {
        scope.launch {
            try {
                patient = if (patientId == -1) {
                    ApiClient.patientApi.getProfile("Bearer $token")
                } else {
                    ApiClient.relativeApi.getFullPatient(patientId, "Bearer $token")
                }
            } catch (e: Exception) {
                error = e.message
            }
        }
    }

    Scaffold(
        topBar = {
            TopAppBar(
                title = { Text(stringResource(R.string.patient_details)) },
                navigationIcon = {
                    IconButton(onClick = onBack) {
                        Icon(Icons.Default.ArrowBack, contentDescription = stringResource(R.string.back))
                    }
                }
            )
        }
    ) { innerPadding ->

        LazyColumn(
            modifier = Modifier
                .padding(innerPadding)
                .padding(16.dp)
                .fillMaxSize()
        ) {
            if (error != null) {
                item {
                    Text("${stringResource(R.string.error)}: $error", color = MaterialTheme.colorScheme.error)
                }
                return@LazyColumn
            }

            patient?.let { p ->

                item {
                    Section(title = stringResource(R.string.general_info)) {
                        Text("${stringResource(R.string.name)}: ${p.name} ${p.surname}")
                        Text("${stringResource(R.string.email)}: ${p.email}")
                        Text("${stringResource(R.string.birthday)}: ${p.birthday}")
                        Text("${stringResource(R.string.sex)}: ${if (p.sex) stringResource(R.string.female) else stringResource(R.string.male)}")
                    }
                }

                item {
                    Section(title = stringResource(R.string.diagnoses)) {
                        if (p.diagnoses.isEmpty()) {
                            Text(stringResource(R.string.no_diagnoses))
                        } else {
                            p.diagnoses.forEach { d ->
                                Card(modifier = Modifier.padding(vertical = 4.dp)) {
                                    Column(Modifier.padding(8.dp)) {
                                        Text("${stringResource(R.string.name)}: ${d.name}", style = MaterialTheme.typography.bodyLarge)
                                        Text("${stringResource(R.string.description)}: ${d.description}")
                                        Text("${stringResource(R.string.recommendations)}: ${d.recommendations}")
                                    }
                                }
                            }
                        }
                    }
                }

                item {
                    Section(title = stringResource(R.string.treatment_plans)) {
                        if (p.treatment_plans.isEmpty()) {
                            Text(stringResource(R.string.no_treatment_plans))
                        } else {
                            p.treatment_plans.forEach { plan ->
                                Card(modifier = Modifier.padding(vertical = 6.dp)) {
                                    Column(Modifier.padding(8.dp)) {
                                        Text("${stringResource(R.string.plan)} ${plan.id}: ${plan.start_date} – ${plan.end_date}")
                                        Text("${stringResource(R.string.doctor)}: ${plan.doctor.name} ${plan.doctor.surname} (${plan.doctor.email})")

                                        Spacer(Modifier.height(6.dp))
                                        Text("${stringResource(R.string.visits)}:")
                                        plan.visits.forEach { visit ->
                                            Text(" • ${visit.date} — ${visit.reason}")
                                            Text("   ${stringResource(R.string.notes)}: ${visit.notes}")
                                        }

                                        Spacer(Modifier.height(6.dp))
                                        if (!plan.prescriptions.isNullOrEmpty()) {
                                            Text("${stringResource(R.string.prescriptions)}:")
                                            plan.prescriptions.forEach { presc ->
                                                Text(" • ${presc.medicine.name}: ${presc.dosage} (${presc.frequency})")
                                                Text("   → ${presc.medicine.description}")
                                            }
                                        }
                                    }
                                }
                            }
                        }
                    }
                }

                item {
                    Spacer(Modifier.height(16.dp))
                    Box(modifier = Modifier.fillMaxWidth(), contentAlignment = Alignment.Center) {
                        Button(onClick = onBack) {
                            Text(stringResource(R.string.back))
                        }
                    }
                }

            } ?: item {
                Box(modifier = Modifier.fillMaxSize(), contentAlignment = Alignment.Center) {
                    CircularProgressIndicator()
                }
            }
        }
    }
}
