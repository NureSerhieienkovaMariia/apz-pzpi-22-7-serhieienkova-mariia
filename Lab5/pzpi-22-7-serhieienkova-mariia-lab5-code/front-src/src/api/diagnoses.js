const API_BASE = 'http://localhost:8087/doctor/api'

export const getDiagnoses = async (accessToken) => {
    const res = await fetch(`${API_BASE}/diagnosis/`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${accessToken}`,
        },
    })

    if (!res.ok) {
        throw new Error(`Diagnoses fetch failed: ${res.status}`)
    }

    return await res.json()
}

export const createDiagnosis = async (accessToken, diagnosisData) => {
    const res = await fetch(`${API_BASE}/diagnosis/`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${accessToken}`,
        },
        body: JSON.stringify(diagnosisData),
    })

    if (!res.ok) {
        throw new Error(`Create diagnosis failed: ${res.status}`)
    }

    return await res.json()
}

export const updateDiagnosis = async (accessToken, id, diagnosisData) => {
    const res = await fetch(`${API_BASE}/diagnosis/${id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${accessToken}`,
        },
        body: JSON.stringify(diagnosisData),
    })

    if (!res.ok) {
        throw new Error(`Update diagnosis failed: ${res.status}`)
    }

    return await res.json()
}

export const deleteDiagnosis = async (accessToken, id) => {
    const res = await fetch(`${API_BASE}/diagnosis/${id}`, {
        method: 'DELETE',
        headers: {
            'Authorization': `Bearer ${accessToken}`,
        },
    })

    if (!res.ok) {
        throw new Error(`Delete diagnosis failed: ${res.status}`)
    }

    return
}

export const attachDiagnosisToPatient = async (accessToken, patientId, diagnosisId) => {
    const res = await fetch(`${API_BASE}/patients/attach-diagnosis`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${accessToken}`,
        },
        body: JSON.stringify({
            patient_id: patientId,
            diagnosis_id: diagnosisId,
        }),
    })

    if (!res.ok) {
        throw new Error(`Attach diagnosis failed: ${res.status}`)
    }

    return await res.json()
}
