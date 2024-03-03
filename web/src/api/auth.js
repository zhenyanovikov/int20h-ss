import { useQueryClient } from "@tanstack/react-query";
import { useGoogleLogin } from "@react-oauth/google";
import { enqueueSnackbar } from "notistack";
import { useTranslation } from "react-i18next";
import useClientStore from "../store/clientStore";
import { CALLBACK_URL } from "../constants/auth";
import apiClient from "./apiClient";

function login({ code }) {
  return apiClient.get(CALLBACK_URL, {
    params: {
      code,
    },
  });
}

export function useLogin() {
  const { t } = useTranslation();
  const setToken = useClientStore((state) => state.setToken);

  return useGoogleLogin({
    flow: "auth-code",
    onSuccess: async (codeResponse) => {
      try {
        const response = await login({ code: codeResponse.code });

        setToken(response.data.accessToken);
      } catch (_) {
        enqueueSnackbar(t("snackbars.login.error"), { variant: "error" });
      }
    },
    onError: () => {
      enqueueSnackbar(t("snackbars.login.error"), { variant: "error" });
    },
  });
}

export function useLogout() {
  const setToken = useClientStore((state) => state.setToken);
  const queryClient = useQueryClient();

  return () => {
    queryClient.clear();
    setToken(null);
  };
}
