import { createI18n } from 'vue-i18n'
import en from './locales/en.json'
import ua from './locales/ua.json'

const savedLocale = localStorage.getItem('locale') || 'ua'

const i18n = createI18n({
    legacy: false,
    locale: savedLocale,
    fallbackLocale: 'en',
    messages: {
        en,
        ua,
    },
})

export default i18n
