<template>
  <div class="medicines-container">
    <h2>{{ $t('medicines') }}</h2>

    <div v-if="error" class="error">
      {{ error }}
    </div>

    <button @click="showCreateForm = true" class="create-button">{{ $t('createMedicine') }}</button>

    <!-- Create form -->
    <div v-if="showCreateForm" class="form-block">
      <h3>{{ $t('createNewMedicine') }}</h3>
      <input v-model="newMedicine.name" :placeholder="$t('name')" />
      <input v-model="newMedicine.description" :placeholder="$t('description')" />
      <button @click="createNewMedicine">{{ $t('save') }}</button>
      <button @click="showCreateForm = false">{{ $t('cancel') }}</button>
    </div>

    <table v-if="medicines.length > 0" class="medicines-table">
      <thead>
      <tr>
        <th>ID</th>
        <th>{{ $t('name') }}</th>
        <th>{{ $t('description') }}</th>
        <th>{{ $t('actions') }}</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="medicine in medicines" :key="medicine.id">
        <td>{{ medicine.id }}</td>
        <td v-if="editingId !== medicine.id">{{ medicine.name }}</td>
        <td v-if="editingId === medicine.id">
          <input v-model="editMedicine.name" />
        </td>

        <td v-if="editingId !== medicine.id">{{ medicine.description }}</td>
        <td v-if="editingId === medicine.id">
          <input v-model="editMedicine.description" />
        </td>

        <td>
          <button v-if="editingId !== medicine.id" @click="startEdit(medicine)">{{ $t('edit') }}</button>
          <button v-if="editingId === medicine.id" @click="saveEdit(medicine.id)">{{ $t('save') }}</button>
          <button v-if="editingId === medicine.id" @click="cancelEdit">{{ $t('cancel') }}</button>
          <button @click="removeMedicine(medicine.id)">{{ $t('delete') }}</button>
        </td>
      </tr>
      </tbody>
    </table>

    <div v-else-if="!error">
      {{ $t('noMedicines') }}
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getMedicines, createMedicine, updateMedicine, deleteMedicine } from '../api/medicines'
import { useRouter } from 'vue-router'

const medicines = ref([])
const error = ref(null)
const router = useRouter()

const showCreateForm = ref(false)
const newMedicine = ref({ name: '', description: '' })

const editingId = ref(null)
const editMedicine = ref({ name: '', description: '' })

const fetchMedicines = async () => {
  error.value = null
  try {
    const token = localStorage.getItem('accessToken')
    if (!token) {
      router.push('/login')
      return
    }

    const response = await getMedicines(token)
    medicines.value = response
  } catch (err) {
    console.error(err)
    error.value = 'Помилка отримання списку ліків: ' + err.message
  }
}

const createNewMedicine = async () => {
  try {
    const token = localStorage.getItem('accessToken')
    await createMedicine(token, newMedicine.value)
    showCreateForm.value = false
    newMedicine.value = { name: '', description: '' }
    fetchMedicines()
  } catch (err) {
    console.error(err)
    error.value = 'Помилка створення ліків: ' + err.message
  }
}

const startEdit = (medicine) => {
  editingId.value = medicine.id
  editMedicine.value = { name: medicine.name, description: medicine.description }
}

const saveEdit = async (id) => {
  try {
    const token = localStorage.getItem('accessToken')
    await updateMedicine(token, id, editMedicine.value)
    editingId.value = null
    fetchMedicines()
  } catch (err) {
    console.error(err)
    error.value = 'Помилка оновлення ліків: ' + err.message
  }
}

const cancelEdit = () => {
  editingId.value = null
}

const removeMedicine = async (id) => {
  try {
    const token = localStorage.getItem('accessToken')
    await deleteMedicine(token, id)
    fetchMedicines()
  } catch (err) {
    console.error(err)
    error.value = 'Помилка видалення ліків: ' + err.message
  }
}

onMounted(() => {
  fetchMedicines()
})
</script>

<style scoped>
.medicines-container {
  max-width: 1000px;
  margin: 30px auto;
  background-color: white;
  padding: 20px 30px;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.medicines-table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 20px;
}

.medicines-table th,
.medicines-table td {
  border: 1px solid #ddd;
  padding: 8px;
  text-align: center;
}

.medicines-table th {
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
