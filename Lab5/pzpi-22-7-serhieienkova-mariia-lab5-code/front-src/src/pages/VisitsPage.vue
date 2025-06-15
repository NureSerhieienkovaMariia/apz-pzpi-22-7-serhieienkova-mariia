<template>
  <div class="page-container visits-container">
    <h2>{{ $t('visits') }}</h2>

    <div class="button-group">
      <button @click="fetchAllVisits">{{ $t('showAllVisits') }}</button>
      <button @click="fetchWeekVisits">{{ $t('showWeekVisits') }}</button>
    </div>

    <table v-if="visits.length > 0">
      <thead>
      <tr>
        <th>ID</th>
        <th>{{ $t('patientFullInfo') }}</th>
        <th>Reason</th>
        <th>Date</th>
        <th>Notes</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="visit in visits" :key="visit.id">
        <td>{{ visit.id }}</td>
        <td>
          <router-link
              v-if="treatmentPlanPatientMap[visit.treatment_plan_id] && patientInfoMap[treatmentPlanPatientMap[visit.treatment_plan_id]]"
              :to="`/patients/${treatmentPlanPatientMap[visit.treatment_plan_id]}/full-info`"
          >
            {{ patientInfoMap[treatmentPlanPatientMap[visit.treatment_plan_id]].name }}
            {{ patientInfoMap[treatmentPlanPatientMap[visit.treatment_plan_id]].surname }}
          </router-link>
          <span v-else>Loading...</span>
        </td>
        <td>{{ visit.reason }}</td>
        <td>{{ formatDate(visit.date) }}</td>
        <td>{{ visit.notes }}</td>
      </tr>
      </tbody>
    </table>

    <p v-else>{{ $t('noVisits') }}</p>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const visits = ref([])
const treatmentPlanPatientMap = ref({})
const patientInfoMap = ref({})

const fetchAllVisits = async () => {
  try {
    const token = localStorage.getItem('accessToken')
    const res = await fetch('http://localhost:8087/doctor/api/visit/', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`,
      },
    })

    if (!res.ok) {
      throw new Error(`Fetch visits failed: ${res.status}`)
    }

    const visitsData = await res.json()
    visits.value = visitsData
    await loadTreatmentPlanPatients(visitsData)
  } catch (err) {
    console.error('Помилка отримання visits:', err)
    alert('Помилка отримання visits: ' + err.message)
  }
}

const fetchWeekVisits = async () => {
  try {
    const token = localStorage.getItem('accessToken')
    const res = await fetch('http://localhost:8087/doctor/api/visit/week', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`,
      },
    })

    if (!res.ok) {
      throw new Error(`Fetch week visits failed: ${res.status}`)
    }

    const visitsData = await res.json()
    visits.value = visitsData
    await loadTreatmentPlanPatients(visitsData)
  } catch (err) {
    console.error('Помилка отримання week visits:', err)
    alert('Помилка отримання week visits: ' + err.message)
  }
}

const loadTreatmentPlanPatients = async (visitsData) => {
  try {
    const token = localStorage.getItem('accessToken')
    const map = {}
    const patientMap = {}

    for (const visit of visitsData) {
      const treatmentPlanId = visit.treatment_plan_id
      if (!map[treatmentPlanId]) {
        const res = await fetch(`http://localhost:8087/doctor/api/treatmentplan/${treatmentPlanId}`, {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`,
          },
        })

        if (res.ok) {
          const data = await res.json()
          const patientId = data.patient_id
          map[treatmentPlanId] = patientId

          if (!patientMap[patientId]) {
            const resPatient = await fetch(`http://localhost:8087/doctor/api/patients/${patientId}`, {
              method: 'GET',
              headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`,
              },
            })

            if (resPatient.ok) {
              const patientData = await resPatient.json()
              patientMap[patientId] = {
                name: patientData.name,
                surname: patientData.surname,
              }
            } else {
              console.warn(`Failed to fetch patient ${patientId}: ${resPatient.status}`)
              patientMap[patientId] = { name: '', surname: '' }
            }
          }
        } else {
          console.warn(`Failed to fetch treatment plan ${treatmentPlanId}: ${res.status}`)
          map[treatmentPlanId] = null
        }
      }
    }

    treatmentPlanPatientMap.value = map
    patientInfoMap.value = patientMap
  } catch (err) {
    console.error('Помилка отримання patient_id або patient info:', err)
  }
}

const formatDate = (timestamp) => {
  const date = new Date(timestamp)
  return date.toLocaleString()
}

onMounted(() => {
  fetchAllVisits()
})

</script>

<style scoped>
.page-container {
  padding: 0 24px;
}

h2 {
  margin-bottom: 16px;
}

.button-group {
  display: flex;
  gap: 8px;
  margin-bottom: 16px;
}

button {
  padding: 4px 10px;
  font-size: 13px;
  cursor: pointer;
  background-color: #4CAF50;
  color: white;
  border: none;
  border-radius: 4px;
}

button:hover {
  background-color: #45a049;
}

table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 16px;
}

th, td {
  border: 1px solid #ddd;
  padding: 8px;
  text-align: left;
}

th {
  background-color: #f2f2f2;
}

tr:nth-child(even) {
  background-color: #f9f9f9;
}

tr:hover {
  background-color: #f1f1f1;
}

p {
  margin-top: 16px;
}

.visits-container {
  max-width: 1000px;
  margin: 30px auto;
  background-color: white;
  padding: 20px 30px;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

</style>
