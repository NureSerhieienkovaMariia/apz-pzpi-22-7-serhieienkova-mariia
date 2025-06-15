package com.example.clinic

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.compose.material3.MaterialTheme
import androidx.navigation.compose.rememberNavController
import com.example.clinic.navigation.AppNavGraph
import com.example.clinic.localization.ProvideLocalizedApp

class MainActivity : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContent {
            ProvideLocalizedApp { _, onChangeLocale ->
                MaterialTheme {
                    val navController = rememberNavController()
                    AppNavGraph(
                        navController = navController,
                    )
                }
            }
        }
    }
}
