import { Suspense } from "react";
import { ErrorBoundary } from "react-error-boundary";
import { GoogleOAuthProvider } from "@react-oauth/google";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { CssBaseline } from "@mui/material";
import { ThemeProvider } from "@mui/material/styles";
import { SnackbarProvider } from "notistack";
import ErrorPage from "./components/pages/ErrorPage/ErrorPage";
import LocalisationFallbackPage from "./components/pages/LocalisationFallbackPage/LocalisationFallbackPage";
import createTheme from "./theme/createTheme";
import Router from "./router/Router";
import useClientStore from "./store/clientStore";
import useChangeLanguageEffect from "./hooks/useChangeLanguageEffect";
import { CLIENT_ID } from "./constants/auth";

const queryClient = new QueryClient();

function App() {
  const themeMode = useClientStore((state) => state.themeMode);

  useChangeLanguageEffect();

  const theme = createTheme({ mode: themeMode });

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <ErrorBoundary FallbackComponent={ErrorPage}>
        <QueryClientProvider client={queryClient}>
          <GoogleOAuthProvider clientId={CLIENT_ID}>
            <SnackbarProvider />
            <Suspense fallback={<LocalisationFallbackPage />}>
              <Router />
            </Suspense>
          </GoogleOAuthProvider>
        </QueryClientProvider>
      </ErrorBoundary>
    </ThemeProvider>
  );
}

export default App;
