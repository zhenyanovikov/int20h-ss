import { useQueryClient, useQuery, useMutation } from "@tanstack/react-query";
import { enqueueSnackbar } from "notistack";
import { useTranslation } from "react-i18next";
import { GROUPS_KEY, GROUPS_URL } from "../constants/group";
import apiClient from "./apiClient";

async function getGroups() {
  const res = await apiClient.get(GROUPS_URL);
  return res.data;
}

async function createGroup(group) {
  const res = await apiClient.post(GROUPS_URL, group);
  return res.data;
}

export function useGetGroups() {
  return useQuery({
    queryKey: [GROUPS_KEY],
    queryFn: getGroups,
  });
}

export function useCreateGroup() {
  const { t } = useTranslation();
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: createGroup,
    onSuccess: () => {
      enqueueSnackbar(t("snackbars.createGroup.success"), {
        variant: "success",
      });
      queryClient.invalidateQueries(GROUPS_KEY);
    },
    onError: () => {
      enqueueSnackbar(t("snackbars.createGroup.error"), {
        variant: "error",
      });
    },
  });
}
