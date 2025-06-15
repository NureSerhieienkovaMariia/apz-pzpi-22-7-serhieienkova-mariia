<template>
  <div class="diagnoses-container">
    <h2>{{ $t('diagnoses') }}</h2>

    <div v-if="error" class="error">
      {{ error }}
    </div>

    <button @click="showCreateForm = true" class="create-button">{{ $t('createDiagnosis') }}</button>

    <!-- Create form -->
    <div v-if="showCreateForm" class="form-block">
      <h3>{{ $t('createNewDiagnosis') }}</h3>
      <input v-model="newDiagnosis.name" :placeholder="$t('name')" />
      <input v-model="newDiagnosis.description" :placeholder="$t('description')" />
      <input v-model="newDiagnosis.recommendations" :placeholder="$t('recommendations')" />
      <button @click="createNewDiagnosis">{{ $t('save') }}</button>
      <button @click="showCreateForm = false">{{ $t('cancel') }}</button>
    </div>

    <table v-if="diagnoses.length > 0" class="diagnoses-table">
      <thead>
      <tr>
        <th>ID</th>
        <th>{{ $t('name') }}</th>
        <th>{{ $t('description') }}</th>
        <th>{{ $t('recommendations') }}</th>
        <th>{{ $t('actions') }}</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="diagnosis in diagnoses" :key="diagnosis.id">
        <td>{{ diagnosis.id }}</td>
        <td v-if="editingId !== diagnosis.id">{{ diagnosis.name }}</td>
        <td v-if="editingId === diagnosis.id">
          <input v-model="editDiagnosis.name" />
        </td>

        <td v-if="editingId !== diagnosis.id">{{ diagnosis.description }}</td>
        <td v-if="editingId === diagnosis.id">
          <input v-model="editDiagnosis.description" />
        </td>

        <td v-if="editingId !== diagnosis.id">{{ diagnosis.recommendations }}</td>
        <td v-if="editingId === diagnosis.id">
          <input v-model="editDiagnosis.recommendations" />
        </td>

        <td>
          <button v-if="editingId !== diagnosis.id" @click="startEdit(diagnosis)">{{ $t('edit') }}</button>
          <button v-if="editingId === diagnosis.id" @click="saveEdit(diagnosis.id)">{{ $t('save') }}</button>
          <button v-if="editingId === diagnosis.id" @click="cancelEdit">{{ $t('cancel') }}</button>
          <button @click="removeDiagnosis(diagnosis.id)">{{ $t('delete') }}</button>
        </td>
      </tr>
      </tbody>
    </table>

    <div v-else-if="!error">
      {{ $t('noDiagnoses') }}
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getDiagnoses, createDiagnosis, updateDiagnosis, deleteDiagnosis } from '../api/diagnoses'
import { useRouter } from 'vue-router'

const diagnoses = ref([])
const error = ref(null)
const router = useRouter()

const showCreateForm = ref(false)
const newDiagnosis = ref({ name: '', description: '', recommendations: '' })

const editingId = ref(null)
const editDiagnosis = ref({ name: '', description: '', recommendations: '' })

const fetchDiagnoses = async () => {
  error.value = null
  try {
    const token = localStorage.getItem('accessToken')
    if (!token) {
      router.push('/login')
      return
    }

    const response = await getDiagnoses(token)
    diagnoses.value = response
  } catch (err) {
    console.error(err)
    error.value = 'Помилка отримання списку діагнозів: ' + err.message
  }
}

const createNewDiagnosis = async () => {
  try {
    const token = localStorage.getItem('accessToken')
    await createDiagnosis(token, newDiagnosis.value)
    showCreateForm.value = false
    newDiagnosis.value = { name: '', description: '', recommendations: '' }
    fetchDiagnoses()
  } catch (err) {
    console.error(err)
    error.value = 'Помилка створення діагнозу: ' + err.message
  }
}

const startEdit = (diagnosis) => {
  editingId.value = diagnosis.id
  editDiagnosis.value = {
    name: diagnosis.name,
    description: diagnosis.description,
    recommendations: diagnosis.recommendations
  }
}

const saveEdit = async (id) => {
  try {
    const token = localStorage.getItem('accessToken')
    await updateDiagnosis(token, id, editDiagnosis.value)
    editingId.value = null
    fetchDiagnoses()
  } catch (err) {
    console.error(err)
    error.value = 'Помилка оновлення діагнозу: ' + err.message
  }
}

const cancelEdit = () => {
  editingId.value = null
}

const removeDiagnosis = async (id) => {
  try {
    const token = localStorage.getItem('accessToken')
    await deleteDiagnosis(token, id)
    fetchDiagnoses()
  } catch (err) {
    console.error(err)
    error.value = 'Помилка видалення діагнозу: ' + err.message
  }
}

onMounted(() => {
  fetchDiagnoses()
})
</script>

<style scoped>
.diagnoses-container {
  max-width: 1000px;
  margin: 30px auto;
  background-color: white;
  padding: 20px 30px;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.diagnoses-table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 20px;
}

.diagnoses-table th,
.diagnoses-table td {
  border: 1px solid #ddd;
  padding: 8px;
  text-align: center;
}

.diagnoses-table th {
  background-color: #f2f2f2;
}

.error {
  color: red;
  margin-bottom: 10px;
}

.create-button {
  padding: 8px 12px;
  margin-bottom: 15px;
  background-color: #4caf50;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.create-button:hover {
  background-color: #45a049;
}

.form-block {
  margin-bottom: 15px;
}

.form-block input {
  display: block;
  margin-bottom: 8px;
  padding: 6px;
  width: 100%;
}
</style>
