<template>
  <div class="page-container notifications-container">
    <h2>{{ $t('notifications') }}</h2>

    <table v-if="notifications.length > 0">
      <thead>
      <tr>
        <th>ID</th>
        <th>Message</th>
        <th>Timestamp</th>
        <th>{{ $t('patientFullInfo') }}</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="notification in sortedNotifications" :key="notification.id">
        <td>{{ notification.id }}</td>
        <td>{{ notification.message }}</td>
        <td>{{ formatDate(notification.timestamp) }}</td>
        <td>
          <router-link
              v-if="notificationPatientMap[notification.indicator_stamp_id] && patientInfoMap[notificationPatientMap[notification.indicator_stamp_id]]"
              :to="`/patients/${notificationPatientMap[notification.indicator_stamp_id]}/full-info`"
          >
            {{ patientInfoMap[notificationPatientMap[notification.indicator_stamp_id]].name }}
            {{ patientInfoMap[notificationPatientMap[notification.indicator_stamp_id]].surname }}
          </router-link>
          <span v-else>Loading...</span>
        </td>
      </tr>
      </tbody>
    </table>

    <p v-else>{{ $t('noNotifications') }}</p>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'

const notifications = ref([])
const notificationPatientMap = ref({})
const patientInfoMap = ref({})

const fetchNotifications = async () => {
  try {
    const res = await fetch('http://localhost:8087/indicators_notification/', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
    })

    if (!res.ok) {
      throw new Error(`Fetch notifications failed: ${res.status}`)
    }

    const data = await res.json()
    notifications.value = data.sort((a, b) => new Date(b.timestamp) - new Date(a.timestamp))
    await loadNotificationPatients(data)
  } catch (err) {
    console.error('Помилка отримання notifications:', err)
    alert('Помилка отримання notifications: ' + err.message)
  }
}

const loadNotificationPatients = async (notificationsData) => {
  try {
    const token = localStorage.getItem('accessToken')
    const map = {}
    const patientMap = {}

    for (const notification of notificationsData) {
      const indicatorStampId = notification.indicator_stamp_id
      if (!map[indicatorStampId]) {
        const res = await fetch(`http://localhost:8087/indicators/${indicatorStampId}`, {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`,
          },
        })

        if (res.ok) {
          const data = await res.json()
          const patientId = data.patient_id
          map[indicatorStampId] = patientId

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
          console.warn(`Failed to fetch indicator ${indicatorStampId}: ${res.status}`)
          map[indicatorStampId] = null
        }
      }
    }

    notificationPatientMap.value = map
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
  fetchNotifications()
})

const sortedNotifications = computed(() => {
  return notifications.value.slice().sort((a, b) => new Date(b.timestamp) - new Date(a.timestamp))
})
</script>

<style scoped>
.page-container {
  padding: 0 24px;
}

h2 {
  margin-bottom: 16px;
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

.notifications-container {
  max-width: 1000px;
  margin: 30px auto;
  background-color: white;
  padding: 20px 30px;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

</style>
