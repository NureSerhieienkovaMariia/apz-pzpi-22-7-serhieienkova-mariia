const API_BASE = 'http://localhost:8087/doctor/auth'

export const signUp = async ({ name, surname, email, password }) => {
    const res = await fetch(`${API_BASE}/sign-up`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ name, surname, email, password }),
    })

    if (!res.ok) {
        throw new Error(`Sign up failed: ${res.status}`)
    }

    return await res.json()
}

export const signIn = async ({ email, password }) => {
    const res = await fetch(`${API_BASE}/sign-in`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password }),
    })

    if (!res.ok) {
        throw new Error(`Sign in failed: ${res.status}`)
    }

    return await res.json()
}

export const getCurrentUser = async (accessToken) => {
    const res = await fetch(`${API_BASE}/current-user`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${accessToken}`,
        },
    })

    if (!res.ok) {
        throw new Error(`Current user failed: ${res.status}`)
    }

    return await res.json()
}
