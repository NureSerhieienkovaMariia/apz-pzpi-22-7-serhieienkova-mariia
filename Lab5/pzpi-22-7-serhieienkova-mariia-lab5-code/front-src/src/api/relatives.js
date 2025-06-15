const API_BASE = 'http://localhost:8087/doctor/api'

export const getRelatives = async (accessToken) => {
    const res = await fetch(`${API_BASE}/relatives/`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${accessToken}`,
        },
    })

    if (!res.ok) {
        throw new Error(`Get relatives failed: ${res.status}`)
    }

    return await res.json()
}

export const getRelativesByPatient = async (accessToken, patientId) => {
    const res = await fetch(`${API_BASE}/relatives/patient/${patientId}`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${accessToken}`,
        },
    })

    if (!res.ok) {
        throw new Error(`Get patient relatives failed: ${res.status}`)
    }

    return await res.json()
}

export const attachRelativeToPatient = async (accessToken, patientId, relativeId, accessToRecords) => {
    const res = await fetch(`${API_BASE}/patients/attach-relative`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${accessToken}`,
        },
        body: JSON.stringify({
            patient_id: patientId,
            relative_id: relativeId,
            access_to_records: accessToRecords,
        }),
    })

    if (!res.ok) {
        throw new Error(`Attach relative failed: ${res.status}`)
    }

    return await res.json()
}
