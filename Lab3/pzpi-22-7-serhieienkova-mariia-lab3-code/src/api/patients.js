const API_BASE = 'http://localhost:8087/doctor/api'

export const getPatients = async (accessToken) => {
    const res = await fetch(`${API_BASE}/patients/`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${accessToken}`,
        },
    })

    if (!res.ok) {
        throw new Error(`Patients fetch failed: ${res.status}`)
    }

    return await res.json()
}
