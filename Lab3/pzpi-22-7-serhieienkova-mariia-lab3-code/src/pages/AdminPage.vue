<template>
  <div class="admin-container">
    <h1>Admin Dashboard</h1>

    <div class="admin-actions">
      <button @click="backupDatabase">Backup Database</button>
      <button @click="restoreDatabase">Restore Database</button>
    </div>

    <div class="admin-section">
      <h2>Patients</h2>
      <table>
        <thead>
          <tr>
            <th>ID</th><th>Name</th><th>Email</th><th>Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="patient in patients" :key="patient.id">
            <td>{{ patient.id }}</td>
            <td>{{ patient.name }} {{ patient.surname }}</td>
            <td>{{ patient.email }}</td>
            <td><button @click="deletePatient(patient.id)">Delete</button></td>
          </tr>
        </tbody>
      </table>
    </div>

    <div class="admin-section">
      <h2>Doctors</h2>
      <table>
        <thead>
          <tr>
            <th>ID</th><th>Name</th><th>Email</th><th>Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="doctor in doctors" :key="doctor.id">
            <td>{{ doctor.id }}</td>
            <td>{{ doctor.name }} {{ doctor.surname }}</td>
            <td>{{ doctor.email }}</td>
            <td><button @click="deleteDoctor(doctor.id)">Delete</button></td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getAllPatients, deletePatientById } from '@/api/patients'
import { getAllDoctors, deleteDoctorById } from '@/api/doctors'

const patients = ref([])
const doctors = ref([])

const loadData = async () => {
  patients.value = await getAllPatients()
  doctors.value = await getAllDoctors()
}

const deletePatient = async (id) => {
  await deletePatientById(id)
  await loadData()
}

const deleteDoctor = async (id) => {
  await deleteDoctorById(id)
  await loadData()
}

const backupDatabase = () => {
  await backupDBInit(id)
  await loadData()
}

const restoreDatabase = () => {
  await restoreDB(id)
  await loadData()
}

onMounted(loadData)
</script>

<style scoped>
.admin-container {
  max-width: 1000px;
  margin: auto;
  padding: 20px;
}

.admin-actions button {
  margin-right: 10px;
  padding: 10px;
}

.admin-section {
  margin-top: 30px;
}

table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 10px;
}

th, td {
  border: 1px solid #ddd;
  padding: 10px;
}

th {
  background-color: #f5f5f5;
}

button {
  padding: 5px 10px;
}
</style>
