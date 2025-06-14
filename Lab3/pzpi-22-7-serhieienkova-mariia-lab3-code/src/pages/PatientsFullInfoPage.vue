<template>
  <div class="full-info-container">
    <h2>{{ $t('patients') }} #{{ patientId }} - {{ patient?.name }} {{ patient?.surname }}</h2>

    <div v-if="error" class="error">{{ error }}</div>

    <div v-if="patient">
      <p><strong>{{ $t('email') }}:</strong> {{ patient.email }}</p>
      <p><strong>{{ $t('birthday') }}:</strong> {{ formatDate(patient.birthday) }}</p>
      <p><strong>{{ $t('sex') }}:</strong> {{ formatSex(patient.sex) }}</p>

      <!-- Diagnoses Block -->
      <div class="full-info-block">
        <h3 class="block-title">{{ $t('diagnoses') }}</h3>

        <div class="attach-diagnosis-form" v-if="allDiagnoses.length > 0">
          <select v-model="selectedDiagnosisId">
            <option disabled value="">{{ $t('selectDiagnosis') }}</option>
            <option v-for="diag in allDiagnoses" :key="diag.id" :value="diag.id">
              {{ diag.name }}
            </option>
          </select>
          <button @click="handleAttachDiagnosis">{{ $t('addDiagnosis') }}</button>
        </div>

        <div v-if="patient && patient.diagnoses.length > 0" class="diagnoses-grid">
          <div v-for="diag in patient.diagnoses" :key="diag.id" class="card">
            <h4>{{ diag.name }}</h4>
            <p>{{ diag.description }}</p>
            <p><em>{{ diag.recommendations }}</em></p>
          </div>
        </div>
        <div v-else>{{ $t('noDiagnoses') }}</div>
      </div>

      <!-- Relatives Block -->
      <div class="full-info-block">
        <h3 class="block-title">{{ $t('relatives') }}</h3>

        <!-- Attach relative form -->
        <div class="attach-relative-form" v-if="allRelatives.length > 0">
          <select v-model.number="selectedRelativeId">
            <option disabled value="">{{ $t('selectRelative') }}</option>
            <option v-for="rel in allRelatives" :key="rel.id" :value="rel.id">
              {{ rel.name }} {{ rel.surname }} ({{ rel.email }})
            </option>
          </select>
          <label>
            <input type="checkbox" v-model="accessToRecords" />
            {{ $t('accessToRecords') }}
          </label>
          <button @click="handleAttachRelative">{{ $t('addRelative') }}</button>
        </div>

        <div v-if="patientRelatives.length > 0" class="relatives-grid">
          <div v-for="rel in patientRelatives" :key="rel.id" class="card">
            <h4>{{ rel.name }} {{ rel.surname }}</h4>
            <p>{{ rel.email }}</p>
          </div>
        </div>
        <div v-else>{{ $t('noRelatives') }}</div>
      </div>

      <!-- Treatment Plans Block -->
      <div class="full-info-block">
        <h3 class="block-title">{{ $t('treatmentPlans') }}</h3>

        <!-- Створення TreatmentPlan -->
        <div class="create-treatment-plan-form">
          <input type="date" v-model="newTreatmentPlanStartDate" class="date-input" />
          <input type="date" v-model="newTreatmentPlanEndDate" class="date-input" />
          <button @click="handleCreateTreatmentPlan">{{ $t('createTreatmentPlan') }}</button>
        </div>

        <!-- Навігація + показ плану -->
        <div v-if="patient.treatment_plans.length > 0">
          <!-- Поточний TreatmentPlan -->
          <div v-if="currentPlan" class="treatment-plan-card">
            <p><strong>{{ $t('doctor') }}:</strong> {{ currentPlan.doctor?.name }} {{ currentPlan.doctor?.surname }}</p>
            <p><strong>{{ $t('startDate') }}:</strong> {{ formatDate(currentPlan.start_date) }}</p>
            <p><strong>{{ $t('endDate') }}:</strong> {{ formatDate(currentPlan.end_date) }}</p>

            <!-- Видалити TreatmentPlan -->
            <button @click="handleDeleteTreatmentPlan(currentPlan.id)" class="delete-treatment-plan-button">
              {{ $t('delete') }}
            </button>

            <!-- Visits -->
            <div class="sub-block">
              <h4>{{ $t('visits') }}</h4>
              <div v-if="currentPlan.visits.length > 0">
                <ul>
                  <li v-for="visit in currentPlan.visits" :key="visit.id">
                    {{ formatDate(visit.date) }} — {{ visit.reason }}<br />
                    <em>{{ visit.notes }}</em>
                  </li>
                </ul>
              </div>
              <div v-else>{{ $t('noVisits') }}</div>
            </div>

            <!-- Prescriptions -->
            <div class="sub-block">
              <h4>{{ $t('prescriptions') }}</h4>
              <div v-if="currentPlan.prescriptions.length > 0">
                <ul>
                  <li v-for="presc in currentPlan.prescriptions" :key="presc.medicine.id">
                    {{ presc.medicine.name }} — {{ presc.dosage }}, {{ presc.frequency }}
                  </li>
                </ul>
              </div>
              <div v-else>{{ $t('noPrescriptions') }}</div>
            </div>
          </div>
          <!-- Стрілочки Prev / Next -->
          <div class="treatment-plan-navigation">
            <button @click="prevTreatmentPlan" :disabled="currentTreatmentPlanIndex.value === 0">⬅️</button>
            <span>{{ currentTreatmentPlanIndex + 1 }} / {{ patient.treatment_plans.length }}</span>
            <button @click="nextTreatmentPlan" :disabled="currentTreatmentPlanIndex.value === patient.treatment_plans.length - 1">➡️</button>
          </div>


          <!-- Create Visit -->
          <div class="create-visit-form">
            <input v-model="getVisitFieldsForPlan(currentPlan.id).reason" type="text" placeholder="Reason" />
            <input v-model="getVisitFieldsForPlan(currentPlan.id).date" type="datetime-local" />
            <textarea v-model="getVisitFieldsForPlan(currentPlan.id).notes" placeholder="Notes"></textarea>
            <button @click="handleCreateVisit(currentPlan.id)">{{ $t('createVisit') }}</button>
          </div>


          <!-- Create Prescription -->
          <div class="create-prescription-form">
            <select v-model="getPrescriptionFieldsForPlan(currentPlan.id).medicineId">
              <option disabled value="">{{ $t('selectMedicine') }}</option>
              <option v-for="med in medicines" :key="med.id" :value="med.id">
                {{ med.name }}
              </option>
            </select>

            <input v-model="getPrescriptionFieldsForPlan(currentPlan.id).dosage" type="text" placeholder="Dosage" />

            <select v-model="getPrescriptionFieldsForPlan(currentPlan.id).freqType">
              <option value="at">at</option>
              <option value="every">every</option>
            </select>

            <template v-if="getPrescriptionFieldsForPlan(currentPlan.id).freqType === 'at'">
              <input v-model="getPrescriptionFieldsForPlan(currentPlan.id).time" type="time" />
            </template>
            <template v-else>
              <input v-model.number="getPrescriptionFieldsForPlan(currentPlan.id).hours" type="number" min="1" placeholder="Hours" />
            </template>

            <button @click="handleCreatePrescription(currentPlan.id)">{{ $t('createPrescription') }}</button>
          </div>

        </div>

        <!-- Якщо TreatmentPlans немає -->
        <div v-else>{{ $t('noTreatmentPlans') }}</div>
      </div>

    </div>

    <div v-else-if="!error">
      Loading...
    </div>

    <router-link to="/patients" class="back-button">{{ $t('backToPatients') }}</router-link>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { getDiagnoses, attachDiagnosisToPatient } from '../api/diagnoses'
import { getRelatives, getRelativesByPatient, attachRelativeToPatient } from '../api/relatives'
import { createTreatmentPlan, deleteTreatmentPlan } from '../api/treatmentplans'
import { useUser } from '../composables/useUser'
import {computed} from 'vue'

const { user } = useUser()

const route = useRoute()
const router = useRouter()

const patientId = route.params.id
const patient = ref(null)
const error = ref(null)

const allDiagnoses = ref([])
const selectedDiagnosisId = ref('')

const patientRelatives = ref([])
const allRelatives = ref([])
const selectedRelativeId = ref('')
const accessToRecords = ref(false)

const newTreatmentPlanStartDate = ref('')
const newTreatmentPlanEndDate = ref('')

const currentTreatmentPlanIndex = ref(0)

const newVisitReason = ref('')
const newVisitDate = ref('')
const newVisitNotes = ref('')

const newPrescriptionMedicineId = ref('')
const newPrescriptionDosage = ref('')
const newPrescriptionFrequency = ref('')

const newPrescriptionFrequencyType = ref('at') // 'at' або 'every'
const newPrescriptionTime = ref('') // для 'at'
const newPrescriptionHours = ref(null) // для 'every'

const medicines = ref([])

// --- State для Visit ---
const newVisitFields = ref({})

const getVisitFieldsForPlan = (planId) => {
  if (!newVisitFields.value[planId]) {
    newVisitFields.value[planId] = {
      reason: '',
      date: '',
      notes: '',
    }
  }
  return newVisitFields.value[planId]
}

// --- State для Prescription ---
const newPrescriptionFields = ref({})

const getPrescriptionFieldsForPlan = (planId) => {
  if (!newPrescriptionFields.value[planId]) {
    newPrescriptionFields.value[planId] = {
      medicineId: '',
      dosage: '',
      freqType: 'at',
      time: '',
      hours: null,
    }
  }
  return newPrescriptionFields.value[planId]
}

// --- Handle Create Visit ---
const handleCreateVisit = async (treatmentPlanId) => {
  const fields = newVisitFields.value[treatmentPlanId]

  try {
    const token = localStorage.getItem('accessToken')
    await fetch('http://localhost:8087/doctor/api/visit/', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`,
      },
      body: JSON.stringify({
        treatment_plan_id: treatmentPlanId,
        reason: fields.reason,
        date: new Date(fields.date).toISOString(),
        notes: fields.notes,
      }),
    })

    await fetchPatientFullInfo()

    newVisitFields.value[treatmentPlanId] = {
      reason: '',
      date: '',
      notes: '',
    }
  } catch (err) {
    console.error('Помилка створення Visit:', err)
    alert('Помилка створення Visit: ' + err.message)
  }
}

// --- Handle Create Prescription ---
const handleCreatePrescription = async (treatmentPlanId) => {
  const fields = newPrescriptionFields.value[treatmentPlanId]

  try {
    const token = localStorage.getItem('accessToken')

    let frequency = ''
    if (fields.freqType === 'at') {
      if (!fields.time) {
        alert('Будь ласка, вкажіть час!')
        return
      }
      frequency = `at ${fields.time}`
    } else {
      if (!fields.hours || fields.hours <= 0) {
        alert('Будь ласка, вкажіть кількість годин!')
        return
      }
      frequency = `every ${fields.hours} hour(s)`
    }

    await fetch('http://localhost:8087/doctor/api/prescription/', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`,
      },
      body: JSON.stringify({
        treatment_plan_id: treatmentPlanId,
        medicine_id: Number(fields.medicineId),
        dosage: fields.dosage,
        frequency: frequency,
      }),
    })

    await fetchPatientFullInfo()

    newPrescriptionFields.value[treatmentPlanId] = {
      medicineId: '',
      dosage: '',
      freqType: 'at',
      time: '',
      hours: null,
    }
  } catch (err) {
    console.error('Помилка створення Prescription:', err)
    alert('Помилка створення Prescription: ' + err.message)
  }
}

const currentPlan = computed(() => {
  return patient.value && patient.value.treatment_plans.length > 0
      ? patient.value.treatment_plans[currentTreatmentPlanIndex.value]
      : null
})

const prevTreatmentPlan = () => {
  if (currentTreatmentPlanIndex.value > 0) {
    currentTreatmentPlanIndex.value--
  }
}

const nextTreatmentPlan = () => {
  if (currentTreatmentPlanIndex.value < patient.value.treatment_plans.length - 1) {
    currentTreatmentPlanIndex.value++
  }
}
const fetchPatientFullInfo = async () => {
  error.value = null
  try {
    const token = localStorage.getItem('accessToken')
    if (!token) {
      router.push('/login')
      return
    }

    const res = await fetch(`http://localhost:8087/doctor/api/patients/${patientId}/full-info`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`,
      },
    })

    if (!res.ok) {
      throw new Error(`Fetch full info failed: ${res.status}`)
    }

    patient.value = await res.json()
    // Страховка treatment_plans
    if (!patient.value.treatment_plans) {
      patient.value.treatment_plans = []
    }
    patient.value.treatment_plans.forEach(plan => {
      if (!plan.visits) {
        plan.visits = []
      }
      if (!plan.prescriptions) {
        plan.prescriptions = []
      }
    })

    // Страховка diagnoses
    if (!patient.value.diagnoses) {
      patient.value.diagnoses = []
    }

    // Страховка relatives
    if (!patient.value.relatives) {
      patient.value.relatives = []
    }
  } catch (err) {
    console.error(err)
    error.value = 'Помилка отримання повної інформації про пацієнта: ' + err.message
  }
}

const fetchAllDiagnoses = async () => {
  try {
    const token = localStorage.getItem('accessToken')
    const response = await getDiagnoses(token)
    allDiagnoses.value = response
  } catch (err) {
    console.error('Помилка отримання діагнозів:', err)
  }
}

const handleAttachDiagnosis = async () => {
  if (!selectedDiagnosisId.value) {
    alert('Будь ласка, виберіть діагноз')
    return
  }

  try {
    const token = localStorage.getItem('accessToken')
    await attachDiagnosisToPatient(token, Number(patientId), Number(selectedDiagnosisId.value))

    await fetchPatientFullInfo()
    selectedDiagnosisId.value = ''
  } catch (err) {
    console.error('Помилка додавання діагнозу:', err)
    alert('Помилка додавання діагнозу: ' + err.message)
  }
}

const fetchPatientRelatives = async () => {
  try {
    const token = localStorage.getItem('accessToken')
    const response = await getRelativesByPatient(token, patientId)
    patientRelatives.value = response
  } catch (err) {
    console.error('Помилка отримання relatives пацієнта:', err)
  }
}

const fetchAllRelatives = async () => {
  try {
    const token = localStorage.getItem('accessToken')
    const response = await getRelatives(token)
    allRelatives.value = response
  } catch (err) {
    console.error('Помилка отримання всіх relatives:', err)
  }
}

const handleAttachRelative = async () => {
  if (selectedRelativeId.value == null || selectedRelativeId.value === '') {
    alert('Будь ласка, виберіть relative')
    return
  }

  try {
    const token = localStorage.getItem('accessToken')
    await attachRelativeToPatient(token, Number(patientId), selectedRelativeId.value, accessToRecords.value)

    await fetchPatientRelatives()
    selectedRelativeId.value = ''
    accessToRecords.value = false
  } catch (err) {
    console.error('Помилка додавання relative:', err)
    alert('Помилка додавання relative: ' + err.message)
  }
}

const handleCreateTreatmentPlan = async () => {
  if (!newTreatmentPlanStartDate.value || !newTreatmentPlanEndDate.value) {
    alert('Будь ласка, заповніть обидві дати')
    return
  }

  try {
    const token = localStorage.getItem('accessToken')
    await createTreatmentPlan(
        token,
        Number(patientId),
        user.value.userId,
        new Date(newTreatmentPlanStartDate.value).toISOString(),
        new Date(newTreatmentPlanEndDate.value).toISOString()
    )
    await fetchPatientFullInfo()

    // Скидуємо форму
    newTreatmentPlanStartDate.value = ''
    newTreatmentPlanEndDate.value = ''
  } catch (err) {
    console.error('Помилка створення TreatmentPlan:', err)
    alert('Помилка створення TreatmentPlan: ' + err.message)
  }
}

const handleDeleteTreatmentPlan = async (treatmentPlanId) => {
  if (!confirm('Ви дійсно хочете видалити TreatmentPlan?')) return

  try {
    const token = localStorage.getItem('accessToken')
    await deleteTreatmentPlan(token, treatmentPlanId)

    await fetchPatientFullInfo()
  } catch (err) {
    console.error('Помилка видалення TreatmentPlan:', err)
    alert('Помилка видалення TreatmentPlan: ' + err.message)
  }
}

const fetchMedicines = async () => {
  try {
    const token = localStorage.getItem('accessToken')
    const res = await fetch('http://localhost:8087/doctor/api/medicine/', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`,
      },
    })

    if (!res.ok) {
      throw new Error(`Fetch medicines failed: ${res.status}`)
    }

    medicines.value = await res.json()
  } catch (err) {
    console.error('Помилка отримання ліків:', err)
  }
}

const formatDate = (dateStr) => {
  if (!dateStr) return '-'
  return new Date(dateStr).toLocaleDateString()
}

const formatSex = (sex) => {
  return sex ? 'Чоловік' : 'Жінка'
}

onMounted(() => {
  fetchPatientFullInfo()
  fetchAllDiagnoses()
  fetchPatientRelatives()
  fetchAllRelatives()
  fetchMedicines()
})
</script>

<style scoped>
.full-info-container {
  max-width: 1000px;
  margin: 30px auto;
  background-color: white;
  padding: 20px 30px;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.error {
  color: red;
  margin-bottom: 10px;
}

.full-info-block {
  margin-bottom: 25px;
}

.block-title {
  margin-bottom: 10px;
  border-bottom: 2px solid #4caf50;
  padding-bottom: 5px;
  font-size: 1.2em;
}

.card, .treatment-plan-card {
  border: 1px solid #ddd;
  padding: 12px 15px;
  border-radius: 6px;
  margin-bottom: 10px;
  background-color: #fafafa;
  box-shadow: 0 1px 4px rgba(0,0,0,0.05);
}

.sub-block {
  margin-top: 10px;
  padding-top: 5px;
  border-top: 1px dashed #ccc;
}

.sub-block h4 {
  margin-bottom: 5px;
}

.back-button {
  display: inline-block;
  margin-top: 20px;
  padding: 6px 12px;
  background-color: #4caf50;
  color: white;
  text-decoration: none;
  border-radius: 4px;
}

.back-button:hover {
  background-color: #45a049;
}

.diagnoses-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 15px;
}

.attach-diagnosis-form {
  display: flex;
  align-items: center;
  margin-bottom: 15px;
}

.attach-diagnosis-form select {
  flex: 1;
  padding: 6px;
  margin-right: 8px;
}

.attach-diagnosis-form button {
  padding: 6px 12px;
  background-color: #4caf50;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.attach-diagnosis-form button:hover {
  background-color: #45a049;
}

.relatives-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 15px;
}

.attach-relative-form {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 15px;
}

.attach-relative-form select {
  flex: 1;
  padding: 6px;
}

.attach-relative-form button {
  padding: 6px 12px;
  background-color: #4caf50;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.attach-relative-form button:hover {
  background-color: #45a049;
}

.create-treatment-plan-form {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 15px;
}

.date-input {
  width: 150px;
  padding: 4px 6px;
}

.create-treatment-plan-form {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 15px;
}

.date-input {
  width: 150px;
  padding: 4px 6px;
}

.treatment-plan-navigation {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 10px;
}

.create-visit-form,
.create-prescription-form {
  display: flex;
  flex-direction: column;
  gap: 8px;
  margin-bottom: 15px;
}
</style>
