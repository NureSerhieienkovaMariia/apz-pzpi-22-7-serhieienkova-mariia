package com.example.clinic.ui

import androidx.compose.foundation.layout.*
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import com.example.clinic.R
import com.example.clinic.network.ApiClient
import kotlinx.coroutines.launch
import com.example.clinic.model.PatientNote
import com.example.clinic.model.PatientNoteCreateRequest

@Composable
fun PatientHealthNotesScreen(
    token: String,
    patientId: Int,
    onBack: () -> Unit
) {
    val scope = rememberCoroutineScope()
    var notes by remember { mutableStateOf<List<PatientNote>>(emptyList()) }
    var newNote by remember { mutableStateOf("") }
    var error by remember { mutableStateOf<String?>(null) }

    fun loadNotes() {
        scope.launch {
            try {
                notes = ApiClient.patientApi.getHealthNotes("Bearer $token")
                error = null
            } catch (e: Exception) {
                error = e.message
            }
        }
    }

    LaunchedEffect(Unit) { loadNotes() }

    Column(modifier = Modifier.padding(16.dp)) {
        Text(stringResource(R.string.my_medical_notes), style = MaterialTheme.typography.headlineSmall)
        Spacer(Modifier.height(8.dp))

        if (error != null) {
            Text("Error: $error", color = MaterialTheme.colorScheme.error)
        }

        notes.forEach {
            Card(Modifier.fillMaxWidth().padding(vertical = 4.dp)) {
                Column(Modifier.padding(8.dp)) {
                    Text(stringResource(R.string.timestamp_label, it.timestamp))
                    Text(stringResource(R.string.note_label, it.note))
                }
            }
        }

        Spacer(Modifier.height(16.dp))
        OutlinedTextField(
            value = newNote,
            onValueChange = { newNote = it },
            label = { Text(stringResource(R.string.new_note)) },
            modifier = Modifier.fillMaxWidth()
        )
        Spacer(Modifier.height(8.dp))
        Button(onClick = {
            scope.launch {
                try {
                    ApiClient.patientApi.createHealthNote(
                        token = "Bearer $token",
                        note = PatientNoteCreateRequest(patient_id = patientId, note = newNote)
                    )
                    newNote = ""
                    loadNotes()
                } catch (e: Exception) {
                    error = e.message
                }
            }
        }, modifier = Modifier.fillMaxWidth()) {
            Text(stringResource(R.string.add_note))
        }

        Spacer(Modifier.height(8.dp))
        Button(onClick = onBack, modifier = Modifier.fillMaxWidth()) {
            Text(stringResource(R.string.back))
        }
    }
}
