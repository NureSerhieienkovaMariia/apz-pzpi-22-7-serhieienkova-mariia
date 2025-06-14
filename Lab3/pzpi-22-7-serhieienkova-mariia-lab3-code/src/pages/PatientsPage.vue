<template>
  <div class="patients-container">
    <h2>{{ $t('patients') }}</h2>

    <div v-if="error" class="error">
      {{ error }}
    </div>

    <table v-if="patients.length > 0" class="patients-table">
      <thead>
      <tr>
        <th>ID</th>
        <th>{{ $t('name') }}</th>
        <th>{{ $t('surname') }}</th>
        <th>{{ $t('email') }}</th>
        <th>{{ $t('birthday') }}</th>
        <th>{{ $t('sex') }}</th>
        <th>{{ $t('actions') }}</th>
        <th>{{ $t('healthNotes') }}</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="patient in patients" :key="patient.id">
        <td>{{ patient.id }}</td>
        <td>{{ patient.name }}</td>
        <td>{{ patient.surname }}</td>
        <td>{{ patient.email }}</td>
        <td>{{ formatDate(patient.birthday) }}</td>
        <td>{{ formatSex(patient.sex) }}</td>
        <td>
          <router-link :to="`/patients/${patient.id}/full-info`" class="details-button">
            {{ $t('details') }}
          </router-link>
        </td>
        <td>
          <button @click="fetchHealthNotes(patient.id)">
            {{ $t('viewNotes') }}
          </button>
        </td>
      </tr>
      </tbody>
    </table>
    <div v-else-if="!error">
      {{ $t('noPatients') }}
    </div>

    <!-- Health Notes Modal -->
    <div v-if="showHealthNotesModal" class="modal-overlay">
      <div class="modal-content">
        <h3>{{ $t('healthNotesForPatient') }} ID {{ selectedPatientId }}</h3>
        <ul>
          <li v-for="note in healthNotes" :key="note.id" style="margin-bottom: 8px;">
            <strong>{{ formatDate(note.timestamp) }}:</strong> {{ note.note }}
          </li>
        </ul>
        <button @click="closeHealthNotesModal">{{ $t('close') }}</button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getPatients } from '../api/patients'
import { useRouter } from 'vue-router'

const patients = ref([])
const error = ref(null)
const router = useRouter()

const showHealthNotesModal = ref(false)
const healthNotes = ref([])
const selectedPatientId = ref(null)

const fetchHealthNotes = async (patientId) => {
  try {
    const token = localStorage.getItem('accessToken')
    const res = await fetch(`http://localhost:8087/doctor/api/healthnote/patient/${patientId}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`,
      },
    })

    if (!res.ok) {
      throw new Error(`Fetch health notes failed: ${res.status}`)
    }

    healthNotes.value = await res.json()
    selectedPatientId.value = patientId
    showHealthNotesModal.value = true // Відкрити modal
  } catch (err) {
    console.error('Помилка отримання health notes:', err)
    alert('Помилка отримання health notes: ' + err.message)
  }
}

const closeHealthNotesModal = () => {
  showHealthNotesModal.value = false
  healthNotes.value = []
  selectedPatientId.value = null
}

const formatDateNote = (timestamp) => {
  const date = new Date(timestamp)
  return date.toLocaleString()
}

const fetchPatients = async () => {
  error.value = null
  try {
    const token = localStorage.getItem('accessToken')
    if (!token) {
      router.push('/login')
      return
    }

    const response = await getPatients(token)
    patients.value = response
  } catch (err) {
    console.error(err)
    error.value = 'Помилка отримання списку пацієнтів: ' + err.message
  }
}

const formatDate = (dateStr) => {
  if (!dateStr) return '-'
  return new Date(dateStr).toLocaleDateString()
}

const formatSex = (sex) => {
  return sex ? 'Жінка' : 'Чоловік'
}

onMounted(() => {
  fetchPatients()
})
</script>

<style scoped>
.patients-container {
  max-width: 800px;
  margin: 30px auto;
  background-color: white;
  padding: 20px 30px;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.patients-table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 20px;
}

.patients-table th,
.patients-table td {
  border: 1px solid #ddd;
  padding: 8px;
  text-align: center;
}

.patients-table th {
  background-color: #f2f2f2;
}

.error {
  color: red;
  margin-bottom: 10px;
}

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0,0,0,0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 999;
}

.modal-content {
  background: white;
  padding: 20px;
  border-radius: 8px;
  width: 400px;
  max-height: 80vh;
  overflow-y: auto;
  box-shadow: 0 2px 8px rgba(0,0,0,0.3);
}

</style>