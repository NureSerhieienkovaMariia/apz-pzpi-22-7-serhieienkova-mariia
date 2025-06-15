import { ref } from 'vue'

const user = ref(JSON.parse(localStorage.getItem('user')) || null)

export const useUser = () => {
    const setUser = (newUser) => {
        user.value = newUser
    }

    const clearUser = () => {
        user.value = null
    }

    return {
        user,
        setUser,
        clearUser,
    }
}
