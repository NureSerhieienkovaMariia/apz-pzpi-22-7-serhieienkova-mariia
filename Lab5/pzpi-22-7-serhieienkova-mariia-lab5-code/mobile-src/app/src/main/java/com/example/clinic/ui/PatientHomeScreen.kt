package com.example.clinic.ui

import androidx.compose.foundation.layout.*
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import com.example.clinic.model.FullPatient
import com.example.clinic.network.ApiClient
import kotlinx.coroutines.launch
import com.example.clinic.ui.components.Section
import com.example.clinic.R

@Composable
fun PatientHomeScreen(
    token: String,
    onLogout: () -> Unit,
    onOpenHealthNotes: (Int) -> Unit
) {
    val scope = rememberCoroutineScope()
    var patient by remember { mutableStateOf<FullPatient?>(null) }
    var error by remember { mutableStateOf<String?>(null) }

    LaunchedEffect(Unit) {
        scope.launch {
            try {
                patient = ApiClient.patientApi.getProfile("Bearer $token")
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
                patient?.let {
                    Text("${it.name} ${it.surname}", style = MaterialTheme.typography.bodyMedium)
                }
                Button(onClick = onLogout) {
                    Text(stringResource(R.string.logout))
                }
            }
        }

        if (error != null) {
            Text("${stringResource(R.string.error)}: $error", color = MaterialTheme.colorScheme.error)
            return
        }

        Button(
            onClick = { patient?.let { onOpenHealthNotes(it.id) } },
            modifier = Modifier.fillMaxWidth()
        ) {
            Text(stringResource(R.string.my_notes))
        }


        patient?.let { p ->
            LazyColumn(
                modifier = Modifier
                    .fillMaxSize()
                    .padding(8.dp)
            ) {
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
            }
        } ?: Box(modifier = Modifier.fillMaxSize(), contentAlignment = Alignment.Center) {
            CircularProgressIndicator()
        }
    }
}
