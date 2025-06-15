package com.example.clinic.localization

import android.app.LocaleManager
import android.os.Build
import androidx.compose.runtime.*
import androidx.compose.ui.platform.LocalContext
import java.util.*

val LocalAppLocale = compositionLocalOf { Locale("en") }

@Composable
fun ProvideLocalizedApp(content: @Composable (Locale, (Locale) -> Unit) -> Unit) {
    val context = LocalContext.current
    var locale by remember { mutableStateOf(Locale.getDefault()) }

    val updateLocale: (Locale) -> Unit = {
        locale = it
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.TIRAMISU) {
            context.getSystemService(LocaleManager::class.java)
                ?.applicationLocales = android.os.LocaleList.forLanguageTags(it.toLanguageTag())
        }
    }

    CompositionLocalProvider(LocalAppLocale provides locale) {
        SideEffect {
            Locale.setDefault(locale)
        }
        content(locale, updateLocale)
    }
}
