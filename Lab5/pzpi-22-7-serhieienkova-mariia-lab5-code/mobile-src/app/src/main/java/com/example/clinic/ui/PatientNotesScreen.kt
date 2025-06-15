package com.example.clinic.ui

import androidx.compose.foundation.layout.*
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import com.example.clinic.R
import com.example.clinic.model.PatientNote
import com.example.clinic.network.ApiClient
import kotlinx.coroutines.launch

@Composable
fun PatientNotesScreen(
    patientId: Int,
    token: String,
    onBack: () -> Unit
) {
    val scope = rememberCoroutineScope()
    var notes by remember { mutableStateOf<List<PatientNote>>(emptyList()) }
    var error by remember { mutableStateOf<String?>(null) }

    LaunchedEffect(patientId) {
        scope.launch {
            try {
                notes = ApiClient.relativeApi.getPatientNotes(patientId, "Bearer $token")
            } catch (e: Exception) {
                error = e.message
            }
        }
    }

    Column(Modifier.padding(16.dp)) {
        Text(stringResource(R.string.patient_notes), style = MaterialTheme.typography.headlineSmall)
        Spacer(Modifier.height(8.dp))

        if (error != null) {
            Text("${stringResource(R.string.error)}: $error", color = MaterialTheme.colorScheme.error)
        }

        if (notes.isEmpty() && error == null) {
            Text(stringResource(R.string.no_notes))
        }

        notes.forEach {
            Card(modifier = Modifier
                .fillMaxWidth()
                .padding(vertical = 4.dp)) {
                Column(Modifier.padding(8.dp)) {
                    Text("${stringResource(R.string.timestamp)}: ${it.timestamp}")
                    Text("${stringResource(R.string.note)}: ${it.note}")
                }
            }
        }

        Spacer(Modifier.height(16.dp))
        Button(onClick = onBack) {
            Text(stringResource(R.string.back))
        }
    }
}
