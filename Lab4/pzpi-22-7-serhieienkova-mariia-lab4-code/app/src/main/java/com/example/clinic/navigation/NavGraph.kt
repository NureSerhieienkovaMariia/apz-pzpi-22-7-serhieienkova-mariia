package com.example.clinic.navigation

import androidx.compose.runtime.Composable
import androidx.navigation.NavHostController
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import androidx.navigation.navArgument
import androidx.navigation.NavType
import com.example.clinic.ui.*

sealed class Screen(val route: String) {
    object Login : Screen("login")
    object SignUp : Screen("sign_up")
    object RelativeHome : Screen("relative_home")
    object PatientHome : Screen("patient_home?token={token}")
    object PatientDetails : Screen("patient_details/{id}?token={token}")
    object PatientNotes : Screen("patient_notes/{id}?token={token}")
    object PatientHealthNotes : Screen("patient_health_notes?token={token}&id={id}")
}

@Composable
fun AppNavGraph(navController: NavHostController) {
    NavHost(navController = navController, startDestination = Screen.Login.route) {

        composable(Screen.Login.route) {
            LoginScreen(
                onNavigateToSignUp = { navController.navigate(Screen.SignUp.route) },
                onLoginSuccess = { token, userType ->
                    if (userType == "patient") {
                        navController.navigate("patient_home?token=$token")
                    } else {
                        navController.navigate("relative_home?token=$token")
                    }
                }
            )
        }

        composable(Screen.SignUp.route) {
            SignUpScreen(
                onNavigateToLogin = { navController.popBackStack() },
                onSignUpSuccess = { token, userType ->
                    if (userType == "patient") {
                        navController.navigate("patient_home?token=$token")
                    } else {
                        navController.navigate("relative_home?token=$token")
                    }
                }
            )
        }

        composable(
            route = Screen.RelativeHome.route + "?token={token}",
            arguments = listOf(navArgument("token") { type = NavType.StringType })
        ) { backStackEntry ->
            val token = backStackEntry.arguments?.getString("token") ?: ""
            RelativeHomeScreen(
                accessToken = token,
                onLogout = {
                    navController.navigate(Screen.Login.route) {
                        popUpTo(Screen.Login.route) { inclusive = true }
                    }
                },
                onViewPatientDetails = { patientId ->
                    navController.navigate("patient_details/$patientId?token=$token")
                },
                onViewPatientNotes = { patientId ->
                    navController.navigate("patient_notes/$patientId?token=$token")
                }
            )
        }

        composable(
            route = "patient_details/{id}?token={token}",
            arguments = listOf(
                navArgument("id") { type = NavType.IntType },
                navArgument("token") { type = NavType.StringType }
            )
        ) { backStackEntry ->
            val id = backStackEntry.arguments?.getInt("id") ?: -1
            val token = backStackEntry.arguments?.getString("token") ?: ""
            PatientDetailsScreen(
                patientId = id,
                token = token,
                onBack = { navController.popBackStack() }
            )
        }

        composable(
            route = "patient_notes/{id}?token={token}",
            arguments = listOf(
                navArgument("id") { type = NavType.IntType },
                navArgument("token") { type = NavType.StringType }
            )
        ) { backStackEntry ->
            val id = backStackEntry.arguments?.getInt("id") ?: -1
            val token = backStackEntry.arguments?.getString("token") ?: ""
            PatientNotesScreen(
                patientId = id,
                token = token,
                onBack = { navController.popBackStack() }
            )
        }

        composable(
            route = Screen.PatientHome.route,
            arguments = listOf(navArgument("token") { type = NavType.StringType })
        ) { backStackEntry ->
            val token = backStackEntry.arguments?.getString("token") ?: ""
            PatientHomeScreen(
                token = token,
                onLogout = {
                    navController.navigate(Screen.Login.route) {
                        popUpTo(Screen.Login.route) { inclusive = true }
                    }
                },
                onOpenHealthNotes = { patientId ->
                    navController.navigate("patient_health_notes?token=$token&id=$patientId")
                }
            )
        }

        composable(
            route = Screen.PatientHealthNotes.route,
            arguments = listOf(
                navArgument("token") { type = NavType.StringType },
                navArgument("id") { type = NavType.IntType }
            )
        ) { backStackEntry ->
            val token = backStackEntry.arguments?.getString("token") ?: ""
            val id = backStackEntry.arguments?.getInt("id") ?: -1
            PatientHealthNotesScreen(
                token = token,
                patientId = id,
                onBack = { navController.popBackStack() }
            )
        }
    }
}
