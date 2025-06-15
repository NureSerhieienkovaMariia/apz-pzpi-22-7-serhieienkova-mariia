const API_BASE = 'http://localhost:8087/doctor/api'

export const getMedicines = async (accessToken) => {
    const res = await fetch(`${API_BASE}/medicine/`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${accessToken}`,
        },
    })

    if (!res.ok) {
        throw new Error(`Medicines fetch failed: ${res.status}`)
    }

    return await res.json()
}

export const createMedicine = async (accessToken, medicineData) => {
    const res = await fetch(`${API_BASE}/medicine/`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${accessToken}`,
        },
        body: JSON.stringify(medicineData),
    })

    if (!res.ok) {
        throw new Error(`Create medicine failed: ${res.status}`)
    }

    return await res.json()
}

export const updateMedicine = async (accessToken, id, medicineData) => {
    const res = await fetch(`${API_BASE}/medicine/${id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${accessToken}`,
        },
        body: JSON.stringify(medicineData),
    })

    if (!res.ok) {
        throw new Error(`Update medicine failed: ${res.status}`)
    }

    return await res.json()
}

export const deleteMedicine = async (accessToken, id) => {
    const res = await fetch(`${API_BASE}/medicine/${id}`, {
        method: 'DELETE',
        headers: {
            'Authorization': `Bearer ${accessToken}`,
        },
    })

    if (!res.ok) {
        throw new Error(`Delete medicine failed: ${res.status}`)
    }

    return
}
