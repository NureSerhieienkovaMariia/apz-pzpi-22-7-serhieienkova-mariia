const API_BASE = 'http://localhost:8087/doctor/api'

export const createTreatmentPlan = async (accessToken, patientId, doctorId, startDate, endDate) => {
    const res = await fetch(`${API_BASE}/treatmentplan/`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${accessToken}`,
        },
        body: JSON.stringify({
            patient_id: patientId,
            doctor_id: doctorId,
            start_date: startDate,
            end_date: endDate,
        }),
    })

    if (!res.ok) {
        throw new Error(`Create treatment plan failed: ${res.status}`)
    }

    return await res.json()
}

export const deleteTreatmentPlan = async (accessToken, treatmentPlanId) => {
    const res = await fetch(`${API_BASE}/treatmentplan/${treatmentPlanId}`, {
        method: 'DELETE',
        headers: {
            'Authorization': `Bearer ${accessToken}`,
        },
    })

    if (!res.ok) {
        throw new Error(`Delete treatment plan failed: ${res.status}`)
    }
}
